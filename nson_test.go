package nson

import (
	"bytes"
	"reflect"
	"testing"
)

func TestEncodeDecode(t *testing.T) {
	mid := NewMessageId()

	message := Message{
		"a": F32(123.123),
		"b": F64(456.456),
		"c": Message{
			"d": F64(789.789),
		},
		"e": I32(1),
		"f": I64(2),
		"g": U32(3),
		"h": U64(4),
		"i": String("aaa"),
		"j": Array{F32(666.777)},
		"k": Bool(false),
		"l": Null{},
		"m": Binary{1, 2, 3, 4, 5, 6},
		"n": Timestamp(12345),
		"p": mid,
	}

	buf := new(bytes.Buffer)

	err := message.Encode(buf)
	if err != nil {
		t.Fatal(err)
	}

	message2, err := Message{}.Decode(buf)
	if err != nil {
		t.Fatal(err)
	}

	// fmt.Println(message)
	// fmt.Println(message2.(Message))

	if !reflect.DeepEqual(message, message2) {
		t.Fatal("message2 not equal message")
	}

	by := []byte{185, 0, 0, 0, 3, 97, 0, 1, 0, 0, 0, 4, 98, 0, 2, 0, 0, 0, 0, 0, 0, 0, 2, 99, 0, 0, 0, 0, 0, 0, 0, 8, 64, 7, 100, 0, 5, 0, 0, 0, 52, 12, 101, 0, 8, 0, 0, 0, 1, 2, 3, 4, 13, 116, 0, 123, 0, 0, 0, 0, 0, 0, 0, 14, 105, 0, 1, 116, 244, 64, 8, 143, 118, 221, 140, 219, 56, 116, 12, 106, 0, 8, 0, 0, 0, 5, 6, 7, 8, 6, 107, 0, 123, 0, 0, 0, 0, 0, 0, 0, 8, 108, 0, 19, 0, 0, 0, 5, 200, 1, 0, 0, 6, 21, 3, 0, 0, 0, 0, 0, 0, 0, 9, 109, 0, 23, 0, 0, 0, 3, 97, 0, 111, 0, 0, 0, 4, 98, 0, 222, 0, 0, 0, 0, 0, 0, 0, 0, 11, 110, 0, 9, 111, 0, 12, 0, 0, 0, 5, 78, 0, 123, 0, 0, 0, 0, 8, 112, 0, 15, 0, 0, 0, 3, 111, 0, 0, 0, 3, 222, 0, 0, 0, 0, 0}
	buf2 := bytes.NewBuffer(by)

	message3, err := Message{}.Decode(buf2)
	if err != nil {
		t.Fatal(err)
	}

	_ = message3
	//fmt.Println(message3)
}
