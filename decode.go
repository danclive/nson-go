package nson

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
)

func ReadMap(reader io.Reader) (Map, error) {
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
	m, err := DecodeMap(dataBuffer)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func ReadArray(reader io.Reader) (Array, error) {
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
	array, err := DecodeArray(dataBuffer)
	if err != nil {
		return nil, err
	}

	return array, nil
}
