package nson

import (
	"bytes"
	"fmt"
)

func EncodeValue(buf *bytes.Buffer, value Value) error {
	if err := buf.WriteByte(byte(value.DataType())); err != nil {
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
	case U8:
		return EncodeU8(v, buf)
	case U16:
		return EncodeU16(v, buf)
	case I8:
		return EncodeI8(v, buf)
	case I16:
		return EncodeI16(v, buf)
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
		return fmt.Errorf("Unsupported type '%X'", value.DataType())
	}
}

func DecodeValue(buf *bytes.Buffer) (Value, error) {
	tag, err := buf.ReadByte()
	if err != nil {
		return nil, err
	}

	return decodeValueWithTag(buf, DataType(tag))
}

func decodeValueWithTag(buf *bytes.Buffer, tag DataType) (Value, error) {
	switch tag {
	case DataTypeF32:
		value, err := DecodeF32(buf)
		if err != nil {
			return nil, err
		}

		return value, nil
	case DataTypeF64:
		value, err := DecodeF64(buf)
		if err != nil {
			return nil, err
		}

		return value, nil
	case DataTypeI32:
		value, err := DecodeI32(buf)
		if err != nil {
			return nil, err
		}

		return value, nil
	case DataTypeI64:
		value, err := DecodeI64(buf)
		if err != nil {
			return nil, err
		}

		return value, nil
	case DataTypeU32:
		value, err := DecodeU32(buf)
		if err != nil {
			return nil, err
		}

		return value, nil
	case DataTypeU64:
		value, err := DecodeU64(buf)
		if err != nil {
			return nil, err
		}

		return value, nil
	case DataTypeU8:
		value, err := DecodeU8(buf)
		if err != nil {
			return nil, err
		}

		return value, nil
	case DataTypeU16:
		value, err := DecodeU16(buf)
		if err != nil {
			return nil, err
		}

		return value, nil
	case DataTypeI8:
		value, err := DecodeI8(buf)
		if err != nil {
			return nil, err
		}

		return value, nil
	case DataTypeI16:
		value, err := DecodeI16(buf)
		if err != nil {
			return nil, err
		}

		return value, nil
	case DataTypeSTRING:
		value, err := DecodeString(buf)
		if err != nil {
			return nil, err
		}

		return value, nil
	case DataTypeARRAY:
		value, err := DecodeArray(buf)
		if err != nil {
			return nil, err
		}

		return value, nil
	case DataTypeBOOL:
		value, err := DecodeBool(buf)
		if err != nil {
			return nil, err
		}

		return value, nil
	case DataTypeNULL:
		value, err := DecodeNull(buf)
		if err != nil {
			return nil, err
		}

		return value, nil
	case DataTypeBINARY:
		value, err := DecodeBinary(buf)
		if err != nil {
			return nil, err
		}

		return value, nil
	case DataTypeTIMESTAMP:
		value, err := DecodeTimestamp(buf)
		if err != nil {
			return nil, err
		}

		return value, nil
	case DataTypeID:
		value, err := DecodeId(buf)
		if err != nil {
			return nil, err
		}

		return value, nil
	case DataTypeMAP:
		value, err := DecodeMap(buf)
		if err != nil {
			return nil, err
		}

		return value, nil
	default:
		return nil, fmt.Errorf("Unsupported type '%X'", tag)
	}
}
