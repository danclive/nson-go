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

	_, err = writer.Write(buffer.Bytes())
	if err != nil {
		return err
	}

	return nil
}

func (a Array) Write(writer io.Writer) error {
	buffer := new(bytes.Buffer)
	err := a.Encode(buffer)
	if err != nil {
		return err
	}

	_, err = writer.Write(buffer.Bytes())
	if err != nil {
		return err
	}

	return nil
}
