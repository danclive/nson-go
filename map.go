package nson

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
)

// Message
func (self Map) DataType() DataType {
	return DataTypeMAP
}

func (self Map) String() string {
	buf := new(bytes.Buffer)

	fmt.Fprintf(buf, "Map{")

	l := len(self)
	i := 0

	for k, v := range self {
		i++

		if i == l {
			fmt.Fprintf(buf, "%v: %v", k, v.String())
		} else {
			fmt.Fprintf(buf, "%v: %v, ", k, v.String())
		}
	}

	fmt.Fprintf(buf, "}")

	return buf.String()
}

func (self *Map) Get(key string) (Value, bool) {
	value, has := (*self)[key]
	return value, has
}

func (self *Map) Contains(key string) bool {
	_, has := (*self)[key]
	return has
}

func (self *Map) Len() int {
	return len(*self)
}

func (self *Map) Insert(key string, value Value) {
	(*self)[key] = value
}

func (self *Map) Remove(key string) bool {
	has := self.Contains(key)

	if has {
		delete(*self, key)
	}

	return has
}

func (self *Map) Extend(other Map) {
	for k, v := range other {
		(*self)[k] = v
	}
}

func (self *Map) GetF32(key string) (float32, error) {
	value, has := self.Get(key)
	if !has {
		return 0, fmt.Errorf("Not Present, key: %v", key)
	}

	if value.DataType() != DataTypeF32 {
		return 0, fmt.Errorf("Unexpected Type, key: %v, value: %v", key, value)
	}

	return float32(value.(F32)), nil
}

func (self *Map) GetF64(key string) (float64, error) {
	value, has := self.Get(key)
	if !has {
		return 0, fmt.Errorf("Not Present, key: %v", key)
	}

	if value.DataType() != DataTypeF64 {
		return 0, fmt.Errorf("Unexpected Type, key: %v, value: %v", key, value)
	}

	return float64(value.(F64)), nil
}

func (self *Map) GetI32(key string) (int32, error) {
	value, has := self.Get(key)
	if !has {
		return 0, fmt.Errorf("Not Present, key: %v", key)
	}

	if value.DataType() != DataTypeI32 {
		return 0, fmt.Errorf("Unexpected Type, key: %v, value: %v", key, value)
	}

	return int32(value.(I32)), nil
}

func (self *Map) GetU32(key string) (uint32, error) {
	value, has := self.Get(key)
	if !has {
		return 0, fmt.Errorf("Not Present, key: %v", key)
	}

	if value.DataType() != DataTypeU32 {
		return 0, fmt.Errorf("Unexpected Type, key: %v, value: %v", key, value)
	}

	return uint32(value.(U32)), nil
}

func (self *Map) GetI64(key string) (int64, error) {
	value, has := self.Get(key)
	if !has {
		return 0, fmt.Errorf("Not Present, key: %v", key)
	}

	if value.DataType() != DataTypeI64 {
		return 0, fmt.Errorf("Unexpected Type, key: %v, value: %v", key, value)
	}

	return int64(value.(I64)), nil
}

func (self *Map) GetU64(key string) (uint64, error) {
	value, has := self.Get(key)
	if !has {
		return 0, fmt.Errorf("Not Present, key: %v", key)
	}

	if value.DataType() != DataTypeU64 {
		return 0, fmt.Errorf("Unexpected Type, key: %v, value: %v", key, value)
	}

	return uint64(value.(U64)), nil
}

func (self *Map) GetString(key string) (string, error) {
	value, has := self.Get(key)
	if !has {
		return "", fmt.Errorf("Not Present, key: %v", key)
	}

	if value.DataType() != DataTypeSTRING {
		return "", fmt.Errorf("Unexpected Type, key: %v, value: %v", key, value)
	}

	return string(value.(String)), nil
}

func (self *Map) GetArray(key string) (Array, error) {
	value, has := self.Get(key)
	if !has {
		return nil, fmt.Errorf("Not Present, key: %v", key)
	}

	if value.DataType() != DataTypeARRAY {
		return nil, fmt.Errorf("Unexpected Type, key: %v, value: %v", key, value)
	}

	return value.(Array), nil
}

func (self *Map) GetMap(key string) (Map, error) {
	value, has := self.Get(key)
	if !has {
		return nil, fmt.Errorf("Not Present, key: %v", key)
	}

	if value.DataType() != DataTypeMAP {
		return nil, fmt.Errorf("Unexpected Type, key: %v, value: %v", key, value)
	}

	return value.(Map), nil
}

func (self *Map) GetBool(key string) (bool, error) {
	value, has := self.Get(key)
	if !has {
		return false, fmt.Errorf("Not Present, key: %v", key)
	}

	if value.DataType() != DataTypeBOOL {
		return false, fmt.Errorf("Unexpected Type, key: %v, value: %v", key, value)
	}

	return bool(value.(Bool)), nil
}

func (self *Map) IsNull(key string) bool {
	value, has := self.Get(key)
	if !has {
		return false
	}

	if value.DataType() != DataTypeNULL {
		return false
	}

	return true
}

func (self *Map) GetBinary(key string) ([]byte, error) {
	value, has := self.Get(key)
	if !has {
		return nil, fmt.Errorf("Not Present, key: %v", key)
	}

	if value.DataType() != DataTypeBINARY {
		return nil, fmt.Errorf("Unexpected Type, key: %v, value: %v", key, value)
	}

	return []byte(value.(Binary)), nil
}

func (self *Map) GetTimestamp(key string) (int64, error) {
	value, has := self.Get(key)
	if !has {
		return 0, fmt.Errorf("Not Present, key: %v", key)
	}

	if value.DataType() != DataTypeTIMESTAMP {
		return 0, fmt.Errorf("Unexpected Type, key: %v, value: %v", key, value)
	}

	return int64(value.(Timestamp)), nil
}

func (self *Map) GetMapId(key string) (Id, error) {
	value, has := self.Get(key)
	if !has {
		return nil, fmt.Errorf("Not Present, key: %v", key)
	}

	if value.DataType() != DataTypeID {
		return nil, fmt.Errorf("Unexpected Type, key: %v, value: %v", key, value)
	}

	return value.(Id), nil
}

func EncodeMap(m Map, buff *bytes.Buffer) error {
	buf := new(bytes.Buffer)

	if err := writeUint32(buf, 0); err != nil {
		return err
	}

	for k, v := range m {
		if err := writeKey(buf, k); err != nil {
			return err
		}

		if err := EncodeValue(buf, v); err != nil {
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

func DecodeMap(buf *bytes.Buffer) (Map, error) {
	l, err := readUint32(buf)
	if err != nil {
		return nil, err
	}

	if l < MIN_NSON_SIZE {
		return nil, errors.New("Invalid map length")
	}

	if l > MAX_NSON_SIZE {
		return nil, errors.New("Invalid map length")
	}

	msg := Map{}

	for {
		len, err := buf.ReadByte()
		if err != nil {
			return nil, err
		}

		if len == 0 {
			break
		}

		b := make([]byte, len-1)
		if _, err := io.ReadFull(buf, b); err != nil {
			return nil, err
		}

		key := string(b)

		value, err := DecodeValue(buf)
		if err != nil {
			return nil, err
		}

		msg[key] = value
	}

	return msg, nil
}
