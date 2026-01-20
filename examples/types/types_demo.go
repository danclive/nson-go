package main

import (
	"fmt"
	"time"

	"github.com/danclive/nson-go"
)

func main() {
	fmt.Println("=== NSON 类型系统演示 ===")

	demo1_StrongTypes()
	demo2_TimestampType()
	demo3_IdType()
	demo4_ComplexExample()
}

// 演示 1: 强类型整数系统
func demo1_StrongTypes() {
	fmt.Println("1. 强类型整数系统")
	fmt.Println("------------------")

	type IntegerTypes struct {
		Int8   int8   `nson:"i8"`
		Int16  int16  `nson:"i16"`
		Int32  int32  `nson:"i32"`
		Int64  int64  `nson:"i64"`
		Uint8  uint8  `nson:"u8"`
		Uint16 uint16 `nson:"u16"`
		Uint32 uint32 `nson:"u32"`
		Uint64 uint64 `nson:"u64"`
	}

	data := IntegerTypes{
		Int8:   -42,
		Int16:  -1234,
		Int32:  -123456,
		Int64:  -987654321,
		Uint8:  200,
		Uint16: 50000,
		Uint32: 4000000000,
		Uint64: 18000000000000000000,
	}

	m, _ := nson.Marshal(data)

	// 验证类型映射
	fmt.Println("类型映射:")
	fmt.Printf("  int8   -> I8   (值: %d)\n", m["i8"].(nson.I8))
	fmt.Printf("  int16  -> I16  (值: %d)\n", m["i16"].(nson.I16))
	fmt.Printf("  int32  -> I32  (值: %d)\n", m["i32"].(nson.I32))
	fmt.Printf("  int64  -> I64  (值: %d)\n", m["i64"].(nson.I64))
	fmt.Printf("  uint8  -> U8   (值: %d)\n", m["u8"].(nson.U8))
	fmt.Printf("  uint16 -> U16  (值: %d)\n", m["u16"].(nson.U16))
	fmt.Printf("  uint32 -> U32  (值: %d)\n", m["u32"].(nson.U32))
	fmt.Printf("  uint64 -> U64  (值: %d)\n\n", m["u64"].(nson.U64))

	// 反序列化验证
	var result IntegerTypes
	nson.Unmarshal(m, &result)
	fmt.Printf("往返转换成功: %v\n\n", data == result)
}

// 演示 2: Timestamp 类型（time.Time）
func demo2_TimestampType() {
	fmt.Println("2. Timestamp 类型（time.Time）")
	fmt.Println("-----------------------------")

	type Event struct {
		Name      string    `nson:"name"`
		CreatedAt time.Time `nson:"created_at"`
		UpdatedAt time.Time `nson:"updated_at"`
	}

	now := time.Now()
	event := Event{
		Name:      "系统启动",
		CreatedAt: now.Add(-24 * time.Hour), // 24 小时前
		UpdatedAt: now,
	}

	m, _ := nson.Marshal(event)

	// time.Time 自动转换为 Timestamp（毫秒时间戳）
	created := m["created_at"].(nson.Timestamp)
	updated := m["updated_at"].(nson.Timestamp)

	fmt.Printf("原始时间: %v\n", now.Format("2006-01-02 15:04:05"))
	fmt.Printf("Timestamp (毫秒): %d\n", updated)
	fmt.Printf("DataType: %v\n\n", created.DataType())

	// 反序列化自动转回 time.Time
	var result Event
	nson.Unmarshal(m, &result)
	fmt.Printf("反序列化成功，时间匹配: %v\n\n",
		event.UpdatedAt.UnixMilli() == result.UpdatedAt.UnixMilli())
}

// 演示 3: Id 类型（唯一标识符）
func demo3_IdType() {
	fmt.Println("3. Id 类型（唯一标识符）")
	fmt.Println("------------------------")

	type Document struct {
		DocID   nson.Id `nson:"doc_id"`
		Title   string  `nson:"title"`
		Content []byte  `nson:"content"` // Binary 类型
	}

	// 创建文档 ID（通常是 12 字节的 ObjectId）
	doc := Document{
		DocID:   nson.Id{0x5f, 0x8d, 0x0a, 0x1e, 0xb2, 0xc3, 0xd4, 0xe5, 0xf6, 0x01, 0x02, 0x03},
		Title:   "测试文档",
		Content: []byte{0x01, 0x02, 0x03, 0x04},
	}

	m, _ := nson.Marshal(doc)

	// Id 和 Binary 是不同的类型
	docID := m["doc_id"].(nson.Id)
	content := m["content"].(nson.Binary)

	fmt.Printf("DocID (Id):      %x\n", docID)
	fmt.Printf("Content (Binary): %x\n", content)
	fmt.Printf("Id DataType:      %v\n", docID.DataType())
	fmt.Printf("Binary DataType:  %v\n\n", content.DataType())
}

// 演示 4: 复杂结构示例
func demo4_ComplexExample() {
	fmt.Println("4. 复杂结构示例")
	fmt.Println("----------------")

	type User struct {
		ID        nson.Id    `nson:"id"`
		Name      string     `nson:"name"`
		Age       int32      `nson:"age"`
		Score     float64    `nson:"score"`
		Active    bool       `nson:"active"`
		Tags      []string   `nson:"tags"`
		CreatedAt time.Time  `nson:"created_at"`
		LastLogin *time.Time `nson:"last_login,omitempty"` // 可选字段
	}

	lastLogin := time.Now().Add(-2 * time.Hour)
	user := User{
		ID:        nson.Id{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c},
		Name:      "Alice",
		Age:       28,
		Score:     95.5,
		Active:    true,
		Tags:      []string{"developer", "golang"},
		CreatedAt: time.Now().Add(-30 * 24 * time.Hour),
		LastLogin: &lastLogin,
	}

	m, _ := nson.Marshal(user)

	fmt.Println("序列化结果包含的类型:")
	fmt.Printf("  ID:        %T (DataType: %v)\n", m["id"], m["id"].(nson.Id).DataType())
	fmt.Printf("  Name:      %T\n", m["name"])
	fmt.Printf("  Age:       %T\n", m["age"])
	fmt.Printf("  Score:     %T\n", m["score"])
	fmt.Printf("  Active:    %T\n", m["active"])
	fmt.Printf("  Tags:      %T\n", m["tags"])
	fmt.Printf("  CreatedAt: %T (DataType: %v)\n", m["created_at"], m["created_at"].(nson.Timestamp).DataType())
	fmt.Printf("  LastLogin: %T\n\n", m["last_login"])

	// 反序列化
	var result User
	nson.Unmarshal(m, &result)

	fmt.Println("反序列化验证:")
	fmt.Printf("  ID 匹配:   %v\n", user.ID == result.ID)
	fmt.Printf("  Name:      %s\n", result.Name)
	fmt.Printf("  Age:       %d\n", result.Age)
	fmt.Printf("  时间匹配:  %v\n", user.CreatedAt.UnixMilli() == result.CreatedAt.UnixMilli())
}
