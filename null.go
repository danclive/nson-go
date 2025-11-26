package nson

import "bytes"

func (self Null) Tag() Tag {
	return TAG_NULL
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
