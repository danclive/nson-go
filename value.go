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
		return DecodeF32(buf)
	case DataTypeF64:
		return DecodeF64(buf)
	case DataTypeI32:
		return DecodeI32(buf)
	case DataTypeI64:
		return DecodeI64(buf)
	case DataTypeU32:
		return DecodeU32(buf)
	case DataTypeU64:
		return DecodeU64(buf)
	case DataTypeU8:
		return DecodeU8(buf)
	case DataTypeU16:
		return DecodeU16(buf)
	case DataTypeI8:
		return DecodeI8(buf)
	case DataTypeI16:
		return DecodeI16(buf)
	case DataTypeSTRING:
		return DecodeString(buf)
	case DataTypeARRAY:
		return DecodeArray(buf)
	case DataTypeBOOL:
		return DecodeBool(buf)
	case DataTypeNULL:
		return DecodeNull(buf)
	case DataTypeBINARY:
		return DecodeBinary(buf)
	case DataTypeTIMESTAMP:
		return DecodeTimestamp(buf)
	case DataTypeID:
		return DecodeId(buf)
	case DataTypeMAP:
		return DecodeMap(buf)
	default:
		return nil, fmt.Errorf("Unsupported type '%X'", tag)
	}
}
