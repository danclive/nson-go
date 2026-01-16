# NSON-Go

高性能的 NSON (Network Serialization Object Notation) 序列化库。

## 特性

- 🚀 **高性能** - Marshal ~400ns/op, Unmarshal ~232ns/op
- 💪 **强类型系统** - 18 种精确类型，8 种整数类型保证精度不丢失
- 📦 **结构体序列化** - 类似 JSON，支持 struct tag、omitempty、嵌入结构体
- ⏱️ **时间类型** - time.Time ↔ Timestamp 自动转换（毫秒精度）
- 🆔 **ID 类型** - nson.Id 唯一标识符，区别于普通 []byte
- 🎯 **类型安全** - 精确类型匹配，编译时类型检查

## 快速开始

### 安装

```bash
go get github.com/danclive/nson-go
```

### 基本使用

```go
package main

import (
    "fmt"
    nson "github.com/danclive/nson-go"
)

type User struct {
    Name     string   `nson:"name"`
    Age      int32    `nson:"age"`
    Email    string   `nson:"email,omitempty"`
    Tags     []string `nson:"tags"`
}

func main() {
    // 序列化
    user := User{
        Name: "Alice",
        Age:  28,
        Tags: []string{"developer", "golang"},
    }

    m, err := nson.Marshal(user)
    if err != nil {
        panic(err)
    }

    // 反序列化
    var result User
    err = nson.Unmarshal(m, &result)
    if err != nil {
        panic(err)
    }

    fmt.Printf("%+v\n", result)
}
```

## 核心功能

### 强类型系统

精确保留整数类型，不丢失精度：

```go
type Data struct {
    Int8   int8   `nson:"i8"`     // -> I8 (1 byte)
    Int16  int16  `nson:"i16"`    // -> I16 (2 bytes)
    Int32  int32  `nson:"i32"`    // -> I32 (4 bytes)
    Int64  int64  `nson:"i64"`    // -> I64 (8 bytes)
    Uint8  uint8  `nson:"u8"`     // -> U8
    Uint16 uint16 `nson:"u16"`    // -> U16
    Uint32 uint32 `nson:"u32"`    // -> U32
    Uint64 uint64 `nson:"u64"`    // -> U64
}
```

### 时间类型

time.Time 自动转换为 Timestamp（毫秒精度）：

```go
type Event struct {
    Name      string    `nson:"name"`
    CreatedAt time.Time `nson:"created_at"`  // -> Timestamp
    UpdatedAt time.Time `nson:"updated_at"`  // -> Timestamp
}

event := Event{
    Name:      "系统启动",
    CreatedAt: time.Now(),
}

m, _ := nson.Marshal(event)
// m["created_at"] 是 nson.Timestamp 类型（毫秒时间戳）
```

### ID 类型

nson.Id 用于唯一标识符，区别于普通二进制数据：

```go
type Document struct {
    DocID   nson.Id `nson:"doc_id"`    // -> Id (专用类型)
    Content []byte  `nson:"content"`   // -> Binary (通用二进制)
}

doc := Document{
    DocID: nson.Id{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c},
}
```

### Struct Tag

```go
type User struct {
    Name     string `nson:"name"`              // 自定义字段名
    Email    string `nson:"email,omitempty"`   // 空值时省略
    Internal string `nson:"-"`                 // 跳过此字段
}
```

### 嵌入结构体

```go
type BaseEntity struct {
    ID        nson.Id   `nson:"id"`
    CreatedAt time.Time `nson:"created_at"`
}

type Article struct {
    BaseEntity                          // 字段会被展开
    Title   string `nson:"title"`
    Content string `nson:"content"`
}
```

### 类型别名支持

NSON 完全支持类型别名（type alias），无需额外转换：

