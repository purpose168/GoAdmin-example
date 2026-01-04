// Package pages 提供页面生成器，用于构建各种管理后台页面
// 本文件实现数据表格页面的生成功能
package pages

import (
	"fmt"

	"github.com/purpose168/GoAdmin/context"
	"github.com/purpose168/GoAdmin/modules/config"
	"github.com/purpose168/GoAdmin/plugins/admin/modules/paginator"
	"github.com/purpose168/GoAdmin/plugins/admin/modules/parameter"
	"github.com/purpose168/GoAdmin/template"
	"github.com/purpose168/GoAdmin/template/icon"
	"github.com/purpose168/GoAdmin/template/types"
	"github.com/purpose168/GoAdmin/template/types/action"
)

// GetTableContent 获取数据表格内容
// 该函数创建并返回一个包含数据表格的面板，用于展示示例数据
//
// 参数:
//
//	ctx: 上下文对象，包含请求信息和配置
//
// 返回值:
//
//	types.Panel: 包含表格内容的面板对象
//	error: 错误信息，如果创建成功则为 nil
//
// 功能说明:
//   - 创建数据表格组件并设置示例数据
//   - 配置表格的表头和主键
//   - 添加 AJAX 按钮操作
//   - 配置分页器
//   - 将表格包装在面板中返回
func GetTableContent(ctx *context.Context) (types.Panel, error) {

	// 获取当前主题的模板组件
	// template.Get 根据配置的主题名称返回对应的模板组件实例
	comp := template.Get(ctx, config.GetTheme())

	// 创建数据表格组件
	// DataTable() 返回一个数据表格构建器，用于配置表格的各种属性
	table := comp.DataTable().
		// 设置表格数据列表
		// InfoItem 是一个结构体，包含 Content 字段用于存储单元格内容
		// 这里设置了两行示例数据，每行包含 id、name、gender、age 四个字段
		SetInfoList([]map[string]types.InfoItem{
			{
				"id":     {Content: "0"},
				"name":   {Content: "杰克"},
				"gender": {Content: "男"},
				"age":    {Content: "20"},
			},
			{
				"id":     {Content: "1"},
				"name":   {Content: "简"},
				"gender": {Content: "女"},
				"age":    {Content: "23"},
			},
		}).
		// 设置主键字段
		// 主键用于标识表格中的每一行数据，通常用于操作按钮传递参数
		SetPrimaryKey("id").
		// 设置表头配置
		// Thead 定义表格的列结构，包括列标题和对应的字段名
		SetThead(types.Thead{
			{Head: "编号", Field: "id"},
			{Head: "姓名", Field: "name"},
			{Head: "性别", Field: "gender"},
			{Head: "年龄", Field: "age"},
		})

	// 创建按钮集合
	// Buttons 类型用于存储表格操作按钮
	allBtns := make(types.Buttons, 0)

	// 添加一个 AJAX 按钮操作
	// GetDefaultButton 创建一个默认样式的按钮
	// 参数说明:
	//   - "点击我": 按钮显示文本
	//   - icon.ArrowLeft: 按钮图标（左箭头）
	//   - action.Ajax: AJAX 动作配置
	//     - "ajax_id": 动作的唯一标识符
	//     - 回调函数：处理 AJAX 请求的逻辑
	//       参数: ctx - 上下文对象
	//       返回: success - 操作是否成功, msg - 返回消息, data - 附加数据
	allBtns = append(allBtns, types.GetDefaultButton("点击我", icon.ArrowLeft, action.Ajax("ajax_id",
		func(ctx *context.Context) (success bool, msg string, data interface{}) {
			// 打印请求参数中的 id 值
			// FormValue 用于获取表单或 URL 参数
			fmt.Println("上下文请求", ctx.FormValue("id"))
			// 返回成功状态、消息和空数据
			return true, "操作成功", nil
		})))

	// 生成按钮的 HTML 内容和 JavaScript 代码
	// Content 方法返回按钮的 HTML 和对应的 JS 代码
	// btns: 按钮的 HTML 内容
	// btnsJs: 按钮的 JavaScript 代码（用于处理点击事件等）
	btns, btnsJs := allBtns.Content(ctx)
	// 将按钮和 JS 代码设置到表格中
	table = table.SetButtons(btns).SetActionJs(btnsJs)

	// 创建回调函数集合
	// Callbacks 用于存储按钮操作的回调函数
	cbs := make(types.Callbacks, 0)
	// 遍历所有按钮，收集它们的回调函数
	for _, btn := range allBtns {
		cbs = append(cbs, btn.GetAction().GetCallbacks())
	}

	// 生成表格的 HTML 内容
	// GetContent 方法返回表格的完整 HTML 字符串
	body := table.GetContent()

	// 返回面板对象
	// Panel 是 GoAdmin 框架中的页面容器，可以包含各种组件
	return types.Panel{
		// 设置面板内容
		// Box 创建一个盒子容器，用于包装表格内容
		Content: comp.Box().
			// 设置盒子主体内容（表格 HTML）
			SetBody(body).
			// 设置无内边距样式
			SetNoPadding().
			// 设置盒子头部（表格标题和操作栏）
			SetHeader(table.GetDataTableHeader()).
			// 添加头部边框
			WithHeadBorder().
			// 设置盒子底部（分页器）
			// paginator.Get 创建分页器组件
			// Config 配置分页参数:
			//   - Size: 每页显示数量（50）
			//   - PageSizeList: 可选的每页显示数量列表
			//   - Param: 从请求 URL 中获取分页参数
			SetFooter(paginator.Get(ctx, paginator.Config{
				Size:         50,
				PageSizeList: []string{"10", "20", "30", "50"},
				Param:        parameter.GetParam(ctx.Request.URL, 10),
			}).GetContent()).
			// 生成盒子的完整 HTML 内容
			GetContent(),
		// 设置面板标题
		Title: "表格",
		// 设置面板描述
		Description: "表格示例",
		// 设置回调函数（用于处理按钮点击等事件）
		Callbacks: cbs,
	}, nil
}
