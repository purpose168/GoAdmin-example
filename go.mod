// 模块(Module)声明：定义当前 Go 项目的模块路径
// 该路径是项目在 Go 模块系统中的唯一标识符
module github.com/purpose168/GoAdmin-example

// Go 版本要求：指定编译此项目所需的 Go 语言最低版本
// 版本 1.24.2 表示需要使用 Go 1.24.2 或更高版本进行编译
go 1.24.2

// 替换指令(Replace Directive)：用于替换依赖包的版本或路径
// 这在开发过程中非常有用，可以指定使用本地路径或特定版本的依赖
replace (
	// 将 GoAdmin 框架替换为特定版本：v0.0.0-20260104141321-fcc00eb84719
	// 这是一个伪版本号，基于提交哈希值生成，用于精确指定代码版本
	github.com/purpose168/GoAdmin => github.com/purpose168/GoAdmin v0.0.0-20260104141321-fcc00eb84719
	// 将 GoAdmin 主题包替换为特定版本：v0.0.0-20260104133356-8e29cafd3a6d
	// 同样使用伪版本号，确保使用指定提交的代码
	github.com/purpose168/GoAdmin-themes => github.com/purpose168/GoAdmin-themes v0.0.0-20260104133356-8e29cafd3a6d
)

// 直接依赖声明(Require Direct Dependencies)：列出项目直接使用的所有外部依赖包
// 这些包在项目代码中被显式导入和使用
require (
	// HTTP 测试库：用于编写 HTTP 服务的集成测试和端到端测试
	// 提供了类似断言的 API，方便测试 HTTP 请求和响应
	github.com/gavv/httpexpect v2.0.0+incompatible
	// Gin Web 框架：高性能的 HTTP Web 框架，类似于 Martini 但性能更好
	// 提供了路由、中间件、JSON 验证等功能，是 Go 社区最流行的 Web 框架之一
	github.com/gin-gonic/gin v1.11.0
	// GORM ORM 库：Go 语言的 Object-Relational Mapping (对象关系映射) 库
	// 提供了友好的 API 来操作数据库，支持 MySQL、PostgreSQL、SQLite 等多种数据库
	github.com/jinzhu/gorm v1.9.16
	// GoAdmin 核心框架：一个基于 Go 语言的后台管理系统框架
	// 提供了完整的后台管理功能，包括权限管理、菜单管理、数据表格等
	github.com/purpose168/GoAdmin v1.2.26
	// GoAdmin 主题包：为 GoAdmin 框架提供 UI 主题和样式
	// 包含了多种预设主题，可以快速美化后台管理界面
	github.com/purpose168/GoAdmin-themes v0.0.48
)

