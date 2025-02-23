package nson

import (
	"bytes"
	"fmt"
)

// Timestamp
func (self Timestamp) Tag() uint8 {
	return TAG_TIMESTAMP
}

func (self Timestamp) String() string {
	return fmt.Sprintf("Timestamp(%v)", uint64(self))
}

func EncodeTimestamp(value Timestamp, buf *bytes.Buffer) error {
	return writeUint64(buf, uint64(value))
}

func DecodeTimestamp(buf *bytes.Buffer) (Timestamp, error) {
	v, err := readUint64(buf)
	if err != nil {
		return 0, err
	}

	return Timestamp(v), nil
}
