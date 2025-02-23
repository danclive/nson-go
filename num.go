package nson

import (
	"bytes"
	"fmt"
)

// Float 32
func (self F32) Tag() uint8 {
	return TAG_F32
}

func (self F32) String() string {
	return fmt.Sprintf("F32(%v)", float32(self))
}

func EncodeF32(value F32, buf *bytes.Buffer) error {
	return writeFloat32(buf, float32(value))
}

func DecodeF32(buf *bytes.Buffer) (F32, error) {
	v, err := readFloat32(buf)
	if err != nil {
		return 0, err
	}

	return F32(v), nil
}

// Float 64
func (self F64) Tag() uint8 {
	return TAG_F64
}

func (self F64) String() string {
	return fmt.Sprintf("F32(%v)", float64(self))
}

func EncodeF64(value F64, buf *bytes.Buffer) error {
	return writeFloat64(buf, float64(value))
}

func DecodeF64(buf *bytes.Buffer) (F64, error) {
	v, err := readFloat64(buf)
	if err != nil {
		return 0, err
	}

	return F64(v), nil
}

// Int 32
func (self I32) Tag() uint8 {
	return TAG_I32
}

func (self I32) String() string {
	return fmt.Sprintf("I32(%v)", int32(self))
}

func EncodeI32(value I32, buf *bytes.Buffer) error {
	return writeInt32(buf, int32(value))
}

func DecodeI32(buf *bytes.Buffer) (I32, error) {
	v, err := readInt32(buf)
	if err != nil {
		return 0, err
	}

	return I32(v), nil
}

// Int 64
func (self I64) Tag() uint8 {
	return TAG_I64
}

func (self I64) String() string {
	return fmt.Sprintf("I64(%v)", int64(self))
}

func EncodeI64(value I64, buf *bytes.Buffer) error {
	return writeInt64(buf, int64(value))
}

func DecodeI64(buf *bytes.Buffer) (I64, error) {
	v, err := readInt64(buf)
	if err != nil {
		return 0, err
	}

	return I64(v), nil
}

// Uint 32
func (self U32) Tag() uint8 {
	return TAG_U32
}

func (self U32) String() string {
	return fmt.Sprintf("U32(%v)", uint32(self))
}

func EncodeU32(value U32, buf *bytes.Buffer) error {
	return writeUint32(buf, uint32(value))
}

func DecodeU32(buf *bytes.Buffer) (U32, error) {
	v, err := readUint32(buf)
	if err != nil {
		return 0, err
	}

	return U32(v), nil
}

// Uint 64
func (self U64) Tag() uint8 {
	return TAG_U64
}

func (self U64) String() string {
	return fmt.Sprintf("U64(%v)", uint64(self))
}

func EncodeU64(value U64, buf *bytes.Buffer) error {
	return writeUint64(buf, uint64(value))
}

func DecodeU64(buf *bytes.Buffer) (U64, error) {
	v, err := readUint64(buf)
	if err != nil {
		return 0, err
	}

	return U64(v), nil
}
