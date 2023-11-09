package dextk

import (
	"errors"
	"fmt"
	"io"
)

var (
	ErrUnsupportedOp = errors.New("unsupported op")
	ErrMalformedOp   = errors.New("malformed op")
)

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

func (r *OpReader) Pos() int {
	return r.pos
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
		pseudo := (value >> 8) & 0xFF

		switch pseudo {
		case 0:
			// Standard noop
		case 1:
			return OpCodePseudoPackedSwitchPayload, nil
		case 2:
			return OpCodePseudoSparseSwitchPayload, nil
		case 3:
			return OpCodePseudoFillArrayDataPayload, nil
		default:
			return OpCodeInvalid, fmt.Errorf("%w: unknown pseudo %x", ErrUnsupportedOp, pseudo)
		}

	}

	return code, nil
}

func (r *OpReader) Skip() error {
	op, err := r.PeekCode()
	if err != nil {
		return err
	}

	size, err := op.Size(r)
	if err != nil {
		return err
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

func init() {
	opConfigsExtra[OpCodePseudoPackedSwitchPayload] = opConfig{
		Name: "pseudo-packed-switch-payload",
		Size: fmtPackedSwitchPayloadSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmtPackedSwitchPayload()
			if err != nil {
				return nil, err
			}

			return OpPseudoPackedSwitchPayload{
				opBase{pos: pos},
				f,
			}, nil
		},
	}
	opConfigsExtra[OpCodePseudoSparseSwitchPayload] = opConfig{
		Name: "pseudo-sparse-switch-payload",
		Size: fmtSparseSwitchPayloadSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmtSparseSwitchPayload()
			if err != nil {
				return nil, err
			}

			return OpPseudoSparseSwitchPayload{
				opBase{pos: pos},
				f,
			}, nil
		},
	}
	opConfigsExtra[OpCodePseudoFillArrayDataPayload] = opConfig{
		Name: "pseudo-fill-array-data-payload",
		Size: fmtFillArrayDataPayloadSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmtFillArrayDataPayload()
			if err != nil {
				return nil, err
			}

			return OpPseudoFillArrayDataPayload{
				opBase{pos: pos},
				f,
			}, nil
		},
	}
}
