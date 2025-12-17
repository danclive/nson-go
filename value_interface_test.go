package nson

import (
	"testing"
)

// TestValueInterfaceMarshaling 测试 nson.Value 接口类型的序列化
func TestValueInterfaceMarshaling(t *testing.T) {
	type TestStruct struct {
		Name  string `nson:"name"`
		Value Value  `nson:"value"`
	}

	tests := []struct {
		name  string
		input TestStruct
	}{
		{
			name: "I32 value",
			input: TestStruct{
				Name:  "test1",
				Value: I32(42),
			},
		},
		{
			name: "F64 value",
			input: TestStruct{
				Name:  "test2",
				Value: F64(3.14),
			},
		},
		{
			name: "String value",
			input: TestStruct{
				Name:  "test3",
				Value: String("hello"),
			},
		},
		{
			name: "Bool value",
			input: TestStruct{
				Name:  "test4",
				Value: Bool(true),
			},
		},
		{
			name: "Array value",
			input: TestStruct{
				Name:  "test5",
				Value: Array{I32(1), I32(2), I32(3)},
			},
		},
		{
			name: "Map value",
			input: TestStruct{
				Name:  "test6",
				Value: Map{"key": String("value")},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Marshal
			m, err := Marshal(tt.input)
			if err != nil {
				t.Fatalf("Marshal failed: %v", err)
			}

			// Unmarshal
			var output TestStruct
			if err := Unmarshal(m, &output); err != nil {
				t.Fatalf("Unmarshal failed: %v", err)
			}

			// Verify
			if output.Name != tt.input.Name {
				t.Errorf("Name mismatch: got %v, want %v", output.Name, tt.input.Name)
			}

			// 比较 Value (需要类型断言)
			if !compareValues(output.Value, tt.input.Value) {
				t.Errorf("Value mismatch: got %v (%T), want %v (%T)",
					output.Value, output.Value, tt.input.Value, tt.input.Value)
			}
		})
	}
}

// compareValues 比较两个 nson.Value 是否相等
func compareValues(a, b Value) bool {
	if a.DataType() != b.DataType() {
		return false
	}

	switch v1 := a.(type) {
	case I32:
		v2, ok := b.(I32)
		return ok && v1 == v2
	case F64:
		v2, ok := b.(F64)
		return ok && v1 == v2
	case String:
		v2, ok := b.(String)
		return ok && v1 == v2
	case Bool:
		v2, ok := b.(Bool)
		return ok && v1 == v2
	case Array:
		v2, ok := b.(Array)
		if !ok || len(v1) != len(v2) {
			return false
		}
		for i := range v1 {
			if !compareValues(v1[i], v2[i]) {
				return false
			}
		}
		return true
	case Map:
		v2, ok := b.(Map)
		if !ok || len(v1) != len(v2) {
			return false
		}
		for key, val1 := range v1 {
			val2, exists := v2[key]
			if !exists || !compareValues(val1, val2) {
				return false
			}
		}
		return true
	default:
		return false
	}
}
