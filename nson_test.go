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
		//"o": UTCDateTime(time.Now().UTC()),
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

	if !reflect.DeepEqual(message, message2) {
		t.Fatal("message2 not equal message")
	}
}
