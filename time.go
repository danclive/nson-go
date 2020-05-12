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

func (self Timestamp) Encode(buf *bytes.Buffer) error {
	return writeUint64(buf, uint64(self))
}

func (self Timestamp) Decode(buf *bytes.Buffer) (Value, error) {
	v, err := readUint64(buf)
	if err != nil {
		return nil, err
	}

	return Timestamp(v), nil
}
