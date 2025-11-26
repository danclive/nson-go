package nson

type Tag uint8

const (
	TAG_BOOL      Tag = 0x01
	TAG_NULL      Tag = 0x02
	TAG_F32       Tag = 0x11
	TAG_F64       Tag = 0x12
	TAG_I32       Tag = 0x13
	TAG_I64       Tag = 0x14
	TAG_U32       Tag = 0x15
	TAG_U64       Tag = 0x16
	TAG_U8        Tag = 0x17
	TAG_U16       Tag = 0x18
	TAG_I8        Tag = 0x19
	TAG_I16       Tag = 0x1A
	TAG_STRING    Tag = 0x21
	TAG_BINARY    Tag = 0x22
	TAG_ARRAY     Tag = 0x31
	TAG_MAP       Tag = 0x32
	TAG_TIMESTAMP Tag = 0x41
	TAG_ID        Tag = 0x42
)
