#!/bin/bash

echo "正在更新go.mod文件..."

# 检查Go是否安装
if ! command -v go &> /dev/null; then
    echo "错误: Go 未安装，请先安装Go"
    exit 1
fi

echo "Go版本: $(go version)"

# 方法1: 使用go mod tidy自动更新依赖
echo "\n=== 方法1: 使用go mod tidy自动更新依赖 ==="
go mod tidy

# 方法2: 更新特定依赖（可选）
echo "\n=== 方法2: 更新主要依赖 ==="
# 更新GoAdmin核心库
go get github.com/purpose168/GoAdmin@latest
# 更新主题库
go get github.com/purpose168/GoAdmin-themes@latest
# 更新gin框架
go get github.com/gin-gonic/gin@latest
# 更新gorm ORM
go get github.com/jinzhu/gorm@latest

# 再次运行tidy确保依赖关系正确
go mod tidy

echo "\n=== 更新完成 ==="
echo "go.mod文件已更新，新内容："
cat go.mod | head -30
