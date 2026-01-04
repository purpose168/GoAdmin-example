// pages 包 - 页面处理器
// 本文件定义表单页面的生成逻辑
// 使用 GoAdmin 框架的表单组件系统
// 作者: GoAdminGroup
// 创建日期: 2024
// 功能: 提供各种表单字段的示例页面，展示GoAdmin表单功能

package pages

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/modules/language"
	form2 "github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
	template2 "github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/icon"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

// GetFormContent 返回表单页面的内容
// 该函数生成并返回一个包含各种表单字段的示例页面
//
// 参数:
//   - ctx: 请求上下文对象，包含请求相关的信息和方法
//
// 返回值:
//   - types.Panel: 页面面板对象，包含页面内容、标题和描述
//   - error: 错误信息，如果生成页面失败则返回错误
//
// 功能说明:
//  1. 创建表单组件和按钮
//  2. 添加各种类型的表单字段（文本、数字、日期、选择框等）
//  3. 配置表单的验证规则和默认值
//  4. 设置表单的标签页分组
//  5. 生成最终的页面面板
//
// 表单字段类型:
//   - 基础输入: Text, Number, Password, Email, Url, Ip
//   - 日期时间: Date, Datetime, DatetimeRange, DateRange
//   - 文件上传: Multifile
//   - 数值输入: Currency, Rate, Slider
//   - 富文本: RichText, Code
//   - 选择控件: SelectBox, Select, SelectSingle, Radio, Checkbox, CheckboxStacked
//   - 开关控件: Switch
//   - 数组控件: Array
//   - 表格控件: Table
//
// 页面布局:
//   - 使用标签页分组，分为三个标签页:
//     1. input: 基础输入字段
//     2. select: 选择类字段
//     3. multi: 多值字段和表格
//
// 使用示例:
//
//	import "github.com/GoAdminGroup/example/pages"
//
//	// 在路由中注册页面
//	eng.HTML("GET", "/admin/form", pages.GetFormContent)
//
// 注意事项:
//   - 该函数展示了GoAdmin表单系统的各种功能
//   - 表单提交地址为 /admin/form/update
//   - 使用了语言包支持多语言
//   - 表单字段可以根据需要增删或修改
//
// GoAdmin表单系统说明:
//   - FormPanel: 表单面板，用于管理表单字段
//   - AddField: 添加单个字段
//   - AddRow: 添加一行多个字段
//   - AddTable: 添加表格字段
//   - SetTabGroups: 设置标签页分组
func GetFormContent(ctx *context.Context) (types.Panel, error) {
	// 获取模板组件实例
	// template2.Get: 根据上下文和主题获取模板组件
	// config.GetTheme(): 获取当前使用的主题（如adminlte）
	components := template2.Get(ctx, config.GetTheme())

	// 创建第一列（占位列）
	// Col(): 创建列组件
	// GetContent(): 获取列的HTML内容
	col1 := components.Col().GetContent()

	// 创建提交按钮
	// Button(): 创建按钮组件
	// SetType("submit"): 设置按钮类型为提交按钮
	// SetContent: 设置按钮显示文本，使用语言包获取"Save"的翻译
	// SetThemePrimary: 设置按钮主题为主要样式（蓝色）
	// SetOrientationRight: 设置按钮靠右对齐
	// SetLoadingText: 设置加载时显示的文本，包含旋转图标
	// GetContent(): 获取按钮的HTML内容
	btn1 := components.Button().SetType("submit").
		SetContent(language.GetFromHtml("Save")).
		SetThemePrimary().
		SetOrientationRight().
		SetLoadingText(icon.Icon("fa-spinner fa-spin", 2) + `保存中`).
		GetContent()

	// 创建重置按钮
	// SetType("reset"): 设置按钮类型为重置按钮
	// SetContent: 设置按钮显示文本，使用语言包获取"Reset"的翻译
	// SetThemeWarning: 设置按钮主题为警告样式（黄色）
	// SetOrientationLeft: 设置按钮靠左对齐
	btn2 := components.Button().SetType("reset").
		SetContent(language.GetFromHtml("Reset")).
		SetThemeWarning().
		SetOrientationLeft().
		GetContent()

	// 创建第二列，包含两个按钮
	// SetSize: 设置列的宽度，SizeMD(8)表示在中等屏幕上占8/12
	// SetContent: 设置列的内容为两个按钮
	col2 := components.Col().SetSize(types.SizeMD(8)).
		SetContent(btn1 + btn2).GetContent()

	// 创建新的表单面板
	// NewFormPanel: 创建一个空的表单面板
	// FormPanel用于管理表单的所有字段和配置
	var panel = types.NewFormPanel()

	// ========== 基础输入字段 ==========

	// 添加姓名字段（文本输入）
	// AddField参数说明:
	//   - 第1个: 字段显示名称
	//   - 第2个: 字段数据库字段名
	//   - 第3个: 字段数据库类型（db.Varchar表示变长字符串）
	//   - 第4个: 表单控件类型（form.Text表示文本输入框）
	panel.AddField("姓名", "name", db.Varchar, form.Text)

	// 添加年龄字段（数字输入）
	// db.Int: 数据库整数类型
	// form.Number: 数字输入框
	panel.AddField("年龄", "age", db.Int, form.Number)

	// 添加主页字段（URL输入）
	// FieldDefault: 设置字段的默认值
	panel.AddField("主页", "homepage", db.Varchar, form.Url).FieldDefault("http://google.com")

	// 添加邮箱字段（邮箱输入）
	// form.Email: 邮箱输入框，会自动验证邮箱格式
	panel.AddField("邮箱", "email", db.Varchar, form.Email).FieldDefault("xxxx@xxx.com")

	// ========== 日期时间字段 ==========

	// 添加生日字段（日期时间输入）
	// form.Date: 日期选择器
	panel.AddField("生日", "birthday", db.Varchar, form.Date).FieldDefault("2010-09-03 18:09:05")

	// 添加时间字段（日期时间输入）
	// form.Datetime: 日期时间选择器
	panel.AddField("时间", "time", db.Varchar, form.Datetime).FieldDefault("2010-09-05")

	// 添加时间范围字段（日期时间范围选择）
	// form.DatetimeRange: 日期时间范围选择器
	panel.AddField("时间范围", "time_range", db.Varchar, form.DatetimeRange)

	// 添加日期范围字段（日期范围选择）
	// form.DateRange: 日期范围选择器
	panel.AddField("日期范围", "date_range", db.Varchar, form.DateRange)

	// ========== 安全字段 ==========

	// 添加密码字段（密码输入）
	// form.Password: 密码输入框，输入内容会被隐藏
	// FieldDivider: 添加分隔线，用于视觉分组
	panel.AddField("密码", "password", db.Varchar, form.Password).FieldDivider("分隔线")

	// 添加IP地址字段（IP输入）
	// form.Ip: IP地址输入框，会自动验证IP格式
	panel.AddField("IP", "ip", db.Varchar, form.Ip)

	// ========== 文件上传字段 ==========

	// 添加证书字段（多文件上传）
	// form.Multifile: 多文件上传组件
	// FieldOptionExt: 设置扩展选项
	// maxFileCount: 最大文件数量限制为10
	panel.AddField("证书", "certificate", db.Varchar, form.Multifile).FieldOptionExt(map[string]interface{}{
		"maxFileCount": 10,
	})

	// ========== 数值字段 ==========

	// 添加金额字段（货币输入）
	// db.Int: 使用整数存储金额（避免浮点数精度问题）
	// form.Currency: 货币输入框，会自动格式化显示
	panel.AddField("金额", "currency", db.Int, form.Currency)

	// 添加评分字段（评分输入）
	// form.Rate: 评分组件，通常用于星级评分
	panel.AddField("评分", "rate", db.Int, form.Rate)

	// 添加奖励字段（滑块输入）
	// form.Slider: 滑块组件，用于选择数值范围
	// FieldOptionExt: 设置滑块选项
	// max: 最大值1000
	// min: 最小值1
	// step: 步长1
	// postfix: 后缀符号"$"
	panel.AddField("奖励", "reward", db.Int, form.Slider).FieldOptionExt(map[string]interface{}{
		"max":     1000,
		"min":     1,
		"step":    1,
		"postfix": "$",
	})

	// ========== 富文本字段 ==========

	// 添加内容字段（富文本编辑器）
	// db.Text: 数据库文本类型（长文本）
	// form.RichText: 富文本编辑器，支持HTML格式
	// FieldDefault: 设置默认的富文本内容（包含HTML标签）
	panel.AddField("内容", "content", db.Text, form.RichText).
		FieldDefault(`<h1>343434</h1><p>34344433434</p><ol><li>23234</li><li>2342342342</li><li>asdfads</li></ol><ul><li>3434334</li><li>34343343434</li><li>44455</li></ul><p><span style="color: rgb(194, 79, 74);">343434</span></p><p><span style="background-color: rgb(194, 79, 74); color: rgb(0, 0, 0);">434434433434</span></p><table border="0" width="100%" cellpadding="0" cellspacing="0"><tbody><tr><td>&nbsp;</td><td>&nbsp;</td><td>&nbsp;</td></tr><tr><td>&nbsp;</td><td>&nbsp;</td><td>&nbsp;</td></tr><tr><td>&nbsp;</td><td>&nbsp;</td><td>&nbsp;</td></tr><tr><td>&nbsp;</td><td>&nbsp;</td><td>&nbsp;</td></tr></tbody></table><p><br></p><p><span style="color: rgb(194, 79, 74);"><br></span></p>`).
		FieldDivider("分隔线 2")

	// 添加代码字段（代码编辑器）
	// form.Code: 代码编辑器，支持语法高亮
	// FieldDefault: 设置默认的Go代码
	panel.AddField("代码", "code", db.Text, form.Code).FieldDefault(`package main

import "fmt"

func main() {
	fmt.Println("hello GoAdmin!")
}
`)

	// ========== 选择控件字段 ==========

	// 添加网站开关字段（开关控件）
	// db.Tinyint: 数据库微整型（0或1）
	// form.Switch: 开关组件，类似iOS的开关按钮
	// FieldHelpMsg: 设置帮助提示信息
	// FieldOptions: 设置选项（0表示关闭，1表示开启）
	panel.AddField("网站", "website", db.Tinyint, form.Switch).
		FieldHelpMsg("关闭后网站将无法访问，但管理系统仍可登录").
		FieldOptions(types.FieldOptions{
			{Value: "0"},
			{Value: "1"},
		})

	// 添加水果字段（下拉选择框）
	// form.SelectBox: 下拉选择框组件
	// FieldOptions: 设置选项列表
	// Text: 选项显示文本
	// Value: 选项实际值
	// FieldDisplay: 设置显示函数，用于自定义显示逻辑
	panel.AddField("水果", "fruit", db.Varchar, form.SelectBox).
		FieldOptions(types.FieldOptions{
			{Text: "苹果", Value: "apple"},
			{Text: "香蕉", Value: "banana"},
			{Text: "西瓜", Value: "watermelon"},
			{Text: "梨", Value: "pear"},
		}).
		FieldDisplay(func(value types.FieldModel) interface{} {
			return []string{"梨"}
		})

	// 添加性别字段（单选按钮）
	// form.Radio: 单选按钮组件
	// Value: "0"表示男性，"1"表示女性
	panel.AddField("性别", "gender", db.Tinyint, form.Radio).
		FieldOptions(types.FieldOptions{
			{Text: "男", Value: "0"},
			{Text: "女", Value: "1"},
		})

	// 添加饮料字段（下拉选择框）
	// form.Select: 下拉选择框组件
	// FieldDefault: 设置默认值为"beer"
	panel.AddField("饮料", "drink", db.Tinyint, form.Select).
		FieldOptions(types.FieldOptions{
			{Text: "啤酒", Value: "beer"},
			{Text: "果汁", Value: "juice"},
			{Text: "水", Value: "water"},
			{Text: "红牛", Value: "red bull"},
		}).FieldDefault("beer")

	// 添加工作经验字段（单选下拉框）
	// form.SelectSingle: 单选下拉框组件
	panel.AddField("工作经验", "experience", db.Tinyint, form.SelectSingle).
		FieldOptions(types.FieldOptions{
			{Text: "两年", Value: "0"},
			{Text: "三年", Value: "1"},
			{Text: "四年", Value: "2"},
			{Text: "五年", Value: "3"},
		}).FieldDefault("beer")

	// 添加零食字段（复选框）
	// form.Checkbox: 复选框组件，支持多选
	panel.AddField("零食", "snacks", db.Varchar, form.Checkbox).
		FieldOptions(types.FieldOptions{
			{Text: "麦片", Value: "0"},
			{Text: "薯片", Value: "1"},
			{Text: "辣条", Value: "2"},
			{Text: "冰淇淋", Value: "3"},
		})

	// 添加猫咪品种字段（堆叠复选框）
	// form.CheckboxStacked: 堆叠复选框组件，选项垂直排列
	panel.AddField("猫咪品种", "cat", db.Varchar, form.CheckboxStacked).
		FieldOptions(types.FieldOptions{
			{Text: "加菲猫", Value: "0"},
			{Text: "英短", Value: "1"},
			{Text: "美短", Value: "2"},
		})

	// ========== 多字段行 ==========

	// 添加一行三个字段（省、市、区）
	// AddRow: 添加一行多个字段
	// FieldRowWidth: 设置字段在行中的宽度（2/12）
	panel.AddRow(func(pa *types.FormPanel) {
		// 省份字段
		panel.AddField("省份", "province", db.Tinyint, form.SelectSingle).
			FieldOptions(types.FieldOptions{
				{Text: "北京", Value: "0"},
				{Text: "上海", Value: "1"},
				{Text: "广东", Value: "2"},
				{Text: "重庆", Value: "3"},
			}).FieldRowWidth(2)
		// 城市字段
		// FieldHeadWidth: 设置标签宽度（2/12）
		// FieldInputWidth: 设置输入框宽度（10/12）
		panel.AddField("城市", "city", db.Tinyint, form.SelectSingle).
			FieldOptions(types.FieldOptions{
				{Text: "北京", Value: "0"},
				{Text: "上海", Value: "1"},
				{Text: "广州", Value: "2"},
				{Text: "深圳", Value: "3"},
			}).FieldRowWidth(3).FieldHeadWidth(2).FieldInputWidth(10)
		// 区域字段
		// FieldHeadWidth: 设置标签宽度（2/12）
		// FieldInputWidth: 设置输入框宽度（9/12）
		panel.AddField("区域", "district", db.Tinyint, form.SelectSingle).
			FieldOptions(types.FieldOptions{
				{Text: "朝阳", Value: "0"},
				{Text: "海珠", Value: "1"},
				{Text: "浦东", Value: "2"},
				{Text: "宝安", Value: "3"},
			}).FieldRowWidth(3).FieldHeadWidth(2).FieldInputWidth(9)
	})

	// ========== 多值和表格字段 ==========

	// 添加员工字段（数组输入）
	// form.Array: 数组组件，支持添加多个值
	panel.AddField("员工", "employee", db.Varchar, form.Array)

	// 添加设置表格字段
	// AddTable: 添加表格组件
	// 表格包含Key和Value两列
	panel.AddTable("设置", "setting", func(panel *types.FormPanel) {
		panel.AddField("键", "key", db.Varchar, form.Text).FieldHideLabel()
		panel.AddField("值", "value", db.Varchar, form.Text).FieldHideLabel()
	})

	// ========== 标签页分组 ==========

	// 设置标签页分组
	// 将所有字段分成三个标签页
	panel.SetTabGroups(types.TabGroups{
		// 第一个标签页: 基础输入字段
		{"name", "age", "homepage", "email", "birthday", "time", "time_range", "date_range", "password", "ip",
			"certificate", "currency", "rate", "reward", "content", "code"},
		// 第二个标签页: 选择类字段
		{"website", "snacks", "fruit", "gender", "cat", "drink", "province", "city", "district", "experience"},
		// 第三个标签页: 多值字段和表格
		{"employee", "setting"},
	})

	// 设置标签页标题
	// SetTabHeaders: 设置每个标签页的标题
	panel.SetTabHeaders("输入", "选择", "多值")

	// 分组字段并生成标签页内容
	// GroupField: 将字段按标签页分组
	// 返回值: fields（标签页内容）, headers（标签页标题）
	fields, headers := panel.GroupField()

	// 创建表单组件
	// Form(): 创建表单组件
	// SetTabHeaders: 设置标签页标题
	// SetTabContents: 设置标签页内容
	// SetPrefix: 设置URL前缀
	// SetUrl: 设置表单提交地址
	// SetTitle: 设置表单标题
	// SetHiddenFields: 设置隐藏字段
	// SetOperationFooter: 设置操作按钮区域
	aform := components.Form().
		SetTabHeaders(headers).
		SetTabContents(fields).
		SetPrefix(config.PrefixFixSlash()).
		SetUrl("/admin/form/update").
		SetTitle("表单").
		SetHiddenFields(map[string]string{
			form2.PreviousKey: "/admin",
		}).
		SetOperationFooter(col1 + col2)

	// 返回页面面板
	// Content: 页面内容，包含表单
	// Title: 页面标题
	// Callbacks: 回调函数
	// Description: 页面描述
	return types.Panel{
		Content: components.Box().
			SetHeader(aform.GetDefaultBoxHeader(true)).
			WithHeadBorder().
			SetBody(aform.GetContent()).
			GetContent(),
		Title:       "表单",
		Callbacks:   panel.Callbacks,
		Description: "表单示例",
	}, nil
}
