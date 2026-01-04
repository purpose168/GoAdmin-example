// models 包 - 数据模型层
// 本文件定义统计数据的模型和操作方法
// 使用 GORM 作为ORM框架，支持数据库持久化

// 创建日期: 2024
// 功能: 提供统计数据的结构定义和模板渲染方法

package models

import (
	"html/template"
	"strconv"
	"time"
)

// Statistics 统计数据模型
// 该结构体用于存储系统运行时的各种统计数据
// 包括CPU使用率、点赞数、销售额和新会员数等关键指标
// 通过GORM自动映射到数据库表
type Statistics struct {
	// ID 主键字段
	// 使用GORM标签指定为主键，映射到数据库的cpu列
	// 注意: 这里ID和CPU字段都映射到同一列，这是一种特殊的设计
	ID uint `gorm:"primary_key,column:cpu"`

	// CPU CPU使用率
	// 表示系统当前的CPU使用情况，通常以百分比形式存储
	// GORM标签: column=cpu 指定数据库列名为cpu
	CPU uint `gorm:"column:cpu"`

	// Likes 点赞数
	// 记录用户点赞的总数，用于衡量内容的受欢迎程度
	// GORM标签: column=likes 指定数据库列名为likes
	Likes uint `gorm:"column:likes"`

	// Sales 销售额
	// 记录产品销售的总金额，用于财务统计和分析
	// GORM标签: column=sales 指定数据库列名为sales
	Sales uint `gorm:"column:sales"`

	// NewMembers 新会员数
	// 记录新增会员的数量，用于用户增长分析
	// GORM标签: column=new_members 指定数据库列名为new_members
	NewMembers uint `gorm:"column:new_members"`

	// CreatedAt 创建时间
	// GORM自动管理的字段，记录数据创建的时间戳
	// 在插入数据时自动填充当前时间
	CreatedAt time.Time

	// UpdatedAt 更新时间
	// GORM自动管理的字段，记录数据最后更新的时间戳
	// 在更新数据时自动填充当前时间
	UpdatedAt time.Time
}

// FirstStatics 获取第一条统计数据
// 该函数从数据库中查询并返回第一条Statistics记录
//
// 返回值:
//   - *Statistics: 指向第一条统计数据的指针，如果数据库中没有记录则返回空结构体
//
// 功能说明:
//  1. 创建一个新的Statistics实例
//  2. 使用GORM的First方法查询第一条记录
//  3. 返回查询结果
//
// 使用示例:
//
//	import "github.com/purpose168/GoAdmin-example/models"
//
//	// 获取统计数据
//	stats := models.FirstStatics()
//	if stats.ID != 0 {
//	    fmt.Printf("CPU使用率: %d%%\n", stats.CPU)
//	}
//
// 注意事项:
//   - 如果数据库中没有记录，返回的Statistics结构体字段将为零值
//   - 该函数不会返回错误，如果查询失败会返回空结构体
//   - 在使用前应检查ID字段是否为0来判断是否有有效数据
//
// GORM说明:
//   - First方法会按主键升序查询第一条记录
//   - 如果没有找到记录，不会返回错误，而是返回空结果
func FirstStatics() *Statistics {
	// 创建一个新的Statistics实例
	// new()函数分配内存并返回指向该类型的指针
	// 所有字段将被初始化为零值
	s := new(Statistics)

	// 使用GORM的First方法查询第一条记录
	// orm是全局的GORM实例，在base.go中初始化
	// First方法会生成SQL: SELECT * FROM statistics ORDER BY id LIMIT 1
	orm.First(s)

	// 返回查询结果
	// 如果查询失败或没有记录，s将保持零值状态
	return s
}

// CPUTmpl 将CPU使用率转换为HTML模板格式
// 该方法将uint类型的CPU值转换为template.HTML类型，用于在HTML模板中安全渲染
//
// 返回值:
//   - template.HTML: 可以在HTML模板中安全使用的HTML字符串
//
// 功能说明:
//  1. 将uint类型的CPU值转换为int类型
//  2. 使用strconv.Itoa将int转换为字符串
//  3. 将字符串转换为template.HTML类型
//
// 使用示例:
//
//	import "github.com/purpose168/GoAdmin-example/models"
//
//	stats := models.FirstStatics()
//	cpuHtml := stats.CPUTmpl()
//	// cpuHtml可以直接在HTML模板中使用
//
// 注意事项:
//   - template.HTML类型会绕过Go的HTML转义，使用时需确保数据安全
//   - 该方法主要用于在GoAdmin模板中显示CPU使用率
//   - 返回的值可以直接嵌入到HTML中
//
// strconv说明:
//   - Itoa是Integer to ASCII的缩写
//   - 将整数转换为十进制字符串表示
func (s *Statistics) CPUTmpl() template.HTML {
	// 将uint类型的CPU值转换为int
	// strconv.Itoa需要int类型参数
	// template.HTML构造函数接受字符串并返回HTML安全类型
	return template.HTML(strconv.Itoa(int(s.CPU)))
}

