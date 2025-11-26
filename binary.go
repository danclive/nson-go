package nson

import (
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
)

func (self Binary) Tag() Tag {
	return TAG_BINARY
}

func (self Binary) String() string {
	return fmt.Sprintf("Binary(%v)", []byte(self))
}

func (self Binary) Hex() string {
	return hex.EncodeToString([]byte(self))
}

func BinaryFromHex(s string) (Binary, error) {
	b, err := hex.DecodeString(s)
	if err != nil {
		return nil, err
	}
	return Binary(b), nil
}

func EncodeBinary(binary Binary, buf *bytes.Buffer) error {
	if err := writeUint32(buf, uint32(len(binary)+4)); err != nil {
		return nil
	}

	if _, err := buf.Write(binary); err != nil {
		return err
	}

	return nil
}

func DecodeBinary(buf *bytes.Buffer) (Binary, error) {
	l, err := readUint32(buf)
	if err != nil {
		return nil, err
	}

	if l < MIN_NSON_SIZE {
		return nil, errors.New("Invalid binary length")
	}

	if l > MAX_NSON_SIZE {
		return nil, errors.New("Invalid binary length")
	}

	b := make([]byte, l-4)
	if _, err := io.ReadFull(buf, b); err != nil {
		return nil, err
	}

	return Binary(b), nil
}
