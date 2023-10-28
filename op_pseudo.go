package dextk

import (
	"fmt"
	"io"
)

type OpPseudoPackedSwitchPayload struct {
	opBase
	FmtPackedSwitchPayload
}

func (o OpPseudoPackedSwitchPayload) Code() OpCode {
	return OpCodePseudoPackedSwitchPayload
}

func (o OpPseudoPackedSwitchPayload) Fmt() Fmt {
	return o.FmtPackedSwitchPayload
}

func (o OpPseudoPackedSwitchPayload) String() string {
	return fmt.Sprintf("0x%x: pseudo-packed-switch-payload %v", o.pos, o.FmtPackedSwitchPayload.String())
}

type FmtPackedSwitchPayload struct {
	FirstKey int32
	Targets  []int32
}

func (f FmtPackedSwitchPayload) internalFmt() {}

func (f FmtPackedSwitchPayload) Size() int {
	return (len(f.Targets) * 2) + 4
}

func (f FmtPackedSwitchPayload) String() string {
	return fmt.Sprintf("FirstKey=%v, Targets=%v", f.FirstKey, f.Targets)
}

func (r *OpReader) readFmtPackedSwitchPayload() (FmtPackedSwitchPayload, error) {
	var res FmtPackedSwitchPayload

	if r.pos+4 > len(r.ops) {
		return res, io.EOF
	}

	if r.ops[r.pos] != 0x0100 {
		return res, fmt.Errorf("%w: unexpected op 0x%x", ErrMalformedOp, r.ops[r.pos])
	}

	size := int(r.ops[r.pos+1])
	res.Targets = make([]int32, size)
	res.FirstKey = int32((uint32(r.ops[r.pos+3]) << 16) | uint32(r.ops[r.pos+2]))
	r.pos += 4

	if r.pos+(size*2) > len(r.ops) {
		return res, io.EOF
	}

	for i := 0; i < size; i++ {
		res.Targets[i] = int32((uint32(r.ops[r.pos+1]) << 16) | uint32(r.ops[r.pos]))
		r.pos += 2
	}

	return res, nil
}

func fmtPackedSwitchPayloadSize(r *OpReader) (int, error) {
	if r.pos+2 > len(r.ops) {
		return 0, io.EOF
	}

	if r.ops[r.pos] != 0x0100 {
		return 0, fmt.Errorf("%w: unexpected op 0x%x", ErrMalformedOp, r.ops[r.pos])
	}

	size := int(r.ops[r.pos+1])

	return (size * 2) + 4, nil
}
