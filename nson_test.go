package nson

import (
	"bytes"
	"reflect"
	"testing"
)

func TestEncodeDecode(t *testing.T) {
	mid := NewId()

	m := Map{
		"a": F32(123.123),
		"b": F64(456.456),
		"c": Map{
			"d": F64(789.789),
		},
		"e": I32(1),
		"f": I64(2),
		"g": U32(3),
		"h": U64(4),
		"i": String("aaa"),
		"j": Array{F32(666.777), String("hello")},
		"k": Bool(false),
		"l": Null{},
		"m": Binary{1, 2, 3, 4, 5, 6},
		"n": Timestamp(12345),
		"p": mid,
	}

	buf := new(bytes.Buffer)

	err := m.Encode(buf)
	if err != nil {
		t.Fatal(err)
	}

	m2, err := Map{}.Decode(buf)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(m, m2) {
		t.Fatal("m2 not equal m")
	}

	aa := Map{
		"aa": String("bb"),
		"cc": Array{I32(1), I32(2), I32(3), I32(4)},
	}

	bufaa := new(bytes.Buffer)
	err = aa.Encode(bufaa)
	if err != nil {
		t.Fatal(err)
	}

}

func TestEncodeDecode2(t *testing.T) {
	m := Map{
		"a": F32(123.123),
		"b": F64(456.456),
		"c": Map{
			"d": F64(789.789),
		},
		"e": I32(1),
		"f": I64(2),
		"g": U32(3),
		"h": U64(4),
		"i": String("aaa"),
		"j": Array{F32(666.777), String("hello")},
		"k": Bool(false),
		"l": Null{},
		"m": Binary{1, 2, 3, 4, 5, 6},
		"n": Timestamp(12345),
		"p": NewId(),
	}

	buff := new(bytes.Buffer)

	m.Write(buff)

	bts := buff.Bytes()

	buff2 := bytes.NewBuffer(bts)

	m2, err := Map{}.Read(buff2)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(m, m2) {
		t.Fatal("m2 not equal m")
	}
}
