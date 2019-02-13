package nson

import (
	"bytes"
	"fmt"
)

func (self Boolean) Tag() uint8 {
	return TAG_BOOLEAN
}

func (self Boolean) String() string {
	if self {
		return fmt.Sprint("True")
	} else {
		return fmt.Sprint("False")
	}
}

func (self Boolean) Encode(buf *bytes.Buffer) error {
	if self {
		if err := buf.WriteByte(0x01); err != nil {
			return err
		}
	} else {
		if err := buf.WriteByte(0x00); err != nil {
			return err
		}
	}

	return nil
}

func (self Boolean) Decode(buf *bytes.Buffer) (Value, error) {
	b, err := buf.ReadByte()
	if err != nil {
		return nil, err
	}

	return Boolean(b == 0x01), nil
}
