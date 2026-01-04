// Package tables 提供数据库表格模型定义
// 本文件实现文章（posts）表格的模型配置，演示表格关联、富文本编辑和 AJAX 表单提交等功能
package tables

import (
	"github.com/purpose168/GoAdmin/context"
	"github.com/purpose168/GoAdmin/modules/db"
	"github.com/purpose168/GoAdmin/plugins/admin/modules/table"
	"github.com/purpose168/GoAdmin/template"
	"github.com/purpose168/GoAdmin/template/types"
	"github.com/purpose168/GoAdmin/template/types/form"
	editType "github.com/purpose168/GoAdmin/template/types/table"
)

// GetPostsTable 获取文章表格模型
// 该函数创建并返回一个配置完整的文章表格模型，用于管理后台的文章信息展示和编辑
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
//   - 配置信息展示字段（列表视图），包括表格关联和自定义显示
//   - 配置表单编辑字段（编辑视图），包括富文本编辑器
//   - 启用 AJAX 表单提交功能
//   - 演示表格关联（JOIN）的使用方法
//
// 核心特性:
//   - 表格关联：通过 FieldJoin 关联 authors 表获取作者信息
//   - 自定义显示：使用 FieldDisplay 创建链接和组合字段
//   - 富文本编辑：使用 form.RichText 支持富文本内容编辑
//   - 文件上传：通过 FieldEnableFileUpload 支持图片等文件上传
//   - AJAX 提交：通过 EnableAjax 实现异步表单提交
func GetPostsTable(ctx *context.Context) (postsTable table.Table) {

	// 创建默认表格模型
	// NewDefaultTable 创建一个使用默认配置的表格实例
	// DefaultConfigWithDriver 指定数据库驱动类型为 "sqlite"
	postsTable = table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("sqlite"))

	// 获取信息展示配置对象
	// GetInfo 返回表格的信息展示配置器，用于配置列表视图的字段
	// SetFilterFormLayout 设置筛选表单的布局为筛选布局
	info := postsTable.GetInfo().SetFilterFormLayout(form.LayoutFilter)

	// 添加 ID 字段
	// 参数说明:
	//   - "ID": 字段显示名称
	//   - "id": 数据库字段名
	//   - db.Int: 字段数据类型（整数）
	// FieldSortable: 设置该字段可排序
	info.AddField("编号", "id", db.Int).FieldSortable()

	// 添加 Title 字段
	// 参数说明:
	//   - "Title": 字段显示名称
	//   - "title": 数据库字段名
	//   - db.Varchar: 字段数据类型（可变长字符串）
	info.AddField("标题", "title", db.Varchar)

	// 添加 AuthorID 字段（自定义显示为链接）
	// 参数说明:
	//   - "AuthorID": 字段显示名称
	//   - "author_id": 数据库字段名
	//   - db.Int: 字段数据类型（整数）
	// FieldDisplay: 使用自定义函数显示字段内容
	//   这里将作者 ID 显示为可点击的链接，点击后在新标签页打开作者详情页
	info.AddField("作者ID", "author_id", db.Int).FieldDisplay(func(value types.FieldModel) interface{} {
		// 创建链接组件
		// template.Default() 获取默认模板组件
		// Link() 创建链接组件
		return template.Default().
			Link().
			// 设置链接 URL
			// /admin/info/authors/detail: 作者详情页路由
			// __goadmin_detail_pk: GoAdmin 框架的主键参数名
			// value.Value: 当前字段的值（作者 ID）
			SetURL("/admin/info/authors/detail?__goadmin_detail_pk=" + value.Value).
			// 设置链接显示内容
			// template.HTML 将字符串转换为 HTML 类型
			SetContent(template.HTML(value.Value)).
			// 在新标签页中打开链接
			OpenInNewTab().
			// 设置新标签页的标题
			// 格式: "作者详情(作者ID)"
			SetTabTitle(template.HTML("作者详情(" + value.Value + ")")).
			// 生成链接的 HTML 内容
			GetContent()
	})

	// 添加 AuthorName 字段（通过 JOIN 关联获取）
	// 参数说明:
	//   - "AuthorName": 字段显示名称
	//   - "name": 数据库字段名（虚拟字段）
	//   - db.Varchar: 字段数据类型
	// FieldDisplay: 使用自定义函数显示字段内容
	//   这里通过 JOIN 查询 authors 表获取作者的 first_name 和 last_name，然后组合成完整姓名
	info.AddField("作者姓名", "name", db.Varchar).FieldDisplay(func(value types.FieldModel) interface{} {
		// 从行数据中获取 first_name 字段值
		// authors_goadmin_join_first_name: JOIN 查询后的字段命名规则
		// 格式: {关联表名}_goadmin_join_{字段名}
		first, _ := value.Row["authors_goadmin_join_first_name"].(string)
		// 从行数据中获取 last_name 字段值
		last, _ := value.Row["authors_goadmin_join_last_name"].(string)
		// 返回组合后的完整姓名
		return first + " " + last
	})

	// 添加 AuthorFirstName 字段（JOIN 关联字段）
	// 参数说明:
	//   - "AuthorFirstName": 字段显示名称
	//   - "first_name": 数据库字段名
	//   - db.Varchar: 字段数据类型
	// FieldJoin: 配置表格关联
	//   - Field: 当前表的关联字段（author_id）
	//   - JoinField: 关联表的关联字段（id）
	//   - Table: 关联表名（authors）
	// FieldHide: 在列表视图中隐藏该字段，但可用于其他字段的显示
	info.AddField("作者名", "first_name", db.Varchar).FieldJoin(types.Join{
		Field:     "author_id",
		JoinField: "id",
		Table:     "authors",
	}).FieldHide()

	// 添加 AuthorLastName 字段（JOIN 关联字段）
	// 参数说明:
	//   - "AuthorLastName": 字段显示名称
	//   - "last_name": 数据库字段名
	//   - db.Varchar: 字段数据类型
	// FieldJoin: 配置表格关联
	//   - Field: 当前表的关联字段（author_id）
	//   - JoinField: 关联表的关联字段（id）
	//   - Table: 关联表名（authors）
	// FieldHide: 在列表视图中隐藏该字段，但可用于其他字段的显示
	info.AddField("作者姓", "last_name", db.Varchar).FieldJoin(types.Join{
		Field:     "author_id",
		JoinField: "id",
		Table:     "authors",
	}).FieldHide()

	// 添加 Description 字段
	// 参数说明:
	//   - "Description": 字段显示名称
	//   - "description": 数据库字段名
	//   - db.Varchar: 字段数据类型（可变长字符串）
	info.AddField("描述", "description", db.Varchar)

	// 添加 Content 字段（可编辑文本域）
	// 参数说明:
	//   - "Content": 字段显示名称
	//   - "content": 数据库字段名
	//   - db.Varchar: 字段数据类型（可变长字符串）
	// FieldEditAble: 设置字段在列表视图中可编辑
	//   editType.Textarea: 使用文本域编辑器
	info.AddField("内容", "content", db.Varchar).FieldEditAble(editType.Textarea)

	// 添加 Date 字段
	// 参数说明:
	//   - "Date": 字段显示名称
	//   - "date": 数据库字段名
	//   - db.Varchar: 字段数据类型（可变长字符串）
	info.AddField("日期", "date", db.Varchar)

	// 设置表格基本信息
	// SetTable: 指定数据库表名
	// SetTitle: 设置表格标题（显示在页面头部）
	// SetDescription: 设置表格描述
	info.SetTable("posts").SetTitle("文章").SetDescription("文章")

	// 获取表单配置对象
	// GetForm 返回表格的表单配置器，用于配置编辑/添加视图的字段
	formList := postsTable.GetForm()

	// 添加 ID 字段到表单
	// 参数说明:
	//   - "ID": 字段显示名称
	//   - "id": 数据库字段名
	//   - db.Int: 字段数据类型
	//   - form.Default: 表单字段类型（默认文本框）
	// FieldNotAllowEdit: 禁止编辑该字段（编辑模式下只读）
	// FieldNotAllowAdd: 禁止添加该字段（新增模式下不显示）
	formList.AddField("编号", "id", db.Int, form.Default).FieldNotAllowEdit().FieldNotAllowAdd()

	// 添加 Title 字段到表单
	// 参数说明:
	//   - "Title": 字段显示名称
	//   - "title": 数据库字段名
	//   - db.Varchar: 字段数据类型
	//   - form.Text: 表单字段类型（文本输入框）
	formList.AddField("标题", "title", db.Varchar, form.Text)

	// 添加 Description 字段到表单
	// 参数说明:
	//   - "Description": 字段显示名称
	//   - "description": 数据库字段名
	//   - db.Varchar: 字段数据类型
	//   - form.Text: 表单字段类型（文本输入框）
	formList.AddField("描述", "description", db.Varchar, form.Text)

	// 添加 Content 字段到表单（富文本编辑器）
	// 参数说明:
	//   - "Content": 字段显示名称
	//   - "content": 数据库字段名
	//   - db.Varchar: 字段数据类型
	//   - form.RichText: 表单字段类型（富文本编辑器）
	// FieldEnableFileUpload: 启用文件上传功能
	//   允许在富文本编辑器中插入图片、视频等文件
	formList.AddField("内容", "content", db.Varchar, form.RichText).FieldEnableFileUpload()

	// 添加 Date 字段到表单
	// 参数说明:
	//   - "Date": 字段显示名称
	//   - "date": 数据库字段名
	//   - db.Varchar: 字段数据类型
	//   - form.Datetime: 表单字段类型（日期时间选择器）
	formList.AddField("日期", "date", db.Varchar, form.Datetime)

	// 启用 AJAX 表单提交
	// EnableAjax 启用异步表单提交功能
	// 参数说明:
	//   - "Success": 提交成功时显示的消息
	//   - "Fail": 提交失败时显示的消息
	// 启用 AJAX 后，表单提交不会刷新页面，而是通过异步请求提交数据
	formList.EnableAjax("提交成功", "提交失败")

	// 设置表单基本信息
	// SetTable: 指定数据库表名
	// SetTitle: 设置表单标题
	// SetDescription: 设置表单描述
	formList.SetTable("posts").SetTitle("文章").SetDescription("文章")

	// 返回配置好的表格模型
	return
}
