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

type Fmt11 struct {
	A uint8
}

func (f Fmt11) internalFmt() {}

func (f Fmt11) Size() int {
	return 1
}

func (f Fmt11) String() string {
	return fmt.Sprintf("a=%v", f.A)
}

type Fmt11x struct{ Fmt11 }
type Fmt10t struct{ Fmt11 }

type Fmt12 struct {
	A uint8
	B uint8
}

type Fmt12x struct{ Fmt12 }
type Fmt11n struct{ Fmt12 }

func (f Fmt12) internalFmt() {}

func (f Fmt12) Size() int {
	return 1
}

func (f Fmt12) String() string {
	return fmt.Sprintf("a=%v, b=%v", f.A, f.B)
}

type Fmt22 struct {
	A uint8
	B uint16
}

type Fmt22x struct{ Fmt22 }
type Fmt21t struct{ Fmt22 }
type Fmt21h struct{ Fmt22 }
type Fmt21c struct{ Fmt22 }

func (f Fmt22) internalFmt() {}

func (f Fmt22) Size() int {
	return 2
}

func (f Fmt22) String() string {
	return fmt.Sprintf("a=%v, b=%v", f.A, f.B)
}

type Fmt21s struct {
	A uint8
	B int16
}

func (f Fmt21s) internalFmt() {}

func (f Fmt21s) Size() int {
	return 2
}

func (f Fmt21s) String() string {
	return fmt.Sprintf("a=%v, b=%v", f.A, f.B)
}

type Fmt23 struct {
	A uint8
	B uint8
	C uint16
}

type Fmt22t struct{ Fmt23 }
type Fmt22c struct{ Fmt23 }

func (f Fmt23) internalFmt() {}

func (f Fmt23) Size() int {
	return 2
}

func (f Fmt23) String() string {
	return fmt.Sprintf("a=%v, b=%v, c=%v", f.A, f.B, f.C)
}

type Fmt32x struct {
	A uint16
	B uint16
}

func (f Fmt32x) internalFmt() {}

func (f Fmt32x) Size() int {
	return 3
}

func (f Fmt32x) String() string {
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

type Fmt31 struct {
	A uint8
	B uint32
}

func (f Fmt31) internalFmt() {}

func (f Fmt31) Size() int {
	return 3
}

func (f Fmt31) String() string {
	return fmt.Sprintf("a=%v, b=%v", f.A, f.B)
}

type Fmt31i struct{ Fmt31 }
type Fmt31t struct{ Fmt31 }
type Fmt31c struct{ Fmt31 }
