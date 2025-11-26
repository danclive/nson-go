# NSON-Go 与 Beacon 项目集成指南

本文档说明如何使用 nson-go 库替代 beacon 项目中的 `AttributeValue` 类型。

## 新增类型

nson-go 已添加以下类型以支持 Matter 协议的数据类型需求：

| NSON 类型 | Go 类型 | Tag 值 | Matter DataType 映射 |
|----------|---------|--------|---------------------|
| `U8`     | `uint8` | 0x17   | `DataTypeUInt8`, `DataTypeEnum8`, `DataTypeBitmap8` |
| `U16`    | `uint16`| 0x18   | `DataTypeUInt16`, `DataTypeEnum16`, `DataTypeBitmap16` |
| `U32`    | `uint32`| 0x15   | `DataTypeUInt32`, `DataTypeBitmap32` |
| `U64`    | `uint64`| 0x16   | `DataTypeUInt64`, `DataTypeBitmap64` |
| `I8`     | `int8`  | 0x19   | `DataTypeInt8` |
| `I16`    | `int16` | 0x1A   | `DataTypeInt16` |
| `I32`    | `int32` | 0x13   | `DataTypeInt32` |
| `I64`    | `int64` | 0x14   | `DataTypeInt64` |
| `F32`    | `float32`| 0x11  | `DataTypeFloat` |
| `F64`    | `float64`| 0x12  | `DataTypeDouble` |
| `Bool`   | `bool`  | 0x01   | `DataTypeBoolean` |
| `String` | `string`| 0x21   | `DataTypeString` |
| `Binary` | `[]byte`| 0x22   | `DataTypeOctetStr` |
| `Array`  | `[]Value`| 0x31  | `DataTypeArray` |
| `Map`    | `map[string]Value`| 0x32 | `DataTypeStruct` |
| `Null`   | `struct{}`| 0x02 | `DataTypeNull` |

## Beacon AttributeValue 到 NSON Value 的映射

### 基本类型映射

```go
// Beacon                    -> NSON
NewBoolValue(true)          -> nson.Bool(true)
NewUInt8Value(255)          -> nson.U8(255)
NewUInt16Value(1234)        -> nson.U16(1234)
NewUInt32Value(12345)       -> nson.U32(12345)
NewUInt64Value(123456)      -> nson.U64(123456)
NewInt8Value(-100)          -> nson.I8(-100)
NewInt16Value(-1000)        -> nson.I16(-1000)
NewInt32Value(-10000)       -> nson.I32(-10000)
NewInt64Value(-100000)      -> nson.I64(-100000)
NewFloat32Value(3.14)       -> nson.F32(3.14)
NewFloat64Value(3.14159)    -> nson.F64(3.14159)
NewStringValue("test")      -> nson.String("test")
NewBytesValue([]byte{1,2})  -> nson.Binary([]byte{1,2})
NewNullValue()              -> nson.Null{}
```

### 枚举和位图类型

```go
// 枚举类型使用对应的整数类型
NewEnum8Value(3)            -> nson.U8(3)
NewEnum16Value(100)         -> nson.U16(100)

// 位图类型使用对应的整数类型
NewBitmap8Value(0b10101010) -> nson.U8(0b10101010)
NewBitmap16Value(0xABCD)    -> nson.U16(0xABCD)
NewBitmap32Value(0x12345678)-> nson.U32(0x12345678)
NewBitmap64Value(0x123456789ABCDEF0) -> nson.U64(0x123456789ABCDEF0)
```

### 复合类型

```go
// Array
NewArrayValue([]any{1, "test", true}) -> nson.Array{
    nson.I32(1),
    nson.String("test"),
    nson.Bool(true),
}

// Struct (使用 Map)
NewStructValue(map[string]any{
    "id": 123,
    "name": "device",
}) -> nson.Map{
    "id": nson.I32(123),
    "name": nson.String("device"),
}
```

