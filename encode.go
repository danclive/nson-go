package nson

import (
	"bytes"
	"io"
)

func (m Map) Write(writer io.Writer) error {
	buffer := new(bytes.Buffer)
	err := m.Encode(buffer)
	if err != nil {
		return err
	}

	return writeAll(writer, buffer.Bytes())
}

func (a Array) Write(writer io.Writer) error {
	buffer := new(bytes.Buffer)
	err := a.Encode(buffer)
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
