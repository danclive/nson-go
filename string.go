package nson

import (
	"bytes"
	"fmt"
)

// String
func (self String) Tag() uint8 {
	return TAG_STRING
}

func (self String) String() string {
	return fmt.Sprintf("String(%v)", string(self))
}

func EncodeString(value String, buf *bytes.Buffer) error {
	return writeString(buf, string(value))
}

func DecodeString(buf *bytes.Buffer) (String, error) {
	v, err := readString(buf)
	if err != nil {
		return "", err
	}

	return String(v), nil
}
