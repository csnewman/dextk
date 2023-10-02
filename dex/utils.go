package dex

import (
	"encoding/binary"
	"io"
)

const reNotImpl = "reverse endian not implemented"

func (r *Reader) toUint(data []byte) uint32 {
	if r.ReverseEndian {
		panic(reNotImpl)
	}

	return binary.LittleEndian.Uint32(data)
}

func (r *Reader) toUshort(data []byte) uint16 {
	if r.ReverseEndian {
		panic(reNotImpl)
	}

	return binary.LittleEndian.Uint16(data)
}

func (r *Reader) readUint(pos uint32) (uint32, error) {
	raw := make([]byte, 4)

	_, err := r.file.ReadAt(raw, int64(pos))
	if err != nil {
		return 0, err
	}

	return r.toUint(raw), nil
}

func (r *Reader) readUshort(pos uint32) (uint16, error) {
	raw := make([]byte, 2)

	_, err := r.file.ReadAt(raw, int64(pos))
	if err != nil {
		return 0, err
	}

	return r.toUshort(raw), nil
}

func (r *Reader) readLeb128(pos uint32) (uint32, uint32, error) {
	if r.ReverseEndian {
		panic(reNotImpl)
	}

	raw := make([]byte, 5)

	n, err := r.file.ReadAt(raw, int64(pos))
	if err != nil && err != io.EOF {
		return 0, 0, err
	}

	value := uint32(0)
	shift := 0

	i := 0
	for i < 5 {
		if n <= i {
			return 0, 0, io.EOF
		}

		b := raw[i]

		more := (b & 0b10000000) != 0
		data := uint32(b & 0b01111111)

		value |= data << shift
		shift += 7
		i++

		if !more {
			break
		}
	}

	return value, uint32(i), nil
}
