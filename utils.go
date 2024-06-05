package dextk

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"math"
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

func (r *Reader) readUleb128(pos uint32) (uint32, uint32, error) {
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

var ErrMUTF8 = errors.New("invalid encoding")

func MUTF8Decode(d []byte, expectedSize int) (String, error) {
	inLen := len(d)
	buf := make([]uint16, 0, expectedSize)

	for i := 0; i < inLen; {
		if d[i] == 0 {
			return String{}, fmt.Errorf("%w: null unexpected", ErrMUTF8)
		} else if d[i] < 0x80 {
			buf = append(buf, uint16(d[i]))
			i++
		} else if d[i]&0xE0 == 0xC0 {
			if i+1 >= inLen {
				return String{}, fmt.Errorf("%w: bytes missing", ErrMUTF8)
			}

			buf = append(buf, ((uint16(d[i])&0x1F)<<6)|(uint16(d[i+1])&0x3F))
			i += 2
		} else if d[i]&0xF0 == 0xE0 {
			if i+2 >= inLen {
				return String{}, fmt.Errorf("%w: bytes missing", ErrMUTF8)
			}

			buf = append(buf, ((uint16(d[i])&0x0F)<<12)|((uint16(d[i+1])&0x3F)<<6)|(uint16(d[i+2])&0x3F))
			i += 3
		} else {
			return String{}, fmt.Errorf("%w: unexpected byte", ErrMUTF8)
		}
	}

	return StringFromUTF16(buf), nil
}

func (r *Reader) readSleb128(pos uint32) (int32, uint32, error) {
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
	lastBit := false

	i := 0
	for i < 5 {
		if n <= i {
			return 0, 0, io.EOF
		}

		b := raw[i]

		more := (b & 0b10000000) != 0
		data := uint32(b & 0b01111111)

		lastBit = (b & 0b1000000) != 0

		value |= data << shift
		shift += 7
		i++

		if !more {
			break
		}
	}

	if lastBit {
		// Set top bits if negative
		value |= (math.MaxUint32 >> shift) << shift
	}

	return int32(value), uint32(i), nil
}
