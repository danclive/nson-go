package nson

import "bytes"

func (self Null) DataType() DataType {
	return DataTypeNULL
}

func (self Null) String() string {
	return "Null"
}

func EncodeNull(buf *bytes.Buffer) error {
	return nil
}

func DecodeNull(buf *bytes.Buffer) (Null, error) {
	return Null{}, nil
}
