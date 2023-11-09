package dextk

import "fmt"

type OpCode int16

type opConfig struct {
	Name   string
	Size   func(r *OpReader) (int, error)
	Reader func(r *OpReader) (Op, error)
}

const (
	OpCodeInvalid                    OpCode = -1
	OpCodePseudoPackedSwitchPayload  OpCode = -10
	OpCodePseudoSparseSwitchPayload  OpCode = -11
	OpCodePseudoFillArrayDataPayload OpCode = -12
)

func (c OpCode) String() string {
	cfg, found := opConfigsExtra[c]
	if found {
		return cfg.Name
	}

	if c >= 0 && c <= 255 {
		cfg := opConfigs[c]

		if cfg.Name != "" {
			return cfg.Name
		}
	}

	return fmt.Sprintf("unknown(0x%x)", int16(c))
}

func (c OpCode) Size(r *OpReader) (int, error) {
	cfg, found := opConfigsExtra[c]
	if found {
		return cfg.Size(r)
	}

	if c >= 0 && c <= 255 {
		cfg := opConfigs[c]

		if cfg.Name != "" {
			return cfg.Size(r)
		}
	}

	return 0, fmt.Errorf("%w: %x", ErrUnsupportedOp, int16(c))
}

func (c OpCode) Read(r *OpReader) (Op, error) {
	cfg, found := opConfigsExtra[c]
	if found {
		return cfg.Reader(r)
	}

	if c >= 0 && c <= 255 {
		cfg := opConfigs[c]

		if cfg.Name != "" {
			return cfg.Reader(r)
		}
	}

	return nil, fmt.Errorf("%w: %x", ErrUnsupportedOp, int16(c))
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
