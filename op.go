package dextk

import "fmt"

type OpCode int16

type opConfig struct {
	Name   string
	Size   func(r *OpReader) (int, error)
	Reader func(r *OpReader) (Op, error)
}

const (
	OpCodeInvalid                   OpCode = -1
	OpCodePseudoPackedSwitchPayload OpCode = -10
)

func (c OpCode) String() string {
	cfg, found := opConfigs[c]
	if !found {
		return fmt.Sprintf("unknown(0x%x)", int16(c))
	}

	return cfg.Name
}

func (c OpCode) Size(r *OpReader) (int, error) {
	cfg, found := opConfigs[c]
	if !found {
		return 0, fmt.Errorf("%w: %x", ErrUnsupportedOp, int16(c))
	}

	return cfg.Size(r)
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
