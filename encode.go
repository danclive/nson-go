package nson

import (
	"bytes"
	"io"
)

func WriteMap(writer io.Writer, m Map) error {
	buffer := new(bytes.Buffer)
	err := EncodeMap(m, buffer)
	if err != nil {
		return err
	}

	return writeAll(writer, buffer.Bytes())
}

func WriteArray(writer io.Writer, a Array) error {
	buffer := new(bytes.Buffer)
	err := EncodeArray(a, buffer)
	if err != nil {
		return err
	}

	return writeAll(writer, buffer.Bytes())
}

func writeAll(writer io.Writer, data []byte) error {
	remaining := data
	for len(remaining) > 0 {
		n, err := writer.Write(remaining)
		if err != nil {
			return err
		}
		remaining = remaining[n:]
	}
	return nil
}
