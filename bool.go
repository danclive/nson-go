package nson

import (
	"bytes"
)

func (self Bool) DataType() DataType {
	return DataTypeBOOL
}

func (self Bool) String() string {
	if self {
		return "True"
	}

	return "False"
}

func EncodeBool(value Bool, buf *bytes.Buffer) error {
	if value {
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

func DecodeBool(buf *bytes.Buffer) (Bool, error) {
	b, err := buf.ReadByte()
	if err != nil {
		return false, err
	}

	return Bool(b == 0x01), nil
}
