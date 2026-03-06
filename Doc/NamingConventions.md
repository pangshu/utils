## Go 语言命名规范 / Go Naming Conventions

本文档整理了Go语言开发中的通用命名规则方式。

---

### 一、基本命名规则

#### 1.1 大小写与可见性

- **首字母大写**：公有，包外可访问，如：`ToString`、`ParseInt`
- **首字母小写**：私有，包内可访问，如：`formatString`、`parseConfig`

#### 1.2 命名风格

- 使用**驼峰命名法**：如：`userName`、`userAge`
- 缩写词保持全大写或全小写：`HTTPCliect`、`url`、`xmlParser`、`ID`
- 避免冗余前缀：如：`utils.Convert.ToString`，而不是`utils.Convert.ConvertToString`

---

### 二、函数命名规则

#### 2。1 动词开头

函数名应以动词或动词短语开头，清晰表达行为：
|前缀|含义|示例|
|------|------|------|
|`Get`|获取值|`GetUserNameByID`|
|`Set`|设置值|`SetUserNameByID`|
|`Add`|添加值|`AddUser`|
|`Del`|删除值|`DelUser`|
|`Is`|判断值|`IsAdmin`|
|`Has`|判断值|`HasPermission`|
|`To`|类型转换|`ToString`、`ToInt`、`ToFloat64`|
|`Parse`|解析值|`ParseInt`|
|`Format`|格式化值|`FormatString`|
|`Check`|检查值|`CheckUser`|
|`List`|列出值|`ListUsers`|
|`Count`|计数值|`CountUsers`|
|`Clear`|清空值|`ClearUsers`|
|`Load`|加载值|`LoadUsers`|
|`Save`|保存值|`SaveUser`|
|`Send`|发送值|`SendEmail`|
|`Receive`|接收值|`ReceiveEmail`|
|`Open`|打开值|`OpenFile`|
|`Close`|关闭值|`CloseFile`|
|`Start`|启动值|`StartServer`|
|`Stop`|停止值|`StopServer`|
|`Run`|运行值|`RunProgram`|
|`Exit`|退出值|`ExitProgram`|
|`Init`|初始化值|`InitConfig`|
|`Must`|不返回error，失败时panic|`MustUsers`|
|`Split`|分割值|`SplitString`|
|`Join`|连接值|`JoinStrings`|
|`Sort`|排序值|`SortUsers`|
|`Filter`|过滤值|`FilterUsers`|
|`Map`|映射值|`MapUsers`|
|`With`|返回带有某配置的新实例|`WithUsers`|
|`Find`|查找值|`FindUser`|
|`FindAll`|查找所有值|`FindAllUsers`|
|`FindOne`|查找一个值|`FindOneUser`|
|`FindFirst`|查找第一个值|`FindFirstUser`|
|`FindLast`|查找最后一个值|`FindLastUser`|
|`FindIndex`|查找索引值|`FindIndexUser`|
|`FindLastIndex`|查找最后一个索引值|`FindLastIndexUser`|
|`FindAllIndex`|查找所有索引值|`FindAllIndexUser`|
|`Remove`|移除值（硬删）|`RemoveUser`|
|`Delete`|删除值（软删，主要用与数据库操作）|`DeleteUser`|
|`Update`|更新值（主要用与数据库操作）|`UpdateUser`|
|`Insert`|插入值（主要用与数据库操作）|`InsertUser`|
|`Replace`|替换值（主要用与数据库操作）|`ReplaceUser`|

#### 2.2 返回值约定
| 模式   | 签名风格                                           | 说明          |
|------|------------------------------------------------|-------------|
| 可能失败 | `func ToInt[T any](v T) (int, error)`          | 返回值 + error |
| 必定成功 | `func MustInt[T any](v T) int`                 | 失败时 panic   |
| 带默认值 | `func ToIntOrDefault[T any](v T, def int) int` | 失败返回默认值     |
| 布尔判断 | `func IsInt[T any](v T) bool`                  | 返回bool      |

#### 2.3 示例
```go
// ToIntOrDefault
func ToIntOrDefault[T any](v T, def int) int {
	if v == nil {
		return def
	}
	return int(v)
}

// ToInt
func ToInt[T any](v T) (int, error) {
	if v == nil {
		return 0, errors.New("value is nil")
	}
	return int(v), nil
}

// MustInt
func MustInt[T any](v T) int {
	if v == nil {
		panic("value is nil")
	}
	return int(v)
}
```
**优点**：语义清晰，功能明确，不会混淆
----