## 使用示例

### 示例 1: Matter 设备基本信息

```go
package main

import (
    "bytes"
    "fmt"
    "github.com/danclive/nson-go"
)

func main() {
    // 创建设备信息（替代 beacon.AttributeValue）
    deviceInfo := nson.Map{
        "vendorId":      nson.U16(0x1234),
        "productId":     nson.U16(0x5678),
        "vendorName":    nson.String("Acme Corp"),
        "productName":   nson.String("Smart Light"),
        "serialNumber":  nson.String("SN-123456"),
        "softwareVersion": nson.U32(0x00010002),
    }

    // 编码为字节流
    var buf bytes.Buffer
    if err := nson.EncodeValue(&buf, deviceInfo); err != nil {
        panic(err)
    }

    fmt.Printf("Encoded size: %d bytes\n", buf.Len())

    // 解码
    decoded, err := nson.DecodeValue(&buf)
    if err != nil {
        panic(err)
    }

    // 类型断言并访问
    info, ok := decoded.(nson.Map)
    if !ok {
        panic("not a map")
    }

    vendorId := info["vendorId"].(nson.U16)
    productName := info["productName"].(nson.String)

    fmt.Printf("Vendor ID: 0x%04X\n", vendorId)
    fmt.Printf("Product: %s\n", productName)
}
```

### 示例 2: OnOff Cluster 状态

```go
// OnOff 集群状态
onOffState := nson.Map{
    "onOff": nson.Bool(true),
    "globalSceneControl": nson.Bool(true),
    "onTime": nson.U16(0),
    "offWaitTime": nson.U16(0),
}

// 编码
var buf bytes.Buffer
nson.EncodeValue(&buf, onOffState)

// 可以通过网络发送 buf.Bytes()
```

### 示例 3: LevelControl Cluster（调光）

```go
// 调光控制
levelControl := nson.Map{
    "currentLevel":  nson.U8(128),      // 当前亮度级别 (0-254)
    "remainingTime": nson.U16(0),       // 剩余过渡时间
    "minLevel":      nson.U8(1),        // 最小级别
    "maxLevel":      nson.U8(254),      // 最大级别
    "currentFrequency": nson.U16(0),    // 当前频率
    "onLevel":       nson.U8(254),      // 开启时的级别
}
```

### 示例 4: 温度传感器数据

```go
// 温度传感器（使用 I16，单位：0.01°C）
tempSensor := nson.Map{
    "measuredValue": nson.I16(2350),    // 23.50°C
    "minMeasuredValue": nson.I16(-5000), // -50.00°C
    "maxMeasuredValue": nson.I16(10000), // 100.00°C
    "tolerance": nson.U16(100),          // ±1.00°C
}

// 读取温度值
if temp, ok := tempSensor["measuredValue"].(nson.I16); ok {
    actualTemp := float64(temp) / 100.0
    fmt.Printf("Temperature: %.2f°C\n", actualTemp)
}
```

### 示例 5: 设备描述符（使用数组和结构体）

```go
// 设备描述符
descriptor := nson.Map{
    "deviceTypeList": nson.Array{
        nson.Map{
            "deviceType": nson.U32(0x0101), // DimmableLight
            "revision":   nson.U16(1),
        },
    },
    "serverList": nson.Array{
        nson.U32(0x0003), // Identify
        nson.U32(0x0006), // OnOff
        nson.U32(0x0008), // LevelControl
        nson.U32(0x0028), // BasicInformation
    },
    "clientList": nson.Array{},
    "partsList":  nson.Array{},
}
```

## 类型转换辅助函数

可以创建辅助函数来简化 beacon DataType 到 nson 类型的转换：

