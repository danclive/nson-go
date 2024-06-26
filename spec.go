package nson

const (
	TAG_F32       uint8 = 0x01
	TAG_F64       uint8 = 0x02
	TAG_I32       uint8 = 0x03
	TAG_I64       uint8 = 0x04
	TAG_U32       uint8 = 0x05
	TAG_U64       uint8 = 0x06
	TAG_STRING    uint8 = 0x07
	TAG_ARRAY     uint8 = 0x08
	TAG_MAP       uint8 = 0x09
	TAG_BOOL      uint8 = 0x0A
	TAG_NULL      uint8 = 0x0B
	TAG_BINARY    uint8 = 0x0C
	TAG_TIMESTAMP uint8 = 0x0D
	TAG_ID        uint8 = 0x0E

	// Deprecated: use TAG_MAP
	TAG_MESSAGE uint8 = TAG_MAP
	// Deprecated: use TAG_ID
	TAG_MESSAGE_ID uint8 = TAG_ID
)
