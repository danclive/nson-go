package nson

import (
	"bytes"
	"testing"
)

func TestEmptyStringInStruct(t *testing.T) {
	type TestStruct struct {
		Name  string `nson:"name"`
		Value string `nson:"value"`
	}

	original := TestStruct{
		Name:  "test",
		Value: "", // 空字符串
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
	t.Logf("Encoded data: %x", data)

	buf2 := bytes.NewBuffer(data)
	m2, err := DecodeMap(buf2)
	if err != nil {
		t.Fatalf("DecodeMap failed: %v", err)
	}

	// Unmarshal
	var decoded TestStruct
	if err := Unmarshal(m2, &decoded); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	// Verify
	if decoded.Name != original.Name {
		t.Errorf("Name mismatch: got %v, want %v", decoded.Name, original.Name)
	}
	if decoded.Value != original.Value {
		t.Errorf("Value mismatch: got %v, want %v", decoded.Value, original.Value)
	}
}
