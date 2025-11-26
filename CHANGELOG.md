# NSON-Go 扩展类型更新说明

## 概述

为了支持 beacon 项目中 Matter 协议的数据类型需求，nson-go 库已添加以下新的数值类型：

- `U8` (uint8) - Tag: 0x17
- `U16` (uint16) - Tag: 0x18
- `I8` (int8) - Tag: 0x19
- `I16` (int16) - Tag: 0x1A

这些类型补充了原有的 U32、U64、I32、I64，形成了完整的整数类型体系。

## 变更内容

### 1. 新增文件

- `extended_types_test.go` - 新类型的单元测试
- `examples/extended_types_demo.go` - 使用示例演示
- `BEACON_INTEGRATION.md` - 与 beacon 项目集成指南

### 2. 修改文件

#### `spec.go`
- 添加 4 个新的 Tag 常量：`TAG_U8`, `TAG_U16`, `TAG_I8`, `TAG_I16`

#### `type.go`
- 添加 4 个新的类型定义：`U8`, `U16`, `I8`, `I16`

#### `num.go`
- 为每个新类型实现了 `Tag()`, `String()`, `Encode*`, `Decode*` 方法

#### `util.go`
- 添加 `writeInt16()` 和 `readInt16()` 函数
- 添加 `writeUint16()` 和 `readUint16()` 函数

#### `value.go`
- 在 `EncodeValue()` 中添加新类型的编码分支
- 在 `decodeValueWithTag()` 中添加新类型的解码分支

## 类型对应关系

| NSON 类型 | Go 类型 | 大小 | Matter DataType 对应 |
|----------|---------|------|---------------------|
| U8       | uint8   | 1字节 | UInt8, Enum8, Bitmap8 |
| U16      | uint16  | 2字节 | UInt16, Enum16, Bitmap16 |
| I8       | int8    | 1字节 | Int8 |
| I16      | int16   | 2字节 | Int16 |

## 使用示例

### 基本使用

```go
import "github.com/danclive/nson-go"

// 创建值
val := nson.U8(255)
fmt.Println(val) // U8(255)

// 编码
var buf bytes.Buffer
nson.EncodeValue(&buf, val)

// 解码
decoded, _ := nson.DecodeValue(&buf)
result := decoded.(nson.U8)
```

### Matter 设备属性示例

```go
// 温度传感器数据（0.01°C 精度）
temperature := nson.I16(2350) // 23.50°C

// 亮度级别（0-254）
brightness := nson.U8(200)

// 厂商 ID
vendorId := nson.U16(0x1234)

// 设备状态位图
statusFlags := nson.U8(0b00000011)
```

### 复杂结构示例

```go
// Matter 设备完整状态
deviceState := nson.Map{
    "vendorId":    nson.U16(0x1234),
    "productId":   nson.U16(0x5678),
    "onOff":       nson.Bool(true),
    "brightness":  nson.U8(200),
    "temperature": nson.I16(2350),
    "humidity":    nson.U8(65),
}
```

## 测试覆盖

所有新类型都通过了以下测试：

1. ✅ 基本编码/解码测试
2. ✅ Map 中使用测试
3. ✅ Array 中使用测试
4. ✅ 边界值测试（min/max）
5. ✅ 与现有类型混合使用测试

测试结果：
```
=== RUN   TestU8
--- PASS: TestU8 (0.00s)
=== RUN   TestU16
--- PASS: TestU16 (0.00s)
=== RUN   TestI8
--- PASS: TestI8 (0.00s)
=== RUN   TestI16
--- PASS: TestI16 (0.00s)
=== RUN   TestMapWithExtendedTypes
--- PASS: TestMapWithExtendedTypes (0.00s)
=== RUN   TestArrayWithExtendedTypes
--- PASS: TestArrayWithExtendedTypes (0.00s)
=== RUN   TestBoundaryValues
--- PASS: TestBoundaryValues (0.00s)
PASS
ok      github.com/danclive/nson-go     0.002s
```

## 向后兼容性

- ✅ 所有现有代码保持不变
- ✅ 现有测试全部通过
- ✅ 新增 Tag 值不与现有 Tag 冲突
- ✅ 二进制格式保持一致（小端序）

## 性能特性

- **紧凑编码**: U8/I8 只需 2 字节（1字节tag + 1字节数据）
- **高效序列化**: 使用 Go 标准库 `encoding/binary`
- **零分配**: 基本类型操作不涉及堆分配
- **小端序**: 适合现代处理器架构

## 与 Beacon 集成

详细的集成指南请参考：`BEACON_INTEGRATION.md`

主要内容包括：
- AttributeValue 到 NSON Value 的映射表
- 类型转换辅助函数
- 实际使用示例
- 迁移策略建议

## 未来扩展

如果需要，可以继续添加：
- `F16` (float16) - 半精度浮点数
- `Decimal` - 定点小数
- 自定义 Tag 范围（0x80-0xFF 保留用于用户扩展）

## 总结

nson-go 现在完全支持 beacon 项目所需的所有基本数据类型，可以作为 `AttributeValue` 的替代方案，提供：

1. ✅ 完整的类型系统支持
2. ✅ 内置的序列化/反序列化
3. ✅ 类型安全的 Go 实现
4. ✅ 高性能的二进制编码
5. ✅ 良好的测试覆盖
6. ✅ 清晰的文档和示例

所有代码已通过测试，可以直接用于生产环境。
