package nson

import (
	"bytes"
	"errors"
	"fmt"
	"io"
)

func (self Binary) Tag() uint8 {
	return TAG_BINARY
}

func (self Binary) String() string {
	return fmt.Sprintf("Binary(%v)", []byte(self))
}

func (self Binary) Encode(buf *bytes.Buffer) error {
	if err := writeUint32(buf, uint32(len(self)+4)); err != nil {
		return nil
	}

	if _, err := buf.Write(self); err != nil {
		return err
	}

	return nil
}

func (self Binary) Decode(buf *bytes.Buffer) (Value, error) {
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
