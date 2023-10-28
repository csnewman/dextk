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

func (r *OpReader) readFmt11x() (Fmt11x, error) {
	var res Fmt11x

	if r.pos+1 > len(r.ops) {
		return res, io.EOF
	}

	res.A = uint8((r.ops[r.pos] >> 8) & 0xFF)

	r.pos += 1

	return res, nil
}

func (r *OpReader) readFmt21c() (Fmt21c, error) {
	var res Fmt21c

	if r.pos+2 > len(r.ops) {
		return res, io.EOF
	}

	res.A = uint8((r.ops[r.pos] >> 8) & 0xFF)
	res.B = r.ops[r.pos+1]

	r.pos += 2

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
