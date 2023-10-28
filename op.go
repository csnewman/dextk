package dextk

import "fmt"

type OpCode int16

type opConfig struct {
	Name   string
	Size   int
	Reader func(r *OpReader) (Op, error)
}

const (
	OpCodeInvalid OpCode = -1
)

func (c OpCode) String() string {
	cfg, found := opConfigs[c]
	if !found {
		return fmt.Sprintf("unknown(0x%x)", int16(c))
	}

	return cfg.Name
}

func (c OpCode) Size() int {
	cfg, found := opConfigs[c]
	if !found {
		return 0
	}

	return cfg.Size
}

func (c OpCode) Read(r *OpReader) (Op, error) {
	cfg, found := opConfigs[c]
	if !found {
		return nil, fmt.Errorf("%w: %x", ErrUnsupportedOp, int16(c))
	}

	return cfg.Reader(r)
}

type Op interface {
	Code() OpCode

	Pos() int

	String() string

	Fmt() Fmt
}

type opBase struct {
	pos int
}

func (o opBase) Pos() int {
	return o.pos
}

type Fmt interface {
	internalFmt()

	Size() int

	String() string
}

type Fmt10x struct {
}

func (f Fmt10x) internalFmt() {}

func (f Fmt10x) Size() int {
	return 1
}

func (f Fmt10x) String() string {
	return ""
}

type Fmt11x struct {
	A uint8
}

func (f Fmt11x) internalFmt() {}

func (f Fmt11x) Size() int {
	return 1
}

func (f Fmt11x) String() string {
	return fmt.Sprintf("a=%v", f.A)
}

type Fmt21c struct {
	A uint8
	B uint16
}

func (f Fmt21c) internalFmt() {}

func (f Fmt21c) Size() int {
	return 2
}

func (f Fmt21c) String() string {
	return fmt.Sprintf("a=%v, b=%v", f.A, f.B)
}

type Fmt35c struct {
	A uint8
	B uint16
	C uint8
	D uint8
	E uint8
	F uint8
	G uint8
}

func (f Fmt35c) internalFmt() {}

func (f Fmt35c) Size() int {
	return 3
}

func (f Fmt35c) String() string {
	return fmt.Sprintf("a=%v, b=%v, c=%v, d=%v, e=%v, f=%v, g=%v", f.A, f.B, f.C, f.D, f.E, f.F, f.G)
}