// LikesTmpl 将点赞数转换为HTML模板格式
// 该方法将uint类型的Likes值转换为template.HTML类型，用于在HTML模板中安全渲染
//
// 返回值:
//   - template.HTML: 可以在HTML模板中安全使用的HTML字符串
//
// 功能说明:
//  1. 将uint类型的Likes值转换为int类型
//  2. 使用strconv.Itoa将int转换为字符串
//  3. 将字符串转换为template.HTML类型
//
// 使用示例:
//
//	import "github.com/purpose168/GoAdmin-example/models"
//
//	stats := models.FirstStatics()
//	likesHtml := stats.LikesTmpl()
//	// likesHtml可以直接在HTML模板中使用，如: <div>点赞数: {likesHtml}</div>
//
// 注意事项:
//   - template.HTML类型会绕过Go的HTML转义，使用时需确保数据安全
//   - 该方法主要用于在GoAdmin模板中显示点赞数
//   - 返回的值可以直接嵌入到HTML中
//
// 应用场景:
//   - 在仪表板页面显示点赞统计
//   - 在信息框组件中显示点赞数
//   - 可以与其他HTML元素组合使用
func (s *Statistics) LikesTmpl() template.HTML {
	// 将uint类型的Likes值转换为int
	// strconv.Itoa将整数转换为字符串
	// template.HTML确保字符串在HTML模板中不会被转义
	return template.HTML(strconv.Itoa(int(s.Likes)))
}

// SalesTmpl 将销售额转换为HTML模板格式
// 该方法将uint类型的Sales值转换为template.HTML类型，用于在HTML模板中安全渲染
//
// 返回值:
//   - template.HTML: 可以在HTML模板中安全使用的HTML字符串
//
// 功能说明:
//  1. 将uint类型的Sales值转换为int类型
//  2. 使用strconv.Itoa将int转换为字符串
//  3. 将字符串转换为template.HTML类型
//
// 使用示例:
//
//	import "github.com/purpose168/GoAdmin-example/models"
//
//	stats := models.FirstStatics()
//	salesHtml := stats.SalesTmpl()
//	// salesHtml可以直接在HTML模板中使用，如: <div>销售额: ¥{salesHtml}</div>
//
// 注意事项:
//   - template.HTML类型会绕过Go的HTML转义，使用时需确保数据安全
//   - 该方法主要用于在GoAdmin模板中显示销售额
//   - 返回的值可以直接嵌入到HTML中
//   - 在实际应用中，可能需要添加货币符号或格式化
//
// 应用场景:
//   - 在仪表板页面显示销售统计
//   - 在信息框组件中显示销售额
//   - 可以与其他HTML元素组合使用，如添加货币符号
func (s *Statistics) SalesTmpl() template.HTML {
	// 将uint类型的Sales值转换为int
	// strconv.Itoa将整数转换为字符串
	// template.HTML确保字符串在HTML模板中不会被转义
	return template.HTML(strconv.Itoa(int(s.Sales)))
}

// NewMembersTmpl 将新会员数转换为HTML模板格式
// 该方法将uint类型的NewMembers值转换为template.HTML类型，用于在HTML模板中安全渲染
//
// 返回值:
//   - template.HTML: 可以在HTML模板中安全使用的HTML字符串
//
// 功能说明:
//  1. 将uint类型的NewMembers值转换为int类型
//  2. 使用strconv.Itoa将int转换为字符串
//  3. 将字符串转换为template.HTML类型
//
// 使用示例:
//
//	import "github.com/purpose168/GoAdmin-example/models"
//
//	stats := models.FirstStatics()
//	membersHtml := stats.NewMembersTmpl()
//	// membersHtml可以直接在HTML模板中使用，如: <div>新会员: {membersHtml}人</div>
//
// 注意事项:
//   - template.HTML类型会绕过Go的HTML转义，使用时需确保数据安全
//   - 该方法主要用于在GoAdmin模板中显示新会员数
//   - 返回的值可以直接嵌入到HTML中
//   - 在实际应用中，可能需要添加单位（如"人"）
//
// 应用场景:
//   - 在仪表板页面显示会员增长统计
//   - 在信息框组件中显示新会员数
//   - 可以与其他HTML元素组合使用，如添加单位或图标
//
// 设计模式:
//   - 这是Go语言的接收者方法（receiver method）
//   - 允许通过Statistics实例直接调用
//   - 提供了数据转换的统一接口
func (s *Statistics) NewMembersTmpl() template.HTML {
	// 将uint类型的NewMembers值转换为int
	// strconv.Itoa将整数转换为字符串
	// template.HTML确保字符串在HTML模板中不会被转义
	return template.HTML(strconv.Itoa(int(s.NewMembers)))
}
