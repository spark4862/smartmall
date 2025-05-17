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