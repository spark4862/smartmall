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