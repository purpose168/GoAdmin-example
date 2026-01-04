// pages 包 - 页面处理器
// 本文件定义仪表板页面的生成逻辑
// 使用 GoAdmin 框架的模板组件系统

// 创建日期: 2024
// 功能: 提供管理后台仪表板页面，展示各种统计信息和图表

package pages

import (
	"html/template"

	"github.com/purpose168/GoAdmin-example/models"
	"github.com/purpose168/GoAdmin-themes/adminlte/components/chart_legend"
	"github.com/purpose168/GoAdmin-themes/adminlte/components/description"
	"github.com/purpose168/GoAdmin-themes/adminlte/components/infobox"
	"github.com/purpose168/GoAdmin-themes/adminlte/components/productlist"
	"github.com/purpose168/GoAdmin-themes/adminlte/components/progress_group"
	"github.com/purpose168/GoAdmin-themes/adminlte/components/smallbox"
	"github.com/purpose168/GoAdmin/context"
	tmpl "github.com/purpose168/GoAdmin/template"
	"github.com/purpose168/GoAdmin/template/chartjs"
	"github.com/purpose168/GoAdmin/template/color"
	"github.com/purpose168/GoAdmin/template/icon"
	"github.com/purpose168/GoAdmin/template/types"
)

// DashboardPage 返回仪表板页面的内容
// 该函数生成并返回管理后台的仪表板页面，包含各种统计信息和图表
//
// 参数:
//   - ctx: 请求上下文对象，包含请求相关的信息和方法
//
// 返回值:
//   - types.Panel: 页面面板对象，包含页面内容、标题和描述
//   - error: 错误信息，如果生成页面失败则返回错误
//
// 功能说明:
//  1. 获取系统统计数据（CPU、点赞数、销售额、新会员数）
//  2. 创建信息框组件显示关键指标
//  3. 创建表格组件显示最新订单
//  4. 创建产品列表组件显示最近添加的产品
//  5. 创建折线图组件显示销售趋势
//  6. 创建进度条组件显示目标完成情况
//  7. 创建饼图组件显示浏览器使用情况
//  8. 创建标签页和弹窗组件
//
// 页面布局:
//   - 第一行: 4个信息框（CPU流量、点赞数、销售额、新会员数）
//   - 第二行: 订单表格和产品列表
//   - 第三行: 销售折线图和目标完成进度条
//   - 第四行: 浏览器使用饼图和标签页/弹窗
//
// 使用示例:
//
//	import "github.com/purpose168/GoAdmin-example/pages"
//
//	// 在路由中注册页面
//	eng.HTML("GET", "/admin", pages.DashboardPage)
//
// 注意事项:
//   - 该函数依赖models包的统计数据
//   - 使用了GoAdmin模板系统的各种组件
//   - 所有数据从数据库中获取
//   - 页面使用AdminLTE主题样式
func DashboardPage(ctx *context.Context) (types.Panel, error) {

	components := tmpl.Default()
	colComp := components.Col()

	// 获取统计数据
	// models.FirstStatics: 从数据库查询第一条统计记录
	// 返回值: 包含CPU使用率、点赞数、销售额、新会员数等统计信息
	// 如果数据库中没有记录，返回零值结构体
	statics := models.FirstStatics()

	/**************************
	 * Info Box
	/**************************/

	// 创建CPU流量信息框
	// statics.CPUTmpl(): 将CPU使用率转换为HTML格式
	// SetText: 设置显示文本为"CPU流量"
	// SetColor: 设置颜色为青色(Aqua)
	// SetNumber: 显示CPU使用率数值
	// SetIcon: 设置图标为齿轮图标
	infobox1 := infobox.New().
		SetText("CPU流量").
		SetColor(color.Aqua).
		SetNumber(statics.CPUTmpl()).
		SetIcon("ion-ios-gear-outline").
		GetContent()

	// 创建点赞数信息框
	// statics.LikesTmpl(): 将点赞数转换为HTML格式
	// SetText: 设置显示文本为"点赞"
	// SetColor: 设置颜色为红色(Red)
	// SetNumber: 显示点赞数，并添加美元符号后缀
	// SetIcon: 设置图标为Google Plus图标
	infobox2 := infobox.New().
		SetText("点赞").
		SetColor(color.Red).
		SetNumber(statics.LikesTmpl() + "<small>$</small>").
		SetIcon(icon.GooglePlus).
		GetContent()

	// 创建销售额信息框
	// statics.SalesTmpl(): 将销售额转换为HTML格式
	// SetText: 设置显示文本为"销售额"
	// SetColor: 设置颜色为绿色(Green)
	// SetNumber: 显示销售额数值
	// SetIcon: 设置图标为购物车图标
	infobox3 := infobox.New().
		SetText("销售额").
		SetColor(color.Green).
		SetNumber(statics.SalesTmpl()).
		SetIcon("ion-ios-cart-outline").
		GetContent()

	// 创建新会员数信息框
	// statics.NewMembersTmpl(): 将新会员数转换为HTML格式
	// SetText: 设置显示文本为"新会员"
	// SetColor: 设置颜色为黄色(Yellow)
	// SetNumber: 显示新会员数数值
	// SetIcon: 设置图标为用户群组图标
	infobox4 := infobox.New().
		SetText("新会员").
		SetColor(color.Yellow).
		SetNumber(statics.NewMembersTmpl()).
		SetIcon("ion-ios-people-outline").
		GetContent()

	var size = types.SizeMD(3).SM(6).XS(12)
	// 设置列的响应式宽度
	// SizeMD(3): 在中等屏幕上占3/12宽度
	// SM(6): 在小屏幕上占6/12宽度
	// XS(12): 在超小屏幕上占12/12宽度（全宽）
	infoboxCol1 := colComp.SetSize(size).SetContent(infobox1).GetContent()
	infoboxCol2 := colComp.SetSize(size).SetContent(infobox2).GetContent()
	infoboxCol3 := colComp.SetSize(size).SetContent(infobox3).GetContent()
	infoboxCol4 := colComp.SetSize(size).SetContent(infobox4).GetContent()
	// 创建第一行，包含4个信息框
	row1 := components.Row().SetContent(infoboxCol1 + infoboxCol2 + infoboxCol3 + infoboxCol4).GetContent()

	/**************************
	 * Box - 订单表格
	/**************************/

	// 创建表格组件显示最新订单
	// SetType("table"): 设置表格类型为标准表格
	// SetInfoList: 设置表格数据，每行是一个map，键是列名，值是单元格内容
	table := components.Table().SetType("table").SetInfoList([]map[string]types.InfoItem{
		{
			"订单ID": {Content: "OR9842"},
			"商品":   {Content: "使命召唤IV"},
			"状态":   {Content: "已发货"},
			"热度":   {Content: "90%"},
		}, {
			"订单ID": {Content: "OR9842"},
			"商品":   {Content: "使命召唤IV"},
			"状态":   {Content: "已发货"},
			"热度":   {Content: "90%"},
		}, {
			"订单ID": {Content: "OR9842"},
			"商品":   {Content: "使命召唤IV"},
			"状态":   {Content: "已发货"},
			"热度":   {Content: "90%"},
		}, {
			"订单ID": {Content: "OR9842"},
			"商品":   {Content: "使命召唤IV"},
			"状态":   {Content: "已发货"},
			"热度":   {Content: "90%"},
		},
	}).SetThead(types.Thead{
		// 设置表头
		{Head: "订单ID"},
		{Head: "商品"},
		{Head: "状态"},
		{Head: "热度"},
	}).GetContent()

	// 创建盒子组件包裹表格
	// WithHeadBorder: 显示头部边框
	// SetHeader: 设置盒子标题
	// SetHeadColor: 设置头部背景色
	// SetBody: 设置盒子内容（表格）
	// SetFooter: 设置底部内容（操作按钮）
	boxInfo := components.Box().
		WithHeadBorder().
		SetHeader("最新订单").
		SetHeadColor("#f7f7f7").
		SetBody(table).
		SetFooter(`<div class="clearfix"><a href="javascript:void(0)" class="btn btn-sm btn-info btn-flat pull-left">处理订单</a><a href="javascript:void(0)" class="btn btn-sm btn-default btn-flat pull-right">查看所有新订单</a> </div>`).
		GetContent()

	// 创建表格列，占8/12宽度
	tableCol := colComp.SetSize(types.SizeMD(8)).SetContent(row1 + boxInfo).GetContent()

	/**************************
	 * Product List - 产品列表
	/**************************/

	// 创建产品列表组件
	// SetData: 设置产品数据，每个产品包含图片、标题、标签、描述等信息
	productList := productlist.New().SetData([]map[string]string{
		{
			"img":         "//adminlte.io/themes/AdminLTE/dist/img/default-50x50.gif",
			"title":       "GoAdmin",
			"has_tabel":   "true",
			"labeltype":   "warning",
			"label":       "免费",
			"description": `一个帮助您构建数据可视化系统的框架`,
		}, {
			"img":         "//adminlte.io/themes/AdminLTE/dist/img/default-50x50.gif",
			"title":       "GoAdmin",
			"has_tabel":   "true",
			"labeltype":   "warning",
			"label":       "免费",
			"description": `一个帮助您构建数据可视化系统的框架`,
		}, {
			"img":         "//adminlte.io/themes/AdminLTE/dist/img/default-50x50.gif",
			"title":       "GoAdmin",
			"has_tabel":   "true",
			"labeltype":   "warning",
			"label":       "免费",
			"description": `一个帮助您构建数据可视化系统的框架`,
		}, {
			"img":         "//adminlte.io/themes/AdminLTE/dist/img/default-50x50.gif",
			"title":       "GoAdmin",
			"has_tabel":   "true",
			"labeltype":   "warning",
			"label":       "免费",
			"description": `一个帮助您构建数据可视化系统的框架`,
		},
	}).GetContent()

	// 创建警告主题的盒子组件包裹产品列表
	// SetTheme("warning"): 设置主题为警告样式（黄色）
	// WithHeadBorder: 显示头部边框
	// SetHeader: 设置盒子标题
	// SetBody: 设置盒子内容（产品列表）
	// SetFooter: 设置底部内容（查看所有产品链接）
	boxWarning := components.Box().SetTheme("warning").WithHeadBorder().SetHeader("最近添加的产品").
		SetBody(productList).
		SetFooter(`<a href="javascript:void(0)" class="uppercase">查看所有产品</a>`).
		GetContent()

	// 创建产品列表列，占4/12宽度
	newsCol := colComp.SetSize(types.SizeMD(4)).SetContent(boxWarning).GetContent()

	// 创建第五行，包含表格列和产品列表列
	row5 := components.Row().SetContent(tableCol + newsCol).GetContent()

	/**************************
	 * Box - 销售折线图和目标完成进度条
	/**************************/

	// 创建折线图组件
	// chartjs.Line(): 创建Chart.js折线图实例
	line := chartjs.Line()

	// 配置折线图
	// SetID: 设置图表ID，用于在HTML中引用
	// SetHeight: 设置图表高度（像素）
	// SetTitle: 设置图表标题
	// SetLabels: 设置X轴标签（月份）
	// AddDataSet: 添加数据集
	// DSData: 设置数据集的数值
	// DSFill: 设置是否填充区域（false表示不填充）
	// DSBorderColor: 设置线条颜色
	// DSLineTension: 设置线条张力（0.1表示轻微曲线）
	lineChart := line.
		SetID("salechart").
		SetHeight(180).
		SetTitle("销售额: 2019年1月1日 - 2019年7月30日").
		SetLabels([]string{"一月", "二月", "三月", "四月", "五月", "六月", "七月"}).
		AddDataSet("电子产品").
		DSData([]float64{65, 59, 80, 81, 56, 55, 40}).
		DSFill(false).
		DSBorderColor("rgb(210, 214, 222)").
		DSLineTension(0.1).
		AddDataSet("数字商品").
		DSData([]float64{28, 48, 40, 19, 86, 27, 90}).
		DSFill(false).
		DSBorderColor("rgba(60,141,188,1)").
		DSLineTension(0.1).
		GetContent()

	// 创建进度条组件的标题
	title := `<p class="text-center"><strong>目标完成情况</strong></p>`

	// 创建第一个进度条组件
	// SetTitle: 设置进度条标题
	// SetColor: 设置进度条颜色
	// SetDenominator: 设置分母（总数）
	// SetMolecular: 设置分子（当前值）
	// SetPercent: 设置百分比
	progressGroup := progress_group.New().
		SetTitle("添加商品到购物车").
		SetColor("#76b2d4").
		SetDenominator(200).
		SetMolecular(160).
		SetPercent(80).
		GetContent()

	// 创建第二个进度条组件
	progressGroup1 := progress_group.New().
		SetTitle("完成购买").
		SetColor("#f17c6e").
		SetDenominator(400).
		SetMolecular(310).
		SetPercent(80).
		GetContent()

	// 创建第三个进度条组件
	progressGroup2 := progress_group.New().
		SetTitle("访问高级页面").
		SetColor("#ace0ae").
		SetDenominator(800).
		SetMolecular(490).
		SetPercent(80).
		GetContent()

	// 创建第四个进度条组件
	progressGroup3 := progress_group.New().
		SetTitle("发送咨询").
		SetColor("#fdd698").
		SetDenominator(500).
		SetMolecular(250).
		SetPercent(50).
		GetContent()

	// 创建内部第一列，包含折线图
	boxInternalCol1 := colComp.SetContent(lineChart).SetSize(types.SizeMD(8)).GetContent()

	// 创建内部第二列，包含4个进度条
	boxInternalCol2 := colComp.
		SetContent(template.HTML(title) + progressGroup + progressGroup1 + progressGroup2 + progressGroup3).
		SetSize(types.SizeMD(4)).
		GetContent()

	// 创建内部第一行，包含折线图和进度条
	boxInternalRow := components.Row().SetContent(boxInternalCol1 + boxInternalCol2).GetContent()

	// 创建描述组件1
	// SetPercent: 设置百分比变化
	// SetNumber: 设置数值
	// SetTitle: 设置标题
	// SetArrow: 设置箭头方向（up表示上升）
	// SetColor: 设置颜色（green表示增长）
	// SetBorder: 设置边框位置（right表示右边框）
	description1 := description.New().
		SetPercent("17").
		SetNumber("¥140,100").
		SetTitle("总收入").
		SetArrow("up").
		SetColor("green").
		SetBorder("right").
		GetContent()

	// 创建描述组件2
	description2 := description.New().
		SetPercent("2").
		SetNumber("440,560").
		SetTitle("总收入").
		SetArrow("down").
		SetColor("red").
		SetBorder("right").
		GetContent()

	// 创建描述组件3
	description3 := description.New().
		SetPercent("12").
		SetNumber("¥140,050").
		SetTitle("总收入").
		SetArrow("up").
		SetColor("green").
		SetBorder("right").
		GetContent()

	// 创建描述组件4
	description4 := description.New().
		SetPercent("1").
		SetNumber("30943").
		SetTitle("总收入").
		SetArrow("up").
		SetColor("green").
		GetContent()

	// 设置小屏幕尺寸
	size2 := types.SizeSM(3).XS(6)

	// 创建内部第三列到第六列，包含4个描述组件
	boxInternalCol3 := colComp.SetContent(description1).SetSize(size2).GetContent()
	boxInternalCol4 := colComp.SetContent(description2).SetSize(size2).GetContent()
	boxInternalCol5 := colComp.SetContent(description3).SetSize(size2).GetContent()
	boxInternalCol6 := colComp.SetContent(description4).SetSize(size2).GetContent()

	// 创建内部第二行，包含4个描述组件
	boxInternalRow2 := components.Row().SetContent(boxInternalCol3 + boxInternalCol4 + boxInternalCol5 + boxInternalCol6).GetContent()

	// 创建盒子组件包裹内部内容
	// WithHeadBorder: 显示头部边框
	// SetHeader: 设置盒子标题
	// SetBody: 设置盒子内容（折线图和进度条）
	// SetFooter: 设置底部内容（描述组件）
	box := components.Box().WithHeadBorder().SetHeader("月度总结报告").
		SetBody(boxInternalRow).
		SetFooter(boxInternalRow2).
		GetContent()

	// 创建盒子列，占12/12宽度（全宽）
	boxcol := colComp.SetContent(box).SetSize(types.SizeMD(12)).GetContent()

	// 创建第二行，包含月度报告盒子
	row2 := components.Row().SetContent(boxcol).GetContent()

	/**************************
	 * Small Box - 小盒子组件
	/**************************/

	// 创建4个小盒子组件
	// SetColor: 设置盒子颜色
	// SetIcon: 设置图标
	// SetUrl: 设置点击跳转的URL
	// SetTitle: 设置标题
	// SetValue: 设置显示的数值
	smallbox1 := smallbox.New().SetColor("blue").SetIcon("ion-ios-gear-outline").SetUrl("/").SetTitle("新用户").SetValue("345￥").GetContent()
	smallbox2 := smallbox.New().SetColor("yellow").SetIcon("ion-ios-cart-outline").SetUrl("/").SetTitle("新用户").SetValue("80%").GetContent()
	smallbox3 := smallbox.New().SetColor("red").SetIcon("fa-user").SetUrl("/").SetTitle("新用户").SetValue("645￥").GetContent()
	smallbox4 := smallbox.New().SetColor("green").SetIcon("ion-ios-cart-outline").SetUrl("/").SetTitle("新用户").SetValue("889￥").GetContent()

	// 创建4个列，每个列包含一个小盒子
	col1 := colComp.SetSize(size).SetContent(smallbox1).GetContent()
	col2 := colComp.SetSize(size).SetContent(smallbox2).GetContent()
	col3 := colComp.SetSize(size).SetContent(smallbox3).GetContent()
	col4 := colComp.SetSize(size).SetContent(smallbox4).GetContent()

	// 创建第三行，包含4个小盒子
	row3 := components.Row().SetContent(col1 + col2 + col3 + col4).GetContent()

	/**************************
	 * Pie Chart - 饼图组件
	/**************************/

	// 创建饼图组件
	// chartjs.Pie(): 创建Chart.js饼图实例
	// SetHeight: 设置图表高度（像素）
	// SetLabels: 设置标签（浏览器名称）
	// SetID: 设置图表ID
	// AddDataSet: 添加数据集
	// DSData: 设置数据集的数值
	// DSBackgroundColor: 设置每个扇区的背景色
	pie := chartjs.Pie().
		SetHeight(170).
		SetLabels([]string{"导航器", "欧朋", "Safari", "火狐", "IE", "Chrome"}).
		SetID("pieChart").
		AddDataSet("浏览器").
		DSData([]float64{100, 300, 600, 400, 500, 700}).
		DSBackgroundColor([]chartjs.Color{
			"rgb(255, 205, 86)", "rgb(54, 162, 235)", "rgb(255, 99, 132)", "rgb(255, 205, 86)", "rgb(54, 162, 235)", "rgb(255, 99, 132)",
		}).
		GetContent()

	// 创建图例组件
	// SetData: 设置图例数据，包含标签和颜色
	legend := chart_legend.New().SetData([]map[string]string{
		{
			"label": " Chrome",
			"color": "red",
		}, {
			"label": " IE",
			"color": "Green",
		}, {
			"label": " 火狐",
			"color": "yellow",
		}, {
			"label": " Safari",
			"color": "blue",
		}, {
			"label": " 欧朋",
			"color": "light-blue",
		}, {
			"label": " 导航器",
			"color": "gray",
		},
	}).GetContent()

	// 创建危险主题的盒子组件包裹饼图和图例
	// SetTheme("danger"): 设置主题为危险样式（红色）
	// WithHeadBorder: 显示头部边框
	// SetHeader: 设置盒子标题
	// SetBody: 设置盒子内容（饼图和图例）
	// SetFooter: 设置底部内容（查看所有用户链接）
	boxDanger := components.Box().SetTheme("danger").WithHeadBorder().SetHeader("浏览器使用情况").
		SetBody(components.Row().
			SetContent(colComp.SetSize(types.SizeMD(8)).
				SetContent(pie).
				GetContent() + colComp.SetSize(types.SizeMD(4)).
				SetContent(legend).
				GetContent()).GetContent()).
		SetFooter(`<p class="text-center"><a href="javascript:void(0)" class="uppercase">查看所有用户</a></p>`).
		GetContent()

	/**************************
	 * Tabs - 标签页组件
	/**************************/

	// 创建标签页组件
	// SetData: 设置标签页数据，每个标签页包含标题和内容
	tabs := components.Tabs().SetData([]map[string]template.HTML{
		{
			"title": "标签页1",
			"content": template.HTML(`<b>如何使用:</b>

                <p>与原始的Bootstrap标签页完全相同，除了您应该使用
                  自定义包装器 <code>.nav-tabs-custom</code> 来实现这种样式。</p>
                一种美妙的宁静占据了我的整个灵魂，
                就像这些我全心全意享受的春天的甜美早晨。
                我独自一人，在这个地方感受到存在的魅力，
                这个地方是为像我这样的灵魂的幸福而创造的。我很快乐，
                我亲爱的朋友，如此沉浸在单纯的宁静存在的精致感觉中，
                以至于我忽视了我的天赋。在目前这一刻，我无法画出一笔
                ；然而我觉得我从未像现在这样成为一个伟大的艺术家。`),
		}, {
			"title": "标签页2",
			"content": template.HTML(`
                欧洲语言属于同一个家族。它们各自的存在是一个神话。
                对于科学、音乐、体育等，欧洲使用相同的词汇。这些语言仅在
                语法、发音和最常见的词汇上有所不同。每个人都意识到为什么
                一种新的通用语言是可取的：人们可以拒绝支付昂贵的翻译费用。为了
                实现这一点，需要统一的语法、发音和更常见的
                词汇。如果几种语言融合，结果语言的语法将比
                个别语言的语法更简单和规则。
              `),
		}, {
			"title": "标签页3",
			"content": template.HTML(`
                Lorem Ipsum 只是印刷和排版行业的虚拟文本。
                自1500年代以来，Lorem Ipsum 一直是行业的标准虚拟文本，
                当时一位不知名的印刷商拿了一个字样盘并将其打乱以制作一个字样样本书。
                它不仅存活了五个世纪，还跨越了电子排版的飞跃，
                基本上保持不变。它在1960年代随着包含Lorem Ipsum段落的Letraset
                表的发布而流行起来，最近又随着桌面出版软件
                如Aldus PageMaker（包括Lorem Ipsum版本）而流行。
                <br><br>
                <b>技术说明：</b>
                <p>Lorem Ipsum 是一种虚拟文本，用于在印刷和排版行业中填充空间。
                它不包含任何实际含义，只是用来展示排版效果。</p>
              `),
		},
	}).GetContent()

	// 创建弹窗触发按钮
	// data-toggle="modal": 设置为模态框触发器
	// data-target: 指定要打开的模态框ID
	buttonTest := `<button type="button" class="btn btn-primary" data-toggle="modal" data-target="#exampleModal" data-whatever="@mdo">为 @mdo 打开弹窗</button>`

	// 创建弹窗表单内容
	// 包含收件人和消息输入框
	popupForm := `<form>
          <div class="form-group">
            <label for="recipient-name" class="col-form-label">收件人:</label>
            <input type="text" class="form-control" id="recipient-name">
          </div>
          <div class="form-group">
            <label for="message-text" class="col-form-label">消息:</label>
            <textarea class="form-control" id="message-text"></textarea>
          </div>
        </form>`

	// 创建弹窗组件
	// SetID: 设置弹窗ID，用于在按钮中引用
	// SetFooter: 设置底部按钮文本
	// SetTitle: 设置弹窗标题
	// SetBody: 设置弹窗内容（表单）
	popup := components.Popup().SetID("exampleModal").
		SetFooter("保存更改").
		SetTitle("这是一个弹窗").
		SetBody(template.HTML(popupForm)).
		GetContent()

	// 创建第五列，包含标签页和弹窗触发按钮
	col5 := colComp.SetSize(types.SizeMD(8)).SetContent(tabs + template.HTML(buttonTest)).GetContent()

	// 创建第六列，包含饼图盒子和弹窗
	col6 := colComp.SetSize(types.SizeMD(4)).SetContent(boxDanger + popup).GetContent()

	// 创建第四行，包含标签页/按钮列和饼图/弹窗列
	row4 := components.Row().SetContent(col5 + col6).GetContent()

	// 返回页面面板
	// Content: 页面内容，按顺序包含row3、row2、row5、row4
	// Title: 页面标题
	// Description: 页面描述
	return types.Panel{
		Content:     row3 + row2 + row5 + row4,
		Title:       "仪表板",
		Description: "仪表板示例",
	}, nil
}