### 三、带默认值功能的命名规范

带默认值的函数用于在转换/操作失败时返回一个安全的默认值，避免调用方处理error。

#### 3。1 常见命名模式
###### 模式一：`XxxOrDefault` 后缀
```go
func ToIntOrDefault[T any](v T, def int) int {}
func ToStringOrDefault[T any](v T, def string) string {}
func ToBoolOrDefault[T any](v T, def bool) bool {}
```
**优点**：语义清晰，功能明确，不会混淆

###### 模式二：`DefaultXxx` 前缀
```go
func DefaultInt[T any](v T) int {}
func DefaultString[T any](v T) string {}
func DefaultBool[T any](v T) bool {}
```
**优点**：简洁，强调“提供默认值”，常用于参数/配置

###### 模式三：`MustXxx` 前缀
```go
func MustInt[T any](v T) int {}
func MustString[T any](v T) string {}
func MustBool[T any](v T) bool {}
```
**优点**：失败panic，强调“必须成功”，常用初始化、配置加载等确认不会失败的场景

###### 模式四：`WithXxx` 前缀
```go
func WithTimeout[T any](v T) NowTime {}
func WithRetry[T any](v T) Retry {}
```
**优点**：返回带有某配置的新实例

#### 3.2 命名模式对比
| 模式             | 示例                     | 语义侧重          | 适用场景    |
|----------------|------------------------|---------------|---------|
| `XxxOrDefault` | `ToIntOrDefault(v, 0)` | 转换/操作失败时返回默认值 | 类型转换    |
| `DefaultXxx`   | `DefaultXxx(v, 0)`     | 提供默认值         | 配置/参数处理 |
| `MustXxx`      | `MustXxx(v)`        | 必须成功，否则panic  | 初始化、测试  |
| `WithXxx`      | `WithTimeout(v)`      | 返回带有某配置的新实例 | 配置新实例   |

---

#### 四、类型检测函数命名
类型检测函数以`Is`开头，返回`bool`类型，用于判断类型是否为某个类型或是否实现了某个接口
```go
func IsString[T any](v T) bool {}
```
---
#### 五、编码与解析函数命名
编码与解析函数以`XxxToYyy`命名，表示将Xxx类型转换为Yyy类型
| 方向             | 前缀                     | 示例|
|----------------|------------------------|----|
| 编码/序列化   | `To`、`Encode`、`Marshal`     | `ToInt`、`ToString`、`ToJson`、`EncodeByte`|
| 解码/反序列化    | `From`、`Decode`、`Unmarshal`    | `FromJson`、`FromBase64`、`DecodeByte`|
| 格式转换 | `XxxToYyy`         | `GbkToUtf8`、`HexToDecimal`、`ColorRGBToHSL`|
---

#### 六、变量与常量命名

###### 6.1 变量
```go
// 好的命名
var userCount int
var isEnabled bool
var maxConnections int

// 不好的命名
var i int
var b bool
var max int
```

###### 6.2 常量
```go
const MaxBufferSize = 1024
const DefaultPort = 8080
const Success = "OK"
```
#### 6.3 循环变量例外
在循环中，变量名可以简短，因为它们的作用域有限，并且通常只在一个循环中使用。故单字母可拉受。
```go
for i := 0; i < 10; i++ {
	fmt.Println(i)
}
for k,v := range m {
	fmt.Println(k,v)
}
```

---

#### 七、接口命名
- 单方法接口以`er`结尾，表示该接口是一个行为或能力，例如`Reader`、`Writer`、`Logger`、`Converter`等。
- 多方法接口使用描述性名词，例如`ReadWriter`、`LoggerWriter`、`FileSystem`等。
```go
type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}
```

---

#### 八、包命名
- 包名应尽量简短，使用小写字母，并且不使用下划线或驼峰命名法。例如，`net`、`http`、`json`、`time`等。
- 包名应与包的用途或功能相关，例如`math`、`crypto`、`os`、`io`等。
- 包名应避免使用缩写，除非缩写是广泛接受的，例如`json`、`xml`、`http`等。
- 包名应避免使用与标准库包名相同的名称，例如`fmt`、`log`、`os`、`io`等。
- 包名应避免使用与第三方包名相同的名称，例如`github.com/gin-gonic/gin`、`github.com/go-sql-driver/mysql`等。
- 包名应避免使用与项目名相同的名称，例如`github.com/myproject/myproject`。

