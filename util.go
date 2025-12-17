package nson

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
)

func writeKey(buf *bytes.Buffer, s string) error {
	if len(s) == 0 || len(s) >= 255 {
		return errors.New("Key len must > 0 and < 255")
	}

	if err := buf.WriteByte(uint8(len(s) + 1)); err != nil {
		return err
	}
	if _, err := buf.WriteString(s); err != nil {
		return err
	}

	return nil
}

func writeString(buf *bytes.Buffer, s string) error {
	if err := binary.Write(buf, binary.LittleEndian, uint32(len(s)+4)); err != nil {
		return err
	}
	if _, err := buf.WriteString(s); err != nil {
		return err
	}
	return nil
}

func writeFloat32(buf *bytes.Buffer, f float32) error {
	return binary.Write(buf, binary.LittleEndian, f)
}

func writeFloat64(buf *bytes.Buffer, f float64) error {
	return binary.Write(buf, binary.LittleEndian, f)
}

func writeInt32(buf *bytes.Buffer, i int32) error {
	return binary.Write(buf, binary.LittleEndian, i)
}

func writeInt64(buf *bytes.Buffer, i int64) error {
	return binary.Write(buf, binary.LittleEndian, i)
}

func writeUint32(buf *bytes.Buffer, u uint32) error {
	return binary.Write(buf, binary.LittleEndian, u)
}

func writeUint64(buf *bytes.Buffer, u uint64) error {
	return binary.Write(buf, binary.LittleEndian, u)
}

func writeInt16(buf *bytes.Buffer, i int16) error {
	return binary.Write(buf, binary.LittleEndian, i)
}

func writeUint16(buf *bytes.Buffer, u uint16) error {
	return binary.Write(buf, binary.LittleEndian, u)
}

// func readCstring(rd *bytes.Buffer) (string, error) {
// 	s, err := rd.ReadString(0x00)
// 	if err != nil {
// 		return "", err
// 	}
// 	return s[:len(s)-1], nil
// }

func readString(rd *bytes.Buffer) (string, error) {
	// Read string length.
	var l uint32
	if err := binary.Read(rd, binary.LittleEndian, &l); err != nil {
		return "", err
	}

	if l < MIN_NSON_SIZE-1 {
		return "", errors.New("Invalid string length")
	}

	if l > MAX_NSON_SIZE {
		return "", errors.New("Invalid string length")
	}

	// Read string.
	b := make([]byte, l-4)
	if _, err := io.ReadFull(rd, b); err != nil {
		return "", err
	}
	return string(b), nil
}

func readFloat32(rd *bytes.Buffer) (float32, error) {
	var f float32
	if err := binary.Read(rd, binary.LittleEndian, &f); err != nil {
		return 0, err
	}
	return f, nil
}

func readFloat64(rd *bytes.Buffer) (float64, error) {
	var f float64
	if err := binary.Read(rd, binary.LittleEndian, &f); err != nil {
		return 0, err
	}
	return f, nil
}

func readInt32(rd *bytes.Buffer) (int32, error) {
	var i int32
	if err := binary.Read(rd, binary.LittleEndian, &i); err != nil {
		return 0, err
	}
	return i, nil
}

func readInt64(rd *bytes.Buffer) (int64, error) {
	var i int64
	if err := binary.Read(rd, binary.LittleEndian, &i); err != nil {
		return 0, err
	}
	return i, nil
}

func readUint32(rd *bytes.Buffer) (uint32, error) {
	var u uint32
	if err := binary.Read(rd, binary.LittleEndian, &u); err != nil {
		return 0, err
	}
	return u, nil
}

func readUint64(rd *bytes.Buffer) (uint64, error) {
	var u uint64
	if err := binary.Read(rd, binary.LittleEndian, &u); err != nil {
		return 0, err
	}
	return u, nil
}

func readInt16(rd *bytes.Buffer) (int16, error) {
	var i int16
	if err := binary.Read(rd, binary.LittleEndian, &i); err != nil {
		return 0, err
	}
	return i, nil
}

func readUint16(rd *bytes.Buffer) (uint16, error) {
	var u uint16
	if err := binary.Read(rd, binary.LittleEndian, &u); err != nil {
		return 0, err
	}
	return u, nil
}
