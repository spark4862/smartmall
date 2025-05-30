如果你在 go.work 中直接使用 use .，Go 会尝试将当前目录作为模块根目录。但如果当前目录没有定义 go.mod 文件或者模块名与实际不符，就会导致问题。

go work use .只会将当前目录添加到 go.work 文件中，不会添加子模块

go work sync 这个子命令，用于将工作区(go.work)中计算出的依赖版本一致性地同步回各个子模块的 go.mod 文件。简而言之，当你在一个多模块工作区里通过 use 或 replace 指定了本地模块路径，或者当 workspace 内各模块的依赖版本因 Minimal Version Selection（MVS）算法被提升时，go work sync 会： 计算整个工作区的“构建列表”——即在 workspace 模式下所有模块实际使用的确切版本。 将这些确切版本写回到每个子模块的 go.mod，确保单独在模块级别执行 go build、go test 时也能使用与 workspace 模式相同的依赖。

使用个git tag为提交添加标签，这样go get就可以使用标签标识版本
go mod tidy 只看当前模块的 require，不会考虑 go.work 中对本地路径或版本的覆盖；它生成的依赖可能就不是你在 workspace 联调时使用的那套
因此，在多模块 workspace 开发后，一定要用 go work sync 来同步，而不是 go mod tidy。
go work sync用于更新子模块中的依赖为最新依赖，保持和工作区依赖一致，如果以来的模块还未更新，用replace,如果已经更新，打最新的标签


monorepo下sync实践
先commit tag
在sync
最后push


uber go style guide: https://github.com/uber-go/guide
protobuf style guide: https://protobuf.dev/programming-guides/style/
MDN HTTP response status codes：https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Status
约定式提交：https://www.conventionalcommits.org/zh-hans/v1.0.0/
语义化版本：https://semver.org/lang/zh-CN/

正确切换go版本
go env -w GOTOOLCHAIN='go1.24.0'

可以尝试删除go.sum和go.work.sum
你把某个依赖模块切换到了自托管的私有仓库或者本地替换（replace），但原来的 go.sum 里还记录着旧仓库的校验和。
这时 Go 发现校验不对，就会报错，删掉 go.sum 让它在下次 go mod tidy 时重新生成，就不会再和旧校验冲突了。

如果work和mod存在于一个目录中，依赖判断会比较麻烦
当你把根目录的 go.mod 删除以后，就不再有一个“更高层级”的模块去挡住（shadow）你的工作区，Go 工具链就会把你当前所在目录和它上层都当成 “没有 go.mod，正在工作区里” 这种场景来处理，这时它才会真正去读 go.work，把 use 进来的子模块都当成本地模块来解析。
当你在一个目录里运行 go build、go mod tidy、go test 等命令时，Go 会一路往上找最近的 go.mod，把它当作“当前模块”的根。
只有当命令在一个“无 go.mod 且在 go.work 覆盖范围内”的目录下时，Go 才会把整个 go.work 当作一个更大范围来解析依赖。


如果保留根目录的 go.mod，go mod tidy 系列命令就永远“只看它”，此时要想让 tidy 理解本地模块，必须在你的 go.mod 里用 replace（或把根当聚合模块）。
go.work 依然对构建、测试、运行有大用处，它可以让你在本地联调多模块而不去拉远程，只是它不参与 go mod 子命令的依赖解析


go run .和go test运行路径即为程序中相对路径查找是用的路径，比如conf.go
因此为测试带来了困难
可以使用
dir := filepath.Dir(filename)
// 假设项目根目录在当前文件的上两级目录
_, filename, _, _ := runtime.Caller(0)
dir := filepath.Dir(filename)
// 假设项目根目录在当前文件的上两级目录
projectRoot := filepath.Join(dir, "..", "..")
_ = os.Chdir(projectRoot)
_ = godotenv.Load()


接收器
https://github.com/uber-go/guide/blob/master/style.md#receivers-and-interfaces
methods with value receiver can be called on pointers as well as values
methods with pointer receiver can only be called on pointers or addressable values

interface can be satisfied by a pointer even if the method has a value receiver
attention: interface requires implementations have the function, 
method with value receiver can be regarded as value's and pointer's method
method with pointer receiver can only be regarded as pointer's method

