// Package tables 提供数据库表格模型定义
// 本文件实现作者（authors）表格的模型配置，包括信息展示和表单编辑功能
package tables

import (
	"github.com/purpose168/GoAdmin/context"
	"github.com/purpose168/GoAdmin/modules/db"
	"github.com/purpose168/GoAdmin/plugins/admin/modules/table"
	"github.com/purpose168/GoAdmin/template/icon"
	"github.com/purpose168/GoAdmin/template/types"
	"github.com/purpose168/GoAdmin/template/types/action"
	"github.com/purpose168/GoAdmin/template/types/form"
)

// GetAuthorsTable 获取作者表格模型
// 该函数创建并返回一个配置完整的作者表格模型，用于管理后台的作者信息展示和编辑
//
// 参数:
//
//	ctx: 上下文对象，包含请求信息和配置
//
// 返回值:
//
//	table.Table: 配置好的表格模型对象
//
// 功能说明:
//   - 创建基于 SQLite 数据库的表格模型
//   - 配置信息展示字段（列表视图）
//   - 配置表单编辑字段（编辑视图）
//   - 添加自定义按钮操作（查看文章列表）
//   - 设置表格标题和描述
func GetAuthorsTable(ctx *context.Context) (authorsTable table.Table) {

	// 创建默认表格模型
	// NewDefaultTable 创建一个使用默认配置的表格实例
	// DefaultConfigWithDriver 指定数据库驱动类型为 "sqlite"
	// 支持的驱动类型: mysql, postgres, sqlite, mssql 等
	authorsTable = table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("sqlite"))

	// 如果需要使用自定义数据库连接，可以使用以下方式：
	// authorsTable = table.NewDefaultTable(ctx, table.DefaultConfigWithDriverAndConnection("mysql", "admin"))
	// 参数说明:
	//   - "mysql": 数据库驱动类型
	//   - "admin": 自定义连接名称（需要在配置文件中定义）

	// 获取信息展示配置对象
	// GetInfo 返回表格的信息展示配置器，用于配置列表视图的字段
	info := authorsTable.GetInfo().SetFilterFormLayout(form.LayoutFilter)

	// 添加 ID 字段
	// 参数说明:
	//   - "ID": 字段显示名称
	//   - "id": 数据库字段名
	//   - db.Int: 字段数据类型（整数）
	// FieldSortable: 设置该字段可排序
	info.AddField("编号", "id", db.Int).FieldSortable()

	// 添加 First Name 字段
	// FieldHide: 在列表视图中隐藏该字段，但仍在表单中显示
	info.AddField("名", "first_name", db.Varchar).FieldHide()

	// 添加 Last Name 字段
	// FieldHide: 在列表视图中隐藏该字段，但仍在表单中显示
	info.AddField("姓", "last_name", db.Varchar).FieldHide()

	// 添加 Name 字段（组合字段）
	// FieldDisplay: 使用自定义函数显示字段内容
	// 参数说明:
	//   - "Name": 字段显示名称
	//   - "name": 数据库字段名（虚拟字段，实际不存储）
	//   - db.Varchar: 字段数据类型
	//   - 回调函数: 接收 FieldModel 参数，返回显示值
	//     FieldModel.Row 包含当前行的所有字段数据
	//     这里将 first_name 和 last_name 组合成完整姓名
	info.AddField("姓名", "name", db.Varchar).FieldDisplay(func(value types.FieldModel) interface{} {
		// 从行数据中获取 first_name 字段值
		// 类型断言将 interface{} 转换为 string
		first, _ := value.Row["first_name"].(string)
		// 从行数据中获取 last_name 字段值
		last, _ := value.Row["last_name"].(string)
		// 返回组合后的完整姓名
		return first + " " + last
	})

	// 添加 Email 字段
	// 参数说明:
	//   - "Email": 字段显示名称
	//   - "email": 数据库字段名
	//   - db.Varchar: 字段数据类型（可变长字符串）
	info.AddField("邮箱", "email", db.Varchar)

	// 添加 Birthdate 字段
	// 参数说明:
	//   - "Birthdate": 字段显示名称
	//   - "birthdate": 数据库字段名
	//   - db.Date: 字段数据类型（日期）
	info.AddField("出生日期", "birthdate", db.Date)

	// 添加 Added 字段
	// 参数说明:
	//   - "Added": 字段显示名称
	//   - "added": 数据库字段名
	//   - db.Timestamp: 字段数据类型（时间戳）
	info.AddField("添加时间", "added", db.Timestamp)

	// 添加自定义按钮操作
	// AddButton 在每行数据中添加一个操作按钮
	// 参数说明:
	//   - ctx: 上下文对象
	//   - "Articles": 按钮显示文本
	//   - icon.Tv: 按钮图标（电视图标）
	//   - action.PopUpWithIframe: 弹出 iframe 窗口的动作
	//     - "/authors/list": 动作路由
	//     - "文章": 窗口标题
	//     - action.IframeData: iframe 数据配置
	//       - Src: iframe 加载的 URL 地址
	//     - "900px": 弹窗宽度
	//     - "560px": 弹窗高度
	info.AddButton(ctx, "文章", icon.Tv,
		action.PopUpWithIframe("/authors/list", "文章", action.IframeData{Src: "/admin/info/posts"}, "900px", "560px"))

	// 设置表格基本信息
	// SetTable: 指定数据库表名
	// SetTitle: 设置表格标题（显示在页面头部）
	// SetDescription: 设置表格描述
	info.SetTable("authors").SetTitle("作者").SetDescription("作者")

	// 获取表单配置对象
	// GetForm 返回表格的表单配置器，用于配置编辑/添加视图的字段
	formList := authorsTable.GetForm()

	// 添加 ID 字段到表单
	// 参数说明:
	//   - "ID": 字段显示名称
	//   - "id": 数据库字段名
	//   - db.Int: 字段数据类型
	//   - form.Default: 表单字段类型（默认文本框）
	// FieldNotAllowEdit: 禁止编辑该字段（编辑模式下只读）
	// FieldNotAllowAdd: 禁止添加该字段（新增模式下不显示）
	formList.AddField("编号", "id", db.Int, form.Default).FieldNotAllowEdit().FieldNotAllowAdd()

	// 添加 First Name 字段到表单
	// 参数说明:
	//   - "First Name": 字段显示名称
	//   - "first_name": 数据库字段名
	//   - db.Varchar: 字段数据类型
	//   - form.Text: 表单字段类型（文本输入框）
	formList.AddField("名", "first_name", db.Varchar, form.Text)

	// 添加 Last Name 字段到表单
	// 参数说明:
	//   - "Last Name": 字段显示名称
	//   - "last_name": 数据库字段名
	//   - db.Varchar: 字段数据类型
	//   - form.Text: 表单字段类型（文本输入框）
	formList.AddField("姓", "last_name", db.Varchar, form.Text)

	// 添加 Email 字段到表单
	// 参数说明:
	//   - "Email": 字段显示名称
	//   - "email": 数据库字段名
	//   - db.Varchar: 字段数据类型
	//   - form.Text: 表单字段类型（文本输入框）
	formList.AddField("邮箱", "email", db.Varchar, form.Text)

	// 添加 Birthdate 字段到表单
	// 参数说明:
	//   - "Birthdate": 字段显示名称
	//   - "birthdate": 数据库字段名
	//   - db.Date: 字段数据类型
	//   - form.Text: 表单字段类型（文本输入框，会自动使用日期选择器）
	formList.AddField("出生日期", "birthdate", db.Date, form.Text)

	// 添加 Added 字段到表单
	// 参数说明:
	//   - "Added": 字段显示名称
	//   - "added": 数据库字段名
	//   - db.Timestamp: 字段数据类型
	//   - form.Text: 表单字段类型（文本输入框，会自动使用日期时间选择器）
	formList.AddField("添加时间", "added", db.Timestamp, form.Text)

	// 设置表单基本信息
	// SetTable: 指定数据库表名
	// SetTitle: 设置表单标题
	// SetDescription: 设置表单描述
	formList.SetTable("authors").SetTitle("作者").SetDescription("作者")

	// 返回配置好的表格模型
	return
}
