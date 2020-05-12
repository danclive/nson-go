package nson

import (
	"bytes"
)

type Value interface {
	Tag() uint8
	String() string
	Encode(*bytes.Buffer) error
	Decode(*bytes.Buffer) (Value, error)
}

type Message map[string]Value

type F32 float32

type F64 float64

type I32 int32

type I64 int64

type U32 uint32

type U64 uint64

type String string

type Array []Value

type Bool bool

type Null struct{}

type Binary []byte

type Timestamp uint64

type MessageId []byte