This rule arises because pointer methods can modify the receiver; invoking them on a value would cause the method to receive a copy of the value, so any modifications would be discarded. The language therefore disallows this mistake. There is a handy exception, though. When the value is addressable, the language takes care of the common case of invoking a pointer method on a value by inserting the address operator automatically. In our example, the variable b is addressable, so we can call its Write method with just b.Write. The compiler will rewrite that to (&b).Write for us.

接口有类型指针和数据指针
当给接口传指针，数据指针的值就是指针的值
当给接口传值，数据指针的值指向值的副本
如果传的值长度较短，可能直接存在数据指针中
值接收器和指针接收器的区别在于是否需要复制，是否可以修改原始值


// Field appears in JSON as key "myName".
Field int `json:"myName"`

// Field appears in JSON as key "myName" and
// the field is omitted from the object if its value is empty,
// as defined above.
Field int `json:"myName,omitempty"`

// Field appears in JSON as key "Field" (the default), but
// the field is skipped if empty.
// Note the leading comma.
Field int `json:",omitempty"`

// Field is ignored by this package.
Field int `json:"-"`

// Field appears in JSON as key "-".
Field int `json:"-,"`

go工作区只对go run有效，对go mod tidy无效

如果某些模块间有版本冲突（比如一个模块依赖 v1.0.0，另一个依赖 v2.0.0），而 Go 无法自动解析，你可以使用 replace 强制选择一个版本，直到冲突解决为止：

go工作区不能用于未发布的本地版本
https://github.com/golang/go/issues/50750
正确的使用方式是先发布版本，然后修改，这样的花依赖的模块不需要replace也可以感知到修改

提交呢? go work sync呢？

go template自动解引
rpc中消息中的结构体字段都以指针存在


git tag relativePath/tag

git tag -d tag

git push origin --delete tag

git push origin tag/--tags

GOPROXY=direct go mod tidy

GOSUMD=off

monorepo中tag需要加上子模块相对路径

团队成员无需手动执行 go get module@version：go run、go build 等命令会自动读取并下载 go.mod 里列好的版本。

可重复性：相同输入应有相同输出
如果你的项目自上次 go mod tidy 以来没有：

新增或删除任何 import

修改过任何版本声明（go.mod）

上游模块的版本或校验和没有发生变化（例如，仓库未重写历史）

那么多次运行 go mod tidy 应该会得到完全相同的 go.mod 和 go.sum（即便文件的时间戳会变，内容不会变）。

注意：如果上游包的维护者强制重写了 git tag 或删除了老版本，下一次 go mod tidy 可能会因为校验和不匹配而报错或自动选择替代版本，从而导致结果有差异。

go mod tidy vs. go get -u 的关系

go mod tidy：

扫描所有源码（包括测试），找出实际用到的依赖；

在 go.mod 中添加因 import 而“漏写”的 require；

删除不再使用的 require；

更新 go.sum，补充或移除对应校验和。

它并不会主动把依赖“升级”到新版本。

go get -u [module@version]：

会尝试将指定模块（或全部模块，若省略）升级到最新次要版本（minor/patch）或某个显式版本；

同时下载该版本及其依赖，并更新 go.mod、go.sum；

说白了：它既做“升级”也会做“下载+校验”。

常见用法：

bash
复制
编辑
# 只清理，不改版本
go mod tidy

# 升级所有依赖到最新小版本
go get -u all
# 然后再 tidy 一次，删掉没用到的依赖
go mod tidy

已下载依赖，go run . 是否可以直接使用？
是的。

go run . 首先会读取 go.mod 和 go.sum，

然后去本地模块缓存里查找对应版本的源码。

若缓存中已有，则直接编译运行；否则才会去远端下载。

这样，你在离线或网络不稳定的情况下，只要依赖以前已经下载过，就能正常执行。

