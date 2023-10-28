package dextk

import "fmt"

const (
	OpCodeNop OpCode = 0x0
	OpCodeMove OpCode = 0x1
	OpCodeMoveFrom16 OpCode = 0x2
	OpCodeMove16 OpCode = 0x3
	OpCodeMoveWide OpCode = 0x4
	OpCodeMoveWideFrom16 OpCode = 0x5
	OpCodeMoveWide16 OpCode = 0x6
	OpCodeMoveObject OpCode = 0x7
	OpCodeMoveObjectFrom16 OpCode = 0x8
	OpCodeMoveObject16 OpCode = 0x9
	OpCodeMoveResult OpCode = 0xa
	OpCodeMoveResultWide OpCode = 0xb
	OpCodeMoveResultObject OpCode = 0xc
	OpCodeMoveException OpCode = 0xd
	OpCodeReturnVoid OpCode = 0xe
	OpCodeReturn OpCode = 0xf
	OpCodeReturnWide OpCode = 0x10
	OpCodeReturnObject OpCode = 0x11
	OpCodeConst4 OpCode = 0x12
	OpCodeConst16 OpCode = 0x13
	OpCodeConst OpCode = 0x14
	OpCodeConstHigh16 OpCode = 0x15
	OpCodeConstWide16 OpCode = 0x16
	OpCodeConstWide32 OpCode = 0x17
	OpCodeConstWideHigh16 OpCode = 0x19
	OpCodeConstString OpCode = 0x1a
	OpCodeConstStringJumbo OpCode = 0x1b
	OpCodeConstClass OpCode = 0x1c
	OpCodeMonitorEnter OpCode = 0x1d
	OpCodeMonitorExit OpCode = 0x1e
	OpCodeCheckCast OpCode = 0x1f
	OpCodeInstanceOf OpCode = 0x20
	OpCodeArrayLength OpCode = 0x21
	OpCodeNewInstance OpCode = 0x22
	OpCodeNewArray OpCode = 0x23
	OpCodeThrow OpCode = 0x27
	OpCodeGoto OpCode = 0x28
	OpCodeIfEqz OpCode = 0x38
	OpCodeIfNez OpCode = 0x39
	OpCodeIfLtz OpCode = 0x3a
	OpCodeIfGez OpCode = 0x3b
	OpCodeIfGtz OpCode = 0x3c
	OpCodeIfLez OpCode = 0x3d
	OpCodeIget OpCode = 0x52
	OpCodeIgetWide OpCode = 0x53
	OpCodeIgetObject OpCode = 0x54
	OpCodeIgetBoolean OpCode = 0x55
	OpCodeIgetByte OpCode = 0x56
	OpCodeIgetChar OpCode = 0x57
	OpCodeIgetShort OpCode = 0x58
	OpCodeIput OpCode = 0x59
	OpCodeIputWide OpCode = 0x5a
	OpCodeIputObject OpCode = 0x5b
	OpCodeIputBoolean OpCode = 0x5c
	OpCodeIputByte OpCode = 0x5d
	OpCodeIputChar OpCode = 0x5e
	OpCodeIputShort OpCode = 0x5f
	OpCodeInvokeVirtual OpCode = 0x6e
	OpCodeInvokeSuper OpCode = 0x6f
	OpCodeInvokeDirect OpCode = 0x70
	OpCodeInvokeStatic OpCode = 0x71
	OpCodeInvokeInterface OpCode = 0x72
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

			return &OpNop {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeMove: {
		Name: "move",
		Size: 1,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt12x()
			if err != nil {
				return nil, err
			}

			return &OpMove {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeMoveFrom16: {
		Name: "move/from16",
		Size: 2,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt22x()
			if err != nil {
				return nil, err
			}

			return &OpMoveFrom16 {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeMove16: {
		Name: "move/16",
		Size: 3,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt32x()
			if err != nil {
				return nil, err
			}

			return &OpMove16 {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeMoveWide: {
		Name: "move-wide",
		Size: 1,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt12x()
			if err != nil {
				return nil, err
			}

			return &OpMoveWide {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeMoveWideFrom16: {
		Name: "move-wide/from16",
		Size: 2,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt22x()
			if err != nil {
				return nil, err
			}

			return &OpMoveWideFrom16 {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeMoveWide16: {
		Name: "move-wide/16",
		Size: 3,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt32x()
			if err != nil {
				return nil, err
			}

			return &OpMoveWide16 {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeMoveObject: {
		Name: "move-object",
		Size: 1,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt12x()
			if err != nil {
				return nil, err
			}

			return &OpMoveObject {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeMoveObjectFrom16: {
		Name: "move-object/from16",
		Size: 2,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt22x()
			if err != nil {
				return nil, err
			}

			return &OpMoveObjectFrom16 {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeMoveObject16: {
		Name: "move-object/16",
		Size: 3,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt32x()
			if err != nil {
				return nil, err
			}

			return &OpMoveObject16 {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeMoveResult: {
		Name: "move-result",
		Size: 1,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt11x()
			if err != nil {
				return nil, err
			}

			return &OpMoveResult {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeMoveResultWide: {
		Name: "move-result-wide",
		Size: 1,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt11x()
			if err != nil {
				return nil, err
			}

			return &OpMoveResultWide {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeMoveResultObject: {
		Name: "move-result-object",
		Size: 1,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt11x()
			if err != nil {
				return nil, err
			}

			return &OpMoveResultObject {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeMoveException: {
		Name: "move-exception",
		Size: 1,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt11x()
			if err != nil {
				return nil, err
			}

			return &OpMoveException {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeReturnVoid: {
		Name: "return-void",
		Size: 1,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt10x()
			if err != nil {
				return nil, err
			}

			return &OpReturnVoid {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeReturn: {
		Name: "return",
		Size: 1,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt11x()
			if err != nil {
				return nil, err
			}

			return &OpReturn {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeReturnWide: {
		Name: "return-wide",
		Size: 1,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt11x()
			if err != nil {
				return nil, err
			}

			return &OpReturnWide {
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

			return &OpReturnObject {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeConst4: {
		Name: "const/4",
		Size: 1,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt11n()
			if err != nil {
				return nil, err
			}

			return &OpConst4 {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeConst16: {
		Name: "const/16",
		Size: 2,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt21s()
			if err != nil {
				return nil, err
			}

			return &OpConst16 {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeConst: {
		Name: "const",
		Size: 3,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt31i()
			if err != nil {
				return nil, err
			}

			return &OpConst {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeConstHigh16: {
		Name: "const/high16",
		Size: 2,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt21h()
			if err != nil {
				return nil, err
			}

			return &OpConstHigh16 {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeConstWide16: {
		Name: "const-wide/16",
		Size: 2,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt21s()
			if err != nil {
				return nil, err
			}

			return &OpConstWide16 {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeConstWide32: {
		Name: "const-wide/32",
		Size: 3,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt31i()
			if err != nil {
				return nil, err
			}

			return &OpConstWide32 {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeConstWideHigh16: {
		Name: "const-wide/high16",
		Size: 2,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt21h()
			if err != nil {
				return nil, err
			}

			return &OpConstWideHigh16 {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeConstString: {
		Name: "const-string",
		Size: 2,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt21c()
			if err != nil {
				return nil, err
			}

			return &OpConstString {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeConstStringJumbo: {
		Name: "const-string/jumbo",
		Size: 3,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt31c()
			if err != nil {
				return nil, err
			}

			return &OpConstStringJumbo {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeConstClass: {
		Name: "const-class",
		Size: 2,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt21c()
			if err != nil {
				return nil, err
			}

			return &OpConstClass {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeMonitorEnter: {
		Name: "monitor-enter",
		Size: 1,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt11x()
			if err != nil {
				return nil, err
			}

			return &OpMonitorEnter {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeMonitorExit: {
		Name: "monitor-exit",
		Size: 1,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt11x()
			if err != nil {
				return nil, err
			}

			return &OpMonitorExit {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeCheckCast: {
		Name: "check-cast",
		Size: 2,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt21c()
			if err != nil {
				return nil, err
			}

			return &OpCheckCast {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeInstanceOf: {
		Name: "instance-of",
		Size: 2,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt22c()
			if err != nil {
				return nil, err
			}

			return &OpInstanceOf {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeArrayLength: {
		Name: "array-length",
		Size: 1,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt12x()
			if err != nil {
				return nil, err
			}

			return &OpArrayLength {
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

			return &OpNewInstance {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeNewArray: {
		Name: "new-array",
		Size: 2,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt22c()
			if err != nil {
				return nil, err
			}

			return &OpNewArray {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeThrow: {
		Name: "throw",
		Size: 1,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt11x()
			if err != nil {
				return nil, err
			}

			return &OpThrow {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeGoto: {
		Name: "goto",
		Size: 1,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt10t()
			if err != nil {
				return nil, err
			}

			return &OpGoto {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeIfEqz: {
		Name: "if-eqz",
		Size: 2,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt21t()
			if err != nil {
				return nil, err
			}

			return &OpIfEqz {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeIfNez: {
		Name: "if-nez",
		Size: 2,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt21t()
			if err != nil {
				return nil, err
			}

			return &OpIfNez {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeIfLtz: {
		Name: "if-ltz",
		Size: 2,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt21t()
			if err != nil {
				return nil, err
			}

			return &OpIfLtz {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeIfGez: {
		Name: "if-gez",
		Size: 2,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt21t()
			if err != nil {
				return nil, err
			}

			return &OpIfGez {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeIfGtz: {
		Name: "if-gtz",
		Size: 2,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt21t()
			if err != nil {
				return nil, err
			}

			return &OpIfGtz {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeIfLez: {
		Name: "if-lez",
		Size: 2,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt21t()
			if err != nil {
				return nil, err
			}

			return &OpIfLez {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeIget: {
		Name: "iget",
		Size: 2,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt22c()
			if err != nil {
				return nil, err
			}

			return &OpIget {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeIgetWide: {
		Name: "iget-wide",
		Size: 2,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt22c()
			if err != nil {
				return nil, err
			}

			return &OpIgetWide {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeIgetObject: {
		Name: "iget-object",
		Size: 2,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt22c()
			if err != nil {
				return nil, err
			}

			return &OpIgetObject {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeIgetBoolean: {
		Name: "iget-boolean",
		Size: 2,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt22c()
			if err != nil {
				return nil, err
			}

			return &OpIgetBoolean {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeIgetByte: {
		Name: "iget-byte",
		Size: 2,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt22c()
			if err != nil {
				return nil, err
			}

			return &OpIgetByte {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeIgetChar: {
		Name: "iget-char",
		Size: 2,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt22c()
			if err != nil {
				return nil, err
			}

			return &OpIgetChar {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeIgetShort: {
		Name: "iget-short",
		Size: 2,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt22c()
			if err != nil {
				return nil, err
			}

			return &OpIgetShort {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeIput: {
		Name: "iput",
		Size: 2,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt22c()
			if err != nil {
				return nil, err
			}

			return &OpIput {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeIputWide: {
		Name: "iput-wide",
		Size: 2,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt22c()
			if err != nil {
				return nil, err
			}

			return &OpIputWide {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeIputObject: {
		Name: "iput-object",
		Size: 2,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt22c()
			if err != nil {
				return nil, err
			}

			return &OpIputObject {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeIputBoolean: {
		Name: "iput-boolean",
		Size: 2,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt22c()
			if err != nil {
				return nil, err
			}

			return &OpIputBoolean {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeIputByte: {
		Name: "iput-byte",
		Size: 2,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt22c()
			if err != nil {
				return nil, err
			}

			return &OpIputByte {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeIputChar: {
		Name: "iput-char",
		Size: 2,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt22c()
			if err != nil {
				return nil, err
			}

			return &OpIputChar {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeIputShort: {
		Name: "iput-short",
		Size: 2,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt22c()
			if err != nil {
				return nil, err
			}

			return &OpIputShort {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeInvokeVirtual: {
		Name: "invoke-virtual",
		Size: 3,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt35c()
			if err != nil {
				return nil, err
			}

			return &OpInvokeVirtual {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeInvokeSuper: {
		Name: "invoke-super",
		Size: 3,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt35c()
			if err != nil {
				return nil, err
			}

			return &OpInvokeSuper {
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

			return &OpInvokeDirect {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeInvokeStatic: {
		Name: "invoke-static",
		Size: 3,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt35c()
			if err != nil {
				return nil, err
			}

			return &OpInvokeStatic {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeInvokeInterface: {
		Name: "invoke-interface",
		Size: 3,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt35c()
			if err != nil {
				return nil, err
			}

			return &OpInvokeInterface {
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

type OpMove struct {
	opBase
	Fmt12x
}

func (o OpMove) Code() OpCode {
	return OpCodeMove
}

func (o OpMove) Fmt() Fmt {
	return o.Fmt12x
}

func (o OpMove) String() string {
	f := o.Fmt12x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: move %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: move", o.pos)
}

type OpMoveFrom16 struct {
	opBase
	Fmt22x
}

func (o OpMoveFrom16) Code() OpCode {
	return OpCodeMoveFrom16
}

func (o OpMoveFrom16) Fmt() Fmt {
	return o.Fmt22x
}

func (o OpMoveFrom16) String() string {
	f := o.Fmt22x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: move/from16 %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: move/from16", o.pos)
}

type OpMove16 struct {
	opBase
	Fmt32x
}

func (o OpMove16) Code() OpCode {
	return OpCodeMove16
}

func (o OpMove16) Fmt() Fmt {
	return o.Fmt32x
}

func (o OpMove16) String() string {
	f := o.Fmt32x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: move/16 %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: move/16", o.pos)
}

type OpMoveWide struct {
	opBase
	Fmt12x
}

func (o OpMoveWide) Code() OpCode {
	return OpCodeMoveWide
}

func (o OpMoveWide) Fmt() Fmt {
	return o.Fmt12x
}

func (o OpMoveWide) String() string {
	f := o.Fmt12x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: move-wide %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: move-wide", o.pos)
}

type OpMoveWideFrom16 struct {
	opBase
	Fmt22x
}

func (o OpMoveWideFrom16) Code() OpCode {
	return OpCodeMoveWideFrom16
}

func (o OpMoveWideFrom16) Fmt() Fmt {
	return o.Fmt22x
}

func (o OpMoveWideFrom16) String() string {
	f := o.Fmt22x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: move-wide/from16 %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: move-wide/from16", o.pos)
}

type OpMoveWide16 struct {
	opBase
	Fmt32x
}

func (o OpMoveWide16) Code() OpCode {
	return OpCodeMoveWide16
}

func (o OpMoveWide16) Fmt() Fmt {
	return o.Fmt32x
}

func (o OpMoveWide16) String() string {
	f := o.Fmt32x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: move-wide/16 %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: move-wide/16", o.pos)
}

type OpMoveObject struct {
	opBase
	Fmt12x
}

func (o OpMoveObject) Code() OpCode {
	return OpCodeMoveObject
}

func (o OpMoveObject) Fmt() Fmt {
	return o.Fmt12x
}

func (o OpMoveObject) String() string {
	f := o.Fmt12x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: move-object %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: move-object", o.pos)
}

type OpMoveObjectFrom16 struct {
	opBase
	Fmt22x
}

func (o OpMoveObjectFrom16) Code() OpCode {
	return OpCodeMoveObjectFrom16
}

func (o OpMoveObjectFrom16) Fmt() Fmt {
	return o.Fmt22x
}

func (o OpMoveObjectFrom16) String() string {
	f := o.Fmt22x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: move-object/from16 %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: move-object/from16", o.pos)
}

type OpMoveObject16 struct {
	opBase
	Fmt32x
}

func (o OpMoveObject16) Code() OpCode {
	return OpCodeMoveObject16
}

func (o OpMoveObject16) Fmt() Fmt {
	return o.Fmt32x
}

func (o OpMoveObject16) String() string {
	f := o.Fmt32x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: move-object/16 %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: move-object/16", o.pos)
}

type OpMoveResult struct {
	opBase
	Fmt11x
}

func (o OpMoveResult) Code() OpCode {
	return OpCodeMoveResult
}

func (o OpMoveResult) Fmt() Fmt {
	return o.Fmt11x
}

func (o OpMoveResult) String() string {
	f := o.Fmt11x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: move-result %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: move-result", o.pos)
}

type OpMoveResultWide struct {
	opBase
	Fmt11x
}

func (o OpMoveResultWide) Code() OpCode {
	return OpCodeMoveResultWide
}

func (o OpMoveResultWide) Fmt() Fmt {
	return o.Fmt11x
}

func (o OpMoveResultWide) String() string {
	f := o.Fmt11x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: move-result-wide %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: move-result-wide", o.pos)
}

type OpMoveResultObject struct {
	opBase
	Fmt11x
}

func (o OpMoveResultObject) Code() OpCode {
	return OpCodeMoveResultObject
}

func (o OpMoveResultObject) Fmt() Fmt {
	return o.Fmt11x
}

func (o OpMoveResultObject) String() string {
	f := o.Fmt11x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: move-result-object %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: move-result-object", o.pos)
}

type OpMoveException struct {
	opBase
	Fmt11x
}

func (o OpMoveException) Code() OpCode {
	return OpCodeMoveException
}

func (o OpMoveException) Fmt() Fmt {
	return o.Fmt11x
}

func (o OpMoveException) String() string {
	f := o.Fmt11x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: move-exception %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: move-exception", o.pos)
}

type OpReturnVoid struct {
	opBase
	Fmt10x
}

func (o OpReturnVoid) Code() OpCode {
	return OpCodeReturnVoid
}

func (o OpReturnVoid) Fmt() Fmt {
	return o.Fmt10x
}

func (o OpReturnVoid) String() string {
	f := o.Fmt10x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: return-void %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: return-void", o.pos)
}

type OpReturn struct {
	opBase
	Fmt11x
}

func (o OpReturn) Code() OpCode {
	return OpCodeReturn
}

func (o OpReturn) Fmt() Fmt {
	return o.Fmt11x
}

func (o OpReturn) String() string {
	f := o.Fmt11x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: return %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: return", o.pos)
}

type OpReturnWide struct {
	opBase
	Fmt11x
}

func (o OpReturnWide) Code() OpCode {
	return OpCodeReturnWide
}

func (o OpReturnWide) Fmt() Fmt {
	return o.Fmt11x
}

func (o OpReturnWide) String() string {
	f := o.Fmt11x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: return-wide %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: return-wide", o.pos)
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

type OpConst4 struct {
	opBase
	Fmt11n
}

func (o OpConst4) Code() OpCode {
	return OpCodeConst4
}

func (o OpConst4) Fmt() Fmt {
	return o.Fmt11n
}

func (o OpConst4) String() string {
	f := o.Fmt11n.String()
	if f != "" {
		return fmt.Sprintf("0x%x: const/4 %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: const/4", o.pos)
}

type OpConst16 struct {
	opBase
	Fmt21s
}

func (o OpConst16) Code() OpCode {
	return OpCodeConst16
}

func (o OpConst16) Fmt() Fmt {
	return o.Fmt21s
}

func (o OpConst16) String() string {
	f := o.Fmt21s.String()
	if f != "" {
		return fmt.Sprintf("0x%x: const/16 %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: const/16", o.pos)
}

type OpConst struct {
	opBase
	Fmt31i
}

func (o OpConst) Code() OpCode {
	return OpCodeConst
}

func (o OpConst) Fmt() Fmt {
	return o.Fmt31i
}

func (o OpConst) String() string {
	f := o.Fmt31i.String()
	if f != "" {
		return fmt.Sprintf("0x%x: const %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: const", o.pos)
}

type OpConstHigh16 struct {
	opBase
	Fmt21h
}

func (o OpConstHigh16) Code() OpCode {
	return OpCodeConstHigh16
}

func (o OpConstHigh16) Fmt() Fmt {
	return o.Fmt21h
}

func (o OpConstHigh16) String() string {
	f := o.Fmt21h.String()
	if f != "" {
		return fmt.Sprintf("0x%x: const/high16 %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: const/high16", o.pos)
}

type OpConstWide16 struct {
	opBase
	Fmt21s
}

func (o OpConstWide16) Code() OpCode {
	return OpCodeConstWide16
}

func (o OpConstWide16) Fmt() Fmt {
	return o.Fmt21s
}

func (o OpConstWide16) String() string {
	f := o.Fmt21s.String()
	if f != "" {
		return fmt.Sprintf("0x%x: const-wide/16 %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: const-wide/16", o.pos)
}

type OpConstWide32 struct {
	opBase
	Fmt31i
}

func (o OpConstWide32) Code() OpCode {
	return OpCodeConstWide32
}

func (o OpConstWide32) Fmt() Fmt {
	return o.Fmt31i
}

func (o OpConstWide32) String() string {
	f := o.Fmt31i.String()
	if f != "" {
		return fmt.Sprintf("0x%x: const-wide/32 %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: const-wide/32", o.pos)
}

type OpConstWideHigh16 struct {
	opBase
	Fmt21h
}

func (o OpConstWideHigh16) Code() OpCode {
	return OpCodeConstWideHigh16
}

func (o OpConstWideHigh16) Fmt() Fmt {
	return o.Fmt21h
}

func (o OpConstWideHigh16) String() string {
	f := o.Fmt21h.String()
	if f != "" {
		return fmt.Sprintf("0x%x: const-wide/high16 %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: const-wide/high16", o.pos)
}

type OpConstString struct {
	opBase
	Fmt21c
}

func (o OpConstString) Code() OpCode {
	return OpCodeConstString
}

func (o OpConstString) Fmt() Fmt {
	return o.Fmt21c
}

func (o OpConstString) String() string {
	f := o.Fmt21c.String()
	if f != "" {
		return fmt.Sprintf("0x%x: const-string %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: const-string", o.pos)
}

type OpConstStringJumbo struct {
	opBase
	Fmt31c
}

func (o OpConstStringJumbo) Code() OpCode {
	return OpCodeConstStringJumbo
}

func (o OpConstStringJumbo) Fmt() Fmt {
	return o.Fmt31c
}

func (o OpConstStringJumbo) String() string {
	f := o.Fmt31c.String()
	if f != "" {
		return fmt.Sprintf("0x%x: const-string/jumbo %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: const-string/jumbo", o.pos)
}

type OpConstClass struct {
	opBase
	Fmt21c
}

func (o OpConstClass) Code() OpCode {
	return OpCodeConstClass
}

func (o OpConstClass) Fmt() Fmt {
	return o.Fmt21c
}

func (o OpConstClass) String() string {
	f := o.Fmt21c.String()
	if f != "" {
		return fmt.Sprintf("0x%x: const-class %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: const-class", o.pos)
}

type OpMonitorEnter struct {
	opBase
	Fmt11x
}

func (o OpMonitorEnter) Code() OpCode {
	return OpCodeMonitorEnter
}

func (o OpMonitorEnter) Fmt() Fmt {
	return o.Fmt11x
}

func (o OpMonitorEnter) String() string {
	f := o.Fmt11x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: monitor-enter %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: monitor-enter", o.pos)
}

type OpMonitorExit struct {
	opBase
	Fmt11x
}

func (o OpMonitorExit) Code() OpCode {
	return OpCodeMonitorExit
}

func (o OpMonitorExit) Fmt() Fmt {
	return o.Fmt11x
}

func (o OpMonitorExit) String() string {
	f := o.Fmt11x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: monitor-exit %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: monitor-exit", o.pos)
}

type OpCheckCast struct {
	opBase
	Fmt21c
}

func (o OpCheckCast) Code() OpCode {
	return OpCodeCheckCast
}

func (o OpCheckCast) Fmt() Fmt {
	return o.Fmt21c
}

func (o OpCheckCast) String() string {
	f := o.Fmt21c.String()
	if f != "" {
		return fmt.Sprintf("0x%x: check-cast %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: check-cast", o.pos)
}

type OpInstanceOf struct {
	opBase
	Fmt22c
}

func (o OpInstanceOf) Code() OpCode {
	return OpCodeInstanceOf
}

func (o OpInstanceOf) Fmt() Fmt {
	return o.Fmt22c
}

func (o OpInstanceOf) String() string {
	f := o.Fmt22c.String()
	if f != "" {
		return fmt.Sprintf("0x%x: instance-of %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: instance-of", o.pos)
}

type OpArrayLength struct {
	opBase
	Fmt12x
}

func (o OpArrayLength) Code() OpCode {
	return OpCodeArrayLength
}

func (o OpArrayLength) Fmt() Fmt {
	return o.Fmt12x
}

func (o OpArrayLength) String() string {
	f := o.Fmt12x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: array-length %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: array-length", o.pos)
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

type OpNewArray struct {
	opBase
	Fmt22c
}

func (o OpNewArray) Code() OpCode {
	return OpCodeNewArray
}

func (o OpNewArray) Fmt() Fmt {
	return o.Fmt22c
}

func (o OpNewArray) String() string {
	f := o.Fmt22c.String()
	if f != "" {
		return fmt.Sprintf("0x%x: new-array %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: new-array", o.pos)
}

type OpThrow struct {
	opBase
	Fmt11x
}

func (o OpThrow) Code() OpCode {
	return OpCodeThrow
}

func (o OpThrow) Fmt() Fmt {
	return o.Fmt11x
}

func (o OpThrow) String() string {
	f := o.Fmt11x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: throw %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: throw", o.pos)
}

type OpGoto struct {
	opBase
	Fmt10t
}

func (o OpGoto) Code() OpCode {
	return OpCodeGoto
}

func (o OpGoto) Fmt() Fmt {
	return o.Fmt10t
}

func (o OpGoto) String() string {
	f := o.Fmt10t.String()
	if f != "" {
		return fmt.Sprintf("0x%x: goto %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: goto", o.pos)
}

type OpIfEqz struct {
	opBase
	Fmt21t
}

func (o OpIfEqz) Code() OpCode {
	return OpCodeIfEqz
}

func (o OpIfEqz) Fmt() Fmt {
	return o.Fmt21t
}

func (o OpIfEqz) String() string {
	f := o.Fmt21t.String()
	if f != "" {
		return fmt.Sprintf("0x%x: if-eqz %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: if-eqz", o.pos)
}

type OpIfNez struct {
	opBase
	Fmt21t
}

func (o OpIfNez) Code() OpCode {
	return OpCodeIfNez
}

func (o OpIfNez) Fmt() Fmt {
	return o.Fmt21t
}

func (o OpIfNez) String() string {
	f := o.Fmt21t.String()
	if f != "" {
		return fmt.Sprintf("0x%x: if-nez %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: if-nez", o.pos)
}

type OpIfLtz struct {
	opBase
	Fmt21t
}

func (o OpIfLtz) Code() OpCode {
	return OpCodeIfLtz
}

func (o OpIfLtz) Fmt() Fmt {
	return o.Fmt21t
}

func (o OpIfLtz) String() string {
	f := o.Fmt21t.String()
	if f != "" {
		return fmt.Sprintf("0x%x: if-ltz %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: if-ltz", o.pos)
}

type OpIfGez struct {
	opBase
	Fmt21t
}

func (o OpIfGez) Code() OpCode {
	return OpCodeIfGez
}

func (o OpIfGez) Fmt() Fmt {
	return o.Fmt21t
}

func (o OpIfGez) String() string {
	f := o.Fmt21t.String()
	if f != "" {
		return fmt.Sprintf("0x%x: if-gez %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: if-gez", o.pos)
}

type OpIfGtz struct {
	opBase
	Fmt21t
}

func (o OpIfGtz) Code() OpCode {
	return OpCodeIfGtz
}

func (o OpIfGtz) Fmt() Fmt {
	return o.Fmt21t
}

func (o OpIfGtz) String() string {
	f := o.Fmt21t.String()
	if f != "" {
		return fmt.Sprintf("0x%x: if-gtz %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: if-gtz", o.pos)
}

type OpIfLez struct {
	opBase
	Fmt21t
}

func (o OpIfLez) Code() OpCode {
	return OpCodeIfLez
}

func (o OpIfLez) Fmt() Fmt {
	return o.Fmt21t
}

func (o OpIfLez) String() string {
	f := o.Fmt21t.String()
	if f != "" {
		return fmt.Sprintf("0x%x: if-lez %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: if-lez", o.pos)
}

type OpIget struct {
	opBase
	Fmt22c
}

func (o OpIget) Code() OpCode {
	return OpCodeIget
}

func (o OpIget) Fmt() Fmt {
	return o.Fmt22c
}

func (o OpIget) String() string {
	f := o.Fmt22c.String()
	if f != "" {
		return fmt.Sprintf("0x%x: iget %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: iget", o.pos)
}

type OpIgetWide struct {
	opBase
	Fmt22c
}

func (o OpIgetWide) Code() OpCode {
	return OpCodeIgetWide
}

func (o OpIgetWide) Fmt() Fmt {
	return o.Fmt22c
}

func (o OpIgetWide) String() string {
	f := o.Fmt22c.String()
	if f != "" {
		return fmt.Sprintf("0x%x: iget-wide %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: iget-wide", o.pos)
}

type OpIgetObject struct {
	opBase
	Fmt22c
}

func (o OpIgetObject) Code() OpCode {
	return OpCodeIgetObject
}

func (o OpIgetObject) Fmt() Fmt {
	return o.Fmt22c
}

func (o OpIgetObject) String() string {
	f := o.Fmt22c.String()
	if f != "" {
		return fmt.Sprintf("0x%x: iget-object %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: iget-object", o.pos)
}

type OpIgetBoolean struct {
	opBase
	Fmt22c
}

func (o OpIgetBoolean) Code() OpCode {
	return OpCodeIgetBoolean
}

func (o OpIgetBoolean) Fmt() Fmt {
	return o.Fmt22c
}

func (o OpIgetBoolean) String() string {
	f := o.Fmt22c.String()
	if f != "" {
		return fmt.Sprintf("0x%x: iget-boolean %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: iget-boolean", o.pos)
}

type OpIgetByte struct {
	opBase
	Fmt22c
}

func (o OpIgetByte) Code() OpCode {
	return OpCodeIgetByte
}

func (o OpIgetByte) Fmt() Fmt {
	return o.Fmt22c
}

func (o OpIgetByte) String() string {
	f := o.Fmt22c.String()
	if f != "" {
		return fmt.Sprintf("0x%x: iget-byte %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: iget-byte", o.pos)
}

type OpIgetChar struct {
	opBase
	Fmt22c
}

func (o OpIgetChar) Code() OpCode {
	return OpCodeIgetChar
}

func (o OpIgetChar) Fmt() Fmt {
	return o.Fmt22c
}

func (o OpIgetChar) String() string {
	f := o.Fmt22c.String()
	if f != "" {
		return fmt.Sprintf("0x%x: iget-char %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: iget-char", o.pos)
}

type OpIgetShort struct {
	opBase
	Fmt22c
}

func (o OpIgetShort) Code() OpCode {
	return OpCodeIgetShort
}

func (o OpIgetShort) Fmt() Fmt {
	return o.Fmt22c
}

func (o OpIgetShort) String() string {
	f := o.Fmt22c.String()
	if f != "" {
		return fmt.Sprintf("0x%x: iget-short %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: iget-short", o.pos)
}

type OpIput struct {
	opBase
	Fmt22c
}

func (o OpIput) Code() OpCode {
	return OpCodeIput
}

func (o OpIput) Fmt() Fmt {
	return o.Fmt22c
}

func (o OpIput) String() string {
	f := o.Fmt22c.String()
	if f != "" {
		return fmt.Sprintf("0x%x: iput %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: iput", o.pos)
}

type OpIputWide struct {
	opBase
	Fmt22c
}

func (o OpIputWide) Code() OpCode {
	return OpCodeIputWide
}

func (o OpIputWide) Fmt() Fmt {
	return o.Fmt22c
}

func (o OpIputWide) String() string {
	f := o.Fmt22c.String()
	if f != "" {
		return fmt.Sprintf("0x%x: iput-wide %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: iput-wide", o.pos)
}

type OpIputObject struct {
	opBase
	Fmt22c
}

func (o OpIputObject) Code() OpCode {
	return OpCodeIputObject
}

func (o OpIputObject) Fmt() Fmt {
	return o.Fmt22c
}

func (o OpIputObject) String() string {
	f := o.Fmt22c.String()
	if f != "" {
		return fmt.Sprintf("0x%x: iput-object %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: iput-object", o.pos)
}

type OpIputBoolean struct {
	opBase
	Fmt22c
}

func (o OpIputBoolean) Code() OpCode {
	return OpCodeIputBoolean
}

func (o OpIputBoolean) Fmt() Fmt {
	return o.Fmt22c
}

func (o OpIputBoolean) String() string {
	f := o.Fmt22c.String()
	if f != "" {
		return fmt.Sprintf("0x%x: iput-boolean %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: iput-boolean", o.pos)
}

type OpIputByte struct {
	opBase
	Fmt22c
}

func (o OpIputByte) Code() OpCode {
	return OpCodeIputByte
}

func (o OpIputByte) Fmt() Fmt {
	return o.Fmt22c
}

func (o OpIputByte) String() string {
	f := o.Fmt22c.String()
	if f != "" {
		return fmt.Sprintf("0x%x: iput-byte %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: iput-byte", o.pos)
}

type OpIputChar struct {
	opBase
	Fmt22c
}

func (o OpIputChar) Code() OpCode {
	return OpCodeIputChar
}

func (o OpIputChar) Fmt() Fmt {
	return o.Fmt22c
}

func (o OpIputChar) String() string {
	f := o.Fmt22c.String()
	if f != "" {
		return fmt.Sprintf("0x%x: iput-char %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: iput-char", o.pos)
}

type OpIputShort struct {
	opBase
	Fmt22c
}

func (o OpIputShort) Code() OpCode {
	return OpCodeIputShort
}

func (o OpIputShort) Fmt() Fmt {
	return o.Fmt22c
}

func (o OpIputShort) String() string {
	f := o.Fmt22c.String()
	if f != "" {
		return fmt.Sprintf("0x%x: iput-short %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: iput-short", o.pos)
}

type OpInvokeVirtual struct {
	opBase
	Fmt35c
}

func (o OpInvokeVirtual) Code() OpCode {
	return OpCodeInvokeVirtual
}

func (o OpInvokeVirtual) Fmt() Fmt {
	return o.Fmt35c
}

func (o OpInvokeVirtual) String() string {
	f := o.Fmt35c.String()
	if f != "" {
		return fmt.Sprintf("0x%x: invoke-virtual %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: invoke-virtual", o.pos)
}

type OpInvokeSuper struct {
	opBase
	Fmt35c
}

func (o OpInvokeSuper) Code() OpCode {
	return OpCodeInvokeSuper
}

func (o OpInvokeSuper) Fmt() Fmt {
	return o.Fmt35c
}

func (o OpInvokeSuper) String() string {
	f := o.Fmt35c.String()
	if f != "" {
		return fmt.Sprintf("0x%x: invoke-super %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: invoke-super", o.pos)
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

type OpInvokeStatic struct {
	opBase
	Fmt35c
}

func (o OpInvokeStatic) Code() OpCode {
	return OpCodeInvokeStatic
}

func (o OpInvokeStatic) Fmt() Fmt {
	return o.Fmt35c
}

func (o OpInvokeStatic) String() string {
	f := o.Fmt35c.String()
	if f != "" {
		return fmt.Sprintf("0x%x: invoke-static %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: invoke-static", o.pos)
}

type OpInvokeInterface struct {
	opBase
	Fmt35c
}

func (o OpInvokeInterface) Code() OpCode {
	return OpCodeInvokeInterface
}

func (o OpInvokeInterface) Fmt() Fmt {
	return o.Fmt35c
}

func (o OpInvokeInterface) String() string {
	f := o.Fmt35c.String()
	if f != "" {
		return fmt.Sprintf("0x%x: invoke-interface %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: invoke-interface", o.pos)
}
