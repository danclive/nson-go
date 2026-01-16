package nson

import (
	"bytes"
	"testing"
)

// 测试类型别名的序列化和反序列化

// 定义一些类型别名（类似 Queen 项目中的 packet.QoS 和 packet.Priority）
type (
	QoS      uint8
	Priority uint8
	Status   int32
	Code     uint16
)

// QoS 常量
const (
	QoS0 QoS = 0
	QoS1 QoS = 1
	QoS2 QoS = 2
)

// Priority 常量
const (
	PriorityLow    Priority = 0
	PriorityNormal Priority = 1
	PriorityHigh   Priority = 2
)

// 测试结构体
type Message struct {
	ID       string            `nson:"id"`
	Topic    string            `nson:"topic"`
	Payload  []byte            `nson:"payload,omitempty"`
	QoS      QoS               `nson:"qos"`
	Priority Priority          `nson:"priority"`
	Status   Status            `nson:"status"`
	Code     Code              `nson:"code,omitempty"`
	Metadata map[string]string `nson:"metadata,omitempty"`
}

// TestTypeAlias_BasicTypes 测试基本类型别名
func TestTypeAlias_BasicTypes(t *testing.T) {
	tests := []struct {
		name     string
		original Message
	}{
		{
			name: "all fields",
			original: Message{
				ID:       "msg-123",
				Topic:    "test/topic",
				Payload:  []byte("hello world"),
				QoS:      QoS1,
				Priority: PriorityHigh,
				Status:   200,
				Code:     100,
				Metadata: map[string]string{"key": "value"},
			},
		},
		{
			name: "minimal fields",
			original: Message{
				ID:       "msg-456",
				Topic:    "test/minimal",
				QoS:      QoS0,
				Priority: PriorityNormal,
				Status:   0,
			},
		},
		{
			name: "with omitempty",
			original: Message{
				ID:       "msg-789",
				Topic:    "test/omit",
				QoS:      QoS2,
				Priority: PriorityLow,
				Status:   404,
				Code:     0, // 应该被 omitempty 省略
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 序列化
			m, err := Marshal(&tt.original)
			if err != nil {
				t.Fatalf("Marshal failed: %v", err)
			}

			// 检查类型
			if qos, ok := m["qos"].(U8); !ok {
				t.Errorf("Expected U8 for qos, got %T", m["qos"])
			} else if uint8(qos) != uint8(tt.original.QoS) {
				t.Errorf("QoS value mismatch: got %d, want %d", qos, tt.original.QoS)
			}

			if priority, ok := m["priority"].(U8); !ok {
				t.Errorf("Expected U8 for priority, got %T", m["priority"])
			} else if uint8(priority) != uint8(tt.original.Priority) {
				t.Errorf("Priority value mismatch: got %d, want %d", priority, tt.original.Priority)
			}

			if status, ok := m["status"].(I32); !ok {
				t.Errorf("Expected I32 for status, got %T", m["status"])
			} else if int32(status) != int32(tt.original.Status) {
				t.Errorf("Status value mismatch: got %d, want %d", status, tt.original.Status)
			}

			// 反序列化
			var decoded Message
			if err := Unmarshal(m, &decoded); err != nil {
				t.Fatalf("Unmarshal failed: %v", err)
			}

			// 验证所有字段
			if decoded.ID != tt.original.ID {
				t.Errorf("ID mismatch: got %q, want %q", decoded.ID, tt.original.ID)
			}
			if decoded.Topic != tt.original.Topic {
				t.Errorf("Topic mismatch: got %q, want %q", decoded.Topic, tt.original.Topic)
			}
			if string(decoded.Payload) != string(tt.original.Payload) {
				t.Errorf("Payload mismatch: got %q, want %q", string(decoded.Payload), string(tt.original.Payload))
			}
			if decoded.QoS != tt.original.QoS {
				t.Errorf("QoS mismatch: got %d, want %d", decoded.QoS, tt.original.QoS)
			}
			if decoded.Priority != tt.original.Priority {
				t.Errorf("Priority mismatch: got %d, want %d", decoded.Priority, tt.original.Priority)
			}
			if decoded.Status != tt.original.Status {
				t.Errorf("Status mismatch: got %d, want %d", decoded.Status, tt.original.Status)
			}
			if decoded.Code != tt.original.Code {
				t.Errorf("Code mismatch: got %d, want %d", decoded.Code, tt.original.Code)
			}
		})
	}
}

