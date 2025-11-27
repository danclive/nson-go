# NSON-Go Examples

本目录包含 NSON-Go 库的各种示例代码，按功能分类组织。

## 目录结构

### marshal/
结构体序列化/反序列化示例

- `marshal_demo.go` - 基本的 Marshal/Unmarshal 使用示例
  - 演示结构体到 nson.Map 的转换
  - 展示 struct tag 的使用
  - omitempty 标签示例
  - 嵌套结构体处理

### types/
类型系统演示

- `types_demo.go` - 类型系统综合演示
  - 8 种整数类型的精确映射（int8→I8, int16→I16 等）
  - time.Time ↔ Timestamp 自动转换（毫秒精度）
  - nson.Id 唯一标识符类型
  - 复杂结构示例

### codec/
编码/解码示例

- `extended_types_demo.go` - 扩展类型使用示例
  - Timestamp 类型
  - Id 类型
  - Binary 类型
  - 特殊类型的编码解码

## 运行示例

```bash
# 运行 Marshal 示例
go run examples/marshal/marshal_demo.go

## 运行示例

```bash
go run examples/marshal/marshal_demo.go      # Marshal 示例
go run examples/types/types_demo.go          # 类型系统示例
go run examples/codec/extended_types_demo.go # 扩展类型示例
``` 支持嵌套结构体和指针
- ✅ struct tag 自定义字段名
- ✅ omitempty 省略空值
- ✅ 完整的往返转换（roundtrip）

### 强类型示例特点
### 示例内容

**marshal/** - 结构体序列化
- struct tag、omitempty、嵌套结构、指针处理

**types/** - 类型系统
- 8 种整数类型精确映射、time.Time/Timestamp、nson.Id

**codec/** - 编解码
- 二进制编码解码、扩展类型使用
