package nson

import (
	"bytes"
	"fmt"
)

func EncodeValue(buf *bytes.Buffer, value Value) error {
	if err := buf.WriteByte(value.Tag()); err != nil {
		return err
	}

	switch v := value.(type) {
	case F32:
		return EncodeF32(v, buf)
	case F64:
		return EncodeF64(v, buf)
	case I32:
		return EncodeI32(v, buf)
	case I64:
		return EncodeI64(v, buf)
	case U32:
		return EncodeU32(v, buf)
	case U64:
		return EncodeU64(v, buf)
	case String:
		return EncodeString(v, buf)
	case Array:
		return EncodeArray(v, buf)
	case Bool:
		return EncodeBool(v, buf)
	case Null:
		return EncodeNull(buf)
	case Binary:
		return EncodeBinary(v, buf)
	case Timestamp:
		return EncodeTimestamp(v, buf)
	case Id:
		return EncodeId(v, buf)
	case Map:
		return EncodeMap(v, buf)
	default:
		return fmt.Errorf("Unsupported type '%X'", value.Tag())
	}
}

func DecodeValue(buf *bytes.Buffer) (Value, error) {
	tag, err := buf.ReadByte()
	if err != nil {
		return nil, err
	}

	return decodeValueWithTag(buf, tag)
}

func decodeValueWithTag(buf *bytes.Buffer, tag uint8) (Value, error) {
	switch tag {
	case TAG_F32:
		value, err := DecodeF32(buf)
		if err != nil {
			return nil, err
		}

		return value, nil
	case TAG_F64:
		value, err := DecodeF64(buf)
		if err != nil {
			return nil, err
		}

		return value, nil
	case TAG_I32:
		value, err := DecodeI32(buf)
		if err != nil {
			return nil, err
		}

		return value, nil
	case TAG_I64:
		value, err := DecodeI64(buf)
		if err != nil {
			return nil, err
		}

		return value, nil
	case TAG_U32:
		value, err := DecodeU32(buf)
		if err != nil {
			return nil, err
		}

		return value, nil
	case TAG_U64:
		value, err := DecodeU64(buf)
		if err != nil {
			return nil, err
		}

		return value, nil
	case TAG_STRING:
		value, err := DecodeString(buf)
		if err != nil {
			return nil, err
		}

		return value, nil
	case TAG_ARRAY:
		value, err := DecodeArray(buf)
		if err != nil {
			return nil, err
		}

		return value, nil
	case TAG_BOOL:
		value, err := DecodeBool(buf)
		if err != nil {
			return nil, err
		}

		return value, nil
	case TAG_NULL:
		value, err := DecodeNull(buf)
		if err != nil {
			return nil, err
		}

		return value, nil
	case TAG_BINARY:
		value, err := DecodeBinary(buf)
		if err != nil {
			return nil, err
		}

		return value, nil
	case TAG_TIMESTAMP:
		value, err := DecodeTimestamp(buf)
		if err != nil {
			return nil, err
		}

		return value, nil
	case TAG_ID:
		value, err := DecodeId(buf)
		if err != nil {
			return nil, err
		}

		return value, nil
	case TAG_MAP:
		value, err := DecodeMap(buf)
		if err != nil {
			return nil, err
		}

		return value, nil
	default:
		return nil, fmt.Errorf("Unsupported type '%X'", tag)
	}
}
