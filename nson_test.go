package nson

import (
	"bytes"
	"fmt"
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

	// fmt.Println(m)
	// fmt.Println(m2.(m))

	if !reflect.DeepEqual(m, m2) {
		t.Fatal("m2 not equal m")
	}

	by := []byte{185, 0, 0, 0, 2, 97, 3, 1, 0, 0, 0, 2, 98, 4, 2, 0, 0, 0, 0, 0, 0, 0, 2, 99, 2, 0, 0, 0, 0, 0, 0, 8, 64, 2, 100, 7, 5, 0, 0, 0, 52, 2, 101, 12, 8, 0, 0, 0, 1, 2, 3, 4, 2, 116, 13, 123, 0, 0, 0, 0, 0, 0, 0, 2, 105, 14, 1, 133, 181, 236, 53, 28, 253, 160, 205, 235, 243, 159, 2, 106, 12, 8, 0, 0, 0, 5, 6, 7, 8, 2, 107, 6, 123, 0, 0, 0, 0, 0, 0, 0, 2, 108, 8, 19, 0, 0, 0, 5, 200, 1, 0, 0, 6, 21, 3, 0, 0, 0, 0, 0, 0, 0, 2, 109, 9, 23, 0, 0, 0, 2, 97, 3, 111, 0, 0, 0, 2, 98, 4, 222, 0, 0, 0, 0, 0, 0, 0, 0, 2, 110, 11, 2, 111, 9, 12, 0, 0, 0, 2, 78, 5, 123, 0, 0, 0, 0, 2, 112, 8, 15, 0, 0, 0, 3, 111, 0, 0, 0, 3, 222, 0, 0, 0, 0, 0}
	buf2 := bytes.NewBuffer(by)

	m3, err := Map{}.Decode(buf2)
	if err != nil {
		t.Fatal(err)
	}

	_ = m3
	// fmt.Println(m3)

	aa := Map{
		"aa": String("bb"),
		"cc": Array{I32(1), I32(2), I32(3), I32(4)},
	}

	bufaa := new(bytes.Buffer)
	err = aa.Encode(bufaa)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(bufaa.Bytes())
}
