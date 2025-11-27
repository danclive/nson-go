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

// Deprecated: Use DataType* constants instead
const (
	TAG_BOOL      = DataTypeBOOL
	TAG_NULL      = DataTypeNULL
	TAG_F32       = DataTypeF32
	TAG_F64       = DataTypeF64
	TAG_I32       = DataTypeI32
	TAG_I64       = DataTypeI64
	TAG_U32       = DataTypeU32
	TAG_U64       = DataTypeU64
	TAG_I8        = DataTypeI8
	TAG_U8        = DataTypeU8
	TAG_I16       = DataTypeI16
	TAG_U16       = DataTypeU16
	TAG_STRING    = DataTypeSTRING
	TAG_BINARY    = DataTypeBINARY
	TAG_ARRAY     = DataTypeARRAY
	TAG_MAP       = DataTypeMAP
	TAG_TIMESTAMP = DataTypeTIMESTAMP
	TAG_ID        = DataTypeID
)
