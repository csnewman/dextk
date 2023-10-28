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

type OpPseudoSparseSwitchPayload struct {
	opBase
	FmtSparseSwitchPayload
}

func (o OpPseudoSparseSwitchPayload) Code() OpCode {
	return OpCodePseudoSparseSwitchPayload
}

func (o OpPseudoSparseSwitchPayload) Fmt() Fmt {
	return o.FmtSparseSwitchPayload
}

func (o OpPseudoSparseSwitchPayload) String() string {
	return fmt.Sprintf("0x%x: pseudo-sparse-switch-payload %v", o.pos, o.FmtSparseSwitchPayload.String())
}

type FmtSparseSwitchPayload struct {
	Keys    []int32
	Targets []int32
}

func (f FmtSparseSwitchPayload) internalFmt() {}

func (f FmtSparseSwitchPayload) Size() int {
	return (len(f.Targets) * 4) + 2
}

func (f FmtSparseSwitchPayload) String() string {
	return fmt.Sprintf("Keys=%v, Targets=%v", f.Keys, f.Targets)
}

func (r *OpReader) readFmtSparseSwitchPayload() (FmtSparseSwitchPayload, error) {
	var res FmtSparseSwitchPayload

	if r.pos+2 > len(r.ops) {
		return res, io.EOF
	}

	if r.ops[r.pos] != 0x0200 {
		return res, fmt.Errorf("%w: unexpected op 0x%x", ErrMalformedOp, r.ops[r.pos])
	}

	size := int(r.ops[r.pos+1])
	res.Keys = make([]int32, size)
	res.Targets = make([]int32, size)
	r.pos += 2

	if r.pos+(size*4) > len(r.ops) {
		return res, io.EOF
	}

	for i := 0; i < size; i++ {
		res.Keys[i] = int32((uint32(r.ops[r.pos+1]) << 16) | uint32(r.ops[r.pos]))
		r.pos += 2
	}

	for i := 0; i < size; i++ {
		res.Targets[i] = int32((uint32(r.ops[r.pos+1]) << 16) | uint32(r.ops[r.pos]))
		r.pos += 2
	}

	return res, nil
}

func fmtSparseSwitchPayloadSize(r *OpReader) (int, error) {
	if r.pos+2 > len(r.ops) {
		return 0, io.EOF
	}

	if r.ops[r.pos] != 0x0200 {
		return 0, fmt.Errorf("%w: unexpected op 0x%x", ErrMalformedOp, r.ops[r.pos])
	}

	size := int(r.ops[r.pos+1])

	return (size * 4) + 2, nil
}

type OpPseudoFillArrayDataPayload struct {
	opBase
	FmtFillArrayDataPayload
}

func (o OpPseudoFillArrayDataPayload) Code() OpCode {
	return OpCodePseudoFillArrayDataPayload
}

func (o OpPseudoFillArrayDataPayload) Fmt() Fmt {
	return o.FmtFillArrayDataPayload
}

func (o OpPseudoFillArrayDataPayload) String() string {
	return fmt.Sprintf("0x%x: pseudo-fill-array-data-payload %v", o.pos, o.FmtFillArrayDataPayload.String())
}

type FmtFillArrayDataPayload struct {
	ElementWidth uint16
	ArraySize    uint32
	Data         []uint8
}

func (f FmtFillArrayDataPayload) internalFmt() {}

func (f FmtFillArrayDataPayload) Size() int {
	return int((uint64(f.ArraySize)*uint64(f.ElementWidth)+1)/2 + 4)
}

func (f FmtFillArrayDataPayload) String() string {
	return fmt.Sprintf("ElementWidth=%v, ArraySize=%v, Data=%v", f.ElementWidth, f.ArraySize, f.Data)
}

func (r *OpReader) readFmtFillArrayDataPayload() (FmtFillArrayDataPayload, error) {
	var res FmtFillArrayDataPayload

	if r.pos+4 > len(r.ops) {
		return res, io.EOF
	}

	if r.ops[r.pos] != 0x0300 {
		return res, fmt.Errorf("%w: unexpected op 0x%x", ErrMalformedOp, r.ops[r.pos])
	}

	res.ElementWidth = r.ops[r.pos+1]
	res.ArraySize = (uint32(r.ops[r.pos+3]) << 16) | uint32(r.ops[r.pos+2])
	r.pos += 4

	readSize := int((uint64(res.ArraySize)*uint64(res.ElementWidth) + 1) / 2)

	if r.pos+readSize > len(r.ops) {
		return res, io.EOF
	}

	res.Data = make([]uint8, readSize*2)

	for i := 0; i < readSize; i++ {
		read := r.ops[r.pos]
		res.Data[i*2] = uint8(read & 0xFF)
		res.Data[i*2+1] = uint8((read >> 8) & 0xFF)
		r.pos += 1
	}

	return res, nil
}

func fmtFillArrayDataPayloadSize(r *OpReader) (int, error) {
	if r.pos+4 > len(r.ops) {
		return 0, io.EOF
	}

	if r.ops[r.pos] != 0x0300 {
		return 0, fmt.Errorf("%w: unexpected op 0x%x", ErrMalformedOp, r.ops[r.pos])
	}

	width := uint64(r.ops[r.pos+1])
	size := uint64((uint32(r.ops[r.pos+3]) << 16) | uint32(r.ops[r.pos+2]))

	return int((width*size+1)/2 + 4), nil
}
