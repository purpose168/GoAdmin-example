# 定义 Go 命令工具
GOCMD = go
# 定义构建命令
GOBUILD = $(GOCMD) build
# 定义安装命令
GOINSTALL = $(GOCMD) install
# 定义测试命令
GOTEST = $(GOCMD) test
# 定义生成的二进制文件名称
BINARY_NAME = goadmin
# 定义 CLI 工具名称
CLI = adm

# 默认目标：启动服务
all: serve

# 启动服务：直接运行当前目录的 Go 程序
serve:
	$(GOCMD) run .

# 构建项目：生成 Linux 平台的二进制文件
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o ./build/$(BINARY_NAME) -v ./

# 生成代码：安装 go-admin CLI 工具并生成代码
generate:
	$(GOINSTALL) github.com/GoAdminGroup/go-admin/adm
	$(CLI) generate -c adm_config.ini

# 执行所有测试：包括黑盒测试和用户验收测试
test: black-box-test user-acceptance-test

# 黑盒测试：准备测试数据后运行黑盒测试
black-box-test: ready-for-data
	$(GOTEST) -v -test.run=TestExampleBlackBox
	make clean

# 用户验收测试：准备测试数据后运行用户验收测试
user-acceptance-test: ready-for-data
	$(GOTEST) -v -test.run=TestExampleUserAcceptance
	make clean

# 准备测试数据：复制数据库文件用于测试
ready-for-data:
	cp admin.db admin_test.db

# 清理测试数据：删除测试数据库文件
clean:
	rm admin_test.db

# 声明伪目标：这些目标不代表实际文件
.PHONY: all serve build generate test black-box-test user-acceptance-test ready-for-data clean