// TestTypeAlias_RoundTrip 测试完整的编码解码流程
func TestTypeAlias_RoundTrip(t *testing.T) {
	original := Message{
		ID:       "round-trip-test",
		Topic:    "test/round-trip",
		Payload:  []byte("test payload"),
		QoS:      QoS1,
		Priority: PriorityHigh,
		Status:   201,
		Code:     999,
		Metadata: map[string]string{
			"version": "1.0",
			"author":  "test",
		},
	}

	// Marshal -> EncodeMap
	m, err := Marshal(&original)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	buf := new(bytes.Buffer)
	if err := EncodeMap(m, buf); err != nil {
		t.Fatalf("EncodeMap failed: %v", err)
	}

	// DecodeMap -> Unmarshal
	m2, err := DecodeMap(bytes.NewBuffer(buf.Bytes()))
	if err != nil {
		t.Fatalf("DecodeMap failed: %v", err)
	}

	var decoded Message
	if err := Unmarshal(m2, &decoded); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	// 验证
	if decoded.ID != original.ID {
		t.Errorf("ID mismatch")
	}
	if decoded.QoS != original.QoS {
		t.Errorf("QoS mismatch: got %d, want %d", decoded.QoS, original.QoS)
	}
	if decoded.Priority != original.Priority {
		t.Errorf("Priority mismatch: got %d, want %d", decoded.Priority, original.Priority)
	}
	if decoded.Status != original.Status {
		t.Errorf("Status mismatch: got %d, want %d", decoded.Status, original.Status)
	}
	if decoded.Code != original.Code {
		t.Errorf("Code mismatch: got %d, want %d", decoded.Code, original.Code)
	}
}

// TestTypeAlias_NestedStructs 测试嵌套结构体中的类型别名
func TestTypeAlias_NestedStructs(t *testing.T) {
	type Inner struct {
		QoS      QoS      `nson:"qos"`
		Priority Priority `nson:"priority"`
	}

	type Outer struct {
		Name  string `nson:"name"`
		Inner Inner  `nson:"inner"`
	}

	original := Outer{
		Name: "outer",
		Inner: Inner{
			QoS:      QoS2,
			Priority: PriorityHigh,
		},
	}

	// Marshal
	m, err := Marshal(&original)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	// Unmarshal
	var decoded Outer
	if err := Unmarshal(m, &decoded); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	// 验证
	if decoded.Name != original.Name {
		t.Errorf("Name mismatch")
	}
	if decoded.Inner.QoS != original.Inner.QoS {
		t.Errorf("Inner.QoS mismatch: got %d, want %d", decoded.Inner.QoS, original.Inner.QoS)
	}
	if decoded.Inner.Priority != original.Inner.Priority {
		t.Errorf("Inner.Priority mismatch: got %d, want %d", decoded.Inner.Priority, original.Inner.Priority)
	}
}

// TestTypeAlias_SliceOfTypeAlias 测试类型别名的切片
func TestTypeAlias_SliceOfTypeAlias(t *testing.T) {
	type Levels struct {
		Priorities []Priority `nson:"priorities"`
		QosList    []QoS      `nson:"qos_list"`
	}

	original := Levels{
		Priorities: []Priority{PriorityLow, PriorityNormal, PriorityHigh},
		QosList:    []QoS{QoS0, QoS1, QoS2},
	}

	// Marshal
	m, err := Marshal(&original)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	// Unmarshal
	var decoded Levels
	if err := Unmarshal(m, &decoded); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	// 验证
	if len(decoded.Priorities) != len(original.Priorities) {
		t.Errorf("Priorities length mismatch")
	}
	for i := range original.Priorities {
		if decoded.Priorities[i] != original.Priorities[i] {
			t.Errorf("Priorities[%d] mismatch: got %d, want %d", i, decoded.Priorities[i], original.Priorities[i])
		}
	}

	if len(decoded.QosList) != len(original.QosList) {
		t.Errorf("QosList length mismatch")
	}
	for i := range original.QosList {
		if decoded.QosList[i] != original.QosList[i] {
			t.Errorf("QosList[%d] mismatch: got %d, want %d", i, decoded.QosList[i], original.QosList[i])
		}
	}
}

// BenchmarkTypeAlias_Marshal 类型别名序列化性能测试
func BenchmarkTypeAlias_Marshal(b *testing.B) {
	msg := Message{
		ID:       "benchmark-test",
		Topic:    "test/benchmark",
		Payload:  []byte("benchmark payload"),
		QoS:      QoS1,
		Priority: PriorityHigh,
		Status:   200,
		Code:     100,
		Metadata: map[string]string{"key": "value"},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := Marshal(&msg)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkTypeAlias_Unmarshal 类型别名反序列化性能测试
func BenchmarkTypeAlias_Unmarshal(b *testing.B) {
	msg := Message{
		ID:       "benchmark-test",
		Topic:    "test/benchmark",
		Payload:  []byte("benchmark payload"),
		QoS:      QoS1,
		Priority: PriorityHigh,
		Status:   200,
		Code:     100,
		Metadata: map[string]string{"key": "value"},
	}

	m, _ := Marshal(&msg)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var decoded Message
		if err := Unmarshal(m, &decoded); err != nil {
			b.Fatal(err)
		}
	}
}
