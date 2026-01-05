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

# ------------------------
# 服务管理
# ------------------------

# 启动服务：直接运行当前目录的 Go 程序
serve:
	@echo "=== 启动服务 ==="
	$(GOCMD) run .

# ------------------------
# 构建管理
# ------------------------

# 构建项目：生成 Linux 平台的二进制文件
build:
	@echo "=== 构建项目 ==="
	@mkdir -p build
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o ./build/$(BINARY_NAME) -v ./

# ------------------------
# 依赖管理
# ------------------------

# 清理模块缓存：清理 Go 模块缓存，解决依赖冲突
mod-clean:
	@echo "=== 清理模块缓存 ==="
	$(GOCMD) clean -modcache

# 修复依赖关系：更新 go.mod 和 go.sum 文件，确保依赖关系正确
mod-tidy:
	@echo "=== 修复依赖关系 ==="
	$(GOCMD) mod tidy

# 生成 vendor 目录：将依赖复制到本地 vendor 目录
mod-vendor:
	@echo "=== 生成 vendor 目录 ==="
	$(GOCMD) mod vendor

# 检查依赖关系：检查依赖关系是否正确
mod-verify:
	@echo "=== 检查依赖关系 ==="
	$(GOCMD) mod verify

# 调试依赖问题：打印依赖图，用于调试依赖问题
mod-graph:
	@echo "=== 打印依赖图 ==="
	$(GOCMD) mod graph

# 更新依赖：更新所有依赖到最新版本
mod-update:
	@echo "=== 更新依赖 ==="
	$(GOCMD) get -u ./...
	$(GOCMD) mod tidy

# ------------------------
# 测试管理
# ------------------------

# 执行所有测试：包括黑盒测试和用户验收测试
test: black-box-test user-acceptance-test

# 黑盒测试：准备测试数据后运行黑盒测试
black-box-test: ready-for-data
	@echo "=== 运行黑盒测试 ==="
	$(GOTEST) -v -test.run=TestExampleBlackBox
	@make clean

# 用户验收测试：准备测试数据后运行用户验收测试
user-acceptance-test: ready-for-data
	@echo "=== 运行用户验收测试 ==="
	$(GOTEST) -v -test.run=TestExampleUserAcceptance
	@make clean

# 准备测试数据：复制数据库文件用于测试
ready-for-data:
	@echo "=== 准备测试数据 ==="
	@cp admin.db admin_test.db

# 清理测试数据：删除测试数据库文件
clean:
	@echo "=== 清理测试数据 ==="
	@rm -f admin_test.db

# ------------------------
# 代码生成
# ------------------------

# 生成代码：安装 go-admin CLI 工具并生成代码
generate:
	@echo "=== 生成代码 ==="
	$(GOINSTALL) github.com/purpose168/GoAdmin-adm
	$(CLI) generate -c adm_config.ini

# ------------------------
# 开发辅助
# ------------------------

# 格式化代码：使用 gofmt 格式化代码
fmt:
	@echo "=== 格式化代码 ==="
	$(GOCMD) fmt ./...

# 检查代码：使用 go vet 检查代码
vet:
	@echo "=== 检查代码 ==="
	$(GOCMD) vet ./...

# 静态分析：使用 staticcheck 进行静态分析
lint:
	@echo "=== 静态分析 ==="
	@which staticcheck > /dev/null 2>&1 || $(GOCMD) install honnef.co/go/tools/cmd/staticcheck@latest
	staticcheck ./...

# ------------------------
# 声明伪目标：这些目标不代表实际文件
# ------------------------

.PHONY: all serve build \
	mod-clean mod-tidy mod-vendor mod-verify mod-graph mod-update \
	test black-box-test user-acceptance-test ready-for-data clean \
	generate fmt vet lint