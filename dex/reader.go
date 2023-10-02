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
	ErrInvalidTypeID   = errors.New("invalid type id")
	ErrInvalidProtoID  = errors.New("invalid proto id")
)

const (
	EndianConst        = uint32(0x12345678)
	ReverseEndianConst = uint32(0x78563412)
)

type Reader struct {
	file          *os.File
	Version       string
	ReverseEndian bool

	mapOff        uint32
	StringIDCount uint32
	stringIDOff   uint32
	TypeIDCount   uint32
	typeIDOff     uint32
	ProtoIDCount  uint32
	protoIDOff    uint32
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
	r.StringIDCount = r.toUint(header[56:60])
	r.stringIDOff = r.toUint(header[60:64])
	r.TypeIDCount = r.toUint(header[64:68])
	r.typeIDOff = r.toUint(header[68:72])
	r.ProtoIDCount = r.toUint(header[72:76])
	r.protoIDOff = r.toUint(header[76:80])

	return r, nil
}

func (r *Reader) Close() error {
	return r.file.Close()
}

func (r *Reader) GetString(id uint32) (string, error) {
	if id >= r.StringIDCount {
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

type Type struct {
	DescriptorStringID uint32
}

func (r *Reader) GetType(id uint32) (Type, error) {
	var (
		res Type
		err error
	)

	if id >= r.TypeIDCount {
		return res, ErrInvalidTypeID
	}

	entryPos := r.typeIDOff + (id * 4)

	res.DescriptorStringID, err = r.readUint(entryPos)
	if err != nil {
		return res, err
	}

	return res, nil
}

type Proto struct {
	ShortyStringID        uint32
	ReturnTypeID          uint32
	ParametersTypeListOff uint32
}

func (r *Reader) GetProto(id uint32) (Proto, error) {
	var (
		res Proto
		err error
	)

	if id >= r.ProtoIDCount {
		return res, ErrInvalidProtoID
	}

	entryPos := r.protoIDOff + (id * 12)

	res.ShortyStringID, err = r.readUint(entryPos)
	if err != nil {
		return res, err
	}

	res.ReturnTypeID, err = r.readUint(entryPos + 4)
	if err != nil {
		return res, err
	}

	res.ParametersTypeListOff, err = r.readUint(entryPos + 8)
	if err != nil {
		return res, err
	}

	return res, nil
}
