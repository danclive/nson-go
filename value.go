package nson

import (
	"bytes"
	"fmt"
)

func decode_value(buf *bytes.Buffer, tag uint8) (Value, error) {
	switch tag {
	case TAG_F32:
		value, err := F32(0).Decode(buf)
		if err != nil {
			return nil, err
		}

		return value, nil
	case TAG_F64:
		value, err := F64(0).Decode(buf)
		if err != nil {
			return nil, err
		}

		return value, nil
	case TAG_I32:
		value, err := I32(0).Decode(buf)
		if err != nil {
			return nil, err
		}

		return value, nil
	case TAG_I64:
		value, err := I64(0).Decode(buf)
		if err != nil {
			return nil, err
		}

		return value, nil
	case TAG_U32:
		value, err := U32(0).Decode(buf)
		if err != nil {
			return nil, err
		}

		return value, nil
	case TAG_U64:
		value, err := U64(0).Decode(buf)
		if err != nil {
			return nil, err
		}

		return value, nil
	case TAG_STRING:
		value, err := String("").Decode(buf)
		if err != nil {
			return nil, err
		}

		return value, nil
	case TAG_ARRAY:
		value, err := Array{}.Decode(buf)
		if err != nil {
			return nil, err
		}

		return value, nil
	case TAG_BOOL:
		value, err := Bool(false).Decode(buf)
		if err != nil {
			return nil, err
		}

		return value, nil
	case TAG_NULL:
		value, err := Null{}.Decode(buf)
		if err != nil {
			return nil, err
		}

		return value, nil
	case TAG_BINARY:
		value, err := Binary{}.Decode(buf)
		if err != nil {
			return nil, err
		}

		return value, nil
	case TAG_TIMESTAMP:
		value, err := Timestamp(0).Decode(buf)
		if err != nil {
			return nil, err
		}

		return value, nil
	case TAG_MESSAGE_ID:
		value, err := MessageId{}.Decode(buf)
		if err != nil {
			return nil, err
		}

		return value, nil
	case TAG_MESSAGE:
		value, err := Message{}.Decode(buf)
		if err != nil {
			return nil, err
		}

		return value, nil
	default:
		return nil, fmt.Errorf("Unsupported type '%X'", tag)
	}
}
