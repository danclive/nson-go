package nson

type DataType uint8

const (
	DataTypeBOOL      DataType = 0x01
	DataTypeNULL      DataType = 0x02
	DataTypeF32       DataType = 0x11
	DataTypeF64       DataType = 0x12
	DataTypeI32       DataType = 0x13
	DataTypeI64       DataType = 0x14
	DataTypeU32       DataType = 0x15
	DataTypeU64       DataType = 0x16
	DataTypeI8        DataType = 0x17
	DataTypeU8        DataType = 0x18
	DataTypeI16       DataType = 0x19
	DataTypeU16       DataType = 0x1A
	DataTypeSTRING    DataType = 0x21
	DataTypeBINARY    DataType = 0x22
	DataTypeARRAY     DataType = 0x31
	DataTypeMAP       DataType = 0x32
	DataTypeTIMESTAMP DataType = 0x41
	DataTypeID        DataType = 0x42
)

func (dt DataType) IsPrimitive() bool {
	return dt >= DataTypeBOOL && dt <= DataTypeU16
}

func (dt DataType) IsComplex() bool {
	return dt == DataTypeARRAY || dt == DataTypeMAP
}

func (dt DataType) IsFixedSize() bool {
	return dt >= DataTypeBOOL && dt <= DataTypeU16
}

func (dt DataType) IsVariableSize() bool {
	return dt == DataTypeSTRING || dt == DataTypeBINARY
}

func (dt DataType) IsSpecial() bool {
	return dt == DataTypeTIMESTAMP || dt == DataTypeID
}

func (dt DataType) Size() int {
	switch dt {
	case DataTypeBOOL:
		return 1
	case DataTypeI8, DataTypeU8:
		return 1
	case DataTypeI16, DataTypeU16:
		return 2
	case DataTypeI32, DataTypeU32, DataTypeF32:
		return 4
	case DataTypeI64, DataTypeU64, DataTypeF64:
		return 8
	case DataTypeTIMESTAMP:
		return 8
	case DataTypeID:
		return 16
	default:
		return -1 // Variable size or unknown
	}
}

func (dt DataType) ZeroValue() Value {
	switch dt {
	case DataTypeBOOL:
		return Bool(false)
	case DataTypeNULL:
		return Null{}
	case DataTypeF32:
		return F32(0)
	case DataTypeF64:
		return F64(0)
	case DataTypeI32:
		return I32(0)
	case DataTypeI64:
		return I64(0)
	case DataTypeU32:
		return U32(0)
	case DataTypeU64:
		return U64(0)
	case DataTypeI8:
		return I8(0)
	case DataTypeU8:
		return U8(0)
	case DataTypeI16:
		return I16(0)
	case DataTypeU16:
		return U16(0)
	case DataTypeSTRING:
		return String("")
	case DataTypeBINARY:
		return Binary([]byte{})
	case DataTypeARRAY:
		return Array([]Value{})
	case DataTypeMAP:
		return Map(make(map[string]Value))
	case DataTypeTIMESTAMP:
		return Timestamp(0)
	case DataTypeID:
		return Id(make([]byte, 12))
	default:
		return nil
	}
}
