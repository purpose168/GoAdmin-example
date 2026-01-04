# 使用最新版本的 Go 语言官方镜像作为基础镜像
# golang:latest 镜像包含了 Go 编译器和运行时环境
FROM golang:latest

# 设置工作目录为 /go/src/app
# 后续的命令将在此目录下执行
WORKDIR /go/src/app

# 将宿主机当前目录下的所有文件复制到容器的 /go/src/app 目录中
# 第一个点(.)表示宿主机当前目录，第二个点(.)表示容器内当前工作目录
COPY . .