FROM golang:1.24.3 AS builder

ARG SVC

WORKDIR /user/src/gomall

ENV GOPROXY=https://goproxy.cn,direct


COPY rpc_gen common ./
# 将本地的 go.mod 和 go.sum 文件复制到容器的工作目录

RUN cd common && go mod download && go mod verify

COPY app/${SVC}/go.mod app/${SVC}/go.sum ./app/${SVC}/
RUN cd app/${SVC}/ && go mod download && go mod verify


COPY app/${SVC} app/${SVC}
RUN cd app/${SVC}/ && CGO_ENABLED=0 go build -v -o /opt/gomall/${SVC}/server
# 用于完全禁用 CGO（即禁止 Go 调用 C 代码）Go 会纯用 Go 实现的库 会编译出不依赖任何外部动态库的二进制文件
# ./... 表示编译当前目录及其子目录下的所有包

# Docker 会为每个指令生成一个中间层，如果某层没有变化，就会直接使用缓存而不是重新执行。依赖（由 go.mod/go.sum 定义）通常比源代码变化频率低得多。这样分开后：
# 当只有源代码变更时，不需要重新下载依赖
# 只有当 go.mod 或 go.sum 变更时，才会触发 go mod download

FROM busybox

ARG SVC

COPY --from=builder /opt/gomall/${SVC}/server /opt/gomall/${SVC}/server
COPY app/${SVC}/conf /opt/gomall/${SVC}/conf

WORKDIR /opt/gomall/${SVC}

CMD ["./server"]