# GoAdmin 示例

这是一个展示如何运行 go-admin 的示例。

以下是运行代码的三种方式。

如果您是 Windows 用户，[go-sqlite-driver](https://github.com/mattn/go-sqlite3) 需要下载 gcc 才能正常工作。

## 使用 Go 模块

要使用 Go 模块，您应该先设置 GO111MODULE=on。

### 步骤 1

```shell
git clone https://github.com/purpose168/GoAdminExample.git
```

### 步骤 2

```shell
cd example
GO111MODULE=on go run .
```

访问: [http://localhost:9033/admin](http://localhost:9033/admin)

## 使用 Docker

### 步骤 1

```shell
git clone https://github.com/purpose168/GoAdminExample.git
```

### 步骤 2

```shell
cd example
docker build -t go-admin-example .
```

### 步骤 3

```shell
docker attach $(docker run -p 9033:9033 -it -d go-admin-example /bin/bash -c "cd /go/src/app && GOPROXY=http://goproxy.cn GO111MODULE=on go run .")
```

访问: [http://localhost:9033/admin](http://localhost:9033/admin)
