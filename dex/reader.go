package dex

import (
	"encoding/binary"
	"errors"
	"fmt"
	"os"
)

var (
	ErrInvalidHeader   = errors.New("invalid header")
	ErrInvalidStringID = errors.New("invalid string id")
)

const (
	EndianConst        = uint32(0x12345678)
	ReverseEndianConst = uint32(0x78563412)
)

type Reader struct {
	file          *os.File
	Version       string
	ReverseEndian bool

	mapOff      uint32
	StringMaxID uint32
	stringIDOff uint32
	typeIDCount uint32
	typeIDOff   uint32
}

func Open(path string) (*Reader, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	header := make([]byte, 112)

	_, err = f.ReadAt(header[0:40], 0)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrInvalidHeader, err)
	}

	// Validate file magic
	if header[0] != 'd' || header[1] != 'e' || header[2] != 'x' || header[3] != '\n' || header[7] != 0 {
		return nil, fmt.Errorf("%w: invalid magic", ErrInvalidHeader)
	}

	ver := string(header[4:7])

	// TODO: checksum, signature

	headerSize := binary.LittleEndian.Uint32(header[36:40])
	if headerSize < 112 {
		return nil, fmt.Errorf("%w: header too small %d", ErrInvalidHeader, headerSize)
	}

	_, err = f.ReadAt(header[40:112], 40)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrInvalidHeader, err)
	}

	endian := binary.LittleEndian.Uint32(header[40:44])
	reverseEndian := endian == ReverseEndianConst

	r := &Reader{
		file:          f,
		Version:       ver,
		ReverseEndian: reverseEndian,
	}

	r.mapOff = r.toUint(header[52:56])
	r.StringMaxID = r.toUint(header[56:60])
	r.stringIDOff = r.toUint(header[60:64])
	r.typeIDCount = r.toUint(header[64:68])
	r.typeIDOff = r.toUint(header[68:72])

	return r, nil
}

func (r *Reader) Close() error {
	return r.file.Close()
}

func (r *Reader) GetString(id uint32) (string, error) {
	if id >= r.StringMaxID {
		return "", ErrInvalidStringID
	}

	idPos := r.stringIDOff + (id * 4)

	strPos, err := r.readUint(idPos)
	if err != nil {
		return "", err
	}

	strSize, n, err := r.readLeb128(strPos)
	if err != nil {
		return "", err
	}

	data := make([]byte, strSize)

	_, err = r.file.ReadAt(data, int64(strPos+n))
	if err != nil {
		return "", err
	}

	return string(data), nil
}
