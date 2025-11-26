package nson

import (
	"bytes"
	"testing"
)

// 测试 U8 类型
func TestU8(t *testing.T) {
	var buf bytes.Buffer
	original := U8(42)

	// 编码
	if err := EncodeValue(&buf, original); err != nil {
		t.Fatalf("Failed to encode U8: %v", err)
	}

	// 解码
	decoded, err := DecodeValue(&buf)
	if err != nil {
		t.Fatalf("Failed to decode U8: %v", err)
	}

	result, ok := decoded.(U8)
	if !ok {
		t.Fatalf("Decoded value is not U8")
	}

	if result != original {
		t.Errorf("Expected %v, got %v", original, result)
	}

	t.Logf("U8 test passed: %v", result)
}

// 测试 U16 类型
func TestU16(t *testing.T) {
	var buf bytes.Buffer
	original := U16(1234)

	// 编码
	if err := EncodeValue(&buf, original); err != nil {
		t.Fatalf("Failed to encode U16: %v", err)
	}

	// 解码
	decoded, err := DecodeValue(&buf)
	if err != nil {
		t.Fatalf("Failed to decode U16: %v", err)
	}

	result, ok := decoded.(U16)
	if !ok {
		t.Fatalf("Decoded value is not U16")
	}

	if result != original {
		t.Errorf("Expected %v, got %v", original, result)
	}

	t.Logf("U16 test passed: %v", result)
}

// 测试 I8 类型
func TestI8(t *testing.T) {
	var buf bytes.Buffer
	original := I8(-42)

	// 编码
	if err := EncodeValue(&buf, original); err != nil {
		t.Fatalf("Failed to encode I8: %v", err)
	}

	// 解码
	decoded, err := DecodeValue(&buf)
	if err != nil {
		t.Fatalf("Failed to decode I8: %v", err)
	}

	result, ok := decoded.(I8)
	if !ok {
		t.Fatalf("Decoded value is not I8")
	}

	if result != original {
		t.Errorf("Expected %v, got %v", original, result)
	}

	t.Logf("I8 test passed: %v", result)
}

// 测试 I16 类型
func TestI16(t *testing.T) {
	var buf bytes.Buffer
	original := I16(-1234)

	// 编码
	if err := EncodeValue(&buf, original); err != nil {
		t.Fatalf("Failed to encode I16: %v", err)
	}

	// 解码
	decoded, err := DecodeValue(&buf)
	if err != nil {
		t.Fatalf("Failed to decode I16: %v", err)
	}

	result, ok := decoded.(I16)
	if !ok {
		t.Fatalf("Decoded value is not I16")
	}

	if result != original {
		t.Errorf("Expected %v, got %v", original, result)
	}

	t.Logf("I16 test passed: %v", result)
}

// 测试在 Map 中使用新类型（模拟 beacon AttributeValue）
func TestMapWithExtendedTypes(t *testing.T) {
	var buf bytes.Buffer

	// 创建一个包含各种类型的 Map，模拟 Matter 设备属性
	original := Map{
		"vendorId":    U16(0x1234),          // VendorID
		"productId":   U16(0x5678),          // ProductID
		"onOff":       Bool(true),           // OnOff 状态
		"level":       U8(128),              // 调光级别 (0-255)
		"temperature": I16(2350),            // 温度 (23.50°C * 100)
		"humidity":    U8(65),               // 湿度百分比
		"vendorName":  String("TestVendor"), // 厂商名称
		"enabled":     Bool(true),           // 启用状态
		"counter":     U32(12345),           // 计数器
		"timestamp":   U64(1234567890),      // 时间戳
		"offset":      I8(-10),              // 偏移量
		"brightness":  F32(0.85),            // 亮度 (0.0-1.0)
	}

	// 编码
	if err := EncodeValue(&buf, original); err != nil {
		t.Fatalf("Failed to encode Map: %v", err)
	}

	// 解码
	decoded, err := DecodeValue(&buf)
	if err != nil {
		t.Fatalf("Failed to decode Map: %v", err)
	}

	result, ok := decoded.(Map)
	if !ok {
		t.Fatalf("Decoded value is not Map")
	}

	// 验证各个字段
	if vendorId, ok := result["vendorId"].(U16); !ok || vendorId != 0x1234 {
		t.Errorf("vendorId mismatch")
	}

	if productId, ok := result["productId"].(U16); !ok || productId != 0x5678 {
		t.Errorf("productId mismatch")
	}

	if onOff, ok := result["onOff"].(Bool); !ok || onOff != true {
		t.Errorf("onOff mismatch")
	}

	if level, ok := result["level"].(U8); !ok || level != 128 {
		t.Errorf("level mismatch")
	}

	if temp, ok := result["temperature"].(I16); !ok || temp != 2350 {
		t.Errorf("temperature mismatch")
	}

	if humidity, ok := result["humidity"].(U8); !ok || humidity != 65 {
		t.Errorf("humidity mismatch")
	}

	if offset, ok := result["offset"].(I8); !ok || offset != -10 {
		t.Errorf("offset mismatch")
	}

	t.Log("Map with extended types test passed")
}

