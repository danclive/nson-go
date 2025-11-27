package nson_test

import (
	"strings"
	"testing"

	nson "github.com/danclive/nson-go"
)

// 测试强类型检查：只允许精确类型匹配
func TestStrictTypeChecking(t *testing.T) {
	tests := []struct {
		name        string
		targetType  string
		value       nson.Value
		shouldError bool
		errorMsg    string
	}{
		// int8 只接受 I8
		{"int8 accepts I8", "int8", nson.I8(-42), false, ""},
		{"int8 rejects I16", "int8", nson.I16(-100), true, "expected I8"},
		{"int8 rejects I32", "int8", nson.I32(-100), true, "expected I8"},
		{"int8 rejects I64", "int8", nson.I64(-100), true, "expected I8"},

		// int16 只接受 I16（精确匹配）
		{"int16 accepts I16", "int16", nson.I16(-1234), false, ""},
		{"int16 rejects I8", "int16", nson.I8(-42), true, "expected I16"},
		{"int16 rejects I32", "int16", nson.I32(-10000), true, "expected I16"},
		{"int16 rejects I64", "int16", nson.I64(-10000), true, "expected I16"},

		// int32 只接受 I32（精确匹配）
		{"int32 accepts I32", "int32", nson.I32(-123456), false, ""},
		{"int32 rejects I8", "int32", nson.I8(-42), true, "expected I32"},
		{"int32 rejects I16", "int32", nson.I16(-1234), true, "expected I32"},
		{"int32 rejects I64", "int32", nson.I64(-1000000), true, "expected I32"},

		// uint8 只接受 U8
		{"uint8 accepts U8", "uint8", nson.U8(200), false, ""},
		{"uint8 rejects U16", "uint8", nson.U16(300), true, "expected U8"},
		{"uint8 rejects U32", "uint8", nson.U32(300), true, "expected U8"},
		{"uint8 rejects U64", "uint8", nson.U64(300), true, "expected U8"},

		// uint16 只接受 U16（精确匹配）
		{"uint16 accepts U16", "uint16", nson.U16(50000), false, ""},
		{"uint16 rejects U8", "uint16", nson.U8(200), true, "expected U16"},
		{"uint16 rejects U32", "uint16", nson.U32(70000), true, "expected U16"},
		{"uint16 rejects U64", "uint16", nson.U64(70000), true, "expected U16"},

		// uint32 只接受 U32（精确匹配）
		{"uint32 accepts U32", "uint32", nson.U32(4000000000), false, ""},
		{"uint32 rejects U8", "uint32", nson.U8(200), true, "expected U32"},
		{"uint32 rejects U16", "uint32", nson.U16(50000), true, "expected U32"},
		{"uint32 rejects U64", "uint32", nson.U64(5000000000), true, "expected U32"},

		// float32 只接受 F32（精确匹配）
		{"float32 accepts F32", "float32", nson.F32(3.14), false, ""},
		{"float32 rejects F64", "float32", nson.F64(3.14159), true, "expected F32"},

		// float64 只接受 F64（精确匹配）
		{"float64 accepts F64", "float64", nson.F64(3.14159), false, ""},
		{"float64 rejects F32", "float64", nson.F32(3.14), true, "expected F64"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := nson.Map{"value": tt.value}

			var err error
			switch tt.targetType {
			case "int8":
				var result struct {
					Value int8 `nson:"value"`
				}
				err = nson.Unmarshal(m, &result)
			case "int16":
				var result struct {
					Value int16 `nson:"value"`
				}
				err = nson.Unmarshal(m, &result)
			case "int32":
				var result struct {
					Value int32 `nson:"value"`
				}
				err = nson.Unmarshal(m, &result)
			case "uint8":
				var result struct {
					Value uint8 `nson:"value"`
				}
				err = nson.Unmarshal(m, &result)
			case "uint16":
				var result struct {
					Value uint16 `nson:"value"`
				}
				err = nson.Unmarshal(m, &result)
			case "uint32":
				var result struct {
					Value uint32 `nson:"value"`
				}
				err = nson.Unmarshal(m, &result)
			case "float32":
				var result struct {
					Value float32 `nson:"value"`
				}
				err = nson.Unmarshal(m, &result)
			case "float64":
				var result struct {
					Value float64 `nson:"value"`
				}
				err = nson.Unmarshal(m, &result)
			}

			if tt.shouldError {
				if err == nil {
					t.Errorf("Expected error containing '%s', but got no error", tt.errorMsg)
				} else if !strings.Contains(err.Error(), tt.errorMsg) {
					t.Errorf("Expected error containing '%s', but got: %v", tt.errorMsg, err)
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error, but got: %v", err)
				}
			}
		})
	}
}

// 测试无符号和有符号整数不能互相转换
func TestSignedUnsignedSeparation(t *testing.T) {
	t.Run("unsigned cannot unmarshal to signed", func(t *testing.T) {
		m := nson.Map{"value": nson.U8(100)}
		var result struct {
			Value int8 `nson:"value"`
		}
		err := nson.Unmarshal(m, &result)
		if err == nil {
			t.Error("Expected error when unmarshaling unsigned to signed type")
		}
	})

	t.Run("signed cannot unmarshal to unsigned", func(t *testing.T) {
		m := nson.Map{"value": nson.I8(-10)}
		var result struct {
			Value uint8 `nson:"value"`
		}
		err := nson.Unmarshal(m, &result)
		if err == nil {
			t.Error("Expected error when unmarshaling signed to unsigned type")
		}
	})
}
