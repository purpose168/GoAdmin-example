// models 包 - 数据模型层
// 本包负责数据库连接管理和ORM初始化
// 使用 GORM 作为ORM框架，支持多种数据库
// 作者: GoAdminGroup
// 创建日期: 2024
// 功能: 提供统一的数据库访问接口和ORM实例

package models

import (
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/jinzhu/gorm"
)

var (
	// orm GORM数据库实例
	// 这是一个全局的ORM实例，用于执行数据库操作
	// 支持链式调用，提供丰富的查询接口
	// 示例: orm.Find(&users), orm.Create(&user)
	orm *gorm.DB

	// err 错误变量
	// 用于存储数据库操作过程中产生的错误
	// Go语言的错误处理机制要求显式检查错误
	err error
)

// Init 初始化数据库ORM实例
// 该函数建立与数据库的连接并初始化GORM实例
//
// 参数:
//   - c: 数据库连接配置对象，包含数据库连接信息
//
// 功能说明:
//  1. 使用SQLite3驱动打开数据库连接
//  2. 从配置对象中获取名为"default"的数据库配置
//  3. 如果初始化失败，程序会panic并终止运行
//
// 使用示例:
//
//	import "github.com/GoAdminGroup/example/models"
//
//	// 在main函数中调用
//	models.Init(eng.SqliteConnection())
//
// 注意事项:
//   - 必须在使用任何数据库操作之前调用此函数
//   - 数据库连接信息必须在配置文件中正确配置
//   - SQLite数据库文件路径由配置决定
//   - 如果连接失败，程序会立即终止，不会继续执行
//
// 错误处理:
//   - 如果数据库连接失败，会触发panic
//   - 常见失败原因: 数据库文件路径错误、权限不足、磁盘空间不足
//   - 建议在生产环境中添加更详细的错误日志
func Init(c db.Connection) {
	// 使用GORM打开SQLite3数据库连接
	// "sqlite3": 指定使用SQLite数据库驱动
	// c.GetDB("default"): 从配置中获取默认数据库连接信息
	// 返回值: *gorm.DB (ORM实例), error (错误信息)
	orm, err = gorm.Open("sqlite3", c.GetDB("default"))

	// 检查数据库初始化是否成功
	// Go语言的标准错误处理模式
	// 如果err不为nil，表示发生了错误
	if err != nil {
		// panic会立即终止程序执行
		// 在生产环境中，建议使用日志记录并优雅退出
		// panic("initialize orm failed") 表示ORM初始化失败
		panic("initialize orm failed")
	}
}
