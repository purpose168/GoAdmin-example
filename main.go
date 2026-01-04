// GoAdmin 示例项目 - 主程序入口
// 本项目演示如何使用 GoAdmin 框架构建管理后台系统

// 创建日期: 2024
// 功能: 提供基于 Gin 框架的管理后台，包含仪表板、表单和表格页面

package main

import (
	"context"   // 上下文包，用于管理请求范围的操作
	"io/ioutil" // 输入/输出工具包，用于文件操作
	"log"       // 日志包，用于记录应用运行时信息
	"net/http"  // HTTP 包，用于处理 HTTP 请求和响应
	"os"        // 操作系统包，用于访问环境变量和文件系统
	"os/signal" // 信号处理包，用于捕获系统信号
	"time"      // 时间包，用于处理时间相关操作

	_ "github.com/purpose168/GoAdmin-themes/sword"              // Sword UI 主题
	_ "github.com/purpose168/GoAdmin/adapter/gin"               // Gin Web 框架适配器
	_ "github.com/purpose168/GoAdmin/modules/db/drivers/sqlite" // SQLite 数据库驱动

	"github.com/gin-gonic/gin"                       // Gin Web 框架，用于处理 HTTP 请求
	"github.com/purpose168/GoAdmin-example/models"   // 模型包，定义数据库表结构
	"github.com/purpose168/GoAdmin-example/pages"    // 页面包，定义管理后台页面
	"github.com/purpose168/GoAdmin-example/tables"   // 表格包，定义数据表格组件
	"github.com/purpose168/GoAdmin/engine"           // 引擎包，负责初始化和运行 GoAdmin
	"github.com/purpose168/GoAdmin/template"         // 模板包，定义页面模板和组件
	"github.com/purpose168/GoAdmin/template/chartjs" // Chart.js 图表组件
)

// main 主函数 - 程序入口点
// 负责启动服务器并初始化整个应用
func main() {
	startServer()
}

// startServer 初始化并启动 GoAdmin 管理后台服务器
// 该函数执行以下操作:
// 1. 配置 Gin 框架为发布模式
// 2. 创建 GoAdmin 引擎实例
// 3. 从 YAML 配置文件加载配置
// 4. 注册数据表生成器
// 5. 设置路由和页面处理器
// 6. 启动 HTTP 服务器
// 7. 实现优雅关闭机制
func startServer() {
	// 设置 Gin 为发布模式，禁用调试日志
	gin.SetMode(gin.ReleaseMode)
	// 丢弃 Gin 的默认输出，避免日志干扰
	gin.DefaultWriter = ioutil.Discard

	// 创建 Gin 路由器实例
	r := gin.Default()

	// 创建 GoAdmin 引擎实例，使用默认配置
	eng := engine.Default()

	// 添加 Chart.js 图表组件支持
	// Chart.js 是一个流行的 JavaScript 图表库，用于数据可视化
	template.AddComp(chartjs.NewChart())

	// 以下是被注释掉的数据库配置示例
	// 实际配置从 config.yml 文件中读取
	//cfg := config.Config{
	//	Databases: config.DatabaseList{
	//		"default": {
	//			Host:       "127.0.0.1",
	//			Port:       "3306",
	//			User:       "root",
	//			Pwd:        "root",
	//			Name:       "go-admin",
	//			MaxIdleCon: 50,
	//			MaxOpenCon: 150,
	//			Driver:     db.DriverMysql,
	//		},
	//	},
	//	UrlPrefix: "admin",
	//	IndexUrl:  "/",
	//	Debug:     true,
	//	Language:  language.CN,
	//}

	// 从 YAML 配置文件加载配置
	// AddConfigFromYAML: 从指定路径读取配置文件
	// AddGenerators: 注册数据表生成器，用于自动生成管理界面
	// AddGenerator: 添加外部表生成器
	// Use: 将 GoAdmin 引擎集成到 Gin 路由器中
	if err := eng.AddConfigFromYAML("./config.yml").
		AddGenerators(tables.Generators).
		AddGenerator("external", tables.GetExternalTable).
		Use(r); err != nil {
		panic(err)
	}

	// 初始化数据库模型
	// 使用 SQLite 数据库连接
	// models.Init: 初始化ORM实例，建立数据库连接
	// eng.SqliteConnection(): 获取SQLite数据库连接配置
	// 注意: 必须在使用任何数据库操作之前调用此函数
	models.Init(eng.SqliteConnection())

	// 设置静态文件路由
	// 将 /uploads 路径映射到本地 ./uploads 目录
	// 用于处理用户上传的文件访问
	r.Static("/uploads", "./uploads")

	// 注册 HTML 页面路由
	// DashboardPage: 仪表板页面，显示系统概览信息
	eng.HTML("GET", "/admin", pages.DashboardPage)
	// GetFormContent: 表单页面，展示各种表单字段类型
	// 包含基础输入、日期时间、文件上传、富文本、选择控件等多种表单组件
	// 使用标签页分组，分为input、select、multi三个标签页
	eng.HTML("GET", "/admin/form", pages.GetFormContent)
	// GetTableContent: 表格页面，用于数据展示和管理
	eng.HTML("GET", "/admin/table", pages.GetTableContent)
	// 自定义模板文件路由
	// 使用 Go 模板引擎渲染 hello.tmpl 文件
	eng.HTMLFile("GET", "/admin/hello", "./html/hello.tmpl", map[string]interface{}{
		"msg": "你好世界",
	})

	// 创建 HTTP 服务器配置
	// Addr: 监听地址和端口，9033 是默认端口
	// Handler: 使用 Gin 路由器作为请求处理器
	srv := &http.Server{
		Addr:    ":9033",
		Handler: r,
	}

	// 在新的 goroutine 中启动服务器
	// 使用 goroutine 可以让服务器在后台运行，不阻塞主线程
	// 这是 Go 语言并发编程的核心特性
	go func() {
		// ListenAndServe 启动 HTTP 服务器
		// 如果端口被占用或其他错误，会返回错误
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("监听: %s\n", err)
		}
	}()

	// 实现优雅关闭机制
	// 创建一个信号通道，用于接收操作系统信号
	quit := make(chan os.Signal, 1)
	// 监听中断信号（Ctrl+C）
	// 当用户按下 Ctrl+C 时，会向通道发送 os.Interrupt 信号
	signal.Notify(quit, os.Interrupt)
	// 阻塞等待退出信号
	<-quit

	// 收到退出信号后，执行优雅关闭
	// 创建一个带有超时的上下文
	// 5 秒超时：如果服务器在 5 秒内没有关闭，将强制关闭
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer 确保在函数返回前调用取消函数，释放资源
	defer cancel()
	// Shutdown 方法优雅地关闭服务器
	// 它会：
	// 1. 停止接受新的连接
	// 2. 等待所有活跃的请求完成（最多等待超时时间）
	// 3. 关闭所有连接
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("服务器关闭:", err)
	}
	log.Println("服务器退出")
}