```go
package beacon

import (
    "fmt"
    "github.com/danclive/nson-go"
)

// DataTypeToNson 将 beacon DataType 和值转换为 nson.Value
func DataTypeToNson(dataType DataType, value any) (nson.Value, error) {
    switch dataType {
    case DataTypeBoolean:
        if v, ok := value.(bool); ok {
            return nson.Bool(v), nil
        }
    case DataTypeUInt8, DataTypeEnum8, DataTypeBitmap8:
        if v, ok := value.(uint8); ok {
            return nson.U8(v), nil
        }
    case DataTypeUInt16, DataTypeEnum16, DataTypeBitmap16:
        if v, ok := value.(uint16); ok {
            return nson.U16(v), nil
        }
    case DataTypeUInt32, DataTypeBitmap32:
        if v, ok := value.(uint32); ok {
            return nson.U32(v), nil
        }
    case DataTypeUInt64, DataTypeBitmap64:
        if v, ok := value.(uint64); ok {
            return nson.U64(v), nil
        }
    case DataTypeInt8:
        if v, ok := value.(int8); ok {
            return nson.I8(v), nil
        }
    case DataTypeInt16:
        if v, ok := value.(int16); ok {
            return nson.I16(v), nil
        }
    case DataTypeInt32:
        if v, ok := value.(int32); ok {
            return nson.I32(v), nil
        }
    case DataTypeInt64:
        if v, ok := value.(int64); ok {
            return nson.I64(v), nil
        }
    case DataTypeFloat:
        if v, ok := value.(float32); ok {
            return nson.F32(v), nil
        }
    case DataTypeDouble:
        if v, ok := value.(float64); ok {
            return nson.F64(v), nil
        }
    case DataTypeString:
        if v, ok := value.(string); ok {
            return nson.String(v), nil
        }
    case DataTypeOctetStr:
        if v, ok := value.([]byte); ok {
            return nson.Binary(v), nil
        }
    case DataTypeNull:
        return nson.Null{}, nil
    }

    return nil, fmt.Errorf("unsupported data type: %v", dataType)
}

// NsonToDataType 从 nson.Value 推断 DataType
func NsonToDataType(value nson.Value) DataType {
    switch value.(type) {
    case nson.Bool:
        return DataTypeBoolean
    case nson.U8:
        return DataTypeUInt8
    case nson.U16:
        return DataTypeUInt16
    case nson.U32:
        return DataTypeUInt32
    case nson.U64:
        return DataTypeUInt64
    case nson.I8:
        return DataTypeInt8
    case nson.I16:
        return DataTypeInt16
    case nson.I32:
        return DataTypeInt32
    case nson.I64:
        return DataTypeInt64
    case nson.F32:
        return DataTypeFloat
    case nson.F64:
        return DataTypeDouble
    case nson.String:
        return DataTypeString
    case nson.Binary:
        return DataTypeOctetStr
    case nson.Array:
        return DataTypeArray
    case nson.Map:
        return DataTypeStruct
    case nson.Null:
        return DataTypeNull
    default:
        return DataTypeNoData
    }
}
```

## 优势

使用 nson-go 替代 AttributeValue 的优势：

1. **序列化支持**: 内置完整的编码/解码功能，可以直接序列化到字节流
2. **更小的内存占用**: 直接使用 Go 原生类型，没有额外的包装
3. **类型安全**: 编译时类型检查
4. **扩展性**: 支持嵌套的 Map 和 Array，适合复杂的数据结构
5. **互操作性**: 标准的二进制格式，便于跨语言使用

## 迁移建议

1. **逐步迁移**: 先在新功能中使用 nson-go，逐步替换现有代码
2. **保持兼容**: 可以同时保留 AttributeValue，提供转换函数
3. **测试覆盖**: 为类型转换编写充分的单元测试
4. **文档更新**: 更新 API 文档，说明类型映射关系

## 性能考虑

- nson 使用小端序（Little Endian）编码
- U8/I8 类型只占用 1 字节（加 1 字节 tag）
- U16/I16 占用 2 字节（加 1 字节 tag）
- 适合嵌入式和低带宽场景