// 测试 Array 包含新类型
func TestArrayWithExtendedTypes(t *testing.T) {
	var buf bytes.Buffer

	// 创建包含各种新类型的数组
	original := Array{
		U8(10),
		U16(1000),
		I8(-5),
		I16(-500),
		Bool(true),
		String("test"),
	}

	// 编码
	if err := EncodeValue(&buf, original); err != nil {
		t.Fatalf("Failed to encode Array: %v", err)
	}

	// 解码
	decoded, err := DecodeValue(&buf)
	if err != nil {
		t.Fatalf("Failed to decode Array: %v", err)
	}

	result, ok := decoded.(Array)
	if !ok {
		t.Fatalf("Decoded value is not Array")
	}

	if len(result) != len(original) {
		t.Fatalf("Array length mismatch: expected %d, got %d", len(original), len(result))
	}

	// 验证元素
	if v, ok := result[0].(U8); !ok || v != 10 {
		t.Errorf("Element 0 mismatch")
	}
	if v, ok := result[1].(U16); !ok || v != 1000 {
		t.Errorf("Element 1 mismatch")
	}
	if v, ok := result[2].(I8); !ok || v != -5 {
		t.Errorf("Element 2 mismatch")
	}
	if v, ok := result[3].(I16); !ok || v != -500 {
		t.Errorf("Element 3 mismatch")
	}

	t.Log("Array with extended types test passed")
}

// 测试边界值
func TestBoundaryValues(t *testing.T) {
	tests := []struct {
		name  string
		value Value
	}{
		{"U8 Min", U8(0)},
		{"U8 Max", U8(255)},
		{"U16 Min", U16(0)},
		{"U16 Max", U16(65535)},
		{"I8 Min", I8(-128)},
		{"I8 Max", I8(127)},
		{"I16 Min", I16(-32768)},
		{"I16 Max", I16(32767)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer

			if err := EncodeValue(&buf, tt.value); err != nil {
				t.Fatalf("Failed to encode %s: %v", tt.name, err)
			}

			decoded, err := DecodeValue(&buf)
			if err != nil {
				t.Fatalf("Failed to decode %s: %v", tt.name, err)
			}

			// 使用类型断言验证
			switch v := tt.value.(type) {
			case U8:
				if result, ok := decoded.(U8); !ok || result != v {
					t.Errorf("%s: expected %v, got %v", tt.name, v, result)
				}
			case U16:
				if result, ok := decoded.(U16); !ok || result != v {
					t.Errorf("%s: expected %v, got %v", tt.name, v, result)
				}
			case I8:
				if result, ok := decoded.(I8); !ok || result != v {
					t.Errorf("%s: expected %v, got %v", tt.name, v, result)
				}
			case I16:
				if result, ok := decoded.(I16); !ok || result != v {
					t.Errorf("%s: expected %v, got %v", tt.name, v, result)
				}
			}
		})
	}
}
