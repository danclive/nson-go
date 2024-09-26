package nson

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
)

func (m Map) Read(reader io.Reader) (Map, error) {
	lengthBytes := make([]byte, 4)
	if _, err := io.ReadFull(reader, lengthBytes); err != nil {
		return nil, err
	}

	fmt.Println(lengthBytes)

	dataLength := binary.LittleEndian.Uint32(lengthBytes)
	if dataLength < 5 || dataLength > MAX_NSON_SIZE {
		return nil, errors.New("invalid data length")
	}

	fullData := make([]byte, dataLength)
	if _, err := io.ReadFull(reader, fullData[4:]); err != nil {
		return nil, err
	}

	copy(fullData[:4], lengthBytes)

	dataBuffer := bytes.NewBuffer(fullData)
	decodedValue, err := Map{}.Decode(dataBuffer)
	if err != nil {
		return nil, err
	}

	return decodedValue.(Map), nil
}

func (a Array) Read(reader io.Reader) (Array, error) {
	lengthBytes := make([]byte, 4)
	if _, err := io.ReadFull(reader, lengthBytes); err != nil {
		return nil, err
	}

	dataLength := binary.LittleEndian.Uint32(lengthBytes)
	if dataLength < 5 || dataLength > MAX_NSON_SIZE {
		return nil, errors.New("invalid data length")
	}

	fullData := make([]byte, dataLength)
	if _, err := io.ReadFull(reader, fullData[4:]); err != nil {
		return nil, err
	}

	copy(fullData[:4], lengthBytes)

	dataBuffer := bytes.NewBuffer(fullData)
	decodedValue, err := Array{}.Decode(dataBuffer)
	if err != nil {
		return nil, err
	}

	return decodedValue.(Array), nil
}
