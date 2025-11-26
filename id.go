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

func (self Id) Tag() Tag {
	return TAG_ID
}

func (self Id) String() string {
	return fmt.Sprintf("Id(%x)", []byte(self))
}

var lastCount = uint32(time.Now().Nanosecond())
var identify = uint32(time.Now().Nanosecond())

// Create unique incrementing Id.
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+
//	|       timestamp       | count |    random     |
//	+---+---+---+---+---+---+---+---+---+---+---+---+
//	  0   1   2   3   4   5   6   7   8   9   10  11
func NewId() Id {
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

	return Id(buf.Bytes())
}

func IdFromHex(s string) (Id, error) {
	if len(s) != 24 {
		return nil, fmt.Errorf("Id hex must be 24 chars: %v", s)
	}

	b, err := hex.DecodeString(s)
	if err != nil {
		return nil, err
	}

	return Id(b), nil
}

func (self Id) Hex() string {
	return hex.EncodeToString([]byte(self))
}

func (self Id) Timestamp() int64 {
	a := uint64(binary.BigEndian.Uint32([]byte(self)[:4]))
	b := uint64(binary.BigEndian.Uint16([]byte(self)[4:6]))

	return int64(a<<16 + b)
}

func (self Id) Time() time.Time {
	ts := self.Timestamp()
	return time.Unix(ts/1000, ts%1000*1000000).UTC()
}

// uint64 to 8 bytes
func Uint64To4Bytes(i uint64) []byte {
	buf := bytes.NewBuffer([]byte{})
	binary.Write(buf, binary.BigEndian, i)
	return buf.Bytes()
}

func EncodeId(value Id, buf *bytes.Buffer) error {
	if len(value) != 12 {
		return fmt.Errorf("Id must be 12 bytes: %v", value)
	}

	if _, err := buf.Write(value); err != nil {
		return err
	}

	return nil
}

func DecodeId(buf *bytes.Buffer) (Id, error) {
	b := make([]byte, 12)
	_, err := io.ReadFull(buf, b)
	if err != nil {
		return nil, err
	}

	return Id(b), nil
}
