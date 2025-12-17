package nson

import (
	"bytes"
	"testing"
)

// 测试嵌套结构体的序列化和反序列化
func TestNestedStructWithSlice(t *testing.T) {
	type Inner struct {
		ID   string `nson:"id"`
		Name string `nson:"name"`
	}

	type Outer struct {
		ID    string  `nson:"id"`
		Items []Inner `nson:"items"`
	}

	original := Outer{
		ID: "outer-1",
		Items: []Inner{
			{ID: "inner-1", Name: "Item1"},
			{ID: "inner-2", Name: "Item2"},
		},
	}

	// Marshal
	m, err := Marshal(original)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	// Encode
	buf := new(bytes.Buffer)
	if err := EncodeMap(m, buf); err != nil {
		t.Fatalf("EncodeMap failed: %v", err)
	}

	// Decode
	data := buf.Bytes()
	buf2 := bytes.NewBuffer(data)
	m2, err := DecodeMap(buf2)
	if err != nil {
		t.Fatalf("DecodeMap failed: %v", err)
	}

	// Unmarshal
	var decoded Outer
	if err := Unmarshal(m2, &decoded); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	// Verify
	if decoded.ID != original.ID {
		t.Errorf("ID mismatch: got %v, want %v", decoded.ID, original.ID)
	}
	if len(decoded.Items) != len(original.Items) {
		t.Errorf("Items length mismatch: got %v, want %v", len(decoded.Items), len(original.Items))
	}
	for i := range original.Items {
		if decoded.Items[i].ID != original.Items[i].ID {
			t.Errorf("Item[%d].ID mismatch: got %v, want %v", i, decoded.Items[i].ID, original.Items[i].ID)
		}
		if decoded.Items[i].Name != original.Items[i].Name {
			t.Errorf("Item[%d].Name mismatch: got %v, want %v", i, decoded.Items[i].Name, original.Items[i].Name)
		}
	}
}