---

#### 九、函数命名
- 函数名应尽量简短，使用小写字母，并且不使用下划线或驼峰命名法。例如，`Println`、`ReadFile`、`Write`、`File`、`ParseJSON`等。
- 函数名应与函数的用途或功能相关，例如`Println`、`ReadFile`、`Write`、`File`、`ParseJSON`等。
- 函数名应避免使用缩写，除非缩写是广泛接受的，例如`Println`、`ReadFile`、`Write`、`File`、`ParseJSON`等。
- 函数名应避免使用与标准库函数名相同的名称，例如`Println`、`ReadFile`、`Write`、`File`、`ParseJSON`等。
- 函数名应避免使用与第三方库函数名相同的名称，例如`github.com/gin-gonic/gin`、`github.com/go-sql-driver/mysql`等。

---

#### 十、文件命名
- 文件名应尽量简短，使用小写字母，并且不使用下划线或驼峰命名法。例如，`main.go`、`server.go`、`client.go`、`config.go`、`logger.go`等。
- 文件名应与文件的内容或功能相关，例如`main.go`、`server.go`、`client.go`、`config.go`、`logger.go`等。
- 文件名应避免使用缩写，除非缩写是广泛接受的，例如`main.go`、`server.go`、`client.go`、`config.go`、`logger.go`等。
- 文件名应避免使用与标准库文件名相同的名称，例如`main.go`、`server.go`、`client.go`、`config.go`、`logger.go`等。
- 文件名应避免使用与第三方库文件名相同的名称，例如`github.com/gin-gonic/gin`、`github.com/go-sql-driver/mysql`等。

---
#### 十一、错误命名
- 错误名应尽量简短，使用小写字母，并且不使用下划线或驼峰命名法。例如，`ErrNotFound`、`ErrInvalidArgument`、`ErrInternal`、`ErrTimeout`、`ErrPermissionDenied`等。
- 错误名应与错误的含义或原因相关，例如`ErrNotFound`、`ErrInvalidArgument`、`ErrInternal`、`ErrTimeout`、`ErrPermissionDenied`等。
- 错误名应避免使用缩写，除非缩写是广泛接受的，例如`ErrNotFound`、`ErrInvalidArgument`、`ErrInternal`、`ErrTimeout`、`ErrPermissionDenied`等。
- 错误名应避免使用与标准库错误名相同的名称，例如`ErrNotFound`、`ErrInvalidArgument`、`ErrInternal`、`ErrTimeout`、`ErrPermissionDenied`等。
```go
// 错误变量以Err开头
var ErrNotFound = errors.New("not found")
var ErrInvalidArgument = errors.New("invalid argument")
var ErrInternal = errors.New("internal error")
var ErrTimeout = errors.New("timeout")

// 自定义错误类型以Error结尾
type ValidationError struct {...}
type PermissionDeniedError struct {...}
```
---
#### 十二、总结
-- 命名是编程中非常重要的一部分，它可以帮助我们更好地理解代码，提高代码的可读性和可维护性。在Go语言中，命名应该遵循一定的规范，以便与其他Go程序员保持一致。
| 类别 | 规则 | 示例 |
| --- | --- | --- |
| 公用函数 | 大写开头 + 动词 | `ParseConfig` |
| 私有函数 | 小写开头 + 动词 | `parseConfig` |
| 类型转换 | `To` + 目标类型 | `ToInt`、`ToString` |
| 反向构造 | `From` + 来源类型 | `FromJson`、`FormBase64` |
| 类型检测 | `Is` + 类型 | `IsInt`、`IsBool` |
| 带默认值 | `To` + 类型 + `OrDefault` | `ToIntOrDefault`、`ToStringOrDefault` |
| 必须成功 | `Must` + 功能名 | `MustParse` |
| 布尔判断 | `Is`/`Has`/`Can` + 功能名 | `IsValid`、`HasKey` |
| 可选参数 | `With` + 参数名 | `WithTimeout`、`WithLogger` |
| 错误变量 | `Err` + 描述 | `ErrNotFound` |
| 自定义错误类型 | 描述 + `Error` | `ValidationError` |
| 接口 | 动词 + `er` | `Reader`、`Converter` |
---