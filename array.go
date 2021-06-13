package nson

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// Array
func (self Array) Tag() uint8 {
	return TAG_ARRAY
}

func (self Array) String() string {
	buf := new(bytes.Buffer)

	fmt.Fprintf(buf, "Array[")

	l := len(self)

	for i, v := range self {
		if i == l-1 {
			fmt.Fprintf(buf, "%v", v.String())
		} else {
			fmt.Fprintf(buf, "%v, ", v.String())
		}
	}

	fmt.Fprintf(buf, "]")

	return buf.String()
}

func (self Array) Encode(buff *bytes.Buffer) error {
	buf := new(bytes.Buffer)

	if err := writeUint32(buf, 0); err != nil {
		return err
	}

	for _, v := range self {
		if err := buf.WriteByte(v.Tag()); err != nil {
			return err
		}

		if err := v.Encode(buf); err != nil {
			return err
		}
	}

	if err := buf.WriteByte(0x00); err != nil {
		return err
	}

	binary.LittleEndian.PutUint32(buf.Bytes(), uint32(buf.Len()))

	if _, err := buf.WriteTo(buff); err != nil {
		return err
	}

	return nil
}

func (self Array) Decode(buf *bytes.Buffer) (Value, error) {
	_, err := readUint32(buf)
	if err != nil {
		return nil, err
	}

	array := Array{}

	for {
		tag, err := buf.ReadByte()
		if err != nil {
			return nil, err
		}

		if tag == 0 {
			break
		}

		value, err := DecodeValue(buf, tag)
		if err != nil {
			return nil, err
		}

		array = append(array, value)
	}

	return array, nil
}

func (self *Array) Push(value Value) {
	*self = append(*self, value)
}

func (self Array) IntoValue() Value {
	return Value(self)
}
