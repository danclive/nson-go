package nson

const (
	TAG_BOOL      uint8 = 0x01
	TAG_NULL      uint8 = 0x02
	TAG_F32       uint8 = 0x11
	TAG_F64       uint8 = 0x12
	TAG_I32       uint8 = 0x13
	TAG_I64       uint8 = 0x14
	TAG_U32       uint8 = 0x15
	TAG_U64       uint8 = 0x16
	TAG_STRING    uint8 = 0x21
	TAG_BINARY    uint8 = 0x22
	TAG_ARRAY     uint8 = 0x31
	TAG_MAP       uint8 = 0x32
	TAG_TIMESTAMP uint8 = 0x41
	TAG_ID        uint8 = 0x42
)