```go
// 定义类型别名（常见于业务代码中）
type QoS uint8
type Priority uint8
type Status int32

const (
    QoS0 QoS = 0
    QoS1 QoS = 1
    QoS2 QoS = 2
)

type Message struct {
    Topic    string   `nson:"topic"`
    QoS      QoS      `nson:"qos"`        // 直接使用类型别名
    Priority Priority `nson:"priority"`    // 直接使用类型别名
    Status   Status   `nson:"status"`      // 直接使用类型别名
}

// 无需中间结构体转换，直接序列化
msg := Message{
    Topic:    "test/topic",
    QoS:      QoS1,      // 自动转换为 U8
    Priority: 1,         // 自动转换为 U8
    Status:   200,       // 自动转换为 I32
}

m, _ := nson.Marshal(&msg)
// QoS, Priority, Status 根据底层类型自动映射到对应的 NSON 类型

var decoded Message
nson.Unmarshal(m, &decoded)
// decoded.QoS == QoS1 (保持原始类型)
```

**优势：**
- ✅ 无需创建中间结构体
- ✅ 避免字段拷贝
- ✅ 保持类型安全
- ✅ 零运行时开销

## 类型映射

| Go 类型 | NSON 类型 | 大小 | 说明 |
|---------|-----------|------|------|
| `bool` | `Bool` | 1B | 布尔值 |
| `int8` | `I8` | 1B | 有符号整数 |
| `int16` | `I16` | 2B | 有符号整数 |
| `int32`, `int` | `I32` | 4B | 有符号整数 |
| `int64` | `I64` | 8B | 有符号整数 |
| `uint8` | `U8` | 1B | 无符号整数 |
| `uint16` | `U16` | 2B | 无符号整数 |
| `uint32`, `uint` | `U32` | 4B | 无符号整数 |
| `uint64` | `U64` | 8B | 无符号整数 |
| `float32` | `F32` | 4B | 浮点数 |
| `float64` | `F64` | 8B | 浮点数 |
| `string` | `String` | 变长 | UTF-8 字符串 |
| `[]byte` | `Binary` | 变长 | 二进制数据 |
| `time.Time` | `Timestamp` | 8B | 毫秒时间戳 |
| `nson.Id` ([]byte) | `Id` | 12B | 唯一标识符 |
| `[]T` | `Array` | 变长 | 数组 |
| `map[string]T` | `Map` | 变长 | 映射 |
| `struct` | `Map` | 变长 | 结构体 |
| `*T` | `Null` / `T` | - | nil 为 Null，否则为值类型 |
| `interface{}` | 对应类型 | - | 根据实际值类型映射 |

## 示例

```bash
# 运行示例
go run examples/marshal/marshal_demo.go      # 结构体序列化
go run examples/types/types_demo.go          # 类型系统演示
go run examples/codec/extended_types_demo.go # 编解码示例
```

## 测试

```bash
go test ./...                                # 运行所有测试
go test ./marshal_test/...                   # 运行 marshal 测试
go test -bench=. -benchmem ./marshal_test/   # 基准测试
```

## 性能

基准测试结果（Go 1.20+）：

```
BenchmarkMarshal-20        399.5 ns/op    572 B/op    11 allocs/op
BenchmarkUnmarshal-20      232.3 ns/op    152 B/op     3 allocs/op
BenchmarkRoundTrip-20      643.2 ns/op    724 B/op    14 allocs/op
```

运行基准测试：

```bash
go test -bench=. -benchmem ./marshal_test/
```

## 与其他格式对比

| 特性 | NSON | JSON | MessagePack | Protocol Buffers |
|------|------|------|-------------|------------------|
| 强类型整数 | ✅ 8 种 | ❌ 仅 number | ⚠️ 部分 | ✅ 是 |
| 浮点类型 | ✅ 2 种 | ❌ 仅 number | ⚠️ 部分 | ✅ 是 |
| 时间类型 | ✅ Timestamp | ❌ 字符串 | ❌ 扩展 | ⚠️ 需定义 |
| ID 类型 | ✅ 内置 | ❌ 字符串 | ❌ 无 | ⚠️ 需定义 |
| Schema | ❌ 无需 | ❌ 无需 | ❌ 无需 | ✅ 需要 |
| 动态类型 | ✅ 是 | ✅ 是 | ✅ 是 | ❌ 否 |
| 可读性 | ⚠️ 二进制 | ✅ 文本 | ❌ 二进制 | ❌ 二进制 |
| 易用性 | ✅ 高 | ✅ 高 | ✅ 高 | ⚠️ 中 |

## 许可证

MIT License - 参见 [LICENSE](./LICENSE) 文件