Protobuf 要求每个 .proto 文件的 Go 包路径必须全局唯一。你的代码中存在两个不同路径的 common.pb.go：
github.com/spark4862/smartmall/app/frontend/hertz_gen/frontend/common
github.com/spark4862/smartmall/rpc_gen/kitex_gen/frontend/common
尽管它们来自同一个 common.proto，但生成的代码位于不同的 Go 包路径，从而引发注册冲突。
核心思想：确保所有服务（Hertz、Kitex 等）引用同一个 common.pb.go，避免重复生成。

当一个请求从客户端到达 Kitex Server 时，NewServerSuite() 会自动从请求中提取 OpenTelemetry 上下文（例如，通过请求头中的 traceparent 信息）。

在这个过程中，NewServerSuite() 会使用由 OpenTelemetryProvider 初始化的 Tracer 来生成新的 Span 或续写当前的追踪链。

OpenTelemetry 提供的上下文机制
分布式追踪（Distributed Tracing）是一种监控技术，主要用于跟踪跨多个服务、应用或组件的请求流

logrus

Promtail 是一个日志收集工具，主要用于与 Loki（由 Grafana 提供的日志聚合系统）配合使用。它从多个来源收集日志文件，并将这些日志推送到 Loki 中进行存储和分析。Promtail 是一个轻量级的代理，可以作为日志收集的“推送”工具，类似于其他日志收集工具，如 Filebeat 或 Fluentd，但专门为 Grafana 的 Loki 构建和优化。

go mod tidy 应该和go mod download && go mod verify配合使用

只读性：构建过程中生成的中间层是只读的
缓存机制：
Docker 会缓存未改变的层
如果 Dockerfile 的某行及其之前的行没有变化，就会重用缓存

优化中间层的实践 
将变化频率低的指令放在前面
使用 .dockerignore：避免不必要的文件被复制
多阶段构建：只保留必要的层到最终镜像
合并指令
清理临时文件：在同一 RUN 指令中创建和删除临时文件

容器的镜像通常由多个只读层组成，每个层都有特定的文件和更改。每个容器启动时，都会在镜像的最上面添加一个写层。这使得容器能够修改文件系统，同时保持镜像的只读性。
UnionFS 允许不同容器共享相同的只读层（例如，多个容器可以共享同一个镜像层），这大大减少了需要占用的磁盘空间
Docker 容器启动时，镜像层是已经构建好的
Docker 镜像由多个层（层叠文件系统）组成，这样每次修改镜像时，只需要创建新的层，而不必重新构建整个镜像

1. 会生成独立文件系统层的指令（实际存储变更）
这些指令会真正影响镜像的文件系统，生成可观察的存储层：

RUN - 执行命令并创建新层

COPY - 复制文件生成新层

ADD - 类似 COPY 但功能更多

WORKDIR - 虽然只改元数据但仍生成层

2. 只生成元数据层的指令（0B层）
这些指令会出现在镜像历史中，但不会增加实际存储空间：

ENV - 设置环境变量

LABEL - 添加元数据

MAINTAINER - 维护者信息(已弃用)

EXPOSE - 声明端口

VOLUME - 声明卷挂载点

USER - 设置用户

ARG - 构建时变量

ONBUILD - 触发器指令

使用 docker history 查看时会显示 SIZE=0B

3. 完全不生成层的指令
这些指令不会出现在最终镜像的层结构中：

FROM - 指定基础镜像

CMD - 容器启动命令

ENTRYPOINT - 入口点

HEALTHCHECK - 健康检查

SHELL - 指定 shell

STOPSIGNAL - 停止信号

特殊情况说明
多阶段构建：COPY --from 会生成新层，但 FROM ... AS 不会

shopt -s globstar 是一个在 Bash 中启用 globstar 扩展的命令。globstar 扩展允许你在文件路径中使用 ** 来递归匹配所有子目录中的文件和文件夹。

解释：
shopt：是 Bash 的一个内部命令，用于设置或取消启用某些 shell 选项。

-s：表示启用指定的选项（相反的，-u 是用来禁用选项的）。



echo app/cart/v0.2.0 app/checkout/v0.2.0 app/email/v0.2.0 app/frontend/v0.2.0 app/order/v0.2.0 app/payment/v0.2.0 app/product/v0.2.0 app/user/v0.2.0 common/v0.2.0 rpc_gen/v0.2.0 | xargs -n 1 git tag