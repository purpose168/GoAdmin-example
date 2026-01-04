// Package tables 提供数据库表格模型定义
// 本文件实现用户（users）表格的模型配置，演示 GoAdmin 框架的丰富功能和高级特性
package tables

import (
	"fmt"

	"github.com/purpose168/GoAdmin/context"
	"github.com/purpose168/GoAdmin/modules/db"
	form2 "github.com/purpose168/GoAdmin/plugins/admin/modules/form"
	"github.com/purpose168/GoAdmin/plugins/admin/modules/table"
	"github.com/purpose168/GoAdmin/template"
	"github.com/purpose168/GoAdmin/template/icon"
	"github.com/purpose168/GoAdmin/template/types"
	"github.com/purpose168/GoAdmin/template/types/action"
	"github.com/purpose168/GoAdmin/template/types/form"
	selection "github.com/purpose168/GoAdmin/template/types/form/select"
	editType "github.com/purpose168/GoAdmin/template/types/table"
)

// GetUserTable 获取用户表格模型
// 该函数创建并返回一个配置完整的用户表格模型，用于管理后台的用户信息展示和编辑
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
//   - 创建基于 SQLite 数据库的表格模型，使用自定义配置
//   - 配置信息展示字段（列表视图），包括多种字段类型和显示方式
//   - 配置表单编辑字段（编辑视图），包括级联选择和表单分组
//   - 演示多种操作按钮（跳转、AJAX、弹窗、Iframe）
//   - 演示表单后置钩子和自定义过滤函数
//
// 核心特性:
//   - 自定义表格配置：通过 table.Config 配置表格的各种属性
//   - 可编辑字段：通过 FieldEditAble 支持列表视图直接编辑
//   - 开关按钮：通过 editType.Switch 实现开关切换
//   - 图片显示：通过 Image 组件显示头像
//   - 级联选择：通过 FieldOnChooseAjax 实现国家-城市级联选择
//   - 表单分组：通过 TabGroups 实现表单标签页分组
//   - 多种操作：Jump、Ajax、PopUp、PopUpWithIframe 等多种操作类型
//   - 表单钩子：通过 SetPostHook 实现表单提交后的自定义处理
func GetUserTable(ctx *context.Context) (userTable table.Table) {

	// 创建自定义配置的表格模型
	// table.Config 允许自定义表格的各种配置选项
	userTable = table.NewDefaultTable(ctx, table.Config{
		// Driver: 指定数据库驱动类型
		// db.DriverSqlite: 使用 SQLite 数据库
		// 其他选项: db.DriverMysql, db.DriverPostgresql, db.DriverMssql
		Driver: db.DriverSqlite,

		// CanAdd: 是否允许添加新记录
		// true: 显示"添加"按钮，允许用户添加新记录
		// false: 隐藏"添加"按钮，禁止添加新记录
		CanAdd: true,

		// Editable: 是否允许编辑记录
		// true: 显示"编辑"按钮，允许用户编辑记录
		// false: 隐藏"编辑"按钮，禁止编辑记录
		Editable: true,

		// Deletable: 是否允许删除记录
		// true: 显示"删除"按钮，允许用户删除记录
		// false: 隐藏"删除"按钮，禁止删除记录
		Deletable: true,

		// Exportable: 是否允许导出数据
		// true: 显示"导出"按钮，允许用户导出数据
		// false: 隐藏"导出"按钮，禁止导出数据
		Exportable: true,

		// Connection: 指定数据库连接名称
		// table.DefaultConnectionName: 使用默认数据库连接
		// 可以在配置文件中定义多个数据库连接，然后在此处指定使用哪个连接
		Connection: table.DefaultConnectionName,

		// PrimaryKey: 配置主键信息
		// Type: 主键数据类型（db.Int 表示整数类型）
		// Name: 主键字段名（table.DefaultPrimaryKeyName 默认为 "id"）
		PrimaryKey: table.PrimaryKey{
			Type: db.Int,
			Name: table.DefaultPrimaryKeyName,
		},
	})

	// 获取信息展示配置对象
	// GetInfo 返回表格的信息展示配置器，用于配置列表视图的字段
	// SetFilterFormLayout 设置筛选表单的布局为筛选布局
	info := userTable.GetInfo().SetFilterFormLayout(form.LayoutFilter)

	// 添加 ID 字段（支持排序）
	// 参数说明:
	//   - "ID": 字段显示名称
	//   - "id": 数据库字段名
	//   - db.Int: 字段数据类型（整数）
	// FieldSortable: 设置该字段可排序（点击表头可按此字段排序）
	info.AddField("编号", "id", db.Int).FieldSortable()

	// 添加 Name 字段（可编辑，支持模糊筛选）
	// 参数说明:
	//   - "Name": 字段显示名称
	//   - "name": 数据库字段名
	//   - db.Varchar: 字段数据类型（可变长字符串）
	// FieldEditAble: 设置字段在列表视图中可编辑
	//   editType.Text: 使用文本框编辑器
	// FieldFilterable: 设置该字段可筛选
	//   types.FilterType{Operator: types.FilterOperatorLike}: 使用模糊匹配筛选（LIKE 操作符）
	info.AddField("姓名", "name", db.Varchar).FieldEditAble(editType.Text).
		FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})

	// 添加 Gender 字段（可编辑开关，支持筛选）
	// 参数说明:
	//   - "Gender": 字段显示名称
	//   - "gender": 数据库字段名
	//   - db.Tinyint: 字段数据类型（微整数，用于存储性别：0=男，1=女）
	// FieldDisplay: 使用自定义函数显示字段内容
	//   根据字段值显示对应的性别文本
	// FieldEditAble: 设置字段在列表视图中可编辑
	//   editType.Switch: 使用开关按钮编辑器
	// FieldEditOptions: 设置开关的选项
	//   Value: 选项值
	//   Text: 选项显示文本（使用 emoji 图标）
	// FieldFilterable: 设置该字段可筛选
	//   FormType: form.SelectSingle: 使用单选下拉框筛选
	// FieldFilterOptions: 设置筛选选项
	info.AddField("性别", "gender", db.Tinyint).FieldDisplay(func(model types.FieldModel) interface{} {
		// 根据字段值返回对应的性别文本
		if model.Value == "0" {
			return "男"
		}
		if model.Value == "1" {
			return "女"
		}
		return "未知"
	}).FieldEditAble(editType.Switch).FieldEditOptions(types.FieldOptions{
		{Value: "0", Text: "👨"},
		{Value: "1", Text: "👩"},
	}).FieldFilterable(types.FilterType{FormType: form.SelectSingle}).FieldFilterOptions(types.FieldOptions{
		{Value: "0", Text: "男"},
		{Value: "1", Text: "女"},
	})

	// 添加自定义列（不对应数据库字段）
	// AddColumn 添加一个虚拟列，不对应数据库字段，用于显示自定义内容
	// 参数说明:
	//   - "personality": 列标识符
	//   - 回调函数: 返回要显示的内容
	//     这里返回固定的文本 "handsome"
	// 使用场景: 显示计算字段、组合字段或自定义内容
	info.AddColumn("个性", func(value types.FieldModel) interface{} {
		return "帅气"
	})

	// 添加列按钮（每行的操作按钮）
	// AddColumnButtons 在每行数据中添加一个操作按钮
	// 参数说明:
	//   - ctx: 上下文对象
	//   - "see more": 按钮显示文本
	//   - types.GetColumnButton: 创建列按钮
	//     - "more": 按钮标识符
	//     - icon.Info: 按钮图标（信息图标）
	//     - action.PopUp: 弹窗动作
	//       - "/see/more/example": 动作路由
	//       - "Detail": 弹窗标题
	//       - 回调函数: 处理弹窗请求的逻辑
	//         返回: success-操作是否成功, msg-返回消息, data-附加数据（HTML 内容）
	info.AddColumnButtons(ctx, "查看更多", types.GetColumnButton("more", icon.Info,
		action.PopUp("/see/more/example", "详情", func(ctx *context.Context) (success bool, msg string, data interface{}) {
			// 返回弹窗显示的 HTML 内容
			return true, "ok", "<h1>详情</h1><p>balabala</p><p>此功能将在 v1.2.7 版本发布</p>"
		})))

	// 添加 Phone 字段（支持筛选）
	// 参数说明:
	//   - "Phone": 字段显示名称
	//   - "phone": 数据库字段名
	//   - db.Varchar: 字段数据类型（可变长字符串）
	// FieldFilterable: 设置该字段可筛选（默认使用精确匹配）
	info.AddField("电话", "phone", db.Varchar).FieldFilterable()

	// 添加 City 字段（支持筛选）
	// 参数说明:
	//   - "City": 字段显示名称
	//   - "city": 数据库字段名
	//   - db.Varchar: 字段数据类型（可变长字符串）
	// FieldFilterable: 设置该字段可筛选（默认使用精确匹配）
	info.AddField("城市", "city", db.Varchar).FieldFilterable()

	// 添加 Avatar 字段（显示图片）
	// 参数说明:
	//   - "Avatar": 字段显示名称
	//   - "avatar": 数据库字段名
	//   - db.Varchar: 字段数据类型（可变长字符串，存储图片 URL）
	// FieldDisplay: 使用自定义函数显示字段内容
	//   template.Default().Image(): 创建图片组件
	//     SetSrc: 设置图片源 URL
	//     SetHeight: 设置图片高度
	//     SetWidth: 设置图片宽度
	//     WithModal: 点击图片时显示模态框（大图预览）
	//     GetContent: 生成图片的 HTML 内容
	info.AddField("头像", "avatar", db.Varchar).FieldDisplay(func(value types.FieldModel) interface{} {
		return template.Default().Image().
			SetSrc(`//quick.go-admin.cn/demo/assets/dist/img/gopher_avatar.png`).
			SetHeight("120").SetWidth("120").WithModal().GetContent()
	})

	// 添加 CreatedAt 字段（时间戳，支持日期范围筛选）
	// 参数说明:
	//   - "CreatedAt": 字段显示名称
	//   - "created_at": 数据库字段名
	//   - db.Timestamp: 字段数据类型（时间戳）
	// FieldFilterable: 设置该字段可筛选
	//   FormType: form.DatetimeRange: 使用日期时间范围选择器筛选
	info.AddField("创建时间", "created_at", db.Timestamp).
		FieldFilterable(types.FilterType{FormType: form.DatetimeRange})

	// 添加 UpdatedAt 字段（可编辑时间戳）
	// 参数说明:
	//   - "UpdatedAt": 字段显示名称
	//   - "updated_at": 数据库字段名
	//   - db.Timestamp: 字段数据类型（时间戳）
	// FieldEditAble: 设置字段在列表视图中可编辑
	//   editType.Datetime: 使用日期时间选择器编辑器
	info.AddField("更新时间", "updated_at", db.Timestamp).FieldEditAble(editType.Datetime)

	// 添加行操作按钮（每行的操作按钮）
	// AddActionButton 在每行数据的操作列中添加一个按钮

	// 添加 Google 跳转按钮
	// action.Jump: 跳转到指定 URL
	//   - "https://google.com": 目标 URL
	info.AddActionButton(ctx, "谷歌", action.Jump("https://google.com"))

	// 添加审核 AJAX 按钮
	// action.Ajax: 发送 AJAX 请求
	//   - "/admin/audit": 请求路由
	//   - 回调函数: 处理 AJAX 请求的逻辑
	//     返回: success-操作是否成功, msg-返回消息, data-附加数据
	info.AddActionButton(ctx, "审核", action.Ajax("/admin/audit",
		func(ctx *context.Context) (success bool, msg string, data interface{}) {
			// 执行审核逻辑
			return true, "成功", ""
		}))

	// 添加预览弹窗按钮
	// action.PopUp: 弹出模态框
	//   - "/admin/preview": 请求路由
	//   - "Preview": 弹窗标题
	//   - 回调函数: 返回弹窗显示的 HTML 内容
	info.AddActionButton(ctx, "预览", action.PopUp("/admin/preview", "预览",
		func(ctx *context.Context) (success bool, msg string, data interface{}) {
			return true, "", "<h2>你好世界</h2>"
		}))

	// 添加全局操作按钮（表格顶部的操作按钮）
	// AddButton 在表格顶部添加一个操作按钮

	// 添加 Google 跳转按钮（全局）
	// 参数说明:
	//   - ctx: 上下文对象
	//   - "google": 按钮显示文本
	//   - icon.Google: 按钮图标（Google 图标）
	//   - action.Jump: 跳转动作
	//     - "https://google.com": 目标 URL
	info.AddButton(ctx, "谷歌", icon.Google, action.Jump("https://google.com"))

	// 添加弹窗示例按钮（全局）
	// 参数说明:
	//   - ctx: 上下文对象
	//   - "popup": 按钮显示文本
	//   - icon.Terminal: 按钮图标（终端图标）
	//   - action.PopUp: 弹窗动作
	//     - "/admin/popup": 请求路由
	//     - "Popup Example": 弹窗标题
	//     - 回调函数: 返回弹窗显示的 HTML 内容
	info.AddButton(ctx, "弹窗", icon.Terminal, action.PopUp("/admin/popup", "弹窗示例",
		func(ctx *context.Context) (success bool, msg string, data interface{}) {
			return true, "", "<h2>你好世界</h2>"
		}))

	// 添加 Iframe 弹窗按钮（全局）
	// 参数说明:
	//   - ctx: 上下文对象
	//   - "iframe": 按钮显示文本
	//   - icon.Tv: 按钮图标（电视图标）
	//   - action.PopUpWithIframe: 弹出 iframe 窗口的动作
	//     - "/admin/iframe": 动作路由
	//     - "Iframe Example": 窗口标题
	//     - action.IframeData: iframe 数据配置
	//       - Src: iframe 加载的 URL 地址
	//     - "900px": 弹窗宽度
	//     - "480px": 弹窗高度
	info.AddButton(ctx, "iframe", icon.Tv, action.PopUpWithIframe("/admin/iframe", "Iframe 示例",
		action.IframeData{Src: "/admin/info/profile/new"}, "900px", "480px"))

	// 添加 AJAX 按钮示例（全局）
	// 参数说明:
	//   - ctx: 上下文对象
	//   - "ajax": 按钮显示文本
	//   - icon.Android: 按钮图标（Android 图标）
	//   - action.Ajax: AJAX 动作
	//     - "/admin/ajax": 请求路由
	//     - 回调函数: 处理 AJAX 请求的逻辑
	info.AddButton(ctx, "ajax", icon.Android, action.Ajax("/admin/ajax",
		func(ctx *context.Context) (success bool, msg string, data interface{}) {
			return true, "成功", ""
		}))

	// 添加批量选择框（表格顶部的批量操作选择框）
	// AddSelectBox 添加一个批量选择框，用于批量操作
	// 参数说明:
	//   - ctx: 上下文对象
	//   - "gender": 选择框标识符
	//   - types.FieldOptions: 选择框选项
	//   - action.FieldFilter: 筛选动作
	//     - "gender": 要筛选的字段名
	// 使用场景: 批量筛选、批量操作等
	info.AddSelectBox(ctx, "gender", types.FieldOptions{
		{Value: "0", Text: "男"},
		{Value: "1", Text: "女"},
	}, action.FieldFilter("gender"))

	// 设置表格基本信息
	// SetTable: 指定数据库表名
	// SetTitle: 设置表格标题（显示在页面头部）
	// SetDescription: 设置表格描述
	info.SetTable("users").SetTitle("用户").SetDescription("用户")

	// 获取表单配置对象
	// GetForm 返回表格的表单配置器，用于配置编辑/添加视图的字段
	formList := userTable.GetForm()

	// 添加 ID 字段到表单
	// 参数说明:
	//   - "ID": 字段显示名称
	//   - "id": 数据库字段名
	//   - db.Int: 字段数据类型
	//   - form.Default: 表单字段类型（默认文本框）
	// FieldNotAllowEdit: 禁止编辑该字段（编辑模式下只读）
	// FieldNotAllowAdd: 禁止添加该字段（新增模式下不显示）
	formList.AddField("编号", "id", db.Int, form.Default).FieldNotAllowEdit().FieldNotAllowAdd()

	// 添加 Ip 字段到表单
	// 参数说明:
	//   - "Ip": 字段显示名称
	//   - "ip": 数据库字段名
	//   - db.Varchar: 字段数据类型
	//   - form.Text: 表单字段类型（文本输入框）
	formList.AddField("IP", "ip", db.Varchar, form.Text)

	// 添加 Name 字段到表单
	// 参数说明:
	//   - "Name": 字段显示名称
	//   - "name": 数据库字段名
	//   - db.Varchar: 字段数据类型
	//   - form.Text: 表单字段类型（文本输入框）
	formList.AddField("姓名", "name", db.Varchar, form.Text)

	// 添加 Gender 字段到表单（单选按钮）
	// 参数说明:
	//   - "Gender": 字段显示名称
	//   - "gender": 数据库字段名
	//   - db.Tinyint: 字段数据类型
	//   - form.Radio: 表单字段类型（单选按钮）
	// FieldOptions: 设置单选按钮选项
	//   Text: 选项显示文本
	//   Value: 选项值
	// FieldDefault: 设置默认值
	formList.AddField("性别", "gender", db.Tinyint, form.Radio).
		FieldOptions(types.FieldOptions{
			{Text: "男", Value: "0"},
			{Text: "女", Value: "1"},
		}).FieldDefault("0")

	// 添加 Phone 字段到表单
	// 参数说明:
	//   - "Phone": 字段显示名称
	//   - "phone": 数据库字段名
	//   - db.Varchar: 字段数据类型
	//   - form.Text: 表单字段类型（文本输入框）
	formList.AddField("电话", "phone", db.Varchar, form.Text)

	// 添加 Country 字段到表单（单选下拉框，支持级联选择）
	// 参数说明:
	//   - "Country": 字段显示名称
	//   - "country": 数据库字段名
	//   - db.Tinyint: 字段数据类型
	//   - form.SelectSingle: 表单字段类型（单选下拉框）
	// FieldOptions: 设置下拉框选项
	// FieldDefault: 设置默认值
	// FieldOnChooseAjax: 设置级联选择（当选择国家时，动态加载城市列表）
	//   - "city": 级联字段名（城市字段）
	//   - "/choose/country": AJAX 请求路由
	//   - 回调函数: 处理 AJAX 请求，返回城市选项列表
	formList.AddField("国家", "country", db.Tinyint, form.SelectSingle).
		FieldOptions(types.FieldOptions{
			{Text: "中国", Value: "0"},
			{Text: "美国", Value: "1"},
			{Text: "英国", Value: "2"},
			{Text: "加拿大", Value: "3"},
		}).FieldDefault("0").FieldOnChooseAjax("city", "/choose/country",
		func(ctx *context.Context) (bool, string, interface{}) {
			// 获取用户选择的国家值
			country := ctx.FormValue("value")
			// 创建城市选项列表
			var data = make(selection.Options, 0)
			// 根据选择的国家返回对应的城市列表
			switch country {
			case "0": // 中国
				data = selection.Options{
					{Text: "北京", ID: "beijing"},
					{Text: "上海", ID: "shangHai"},
					{Text: "广州", ID: "guangZhou"},
					{Text: "深圳", ID: "shenZhen"},
				}
			case "1": // 美国
				data = selection.Options{
					{Text: "洛杉矶", ID: "los angeles"},
					{Text: "华盛顿特区", ID: "washington, dc"},
					{Text: "纽约", ID: "new york"},
					{Text: "拉斯维加斯", ID: "las vegas"},
				}
			case "2": // 英国
				data = selection.Options{
					{Text: "伦敦", ID: "london"},
					{Text: "剑桥", ID: "cambridge"},
					{Text: "曼彻斯特", ID: "manchester"},
					{Text: "利物浦", ID: "liverpool"},
				}
			case "3": // 加拿大
				data = selection.Options{
					{Text: "温哥华", ID: "vancouver"},
					{Text: "多伦多", ID: "toronto"},
				}
			default: // 默认（中国）
				data = selection.Options{
					{Text: "北京", ID: "beijing"},
					{Text: "上海", ID: "shangHai"},
					{Text: "广州", ID: "guangZhou"},
					{Text: "深圳", ID: "shenZhen"},
				}
			}
			// 返回成功状态、消息和城市选项列表
			return true, "ok", data
		})

	// 添加 City 字段到表单（单选下拉框，动态初始化）
	// 参数说明:
	//   - "City": 字段显示名称
	//   - "city": 数据库字段名
	//   - db.Varchar: 字段数据类型
	//   - form.SelectSingle: 表单字段类型（单选下拉框）
	// FieldOptionInitFn: 设置选项初始化函数
	//   根据当前值动态生成选项列表
	//   这里创建一个只包含当前值的选项，并设置为选中状态
	formList.AddField("城市", "city", db.Varchar, form.SelectSingle).
		FieldOptionInitFn(func(val types.FieldModel) types.FieldOptions {
			return types.FieldOptions{
				{Value: val.Value, Text: val.Value, Selected: true},
			}
		})

	// 添加 Custom Field 字段到表单（自定义字段，带后置过滤函数）
	// 参数说明:
	//   - "Custom Field": 字段显示名称
	//   - "role": 数据库字段名
	//   - db.Varchar: 字段数据类型
	//   - form.Text: 表单字段类型（文本输入框）
	// FieldPostFilterFn: 设置后置过滤函数（表单提交后执行）
	//   value.PostFieldModel: 包含提交的字段值和相关信息
	//   返回值: 过滤后的值（返回空字符串表示不保存）
	formList.AddField("自定义字段", "role", db.Varchar, form.Text).
		FieldPostFilterFn(func(value types.PostFieldModel) interface{} {
			// 打印提交的字段值（用于调试）
			fmt.Println("user custom field", value)
			// 返回空字符串，表示不保存该字段的值
			return ""
		})

	// 添加 UpdatedAt 字段到表单
	// 参数说明:
	//   - "UpdatedAt": 字段显示名称
	//   - "updated_at": 数据库字段名
	//   - db.Timestamp: 字段数据类型
	//   - form.Default: 表单字段类型（默认文本框）
	// FieldNotAllowAdd: 禁止添加该字段（新增模式下不显示）
	formList.AddField("更新时间", "updated_at", db.Timestamp, form.Default).FieldNotAllowAdd()

	// 添加 CreatedAt 字段到表单
	// 参数说明:
	//   - "CreatedAt": 字段显示名称
	//   - "created_at": 数据库字段名
	//   - db.Timestamp: 字段数据类型
	//   - form.Default: 表单字段类型（默认文本框）
	// FieldNotAllowAdd: 禁止添加该字段（新增模式下不显示）
	formList.AddField("创建时间", "created_at", db.Timestamp, form.Default).FieldNotAllowAdd()

	// 设置表单分组（标签页）
	// SetTabGroups 将表单字段分组到不同的标签页
	// types.NewTabGroups: 创建第一个标签页组
	//   参数: 要包含在第一个标签页中的字段名
	// AddGroup: 添加第二个标签页组
	//   参数: 要包含在第二个标签页中的字段名
	// SetTabHeaders: 设置标签页的标题
	//   参数: 各个标签页的标题
	userTable.GetForm().SetTabGroups(types.
		NewTabGroups("id", "ip", "name", "gender", "country", "city").
		AddGroup("phone", "role", "created_at", "updated_at")).
		SetTabHeaders("档案1", "档案2")

	// 设置表单基本信息
	// SetTable: 指定数据库表名
	// SetTitle: 设置表单标题
	// SetDescription: 设置表单描述
	formList.SetTable("users").SetTitle("用户").SetDescription("用户")

	// 设置表单后置钩子
	// SetPostHook 设置表单提交后的回调函数
	// values form2.Values: 表单提交的所有字段值
	// 返回值: error: 如果返回错误，表单提交失败；如果返回 nil，表单提交成功
	// 使用场景: 数据验证、数据处理、发送通知等
	formList.SetPostHook(func(values form2.Values) error {
		// 打印表单提交的值（用于调试）
		fmt.Println("userTable.GetForm().PostHook", values)
		// 返回 nil 表示表单提交成功
		return nil
	})

	// 返回配置好的表格模型
	return
}
