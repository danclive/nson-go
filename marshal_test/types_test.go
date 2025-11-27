package nson_test

import (
	"testing"

	nson "github.com/danclive/nson-go"
)

// 测试所有整数类型的精确映射
func TestMarshalIntegerTypes(t *testing.T) {
	type AllIntTypes struct {
		I8Val  int8   `nson:"i8"`
		I16Val int16  `nson:"i16"`
		I32Val int32  `nson:"i32"`
		I64Val int64  `nson:"i64"`
		U8Val  uint8  `nson:"u8"`
		U16Val uint16 `nson:"u16"`
		U32Val uint32 `nson:"u32"`
		U64Val uint64 `nson:"u64"`
	}

	data := AllIntTypes{
		I8Val:  -42,
		I16Val: -1234,
		I32Val: -123456,
		I64Val: -9876543210,
		U8Val:  200,
		U16Val: 50000,
		U32Val: 4000000000,
		U64Val: 18000000000000000000,
	}

	m, err := nson.Marshal(data)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	// 验证每个字段的 NSON 类型
	if v, ok := m["i8"].(nson.I8); !ok {
		t.Errorf("i8 should be I8 type, got %T", m["i8"])
	} else if int8(v) != -42 {
		t.Errorf("i8 value mismatch: expected -42, got %d", v)
	}

	if v, ok := m["i16"].(nson.I16); !ok {
		t.Errorf("i16 should be I16 type, got %T", m["i16"])
	} else if int16(v) != -1234 {
		t.Errorf("i16 value mismatch: expected -1234, got %d", v)
	}

	if v, ok := m["i32"].(nson.I32); !ok {
		t.Errorf("i32 should be I32 type, got %T", m["i32"])
	} else if int32(v) != -123456 {
		t.Errorf("i32 value mismatch: expected -123456, got %d", v)
	}

	if v, ok := m["i64"].(nson.I64); !ok {
		t.Errorf("i64 should be I64 type, got %T", m["i64"])
	} else if int64(v) != -9876543210 {
		t.Errorf("i64 value mismatch: expected -9876543210, got %d", v)
	}

	if v, ok := m["u8"].(nson.U8); !ok {
		t.Errorf("u8 should be U8 type, got %T", m["u8"])
	} else if uint8(v) != 200 {
		t.Errorf("u8 value mismatch: expected 200, got %d", v)
	}

	if v, ok := m["u16"].(nson.U16); !ok {
		t.Errorf("u16 should be U16 type, got %T", m["u16"])
	} else if uint16(v) != 50000 {
		t.Errorf("u16 value mismatch: expected 50000, got %d", v)
	}

	if v, ok := m["u32"].(nson.U32); !ok {
		t.Errorf("u32 should be U32 type, got %T", m["u32"])
	} else if uint32(v) != 4000000000 {
		t.Errorf("u32 value mismatch: expected 4000000000, got %d", v)
	}

	if v, ok := m["u64"].(nson.U64); !ok {
		t.Errorf("u64 should be U64 type, got %T", m["u64"])
	} else if uint64(v) != 18000000000000000000 {
		t.Errorf("u64 value mismatch: expected 18000000000000000000, got %d", v)
	}
}

// 测试整数类型的反序列化
func TestUnmarshalIntegerTypes(t *testing.T) {
	type AllIntTypes struct {
		I8Val  int8   `nson:"i8"`
		I16Val int16  `nson:"i16"`
		I32Val int32  `nson:"i32"`
		I64Val int64  `nson:"i64"`
		U8Val  uint8  `nson:"u8"`
		U16Val uint16 `nson:"u16"`
		U32Val uint32 `nson:"u32"`
		U64Val uint64 `nson:"u64"`
	}

	m := nson.Map{
		"i8":  nson.I8(-42),
		"i16": nson.I16(-1234),
		"i32": nson.I32(-123456),
		"i64": nson.I64(-9876543210),
		"u8":  nson.U8(200),
		"u16": nson.U16(50000),
		"u32": nson.U32(4000000000),
		"u64": nson.U64(18000000000000000000),
	}

	var data AllIntTypes
	if err := nson.Unmarshal(m, &data); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if data.I8Val != -42 {
		t.Errorf("I8Val mismatch: expected -42, got %d", data.I8Val)
	}
	if data.I16Val != -1234 {
		t.Errorf("I16Val mismatch: expected -1234, got %d", data.I16Val)
	}
	if data.I32Val != -123456 {
		t.Errorf("I32Val mismatch: expected -123456, got %d", data.I32Val)
	}
	if data.I64Val != -9876543210 {
		t.Errorf("I64Val mismatch: expected -9876543210, got %d", data.I64Val)
	}
	if data.U8Val != 200 {
		t.Errorf("U8Val mismatch: expected 200, got %d", data.U8Val)
	}
	if data.U16Val != 50000 {
		t.Errorf("U16Val mismatch: expected 50000, got %d", data.U16Val)
	}
	if data.U32Val != 4000000000 {
		t.Errorf("U32Val mismatch: expected 4000000000, got %d", data.U32Val)
	}
	if data.U64Val != 18000000000000000000 {
		t.Errorf("U64Val mismatch: expected 18000000000000000000, got %d", data.U64Val)
	}
}