// 间接依赖声明(Require Indirect Dependencies)：列出项目间接使用的依赖包
// 这些包被直接依赖包所使用，通过 // indirect 标记
// Go 工具链会自动管理这些依赖，无需手动维护
require (
	// Edwards25519 椭圆曲线加密库：实现了 Ed25519 签名算法
	// 用于密码学操作，提供高性能的数字签名功能
	filippo.io/edwards25519 v1.1.0 // indirect
	// Excelize Excel 文件处理库：用于读写 Excel 文件
	// 支持 .xlsx 格式，可以创建、修改和导出 Excel 电子表格
	github.com/360EntSecGroup-Skylar/excelize v1.4.1 // indirect
	// GoAdmin HTML 模板库：提供 HTML 模板渲染功能
	// 用于生成动态 HTML 内容，支持模板继承和组件复用
	github.com/purpose168/GoAdmin-html v0.0.1 // indirect
	// 快速随机数生成器：提供高性能的伪随机数生成
	// 比标准库的 math/rand 更快，适合性能敏感的场景
	github.com/NebulousLabs/fastrand v0.0.0-20181203155948-6fb6489aac4e // indirect
	// 表单处理库：用于解析和编码 HTML 表单数据
	// 支持 multipart/form-data 和 application/x-www-form-urlencoded 格式
	github.com/ajg/form v1.5.1 // indirect
	// Brotli 压缩库：实现了 Brotli 压缩算法
	// Brotli 是一种高效的压缩格式，比 Gzip 压缩率更高
	github.com/andybalholm/brotli v1.1.0 // indirect
	// FastHTTP 路由器：为 FastHTTP 框架提供高性能的路由功能
	// FastHTTP 是一个高性能的 HTTP 服务器实现，比标准库的 net/http 更快
	github.com/buaazp/fasthttprouter v0.1.1 // indirect
	// Sonic JSON 库：字节跳动开源的高性能 JSON 序列化/反序列化库
	// 使用 SIMD 指令优化，性能远超标准库的 encoding/json
	github.com/bytedance/sonic v1.14.0 // indirect
	// Sonic 加载器：用于动态加载和编译 Sonic 的代码
	// 支持热重载和动态编译功能
	github.com/bytedance/sonic/loader v0.3.0 // indirect
	// Base64 编码库：提供 Base64 编码和解码功能
	// 用于数据的编码传输，常用于 HTTP Basic 认证等场景
	github.com/cloudwego/base64x v0.1.6 // indirect
	// Go-Spew 深度打印库：用于调试时打印复杂数据结构
	// 可以递归打印任意深度的数据结构，支持格式化输出
	github.com/davecgh/go-spew v1.1.1 // indirect
	// SQL Server 驱动：Microsoft SQL Server 数据库的 Go 语言驱动
	// 用于连接和操作 SQL Server 数据库，支持完整的 T-SQL 功能
	github.com/denisenkom/go-mssqldb v0.0.0-20200206145737-bbfc9a55622e // indirect
	// 结构体工具库：提供结构体相关的实用函数
	// 可以将结构体转换为 map、获取结构体字段信息等
	github.com/fatih/structs v1.1.0 // indirect
	// MIME 类型检测库：用于检测文件的 MIME 类型
	// 通过文件扩展名或文件内容来识别文件类型
	github.com/gabriel-vasile/mimetype v1.4.8 // indirect
	// Server-Sent Events (SSE) 库：用于实现服务器推送事件
	// SSE 是一种单向服务器推送技术，用于实时更新客户端
	github.com/gin-contrib/sse v1.1.0 // indirect
	// 本地化库：提供多语言支持的数据
	// 包含各种语言的日期、数字、货币等格式化规则
	github.com/go-playground/locales v0.14.1 // indirect
	// 通用翻译器：用于文本的国际化(i18n)和本地化(l10n)
	// 支持多种语言的翻译和格式化
	github.com/go-playground/universal-translator v0.18.1 // indirect
	// 数据验证库：提供强大的结构体验证功能
	// 支持自定义验证规则、错误消息国际化等
	github.com/go-playground/validator/v10 v10.27.0 // indirect
	// MySQL 驱动：MySQL 数据库的 Go 语言驱动
	// 用于连接和操作 MySQL 数据库，支持完整的 MySQL 协议
	github.com/go-sql-driver/mysql v1.8.1 // indirect
	// Go-JSON 库：另一个高性能的 JSON 序列化/反序列化库
	// 提供了流式处理、自定义编解码等功能
	github.com/goccy/go-json v0.10.5 // indirect
	// Go-YAML 库：用于解析和生成 YAML 格式的数据
	// YAML 是一种人类可读的数据序列化格式
	github.com/goccy/go-yaml v1.18.0 // indirect
	// Civil 日期时间库：提供"民用"日期时间类型
	// Civil 类型表示没有时区信息的日期和时间
	github.com/golang-sql/civil v0.0.0-20220223132316-b832511892a9 // indirect
	// Snappy 压缩库：Google 开源的快速压缩/解压缩库
	// Snappy 是一种快速的压缩算法，适合实时压缩场景
	github.com/golang/snappy v1.0.0 // indirect
	// URL 查询字符串库：用于构建和解析 URL 查询参数
	// 提供了类型安全的 API 来处理 URL 查询字符串
	github.com/google/go-querystring v1.1.0 // indirect
	// UUID 生成库：用于生成和解析 UUID (通用唯一标识符)
	// UUID 是 128 位的唯一标识符，广泛用于分布式系统
	github.com/google/uuid v1.6.0 // indirect
	// GopherJS 编译器：将 Go 代码编译为 JavaScript
	// 允许在浏览器中运行 Go 代码
	github.com/gopherjs/gopherjs v1.17.2 // indirect
	// WebSocket 库：实现 WebSocket 协议的 Go 库
	// WebSocket 是一种全双工通信协议，用于实时双向通信
	github.com/gorilla/websocket v1.5.1 // indirect
	// 字符串插值库：支持类似 Python 的字符串插值语法
	// 可以在字符串中嵌入变量和表达式
	github.com/imkira/go-interpol v1.1.0 // indirect
	// 词形变化库：用于英语单词的单复数转换
	// 常用于 ORM 框架中自动生成表名
	github.com/jinzhu/inflection v1.0.0 // indirect
	// JSON 迭代器库：高性能的 JSON 解析库
	// 使用迭代器模式，内存占用更小，适合处理大型 JSON 数据
	github.com/json-iterator/go v1.1.12 // indirect
	// Go Local Storage (GLS) 库：实现 Goroutine 本地存储
	// 类似于线程局部存储(TLS)，但用于 Goroutine
	github.com/jtolds/gls v4.20.0+incompatible // indirect
	// 压缩库：提供多种压缩算法的实现
	// 包括 zlib、flate、snappy 等压缩格式
	github.com/klauspost/compress v1.17.6 // indirect
	// CPUID 库：用于检测 CPU 特性和指令集支持
	// 可以查询 CPU 的型号、缓存大小、支持的指令集等信息
	github.com/klauspost/cpuid/v2 v2.3.0 // indirect
	// URN 解析库：用于解析和验证 URN (统一资源名称)
	// URN 是一种资源标识符，类似于 URL 但不包含位置信息
	github.com/leodido/go-urn v1.4.0 // indirect
	// PostgreSQL 驱动：PostgreSQL 数据库的 Go 语言驱动
	// 用于连接和操作 PostgreSQL 数据库，支持完整的 PostgreSQL 协议
	github.com/lib/pq v1.10.9 // indirect
	// 终端颜色库：用于在终端中输出彩色文本
	// 支持 ANSI 颜色代码，可以设置前景色、背景色等
	github.com/mattn/go-colorable v0.1.13 // indirect
	// 终端检测库：用于检测输出是否为终端
	// 判断标准输出/错误是否连接到终端设备
	github.com/mattn/go-isatty v0.0.20 // indirect
	// SQLite3 驱动：SQLite3 数据库的 Go 语言驱动
	// SQLite 是一个轻量级的嵌入式数据库，无需独立的服务器进程
	github.com/mattn/go-sqlite3 v2.0.3+incompatible // indirect
	// ANSI 颜色库：用于在终端中输出带 ANSI 转义序列的彩色文本
	// 提供了简单的 API 来生成彩色终端输出
	github.com/mgutz/ansi v0.0.0-20170206155736-9520e82c474b // indirect
	// 现代并发库：提供现代的并发编程工具
	// 包含一些并发原语和实用函数
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	// 现代反射库：提供增强的反射功能
	// 比标准库的 reflect 包性能更好，API 更友好
	github.com/modern-go/reflect2 v1.0.2 // indirect
	// 深拷贝库：用于对任意对象进行深拷贝
	// 可以递归复制对象的所有字段，包括嵌套的结构体和切片
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826 // indirect
	// HTTP2Curl 库：将 HTTP/2 请求转换为 curl 命令
	// 用于调试和记录 HTTP 请求，方便重现问题
	github.com/moul/http2curl v1.0.0 // indirect
	// TOML 解析库：用于解析和生成 TOML 格式的配置文件
	// TOML 是一种简洁的配置文件格式，易于阅读和编写
	github.com/pelletier/go-toml/v2 v2.2.4 // indirect
	// 差异比较库：用于比较两个数据结构的差异
	// 可以生成详细的差异报告，适合测试断言
	github.com/pmezard/go-difflib v1.0.0 // indirect
	// QPACK 库：HTTP/3 和 QUIC 协议的头部压缩库
	// QPACK 是 HPACK 的改进版本，用于压缩 HTTP 头部
	github.com/quic-go/qpack v0.5.1 // indirect
	// QUIC-Go 库：QUIC 协议的 Go 语言实现
	// QUIC 是一种基于 UDP 的传输协议，是 HTTP/3 的基础
	github.com/quic-go/quic-go v0.54.0 // indirect
	// Agouti 测试库：用于 Web 应用的端到端测试
	// 基于 WebDriver 协议，可以模拟浏览器操作
	github.com/sclevine/agouti v3.0.0+incompatible // indirect
	// 差异库：用于计算文本或数据结构的差异
	// 类似于 Unix 的 diff 命令，但提供了更丰富的 API
	github.com/sergi/go-diff v1.2.0 // indirect
	// 断言库：提供丰富的断言函数用于测试
	// 支持相等性检查、异常检查、集合比较等
	github.com/smarty/assertions v1.16.0 // indirect
	// 测试工具库：Go 语言最流行的测试辅助库
	// 提供了断言、模拟、套件管理等测试功能
	github.com/stretchr/testify v1.11.1 // indirect
	// LevelDB 库：Google 开源的键值对存储引擎
	// LevelDB 是一个快速的嵌入式数据库，支持有序存储
	github.com/syndtr/goleveldb v1.0.0 // indirect
	// Minify 库：用于压缩和优化 HTML、CSS、JavaScript 等文件
	// 通过移除空白、缩短变量名等方式减小文件体积
	github.com/tdewolff/minify/v2 v2.20.14 // indirect
	// Parse 库：用于解析 HTML、CSS、JavaScript 等文件
	// 提供了词法分析和语法分析功能
	github.com/tdewolff/parse/v2 v2.7.8 // indirect
	// Go-ASM 库：Go 语言的汇编器
	// 用于生成和操作机器码，支持多种架构
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	// Ugorji Codec 库：高性能的编解码库
	// 支持 JSON、CBOR、MessagePack 等多种格式
	github.com/ugorji/go/codec v1.3.0 // indirect
	// 字节缓冲池：用于重用字节缓冲区，减少内存分配
	// 可以显著提高高并发场景下的性能
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	// FastHTTP 库：高性能的 HTTP 服务器和客户端实现
	// 比标准库的 net/http 性能更高，内存占用更少
	github.com/valyala/fasthttp v1.52.0 // indirect
	// JSON Pointer 库：实现 JSON Pointer 规范
	// JSON Pointer 是一种引用 JSON 文档中特定值的方法
	github.com/xeipuuv/gojsonpointer v0.0.0-20180127040702-4e3ac2762d5f // indirect
	// JSON Reference 库：实现 JSON Reference 规范
	// JSON Reference 用于引用其他 JSON 文档或片段
	github.com/xeipuuv/gojsonreference v0.0.0-20180127040603-bd5ef7bd5415 // indirect
	// JSON Schema 库：实现 JSON Schema 规范
	// JSON Schema 用于验证 JSON 数据的结构和类型
	github.com/xeipuuv/gojsonschema v1.2.0 // indirect
	// JSONPath 库：实现 JSONPath 查询语言
	// JSONPath 类似于 XPath，用于查询 JSON 数据
	github.com/yalp/jsonpath v0.0.0-20180802001716-5cc68e5049a0 // indirect
	// JSON Diff 库：用于比较两个 JSON 对象的差异
	// 可以生成详细的差异报告，适合测试和调试
	github.com/yudai/gojsondiff v1.0.0 // indirect
	// LCS (最长公共子序列) 库：用于计算两个序列的最长公共子序列
	// 常用于文本比较、版本控制等场景
	github.com/yudai/golcs v0.0.0-20170316035057-ecda9a501e82 // indirect
	// Uber Mock 库：Uber 开源的模拟(mock)框架
	// 用于单元测试中模拟依赖对象
	go.uber.org/mock v0.5.0 // indirect
	// Uber Multierr 库：用于组合和处理多个错误
	// 可以将多个错误组合成一个错误，方便错误处理
	go.uber.org/multierr v1.11.0 // indirect
	// Uber Zap 库：Uber 开源的高性能结构化日志库
	// 提供了结构化日志、日志分级、性能优化等功能
	go.uber.org/zap v1.27.1 // indirect
	// Go Arch 库：提供 Go 程序和包的架构信息
	// 可以解析 Go 源代码，获取包结构、类型信息等
	golang.org/x/arch v0.20.0 // indirect
	// Go Crypto 库：Go 语言的密码学扩展库
	// 提供了各种加密算法、哈希函数、随机数生成等
	golang.org/x/crypto v0.46.0 // indirect
	// Go Mod 库：Go 模块系统的工具库
	// 提供了模块解析、版本查询等功能
	golang.org/x/mod v0.30.0 // indirect
	// Go Net 库：Go 网络扩展库
	// 提供了各种网络协议的实现，如 HTTP/2、WebSocket 等
	golang.org/x/net v0.47.0 // indirect
	// Go Sync 库：Go 并发扩展库
	// 提供了额外的并发原语，如 ErrGroup、SingleFlight 等
	golang.org/x/sync v0.19.0 // indirect
	// Go Sys 库：Go 系统调用扩展库
	// 提供了跨平台的系统调用接口，如文件系统、进程管理等
	golang.org/x/sys v0.39.0 // indirect
	// Go Text 库：Go 文本处理扩展库
	// 提供了字符编码、文本搜索、文本格式化等功能
	golang.org/x/text v0.32.0 // indirect
	// Go Tools 库：Go 工具链扩展库
	// 提供了代码分析、代码生成、静态检查等工具
	golang.org/x/tools v0.39.0 // indirect
	// Protocol Buffers 库：Google 的数据序列化格式
	// Protocol Buffers 是一种高效的二进制序列化格式
	google.golang.org/protobuf v1.36.9 // indirect
	// INI 配置文件库：用于解析和生成 INI 格式的配置文件
	// INI 是一种简单的配置文件格式，常用于 Windows 应用
	gopkg.in/ini.v1 v1.67.0 // indirect
	// Lumberjack 日志轮转库：用于日志文件的轮转和压缩
	// 可以自动切割、压缩和删除旧的日志文件
	gopkg.in/natefinch/lumberjack.v2 v2.2.1 // indirect
	// YAML v2 库：YAML 格式的解析库（版本 2）
	// YAML 是一种人类可读的数据序列化格式
	gopkg.in/yaml.v2 v2.4.0 // indirect
	// YAML v3 库：YAML 格式的解析库（版本 3）
	// 版本 3 相比版本 2 有一些改进和变化
	gopkg.in/yaml.v3 v3.0.1 // indirect
	// XORM Builder 库：XORM ORM 框架的 SQL 构建器
	// 用于构建 SQL 查询语句，支持链式调用
	xorm.io/builder v0.3.13 // indirect
	// XORM 库：Go 语言的 ORM 框架
	// 提供了简单的 API 来操作数据库，支持多种数据库引擎
	xorm.io/xorm v1.3.11 // indirect
)
