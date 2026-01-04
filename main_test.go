// GoAdmin 示例项目 - 测试文件
// 本文件包含 GoAdmin 框架的测试用例
// 作者: GoAdminGroup
// 创建日期: 2024
// 功能: 提供黑盒测试和用户验收测试

package main

import (
	"log"
	"testing"

	"github.com/GoAdminGroup/example/tables"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/tests"
	"github.com/GoAdminGroup/go-admin/tests/common"
	"github.com/GoAdminGroup/go-admin/tests/frameworks/gin"
	"github.com/GoAdminGroup/go-admin/tests/web"
	"github.com/gavv/httpexpect"
)

// TestExampleBlackBox 黑盒测试
// 测试 GoAdmin 框架的核心功能，不依赖内部实现细节
// 这种测试方式更接近真实用户的使用场景
func TestExampleBlackBox(t *testing.T) {
	// 使用 Gin 框架的测试处理器
	// 配置 SQLite 数据库，使用临时文件 admin_test.db
	tests.BlackBoxTestSuit(t, gin.NewHandler, config.DatabaseList{
		"default": config.Database{
			File:   "./admin_test.db",
			Driver: "sqlite",
		},
	}, tables.Generators, func(cfg config.DatabaseList) {
		// 框架的数据清理器
		// 在每个测试用例执行前清理数据库，确保测试独立性
		tests.Cleaner(cfg)
		// 清理您自己的数据：
		// ...
	}, func(e *httpexpect.Expect) {
		// 框架的测试用例
		// 执行框架提供的通用测试，验证基本功能
		common.Test(e)
		// 编写您自己的 API 测试，例如：
		// 更多用法: https://github.com/gavv/httpexpect
		// e.POST("/signin").Expect().Status(http.StatusOK)
	})
}

// TestExampleUserAcceptance 用户验收测试
// 使用真实的浏览器（通过 chromedriver）模拟用户操作
// 这种测试方式可以验证用户界面的实际行为
func TestExampleUserAcceptance(t *testing.T) {
	web.UserAcceptanceTestSuit(t, func(t *testing.T, page *web.Page) {
		// 基于 chromedriver 编写测试用例，例如：
		// 更多用法: https://github.com/sclevine/agouti
		// 导航到管理后台页面
		page.NavigateTo("http://127.0.0.1:9033/admin")
		// 验证页面是否包含特定文本
		//page.Contain("username")
		// 模拟点击操作
		//page.Click("")
	}, func(quit chan struct{}) {
		// 启动服务器：
		// ....
		// 在新的 goroutine 中启动服务器，避免阻塞测试
		go startServer()
		// 等待退出信号
		<-quit
		log.Print("test quit")
	}, true) // 如果 local 参数为 true，将不会使用无头模式，测试完成后窗口不会关闭
}
