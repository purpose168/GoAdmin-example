// Package tables 提供数据库表格模型定义
// 本文件实现用户档案（profile）表格的模型配置，演示多种字段类型和显示方式
package tables

import (
	"path/filepath"
	"strings"

	"github.com/purpose168/GoAdmin/context"
	"github.com/purpose168/GoAdmin/modules/db"
	"github.com/purpose168/GoAdmin/plugins/admin/modules/table"
	"github.com/purpose168/GoAdmin/template/types"
	"github.com/purpose168/GoAdmin/template/types/form"
)

// GetProfileTable 获取用户档案表格模型
// 该函数创建并返回一个配置完整的用户档案表格模型，用于管理后台的用户档案信息展示和编辑
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
//   - 配置信息展示字段（列表视图），包括多种字段类型和显示方式
//   - 配置表单编辑字段（编辑视图）
//   - 演示各种字段类型的使用方法
//
// 核心特性:
//   - 字段筛选：通过 FieldFilterable 支持字段筛选
//   - 字段复制：通过 FieldCopyable 支持一键复制字段内容
//   - 布尔字段：通过 FieldBool 显示布尔值
//   - 轮播图：通过 FieldCarousel 显示图片轮播
//   - 状态点：通过 FieldDot 显示带颜色标记的状态
//   - 进度条：通过 FieldProgressBar 显示进度条
//   - 文件下载：通过 FieldDownLoadable 支持文件下载
//   - 文件大小：通过 FieldFileSize 显示文件大小
func GetProfileTable(ctx *context.Context) table.Table {

	// 创建默认表格模型
	// NewDefaultTable 创建一个使用默认配置的表格实例
	// DefaultConfigWithDriver 指定数据库驱动类型为 "sqlite"
	profile := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("sqlite"))

	// 获取信息展示配置对象
	// GetInfo 返回表格的信息展示配置器，用于配置列表视图的字段
	// HideFilterArea: 隐藏筛选区域（不显示默认的筛选表单）
	// SetFilterFormLayout: 设置筛选表单的布局为筛选布局
	info := profile.GetInfo().HideFilterArea().SetFilterFormLayout(form.LayoutFilter)

	// 添加 ID 字段（支持筛选）
	// 参数说明:
	//   - "ID": 字段显示名称
	//   - "id": 数据库字段名
	//   - db.Int: 字段数据类型（整数）
	// FieldFilterable: 设置该字段可筛选（在筛选表单中显示）
	info.AddField("编号", "id", db.Int).FieldFilterable()

	// 添加 UUID 字段（支持复制）
	// 参数说明:
	//   - "UUID": 字段显示名称
	//   - "uuid": 数据库字段名
	//   - db.Varchar: 字段数据类型（可变长字符串）
	// FieldCopyable: 设置该字段可复制（显示复制按钮，点击后复制到剪贴板）
	info.AddField("UUID", "uuid", db.Varchar).FieldCopyable()

	// 添加 Pass 字段（布尔字段）
	// 参数说明:
	//   - "Pass": 字段显示名称
	//   - "pass": 数据库字段名
	//   - db.Tinyint: 字段数据类型（微整数，通常用于存储布尔值）
	// FieldBool: 将字段显示为布尔值
	//   - "1": true 时的显示文本
	//   - "0": false 时的显示文本
	info.AddField("通过", "pass", db.Tinyint).FieldBool("1", "0")

	// 添加 Photos 字段（轮播图）
	// 参数说明:
	//   - "Photos": 字段显示名称
	//   - "photos": 数据库字段名
	//   - db.Varchar: 字段数据类型（可变长字符串，存储逗号分隔的图片 URL）
	// FieldCarousel: 将字段显示为图片轮播
	//   - 回调函数: 将逗号分隔的字符串转换为字符串切片
	//     strings.Split: 按逗号分割字符串
	//   - 150: 轮播图宽度（像素）
	//   - 100: 轮播图高度（像素）
	info.AddField("照片", "photos", db.Varchar).FieldCarousel(func(value string) []string {
		// 将逗号分隔的图片 URL 切片转换为字符串数组
		return strings.Split(value, ",")
	}, 150, 100)

	// 添加 Finish State 字段（带状态点的自定义显示）
	// 参数说明:
	//   - "Finish State": 字段显示名称
	//   - "finish_state": 数据库字段名
	//   - db.Tinyint: 字段数据类型（微整数）
	// FieldDisplay: 使用自定义函数显示字段内容
	//   根据字段值显示不同的步骤名称
	// FieldDot: 为字段添加带颜色标记的状态点
	//   - map[string]types.FieldDotColor: 状态值与颜色的映射
	//     types.FieldDotColorDanger: 红色（危险）
	//     types.FieldDotColorInfo: 蓝色（信息）
	//     types.FieldDotColorPrimary: 主色调
	//   - types.FieldDotColorDanger: 默认颜色（当状态不在映射中时使用）
	info.AddField("完成状态", "finish_state", db.Tinyint).
		FieldDisplay(func(value types.FieldModel) interface{} {
			// 根据字段值返回对应的步骤名称
			if value.Value == "0" {
				return "步骤1"
			}
			if value.Value == "1" {
				return "步骤2"
			}
			if value.Value == "2" {
				return "步骤3"
			}
			// 未知状态
			return "未知"
		}).
		FieldDot(map[string]types.FieldDotColor{
			"步骤1": types.FieldDotColorDanger,
			"步骤2": types.FieldDotColorInfo,
			"步骤3": types.FieldDotColorPrimary,
		}, types.FieldDotColorDanger)

	// 添加 Progress 字段（进度条）
	// 参数说明:
	//   - "Progress": 字段显示名称
	//   - "finish_progress": 数据库字段名
	//   - db.Int: 字段数据类型（整数，表示百分比 0-100）
	// FieldProgressBar: 将字段显示为进度条
	info.AddField("进度", "finish_progress", db.Int).FieldProgressBar()

	// 添加 Resume 字段（文件下载）
	// 参数说明:
	//   - "Resume": 字段显示名称
	//   - "resume": 数据库字段名
	//   - db.Varchar: 字段数据类型（可变长字符串，存储文件路径）
	// FieldDisplay: 使用自定义函数显示字段内容
	//   filepath.Base: 从完整路径中提取文件名
	// FieldDownLoadable: 设置字段可下载
	//   - "http://yinyanghu.github.io/files/": 文件下载的基础 URL
	//     完整下载路径 = 基础 URL + 字段值
	info.AddField("简历", "resume", db.Varchar).
		FieldDisplay(func(value types.FieldModel) interface{} {
			// 从完整路径中提取文件名
			// 例如: "/path/to/resume.pdf" -> "resume.pdf"
			return filepath.Base(value.Value)
		}).
		FieldDownLoadable("http://yinyanghu.github.io/files/")

	// 添加 FileSize 字段（文件大小）
	// 参数说明:
	//   - "FileSize": 字段显示名称
	//   - "resume_size": 数据库字段名
	//   - db.Int: 字段数据类型（整数，表示字节数）
	// FieldFileSize: 将字节数转换为人类可读的文件大小格式
	//   例如: 1024 -> "1 KB", 1048576 -> "1 MB"
	info.AddField("文件大小", "resume_size", db.Int).FieldFileSize()

	// 设置表格基本信息
	// SetTable: 指定数据库表名
	// SetTitle: 设置表格标题（显示在页面头部）
	// SetDescription: 设置表格描述
	info.SetTable("profile").SetTitle("用户档案").SetDescription("用户档案")

	// 获取表单配置对象
	// GetForm 返回表格的表单配置器，用于配置编辑/添加视图的字段
	formList := profile.GetForm()

	// 添加 UUID 字段到表单
	// 参数说明:
	//   - "UUID": 字段显示名称
	//   - "uuid": 数据库字段名
	//   - db.Varchar: 字段数据类型
	//   - form.Text: 表单字段类型（文本输入框）
	formList.AddField("UUID", "uuid", db.Varchar, form.Text)

	// 添加 Photos 字段到表单
	// 参数说明:
	//   - "Photos": 字段显示名称
	//   - "photos": 数据库字段名
	//   - db.Varchar: 字段数据类型
	//   - form.Text: 表单字段类型（文本输入框）
	//   注意: 这里使用文本框输入逗号分隔的图片 URL
	formList.AddField("照片", "photos", db.Varchar, form.Text)

	// 添加 Resume 字段到表单
	// 参数说明:
	//   - "Resume": 字段显示名称
	//   - "resume": 数据库字段名
	//   - db.Varchar: 字段数据类型
	//   - form.Text: 表单字段类型（文本输入框）
	formList.AddField("简历", "resume", db.Varchar, form.Text)

	// 添加 FileSize 字段到表单
	// 参数说明:
	//   - "FileSize": 字段显示名称
	//   - "resume_size": 数据库字段名
	//   - db.Int: 字段数据类型
	//   - form.Number: 表单字段类型（数字输入框）
	formList.AddField("文件大小", "resume_size", db.Int, form.Number)

	// 添加 Finish State 字段到表单
	// 参数说明:
	//   - "Finish State": 字段显示名称
	//   - "finish_state": 数据库字段名
	//   - db.Tinyint: 字段数据类型
	//   - form.Number: 表单字段类型（数字输入框）
	formList.AddField("完成状态", "finish_state", db.Tinyint, form.Number)

	// 添加 Progress 字段到表单
	// 参数说明:
	//   - "Progress": 字段显示名称
	//   - "finish_progress": 数据库字段名
	//   - db.Int: 字段数据类型
	//   - form.Number: 表单字段类型（数字输入框）
	formList.AddField("进度", "finish_progress", db.Int, form.Number)

	// 添加 Pass 字段到表单
	// 参数说明:
	//   - "Pass": 字段显示名称
	//   - "pass": 数据库字段名
	//   - db.Tinyint: 字段数据类型
	//   - form.Number: 表单字段类型（数字输入框）
	formList.AddField("通过", "pass", db.Tinyint, form.Number)

	// 设置表单基本信息
	// SetTable: 指定数据库表名
	// SetTitle: 设置表单标题
	// SetDescription: 设置表单描述
	formList.SetTable("profile").SetTitle("用户档案").SetDescription("用户档案")

	// 返回配置好的表格模型
	return profile
}
