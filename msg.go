package nson

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// Message
func (self Message) Tag() uint8 {
	return TAG_MESSAGE
}

func (self Message) String() string {
	buf := new(bytes.Buffer)

	fmt.Fprintf(buf, "Message{")

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

func (self Message) Encode(buff *bytes.Buffer) error {
	buf := new(bytes.Buffer)

	if err := writeInt32(buf, 0); err != nil {
		return err
	}

	for k, v := range self {
		if err := buf.WriteByte(v.Tag()); err != nil {
			return err
		}

		if err := writeCstring(buf, k); err != nil {
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

func (self Message) Decode(buf *bytes.Buffer) (Value, error) {
	_, err := readInt32(buf)
	if err != nil {
		return nil, err
	}

	msg := Message{}

	for {
		tag, err := buf.ReadByte()
		if err != nil {
			return nil, err
		}

		if tag == 0 {
			break
		}

		key, value, err := decode_value(buf, tag)
		if err != nil {
			return nil, err
		}

		msg[key] = value
	}

	return msg, nil
}
