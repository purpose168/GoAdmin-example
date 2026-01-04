// Package tables 提供数据库表格模型定义
// 本文件实现外部数据源表格的模型配置，演示如何从非数据库来源获取数据
package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/parameter"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

// GetExternalTable 获取外部数据源表格模型
// 该函数创建并返回一个从外部数据源获取数据的表格模型
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
//   - 创建使用外部数据源的表格模型
//   - 配置信息展示字段（列表视图），使用自定义数据获取函数
//   - 配置表单编辑字段（编辑视图）
//   - 配置详情视图（详情页面），使用自定义数据获取函数
//   - 演示如何从 API、缓存或其他非数据库来源获取数据
//
// 使用场景:
//   - 数据来自第三方 API
//   - 数据存储在缓存中（如 Redis）
//   - 数据需要实时计算或处理
//   - 数据来自多个数据源的聚合
func GetExternalTable(ctx *context.Context) (externalTable table.Table) {

	// 创建默认表格模型
	// NewDefaultTable 创建一个使用默认配置的表格实例
	// DefaultConfig 使用默认配置，不指定数据库驱动
	// 因为这个表格的数据来自外部数据源，不需要数据库连接
	externalTable = table.NewDefaultTable(ctx, table.DefaultConfig())

	// 获取信息展示配置对象
	// GetInfo 返回表格的信息展示配置器，用于配置列表视图的字段
	// SetFilterFormLayout 设置筛选表单的布局为筛选布局
	info := externalTable.GetInfo().SetFilterFormLayout(form.LayoutFilter)

	// 添加 ID 字段
	// 参数说明:
	//   - "ID": 字段显示名称
	//   - "id": 数据字段名（对应外部数据中的键名）
	//   - db.Int: 字段数据类型（整数）
	// FieldSortable: 设置该字段可排序
	info.AddField("编号", "id", db.Int).FieldSortable()

	// 添加 Title 字段
	// 参数说明:
	//   - "Title": 字段显示名称
	//   - "title": 数据字段名（对应外部数据中的键名）
	//   - db.Varchar: 字段数据类型（可变长字符串）
	info.AddField("标题", "title", db.Varchar)

	// 设置表格基本信息和数据获取函数
	// SetTable: 指定表名标识符（用于路由和权限控制，不对应真实数据库表）
	// SetTitle: 设置表格标题（显示在页面头部）
	// SetDescription: 设置表格描述
	// SetGetDataFn: 设置自定义数据获取函数（核心功能）
	//   参数: param - 请求参数对象，包含分页、排序、筛选等信息
	//   返回值:
	//     - []map[string]interface{}: 数据列表，每个 map 代表一行数据
	//     - int: 总记录数（用于分页计算）
	//
	// 在实际应用中，这里可以调用 API、查询缓存或执行其他数据获取逻辑
	info.SetTable("external").
		SetTitle("外部数据").
		SetDescription("外部数据").
		SetGetDataFn(func(param parameter.Parameters) ([]map[string]interface{}, int) {
			// 返回模拟的外部数据
			// 在实际应用中，这里应该调用外部 API 或其他数据源
			// 例如: api.GetExternalData(param.Page, param.PageSize, param.SortField)
			return []map[string]interface{}{
				{
					"id":    10,
					"title": "这是一个标题",
				}, {
					"id":    11,
					"title": "这是一个标题2",
				}, {
					"id":    12,
					"title": "这是一个标题3",
				}, {
					"id":    13,
					"title": "这是一个标题4",
				},
			}, 10 // 总记录数，用于分页计算
		})

	// 获取表单配置对象
	// GetForm 返回表格的表单配置器，用于配置编辑/添加视图的字段
	formList := externalTable.GetForm()

	// 添加 ID 字段到表单
	// 参数说明:
	//   - "ID": 字段显示名称
	//   - "id": 数据字段名
	//   - db.Int: 字段数据类型
	//   - form.Default: 表单字段类型（默认文本框）
	// FieldNotAllowEdit: 禁止编辑该字段（编辑模式下只读）
	// FieldNotAllowAdd: 禁止添加该字段（新增模式下不显示）
	formList.AddField("编号", "id", db.Int, form.Default).FieldNotAllowEdit().FieldNotAllowAdd()

	// 添加 Title 字段到表单
	// 参数说明:
	//   - "Title": 字段显示名称
	//   - "title": 数据字段名
	//   - db.Varchar: 字段数据类型
	//   - form.Text: 表单字段类型（文本输入框）
	formList.AddField("标题", "title", db.Varchar, form.Text)

	// 设置表单基本信息
	// SetTable: 指定表名标识符
	// SetTitle: 设置表单标题
	// SetDescription: 设置表单描述
	formList.SetTable("external").SetTitle("外部数据").SetDescription("外部数据")

	// 获取详情视图配置对象
	// GetDetail 返回表格的详情视图配置器，用于配置详情页面的字段和内容
	// 详情视图用于展示单条记录的详细信息
	detail := externalTable.GetDetail()

	// 设置详情视图基本信息和数据获取函数
	// SetTable: 指定表名标识符
	// SetTitle: 设置详情页面标题
	// SetDescription: 设置详情页面描述
	// SetGetDataFn: 设置自定义数据获取函数
	//   参数: param - 请求参数对象，包含当前记录的 ID 等信息
	//   返回值:
	//     - []map[string]interface{}: 数据列表（详情视图通常只返回一条记录）
	//     - int: 记录数（详情视图通常为 1）
	//
	// 在实际应用中，这里应该根据 ID 从外部数据源获取单条记录的详细信息
	detail.SetTable("external").
		SetTitle("外部数据").
		SetDescription("外部数据").
		SetGetDataFn(func(param parameter.Parameters) ([]map[string]interface{}, int) {
			// 返回模拟的单条记录详情数据
			// 在实际应用中，这里应该根据 param.Id 从外部数据源获取单条记录
			// 例如: api.GetExternalDetail(param.Id)
			return []map[string]interface{}{
				{
					"id":    10,
					"title": "这是一个标题",
				},
			}, 1 // 记录数，详情视图通常为 1
		})

	// 返回配置好的表格模型
	return
}
