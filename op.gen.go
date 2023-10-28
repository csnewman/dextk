package dextk

import "fmt"

const (
	OpCodeNop          OpCode = 0x1
	OpCodeReturnObject OpCode = 0x11
	OpCodeNewInstance  OpCode = 0x22
	OpCodeInvokeDirect OpCode = 0x70
)

var opConfigs = map[OpCode]opConfig{
	OpCodeNop: {
		Name: "nop",
		Size: 1,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt10x()
			if err != nil {
				return nil, err
			}

			return &OpNop{
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeReturnObject: {
		Name: "return-object",
		Size: 1,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt11x()
			if err != nil {
				return nil, err
			}

			return &OpReturnObject{
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeNewInstance: {
		Name: "new-instance",
		Size: 2,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt21c()
			if err != nil {
				return nil, err
			}

			return &OpNewInstance{
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeInvokeDirect: {
		Name: "invoke-direct",
		Size: 3,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt35c()
			if err != nil {
				return nil, err
			}

			return &OpInvokeDirect{
				opBase{pos: pos},
				f,
			}, nil
		},
	},
}

type OpNop struct {
	opBase
	Fmt10x
}

func (o OpNop) Code() OpCode {
	return OpCodeNop
}

func (o OpNop) Fmt() Fmt {
	return o.Fmt10x
}

func (o OpNop) String() string {
	f := o.Fmt10x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: nop %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: nop", o.pos)
}

type OpReturnObject struct {
	opBase
	Fmt11x
}

func (o OpReturnObject) Code() OpCode {
	return OpCodeReturnObject
}

func (o OpReturnObject) Fmt() Fmt {
	return o.Fmt11x
}

func (o OpReturnObject) String() string {
	f := o.Fmt11x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: return-object %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: return-object", o.pos)
}

type OpNewInstance struct {
	opBase
	Fmt21c
}

func (o OpNewInstance) Code() OpCode {
	return OpCodeNewInstance
}

func (o OpNewInstance) Fmt() Fmt {
	return o.Fmt21c
}

func (o OpNewInstance) String() string {
	f := o.Fmt21c.String()
	if f != "" {
		return fmt.Sprintf("0x%x: new-instance %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: new-instance", o.pos)
}

type OpInvokeDirect struct {
	opBase
	Fmt35c
}

func (o OpInvokeDirect) Code() OpCode {
	return OpCodeInvokeDirect
}

func (o OpInvokeDirect) Fmt() Fmt {
	return o.Fmt35c
}

func (o OpInvokeDirect) String() string {
	f := o.Fmt35c.String()
	if f != "" {
		return fmt.Sprintf("0x%x: invoke-direct %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: invoke-direct", o.pos)
}
