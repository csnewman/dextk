package dextk

import (
	"errors"
	"fmt"
	"io"
)

var ErrUnsupportedOp = errors.New("unsupported op")

type OpReader struct {
	ops []uint16
	pos int
}

func NewOpReader(ops []uint16) *OpReader {
	return &OpReader{
		ops: ops,
		pos: 0,
	}
}

func (r *OpReader) HasMore() bool {
	return r.pos < len(r.ops)
}

func (r *OpReader) Seek(pos int) {
	r.pos = pos
}

func (r *OpReader) PeekCode() (OpCode, error) {
	if r.pos >= len(r.ops) {
		return OpCodeInvalid, io.EOF
	}

	value := r.ops[r.pos]
	code := OpCode(value & 0xFF)

	if code == OpCodeNop {
		panic("todo: check pseudo")
	}

	return code, nil
}

func (r *OpReader) Skip() error {
	op, err := r.PeekCode()
	if err != nil {
		return err
	}

	size := op.Size()
	if size <= 0 {
		return fmt.Errorf("%w: %x", ErrUnsupportedOp, op)
	}

	r.pos += size

	return nil
}

func (r *OpReader) Read() (Op, error) {
	op, err := r.PeekCode()
	if err != nil {
		return nil, err
	}

	return op.Read(r)
}

func (r *OpReader) readFmt10x() (Fmt10x, error) {
	var res Fmt10x

	if r.pos+1 > len(r.ops) {
		return res, io.EOF
	}

	r.pos += 1

	return res, nil
}

func (r *OpReader) readFmt11() (Fmt11, error) {
	var res Fmt11

	if r.pos+1 > len(r.ops) {
		return res, io.EOF
	}

	res.A = uint8((r.ops[r.pos] >> 8) & 0xFF)

	r.pos += 1

	return res, nil
}

func (r *OpReader) readFmt11x() (Fmt11x, error) {
	f, e := r.readFmt11()
	return Fmt11x{f}, e
}

func (r *OpReader) readFmt10t() (Fmt10t, error) {
	f, e := r.readFmt11()
	return Fmt10t{f}, e
}

func (r *OpReader) readFmt12() (Fmt12, error) {
	var res Fmt12

	if r.pos+1 > len(r.ops) {
		return res, io.EOF
	}

	lower := r.ops[r.pos]
	res.A = uint8((lower >> 8) & 0xF)
	res.B = uint8((lower >> 12) & 0xF)

	r.pos += 1

	return res, nil
}

func (r *OpReader) readFmt12x() (Fmt12x, error) {
	f, e := r.readFmt12()
	return Fmt12x{f}, e
}

func (r *OpReader) readFmt11n() (Fmt11n, error) {
	f, e := r.readFmt12()
	return Fmt11n{f}, e
}

func (r *OpReader) readFmt22() (Fmt22, error) {
	var res Fmt22

	if r.pos+2 > len(r.ops) {
		return res, io.EOF
	}

	res.A = uint8((r.ops[r.pos] >> 8) & 0xFF)
	res.B = r.ops[r.pos+1]

	r.pos += 2

	return res, nil
}

func (r *OpReader) readFmt22x() (Fmt22x, error) {
	f, e := r.readFmt22()
	return Fmt22x{f}, e
}

func (r *OpReader) readFmt21s() (Fmt21s, error) {
	f, e := r.readFmt22()
	return Fmt21s{
		A: f.A,
		B: int16(f.B),
	}, e
}

func (r *OpReader) readFmt21t() (Fmt21t, error) {
	f, e := r.readFmt22()
	return Fmt21t{f}, e
}

func (r *OpReader) readFmt21h() (Fmt21h, error) {
	f, e := r.readFmt22()
	return Fmt21h{f}, e
}

func (r *OpReader) readFmt21c() (Fmt21c, error) {
	f, e := r.readFmt22()
	return Fmt21c{f}, e
}

func (r *OpReader) readFmt23() (Fmt23, error) {
	var res Fmt23

	if r.pos+2 > len(r.ops) {
		return res, io.EOF
	}

	lower := r.ops[r.pos]
	res.A = uint8((lower >> 8) & 0xF)
	res.B = uint8((lower >> 12) & 0xF)
	res.C = r.ops[r.pos]

	r.pos += 2

	return res, nil
}

func (r *OpReader) readFmt22t() (Fmt22t, error) {
	f, e := r.readFmt23()
	return Fmt22t{f}, e
}

func (r *OpReader) readFmt22c() (Fmt22c, error) {
	f, e := r.readFmt23()
	return Fmt22c{f}, e
}

func (r *OpReader) readFmt32x() (Fmt32x, error) {
	var res Fmt32x

	if r.pos+3 > len(r.ops) {
		return res, io.EOF
	}

	res.A = r.ops[r.pos+1]
	res.B = r.ops[r.pos+2]

	r.pos += 3

	return res, nil
}

func (r *OpReader) readFmt35c() (Fmt35c, error) {
	var res Fmt35c

	if r.pos+3 > len(r.ops) {
		return res, io.EOF
	}

	lower := r.ops[r.pos]
	res.A = uint8((lower >> 12) & 0xF)
	res.G = uint8((lower >> 8) & 0xF)

	res.B = r.ops[r.pos+1]

	upper := r.ops[r.pos+2]
	res.C = uint8(upper & 0xF)
	res.D = uint8((upper >> 4) & 0xF)
	res.E = uint8((upper >> 8) & 0xF)
	res.F = uint8((upper >> 12) & 0xF)

	r.pos += 3

	return res, nil
}

func (r *OpReader) readFmt31() (Fmt31, error) {
	var res Fmt31

	if r.pos+3 > len(r.ops) {
		return res, io.EOF
	}

	lower := r.ops[r.pos]
	res.A = uint8((lower >> 8) & 0xFF)
	res.B = (uint32(r.ops[r.pos+2]) << 16) | uint32(r.ops[r.pos+1])

	r.pos += 3

	return res, nil
}

func (r *OpReader) readFmt31i() (Fmt31i, error) {
	f, e := r.readFmt31()
	return Fmt31i{f}, e
}

func (r *OpReader) readFmt31t() (Fmt31t, error) {
	f, e := r.readFmt31()
	return Fmt31t{f}, e
}

func (r *OpReader) readFmt31c() (Fmt31c, error) {
	f, e := r.readFmt31()
	return Fmt31c{f}, e
}