// 测试精确类型匹配（不允许类型转换）
func TestUnmarshalIntegerTypeConversion(t *testing.T) {
	// 测试类型不匹配会报错
	t.Run("I8 to int16 should fail", func(t *testing.T) {
		m := nson.Map{"value": nson.I8(42)}
		var data struct {
			Value int16 `nson:"value"`
		}
		err := nson.Unmarshal(m, &data)
		if err == nil {
			t.Error("Expected error when unmarshaling I8 to int16, but got none")
		}
	})

	t.Run("I8 to int32 should fail", func(t *testing.T) {
		m := nson.Map{"value": nson.I8(42)}
		var data struct {
			Value int32 `nson:"value"`
		}
		err := nson.Unmarshal(m, &data)
		if err == nil {
			t.Error("Expected error when unmarshaling I8 to int32, but got none")
		}
	})

	t.Run("I8 to int64 should fail", func(t *testing.T) {
		m := nson.Map{"value": nson.I8(42)}
		var data struct {
			Value int64 `nson:"value"`
		}
		err := nson.Unmarshal(m, &data)
		if err == nil {
			t.Error("Expected error when unmarshaling I8 to int64, but got none")
		}
	})
}

// 测试完整的往返转换，确保类型不丢失
func TestIntegerTypesRoundTrip(t *testing.T) {
	type TestData struct {
		Small   int8   `nson:"small"`
		Medium  int16  `nson:"medium"`
		Large   int32  `nson:"large"`
		Huge    int64  `nson:"huge"`
		USmall  uint8  `nson:"usmall"`
		UMedium uint16 `nson:"umedium"`
		ULarge  uint32 `nson:"ularge"`
		UHuge   uint64 `nson:"uhuge"`
	}

	original := TestData{
		Small:   -100,
		Medium:  -30000,
		Large:   -2000000000,
		Huge:    -9000000000000000000,
		USmall:  250,
		UMedium: 60000,
		ULarge:  4000000000,
		UHuge:   18000000000000000000,
	}

	// Marshal
	m, err := nson.Marshal(original)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	// 验证 NSON 类型
	if _, ok := m["small"].(nson.I8); !ok {
		t.Errorf("small should be I8, got %T", m["small"])
	}
	if _, ok := m["medium"].(nson.I16); !ok {
		t.Errorf("medium should be I16, got %T", m["medium"])
	}
	if _, ok := m["large"].(nson.I32); !ok {
		t.Errorf("large should be I32, got %T", m["large"])
	}
	if _, ok := m["huge"].(nson.I64); !ok {
		t.Errorf("huge should be I64, got %T", m["huge"])
	}
	if _, ok := m["usmall"].(nson.U8); !ok {
		t.Errorf("usmall should be U8, got %T", m["usmall"])
	}
	if _, ok := m["umedium"].(nson.U16); !ok {
		t.Errorf("umedium should be U16, got %T", m["umedium"])
	}
	if _, ok := m["ularge"].(nson.U32); !ok {
		t.Errorf("ularge should be U32, got %T", m["ularge"])
	}
	if _, ok := m["uhuge"].(nson.U64); !ok {
		t.Errorf("uhuge should be U64, got %T", m["uhuge"])
	}

	// Unmarshal
	var result TestData
	if err := nson.Unmarshal(m, &result); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	// 验证值
	if result.Small != original.Small {
		t.Errorf("Small mismatch: expected %d, got %d", original.Small, result.Small)
	}
	if result.Medium != original.Medium {
		t.Errorf("Medium mismatch: expected %d, got %d", original.Medium, result.Medium)
	}
	if result.Large != original.Large {
		t.Errorf("Large mismatch: expected %d, got %d", original.Large, result.Large)
	}
	if result.Huge != original.Huge {
		t.Errorf("Huge mismatch: expected %d, got %d", original.Huge, result.Huge)
	}
	if result.USmall != original.USmall {
		t.Errorf("USmall mismatch: expected %d, got %d", original.USmall, result.USmall)
	}
	if result.UMedium != original.UMedium {
		t.Errorf("UMedium mismatch: expected %d, got %d", original.UMedium, result.UMedium)
	}
	if result.ULarge != original.ULarge {
		t.Errorf("ULarge mismatch: expected %d, got %d", original.ULarge, result.ULarge)
	}
	if result.UHuge != original.UHuge {
		t.Errorf("UHuge mismatch: expected %d, got %d", original.UHuge, result.UHuge)
	}
}

