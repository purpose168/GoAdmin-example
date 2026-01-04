// Package tables 提供数据库表格模型定义和生成器映射
// 本文件定义了所有表格模型的生成器映射，用于路由到对应的表格处理函数
package tables

import "github.com/purpose168/GoAdmin/plugins/admin/modules/table"

// Generators 表格生成器映射表
//
// 该映射表将 URL 前缀映射到对应的表格生成函数
// GoAdmin 框架通过此映射表来路由请求到正确的表格处理器
//
// 映射规则:
//
//	键(key): 表格的 URL 前缀，用于构建访问 URL
//	值(value): 表格生成函数，返回 table.Table 对象
//
// URL 构建规则:
//
//	http://{{config.Domain}}:{{Port}}/{{config.Prefix}}/info/{{key}}
//
// 参数说明:
//   - config.Domain: 配置的域名
//   - Port: 服务端口号
//   - config.Prefix: 配置的 URL 前缀（默认为 "admin"）
//   - key: 映射表中的键名
//
// 示例:
//
//	"users"   => http://localhost:9033/admin/info/users
//	"posts"   => http://localhost:9033/admin/info/posts
//	"authors" => http://localhost:9033/admin/info/authors
//
// 使用说明:
//  1. 添加新表格时，需要在对应的文件中实现表格生成函数
//  2. 在此映射表中添加新的键值对
//  3. 框架会自动根据 URL 前缀路由到对应的生成函数
//
// 注意事项:
//   - 键名必须唯一，不能重复
//   - 键名建议使用小写字母和下划线
//   - 生成函数必须符合 table.Generator 类型签名
//   - 生成函数接收 context.Context 参数，返回 table.Table 对象
var Generators = map[string]table.Generator{
	// "posts" 前缀映射到 GetPostsTable 函数
	// 访问路径: /admin/info/posts
	// 功能: 文章管理表格，支持富文本编辑、表格关联等功能
	"posts": GetPostsTable,

	// "users" 前缀映射到 GetUserTable 函数
	// 访问路径: /admin/info/users
	// 功能: 用户管理表格
	"users": GetUserTable,

	// "authors" 前缀映射到 GetAuthorsTable 函数
	// 访问路径: /admin/info/authors
	// 功能: 作者管理表格，支持自定义按钮和组合字段显示
	"authors": GetAuthorsTable,

	// "profile" 前缀映射到 GetProfileTable 函数
	// 访问路径: /admin/info/profile
	// 功能: 用户档案表格，演示多种字段类型（轮播图、进度条、状态点等）
	"profile": GetProfileTable,
}
