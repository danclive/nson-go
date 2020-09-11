package nson

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io"
	"math/rand"
	"sync/atomic"
	"time"
)

func (self MessageId) Tag() uint8 {
	return TAG_MESSAGE_ID
}

func (self MessageId) String() string {
	return fmt.Sprintf("MessageId(%x)", []byte(self))
}

func (self MessageId) Encode(buf *bytes.Buffer) error {
	if len(self) != 12 {
		return fmt.Errorf("MessageId must be 12 bytes: %v", self)
	}

	if _, err := buf.Write(self); err != nil {
		return err
	}

	return nil
}

func (self MessageId) Decode(buf *bytes.Buffer) (Value, error) {
	b := make([]byte, 12)
	_, err := io.ReadFull(buf, b)
	if err != nil {
		return nil, err
	}

	return MessageId(b), nil
}

var lastCount = uint32(time.Now().Nanosecond())
var identify = uint32(time.Now().Nanosecond())

// Create unique incrementing MessageId.
//
//   +---+---+---+---+---+---+---+---+---+---+---+---+
//   |       timestamp       | count |    random     |
//   +---+---+---+---+---+---+---+---+---+---+---+---+
//     0   1   2   3   4   5   6   7   8   9   10  11
func NewMessageId() MessageId {
	buf := new(bytes.Buffer)

	// timestamp
	now := time.Now().UnixNano() / 1000000
	now_bytes := Uint64To4Bytes(uint64(now))
	binary.Write(buf, binary.BigEndian, now_bytes[2:])

	// count
	cnt := atomic.AddUint32(&lastCount, 1) % 65536
	binary.Write(buf, binary.BigEndian, uint16(cnt))

	// random
	random := rand.Uint32()
	binary.Write(buf, binary.BigEndian, random)

	return MessageId(buf.Bytes())
}

func MessageIdFromHex(s string) (MessageId, error) {
	if len(s) != 24 {
		return nil, fmt.Errorf("MessageId hex must be 24 chars: %v", s)
	}

	b, err := hex.DecodeString(s)
	if err != nil {
		return nil, err
	}

	return MessageId(b), nil
}

func (self MessageId) Hex() string {
	return hex.EncodeToString([]byte(self))
}

func (self MessageId) Timestamp() int64 {
	a := uint64(binary.BigEndian.Uint32([]byte(self)[:4]))
	b := uint64(binary.BigEndian.Uint16([]byte(self)[4:6]))

	return int64(a<<16 + b)
}

func (self MessageId) Time() time.Time {
	return time.Unix(self.Timestamp(), 0).UTC()
}

func (self MessageId) Counter() uint16 {
	return uint16(binary.BigEndian.Uint16([]byte(self)[6:8]))
}

func (self MessageId) Random() uint32 {
	return uint32(binary.BigEndian.Uint32([]byte(self)[8:]))
}

// uint64 to 8 bytes
func Uint64To4Bytes(i uint64) []byte {
	buf := bytes.NewBuffer([]byte{})
	binary.Write(buf, binary.BigEndian, i)
	return buf.Bytes()
}