// 测试数组中的类型保持
func TestIntegerTypesInArray(t *testing.T) {
	type TestData struct {
		I8Array  []int8   `nson:"i8_array"`
		I16Array []int16  `nson:"i16_array"`
		U16Array []uint16 `nson:"u16_array"`
		U32Array []uint32 `nson:"u32_array"`
	}

	data := TestData{
		I8Array:  []int8{-1, -2, -3},
		I16Array: []int16{-1000, -2000, -3000},
		U16Array: []uint16{10000, 20000, 30000},
		U32Array: []uint32{100000, 200000, 300000},
	}

	m, err := nson.Marshal(data)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	// 验证数组元素的类型
	if arr, ok := m["i8_array"].(nson.Array); ok {
		for i, v := range arr {
			if _, ok := v.(nson.I8); !ok {
				t.Errorf("i8_array[%d] should be I8, got %T", i, v)
			}
		}
	} else {
		t.Error("i8_array should be Array")
	}

	if arr, ok := m["i16_array"].(nson.Array); ok {
		for i, v := range arr {
			if _, ok := v.(nson.I16); !ok {
				t.Errorf("i16_array[%d] should be I16, got %T", i, v)
			}
		}
	} else {
		t.Error("i16_array should be Array")
	}

	if arr, ok := m["u16_array"].(nson.Array); ok {
		for i, v := range arr {
			if _, ok := v.(nson.U16); !ok {
				t.Errorf("u16_array[%d] should be U16, got %T", i, v)
			}
		}
	} else {
		t.Error("u16_array should be Array")
	}

	if arr, ok := m["u32_array"].(nson.Array); ok {
		for i, v := range arr {
			if _, ok := v.(nson.U32); !ok {
				t.Errorf("u32_array[%d] should be U32, got %T", i, v)
			}
		}
	} else {
		t.Error("u32_array should be Array")
	}

	// 反序列化测试
	var result TestData
	if err := nson.Unmarshal(m, &result); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if len(result.I8Array) != 3 || result.I8Array[0] != -1 {
		t.Errorf("I8Array mismatch: %v", result.I8Array)
	}
	if len(result.I16Array) != 3 || result.I16Array[0] != -1000 {
		t.Errorf("I16Array mismatch: %v", result.I16Array)
	}
	if len(result.U16Array) != 3 || result.U16Array[0] != 10000 {
		t.Errorf("U16Array mismatch: %v", result.U16Array)
	}
	if len(result.U32Array) != 3 || result.U32Array[0] != 100000 {
		t.Errorf("U32Array mismatch: %v", result.U32Array)
	}
}

// 基准测试：验证类型映射不影响性能
func BenchmarkMarshalWithPreciseTypes(b *testing.B) {
	type TestData struct {
		I8Val  int8   `nson:"i8"`
		I16Val int16  `nson:"i16"`
		I32Val int32  `nson:"i32"`
		I64Val int64  `nson:"i64"`
		U8Val  uint8  `nson:"u8"`
		U16Val uint16 `nson:"u16"`
		U32Val uint32 `nson:"u32"`
		U64Val uint64 `nson:"u64"`
	}

	data := TestData{
		I8Val:  -42,
		I16Val: -1234,
		I32Val: -123456,
		I64Val: -9876543210,
		U8Val:  200,
		U16Val: 50000,
		U32Val: 4000000000,
		U64Val: 18000000000000000000,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := nson.Marshal(data)
		if err != nil {
			b.Fatal(err)
		}
	}
}
