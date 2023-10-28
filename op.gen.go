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
	OpCodeGoto16 OpCode = 0x29
	OpCodeGoto32 OpCode = 0x2a
	OpCodePackedSwitch OpCode = 0x2b
	OpCodeCmplFloat OpCode = 0x2d
	OpCodeCmpgFloat OpCode = 0x2e
	OpCodeCmplDouble OpCode = 0x2f
	OpCodeCmpgDouble OpCode = 0x30
	OpCodeCmpLong OpCode = 0x31
	OpCodeIfEq OpCode = 0x32
	OpCodeIfNe OpCode = 0x33
	OpCodeIfLt OpCode = 0x34
	OpCodeIfGe OpCode = 0x35
	OpCodeIfGt OpCode = 0x36
	OpCodeIfLe OpCode = 0x37
	OpCodeIfEqz OpCode = 0x38
	OpCodeIfNez OpCode = 0x39
	OpCodeIfLtz OpCode = 0x3a
	OpCodeIfGez OpCode = 0x3b
	OpCodeIfGtz OpCode = 0x3c
	OpCodeIfLez OpCode = 0x3d
	OpCodeAget OpCode = 0x44
	OpCodeAgetWide OpCode = 0x45
	OpCodeAgetObject OpCode = 0x46
	OpCodeAgetBoolean OpCode = 0x47
	OpCodeAgetByte OpCode = 0x48
	OpCodeAgetChar OpCode = 0x49
	OpCodeAgetShort OpCode = 0x4a
	OpCodeAput OpCode = 0x4b
	OpCodeAputWide OpCode = 0x4c
	OpCodeAputObject OpCode = 0x4d
	OpCodeAputBoolean OpCode = 0x4e
	OpCodeAputByte OpCode = 0x4f
	OpCodeAputChar OpCode = 0x50
	OpCodeAputShort OpCode = 0x51
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
	OpCodeSget OpCode = 0x60
	OpCodeSgetWide OpCode = 0x61
	OpCodeSgetObject OpCode = 0x62
	OpCodeSgetBoolean OpCode = 0x63
	OpCodeSgetByte OpCode = 0x64
	OpCodeSgetChar OpCode = 0x65
	OpCodeSgetShort OpCode = 0x66
	OpCodeSput OpCode = 0x67
	OpCodeSputWide OpCode = 0x68
	OpCodeSputObject OpCode = 0x69
	OpCodeSputBoolean OpCode = 0x6a
	OpCodeSputByte OpCode = 0x6b
	OpCodeSputChar OpCode = 0x6c
	OpCodeSputShort OpCode = 0x6d
	OpCodeInvokeVirtual OpCode = 0x6e
	OpCodeInvokeSuper OpCode = 0x6f
	OpCodeInvokeDirect OpCode = 0x70
	OpCodeInvokeStatic OpCode = 0x71
	OpCodeInvokeInterface OpCode = 0x72
	OpCodeInvokeVirtualRange OpCode = 0x74
	OpCodeInvokeSuperRange OpCode = 0x75
	OpCodeInvokeDirectRange OpCode = 0x76
	OpCodeInvokeStaticRange OpCode = 0x77
	OpCodeInvokeInterfaceRange OpCode = 0x78
	OpCodeAddIntLit8 OpCode = 0xd8
	OpCodeRsubIntLit8 OpCode = 0xd9
	OpCodeMulIntLit8 OpCode = 0xda
	OpCodeDivIntLit8 OpCode = 0xdb
	OpCodeRemIntLit8 OpCode = 0xdc
	OpCodeAndIntLit8 OpCode = 0xdd
	OpCodeOrIntLit8 OpCode = 0xde
	OpCodeXorIntLit8 OpCode = 0xdf
	OpCodeShlIntLit8 OpCode = 0xe0
	OpCodeShrIntLit8 OpCode = 0xe1
	OpCodeUshrIntLit8 OpCode = 0xe2
)

var opConfigs = map[OpCode]opConfig{
	OpCodeNop: {
		Name: "nop",
		Size: fmt10xSize,
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
		Size: fmt12xSize,
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
		Size: fmt22xSize,
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
		Size: fmt32xSize,
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
		Size: fmt12xSize,
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
		Size: fmt22xSize,
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
		Size: fmt32xSize,
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
		Size: fmt12xSize,
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
		Size: fmt22xSize,
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
		Size: fmt32xSize,
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
		Size: fmt11xSize,
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
		Size: fmt11xSize,
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
		Size: fmt11xSize,
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
		Size: fmt11xSize,
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
		Size: fmt10xSize,
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
		Size: fmt11xSize,
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
		Size: fmt11xSize,
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
		Size: fmt11xSize,
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
		Size: fmt11nSize,
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
		Size: fmt21sSize,
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
		Size: fmt31iSize,
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
		Size: fmt21hSize,
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
		Size: fmt21sSize,
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
		Size: fmt31iSize,
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
		Size: fmt21hSize,
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
		Size: fmt21cSize,
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
		Size: fmt31cSize,
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
		Size: fmt21cSize,
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
		Size: fmt11xSize,
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
		Size: fmt11xSize,
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
		Size: fmt21cSize,
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
		Size: fmt22cSize,
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
		Size: fmt12xSize,
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
		Size: fmt21cSize,
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
		Size: fmt22cSize,
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
		Size: fmt11xSize,
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
		Size: fmt10tSize,
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
	OpCodeGoto16: {
		Name: "goto/16",
		Size: fmt20tSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt20t()
			if err != nil {
				return nil, err
			}

			return &OpGoto16 {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeGoto32: {
		Name: "goto/32",
		Size: fmt30tSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt30t()
			if err != nil {
				return nil, err
			}

			return &OpGoto32 {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodePackedSwitch: {
		Name: "packed-switch",
		Size: fmt31tSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt31t()
			if err != nil {
				return nil, err
			}

			return &OpPackedSwitch {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeCmplFloat: {
		Name: "cmpl-float",
		Size: fmt23xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt23x()
			if err != nil {
				return nil, err
			}

			return &OpCmplFloat {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeCmpgFloat: {
		Name: "cmpg-float",
		Size: fmt23xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt23x()
			if err != nil {
				return nil, err
			}

			return &OpCmpgFloat {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeCmplDouble: {
		Name: "cmpl-double",
		Size: fmt23xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt23x()
			if err != nil {
				return nil, err
			}

			return &OpCmplDouble {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeCmpgDouble: {
		Name: "cmpg-double",
		Size: fmt23xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt23x()
			if err != nil {
				return nil, err
			}

			return &OpCmpgDouble {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeCmpLong: {
		Name: "cmp-long",
		Size: fmt23xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt23x()
			if err != nil {
				return nil, err
			}

			return &OpCmpLong {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeIfEq: {
		Name: "if-eq",
		Size: fmt22tSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt22t()
			if err != nil {
				return nil, err
			}

			return &OpIfEq {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeIfNe: {
		Name: "if-ne",
		Size: fmt22tSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt22t()
			if err != nil {
				return nil, err
			}

			return &OpIfNe {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeIfLt: {
		Name: "if-lt",
		Size: fmt22tSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt22t()
			if err != nil {
				return nil, err
			}

			return &OpIfLt {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeIfGe: {
		Name: "if-ge",
		Size: fmt22tSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt22t()
			if err != nil {
				return nil, err
			}

			return &OpIfGe {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeIfGt: {
		Name: "if-gt",
		Size: fmt22tSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt22t()
			if err != nil {
				return nil, err
			}

			return &OpIfGt {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeIfLe: {
		Name: "if-le",
		Size: fmt22tSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt22t()
			if err != nil {
				return nil, err
			}

			return &OpIfLe {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeIfEqz: {
		Name: "if-eqz",
		Size: fmt21tSize,
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
		Size: fmt21tSize,
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
		Size: fmt21tSize,
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
		Size: fmt21tSize,
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
		Size: fmt21tSize,
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
		Size: fmt21tSize,
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
	OpCodeAget: {
		Name: "aget",
		Size: fmt23xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt23x()
			if err != nil {
				return nil, err
			}

			return &OpAget {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeAgetWide: {
		Name: "aget-wide",
		Size: fmt23xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt23x()
			if err != nil {
				return nil, err
			}

			return &OpAgetWide {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeAgetObject: {
		Name: "aget-object",
		Size: fmt23xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt23x()
			if err != nil {
				return nil, err
			}

			return &OpAgetObject {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeAgetBoolean: {
		Name: "aget-boolean",
		Size: fmt23xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt23x()
			if err != nil {
				return nil, err
			}

			return &OpAgetBoolean {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeAgetByte: {
		Name: "aget-byte",
		Size: fmt23xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt23x()
			if err != nil {
				return nil, err
			}

			return &OpAgetByte {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeAgetChar: {
		Name: "aget-char",
		Size: fmt23xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt23x()
			if err != nil {
				return nil, err
			}

			return &OpAgetChar {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeAgetShort: {
		Name: "aget-short",
		Size: fmt23xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt23x()
			if err != nil {
				return nil, err
			}

			return &OpAgetShort {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeAput: {
		Name: "aput",
		Size: fmt23xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt23x()
			if err != nil {
				return nil, err
			}

			return &OpAput {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeAputWide: {
		Name: "aput-wide",
		Size: fmt23xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt23x()
			if err != nil {
				return nil, err
			}

			return &OpAputWide {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeAputObject: {
		Name: "aput-object",
		Size: fmt23xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt23x()
			if err != nil {
				return nil, err
			}

			return &OpAputObject {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeAputBoolean: {
		Name: "aput-boolean",
		Size: fmt23xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt23x()
			if err != nil {
				return nil, err
			}

			return &OpAputBoolean {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeAputByte: {
		Name: "aput-byte",
		Size: fmt23xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt23x()
			if err != nil {
				return nil, err
			}

			return &OpAputByte {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeAputChar: {
		Name: "aput-char",
		Size: fmt23xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt23x()
			if err != nil {
				return nil, err
			}

			return &OpAputChar {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeAputShort: {
		Name: "aput-short",
		Size: fmt23xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt23x()
			if err != nil {
				return nil, err
			}

			return &OpAputShort {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeIget: {
		Name: "iget",
		Size: fmt22cSize,
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
		Size: fmt22cSize,
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
		Size: fmt22cSize,
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
		Size: fmt22cSize,
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
		Size: fmt22cSize,
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
		Size: fmt22cSize,
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
		Size: fmt22cSize,
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
		Size: fmt22cSize,
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
		Size: fmt22cSize,
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
		Size: fmt22cSize,
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
		Size: fmt22cSize,
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
		Size: fmt22cSize,
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
		Size: fmt22cSize,
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
		Size: fmt22cSize,
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
	OpCodeSget: {
		Name: "sget",
		Size: fmt21cSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt21c()
			if err != nil {
				return nil, err
			}

			return &OpSget {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeSgetWide: {
		Name: "sget-wide",
		Size: fmt21cSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt21c()
			if err != nil {
				return nil, err
			}

			return &OpSgetWide {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeSgetObject: {
		Name: "sget-object",
		Size: fmt21cSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt21c()
			if err != nil {
				return nil, err
			}

			return &OpSgetObject {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeSgetBoolean: {
		Name: "sget-boolean",
		Size: fmt21cSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt21c()
			if err != nil {
				return nil, err
			}

			return &OpSgetBoolean {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeSgetByte: {
		Name: "sget-byte",
		Size: fmt21cSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt21c()
			if err != nil {
				return nil, err
			}

			return &OpSgetByte {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeSgetChar: {
		Name: "sget-char",
		Size: fmt21cSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt21c()
			if err != nil {
				return nil, err
			}

			return &OpSgetChar {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeSgetShort: {
		Name: "sget-short",
		Size: fmt21cSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt21c()
			if err != nil {
				return nil, err
			}

			return &OpSgetShort {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeSput: {
		Name: "sput",
		Size: fmt21cSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt21c()
			if err != nil {
				return nil, err
			}

			return &OpSput {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeSputWide: {
		Name: "sput-wide",
		Size: fmt21cSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt21c()
			if err != nil {
				return nil, err
			}

			return &OpSputWide {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeSputObject: {
		Name: "sput-object",
		Size: fmt21cSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt21c()
			if err != nil {
				return nil, err
			}

			return &OpSputObject {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeSputBoolean: {
		Name: "sput-boolean",
		Size: fmt21cSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt21c()
			if err != nil {
				return nil, err
			}

			return &OpSputBoolean {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeSputByte: {
		Name: "sput-byte",
		Size: fmt21cSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt21c()
			if err != nil {
				return nil, err
			}

			return &OpSputByte {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeSputChar: {
		Name: "sput-char",
		Size: fmt21cSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt21c()
			if err != nil {
				return nil, err
			}

			return &OpSputChar {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeSputShort: {
		Name: "sput-short",
		Size: fmt21cSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt21c()
			if err != nil {
				return nil, err
			}

			return &OpSputShort {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeInvokeVirtual: {
		Name: "invoke-virtual",
		Size: fmt35cSize,
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
		Size: fmt35cSize,
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
		Size: fmt35cSize,
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
		Size: fmt35cSize,
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
		Size: fmt35cSize,
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
	OpCodeInvokeVirtualRange: {
		Name: "invoke-virtual/range",
		Size: fmt3rcSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt3rc()
			if err != nil {
				return nil, err
			}

			return &OpInvokeVirtualRange {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeInvokeSuperRange: {
		Name: "invoke-super/range",
		Size: fmt3rcSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt3rc()
			if err != nil {
				return nil, err
			}

			return &OpInvokeSuperRange {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeInvokeDirectRange: {
		Name: "invoke-direct/range",
		Size: fmt3rcSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt3rc()
			if err != nil {
				return nil, err
			}

			return &OpInvokeDirectRange {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeInvokeStaticRange: {
		Name: "invoke-static/range",
		Size: fmt3rcSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt3rc()
			if err != nil {
				return nil, err
			}

			return &OpInvokeStaticRange {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeInvokeInterfaceRange: {
		Name: "invoke-interface/range",
		Size: fmt3rcSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt3rc()
			if err != nil {
				return nil, err
			}

			return &OpInvokeInterfaceRange {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeAddIntLit8: {
		Name: "add-int/lit8",
		Size: fmt22bSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt22b()
			if err != nil {
				return nil, err
			}

			return &OpAddIntLit8 {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeRsubIntLit8: {
		Name: "rsub-int/lit8",
		Size: fmt22bSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt22b()
			if err != nil {
				return nil, err
			}

			return &OpRsubIntLit8 {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeMulIntLit8: {
		Name: "mul-int/lit8",
		Size: fmt22bSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt22b()
			if err != nil {
				return nil, err
			}

			return &OpMulIntLit8 {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeDivIntLit8: {
		Name: "div-int/lit8",
		Size: fmt22bSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt22b()
			if err != nil {
				return nil, err
			}

			return &OpDivIntLit8 {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeRemIntLit8: {
		Name: "rem-int/lit8",
		Size: fmt22bSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt22b()
			if err != nil {
				return nil, err
			}

			return &OpRemIntLit8 {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeAndIntLit8: {
		Name: "and-int/lit8",
		Size: fmt22bSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt22b()
			if err != nil {
				return nil, err
			}

			return &OpAndIntLit8 {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeOrIntLit8: {
		Name: "or-int/lit8",
		Size: fmt22bSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt22b()
			if err != nil {
				return nil, err
			}

			return &OpOrIntLit8 {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeXorIntLit8: {
		Name: "xor-int/lit8",
		Size: fmt22bSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt22b()
			if err != nil {
				return nil, err
			}

			return &OpXorIntLit8 {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeShlIntLit8: {
		Name: "shl-int/lit8",
		Size: fmt22bSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt22b()
			if err != nil {
				return nil, err
			}

			return &OpShlIntLit8 {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeShrIntLit8: {
		Name: "shr-int/lit8",
		Size: fmt22bSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt22b()
			if err != nil {
				return nil, err
			}

			return &OpShrIntLit8 {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeUshrIntLit8: {
		Name: "ushr-int/lit8",
		Size: fmt22bSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt22b()
			if err != nil {
				return nil, err
			}

			return &OpUshrIntLit8 {
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

type OpGoto16 struct {
	opBase
	Fmt20t
}

func (o OpGoto16) Code() OpCode {
	return OpCodeGoto16
}

func (o OpGoto16) Fmt() Fmt {
	return o.Fmt20t
}

func (o OpGoto16) String() string {
	f := o.Fmt20t.String()
	if f != "" {
		return fmt.Sprintf("0x%x: goto/16 %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: goto/16", o.pos)
}

type OpGoto32 struct {
	opBase
	Fmt30t
}

func (o OpGoto32) Code() OpCode {
	return OpCodeGoto32
}

func (o OpGoto32) Fmt() Fmt {
	return o.Fmt30t
}

func (o OpGoto32) String() string {
	f := o.Fmt30t.String()
	if f != "" {
		return fmt.Sprintf("0x%x: goto/32 %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: goto/32", o.pos)
}

type OpPackedSwitch struct {
	opBase
	Fmt31t
}

func (o OpPackedSwitch) Code() OpCode {
	return OpCodePackedSwitch
}

func (o OpPackedSwitch) Fmt() Fmt {
	return o.Fmt31t
}

func (o OpPackedSwitch) String() string {
	f := o.Fmt31t.String()
	if f != "" {
		return fmt.Sprintf("0x%x: packed-switch %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: packed-switch", o.pos)
}

type OpCmplFloat struct {
	opBase
	Fmt23x
}

func (o OpCmplFloat) Code() OpCode {
	return OpCodeCmplFloat
}

func (o OpCmplFloat) Fmt() Fmt {
	return o.Fmt23x
}

func (o OpCmplFloat) String() string {
	f := o.Fmt23x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: cmpl-float %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: cmpl-float", o.pos)
}

type OpCmpgFloat struct {
	opBase
	Fmt23x
}

func (o OpCmpgFloat) Code() OpCode {
	return OpCodeCmpgFloat
}

func (o OpCmpgFloat) Fmt() Fmt {
	return o.Fmt23x
}

func (o OpCmpgFloat) String() string {
	f := o.Fmt23x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: cmpg-float %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: cmpg-float", o.pos)
}

type OpCmplDouble struct {
	opBase
	Fmt23x
}

func (o OpCmplDouble) Code() OpCode {
	return OpCodeCmplDouble
}

func (o OpCmplDouble) Fmt() Fmt {
	return o.Fmt23x
}

func (o OpCmplDouble) String() string {
	f := o.Fmt23x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: cmpl-double %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: cmpl-double", o.pos)
}

type OpCmpgDouble struct {
	opBase
	Fmt23x
}

func (o OpCmpgDouble) Code() OpCode {
	return OpCodeCmpgDouble
}

func (o OpCmpgDouble) Fmt() Fmt {
	return o.Fmt23x
}

func (o OpCmpgDouble) String() string {
	f := o.Fmt23x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: cmpg-double %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: cmpg-double", o.pos)
}

type OpCmpLong struct {
	opBase
	Fmt23x
}

func (o OpCmpLong) Code() OpCode {
	return OpCodeCmpLong
}

func (o OpCmpLong) Fmt() Fmt {
	return o.Fmt23x
}

func (o OpCmpLong) String() string {
	f := o.Fmt23x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: cmp-long %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: cmp-long", o.pos)
}

type OpIfEq struct {
	opBase
	Fmt22t
}

func (o OpIfEq) Code() OpCode {
	return OpCodeIfEq
}

func (o OpIfEq) Fmt() Fmt {
	return o.Fmt22t
}

func (o OpIfEq) String() string {
	f := o.Fmt22t.String()
	if f != "" {
		return fmt.Sprintf("0x%x: if-eq %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: if-eq", o.pos)
}

type OpIfNe struct {
	opBase
	Fmt22t
}

func (o OpIfNe) Code() OpCode {
	return OpCodeIfNe
}

func (o OpIfNe) Fmt() Fmt {
	return o.Fmt22t
}

func (o OpIfNe) String() string {
	f := o.Fmt22t.String()
	if f != "" {
		return fmt.Sprintf("0x%x: if-ne %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: if-ne", o.pos)
}

type OpIfLt struct {
	opBase
	Fmt22t
}

func (o OpIfLt) Code() OpCode {
	return OpCodeIfLt
}

func (o OpIfLt) Fmt() Fmt {
	return o.Fmt22t
}

func (o OpIfLt) String() string {
	f := o.Fmt22t.String()
	if f != "" {
		return fmt.Sprintf("0x%x: if-lt %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: if-lt", o.pos)
}

type OpIfGe struct {
	opBase
	Fmt22t
}

func (o OpIfGe) Code() OpCode {
	return OpCodeIfGe
}

func (o OpIfGe) Fmt() Fmt {
	return o.Fmt22t
}

func (o OpIfGe) String() string {
	f := o.Fmt22t.String()
	if f != "" {
		return fmt.Sprintf("0x%x: if-ge %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: if-ge", o.pos)
}

type OpIfGt struct {
	opBase
	Fmt22t
}

func (o OpIfGt) Code() OpCode {
	return OpCodeIfGt
}

func (o OpIfGt) Fmt() Fmt {
	return o.Fmt22t
}

func (o OpIfGt) String() string {
	f := o.Fmt22t.String()
	if f != "" {
		return fmt.Sprintf("0x%x: if-gt %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: if-gt", o.pos)
}

type OpIfLe struct {
	opBase
	Fmt22t
}

func (o OpIfLe) Code() OpCode {
	return OpCodeIfLe
}

func (o OpIfLe) Fmt() Fmt {
	return o.Fmt22t
}

func (o OpIfLe) String() string {
	f := o.Fmt22t.String()
	if f != "" {
		return fmt.Sprintf("0x%x: if-le %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: if-le", o.pos)
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

type OpAget struct {
	opBase
	Fmt23x
}

func (o OpAget) Code() OpCode {
	return OpCodeAget
}

func (o OpAget) Fmt() Fmt {
	return o.Fmt23x
}

func (o OpAget) String() string {
	f := o.Fmt23x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: aget %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: aget", o.pos)
}

type OpAgetWide struct {
	opBase
	Fmt23x
}

func (o OpAgetWide) Code() OpCode {
	return OpCodeAgetWide
}

func (o OpAgetWide) Fmt() Fmt {
	return o.Fmt23x
}

func (o OpAgetWide) String() string {
	f := o.Fmt23x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: aget-wide %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: aget-wide", o.pos)
}

type OpAgetObject struct {
	opBase
	Fmt23x
}

func (o OpAgetObject) Code() OpCode {
	return OpCodeAgetObject
}

func (o OpAgetObject) Fmt() Fmt {
	return o.Fmt23x
}

func (o OpAgetObject) String() string {
	f := o.Fmt23x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: aget-object %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: aget-object", o.pos)
}

type OpAgetBoolean struct {
	opBase
	Fmt23x
}

func (o OpAgetBoolean) Code() OpCode {
	return OpCodeAgetBoolean
}

func (o OpAgetBoolean) Fmt() Fmt {
	return o.Fmt23x
}

func (o OpAgetBoolean) String() string {
	f := o.Fmt23x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: aget-boolean %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: aget-boolean", o.pos)
}

type OpAgetByte struct {
	opBase
	Fmt23x
}

func (o OpAgetByte) Code() OpCode {
	return OpCodeAgetByte
}

func (o OpAgetByte) Fmt() Fmt {
	return o.Fmt23x
}

func (o OpAgetByte) String() string {
	f := o.Fmt23x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: aget-byte %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: aget-byte", o.pos)
}

type OpAgetChar struct {
	opBase
	Fmt23x
}

func (o OpAgetChar) Code() OpCode {
	return OpCodeAgetChar
}

func (o OpAgetChar) Fmt() Fmt {
	return o.Fmt23x
}

func (o OpAgetChar) String() string {
	f := o.Fmt23x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: aget-char %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: aget-char", o.pos)
}

type OpAgetShort struct {
	opBase
	Fmt23x
}

func (o OpAgetShort) Code() OpCode {
	return OpCodeAgetShort
}

func (o OpAgetShort) Fmt() Fmt {
	return o.Fmt23x
}

func (o OpAgetShort) String() string {
	f := o.Fmt23x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: aget-short %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: aget-short", o.pos)
}

type OpAput struct {
	opBase
	Fmt23x
}

func (o OpAput) Code() OpCode {
	return OpCodeAput
}

func (o OpAput) Fmt() Fmt {
	return o.Fmt23x
}

func (o OpAput) String() string {
	f := o.Fmt23x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: aput %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: aput", o.pos)
}

type OpAputWide struct {
	opBase
	Fmt23x
}

func (o OpAputWide) Code() OpCode {
	return OpCodeAputWide
}

func (o OpAputWide) Fmt() Fmt {
	return o.Fmt23x
}

func (o OpAputWide) String() string {
	f := o.Fmt23x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: aput-wide %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: aput-wide", o.pos)
}

type OpAputObject struct {
	opBase
	Fmt23x
}

func (o OpAputObject) Code() OpCode {
	return OpCodeAputObject
}

func (o OpAputObject) Fmt() Fmt {
	return o.Fmt23x
}

func (o OpAputObject) String() string {
	f := o.Fmt23x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: aput-object %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: aput-object", o.pos)
}

type OpAputBoolean struct {
	opBase
	Fmt23x
}

func (o OpAputBoolean) Code() OpCode {
	return OpCodeAputBoolean
}

func (o OpAputBoolean) Fmt() Fmt {
	return o.Fmt23x
}

func (o OpAputBoolean) String() string {
	f := o.Fmt23x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: aput-boolean %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: aput-boolean", o.pos)
}

type OpAputByte struct {
	opBase
	Fmt23x
}

func (o OpAputByte) Code() OpCode {
	return OpCodeAputByte
}

func (o OpAputByte) Fmt() Fmt {
	return o.Fmt23x
}

func (o OpAputByte) String() string {
	f := o.Fmt23x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: aput-byte %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: aput-byte", o.pos)
}

type OpAputChar struct {
	opBase
	Fmt23x
}

func (o OpAputChar) Code() OpCode {
	return OpCodeAputChar
}

func (o OpAputChar) Fmt() Fmt {
	return o.Fmt23x
}

func (o OpAputChar) String() string {
	f := o.Fmt23x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: aput-char %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: aput-char", o.pos)
}

type OpAputShort struct {
	opBase
	Fmt23x
}

func (o OpAputShort) Code() OpCode {
	return OpCodeAputShort
}

func (o OpAputShort) Fmt() Fmt {
	return o.Fmt23x
}

func (o OpAputShort) String() string {
	f := o.Fmt23x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: aput-short %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: aput-short", o.pos)
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

type OpSget struct {
	opBase
	Fmt21c
}

func (o OpSget) Code() OpCode {
	return OpCodeSget
}

func (o OpSget) Fmt() Fmt {
	return o.Fmt21c
}

func (o OpSget) String() string {
	f := o.Fmt21c.String()
	if f != "" {
		return fmt.Sprintf("0x%x: sget %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: sget", o.pos)
}

type OpSgetWide struct {
	opBase
	Fmt21c
}

func (o OpSgetWide) Code() OpCode {
	return OpCodeSgetWide
}

func (o OpSgetWide) Fmt() Fmt {
	return o.Fmt21c
}

func (o OpSgetWide) String() string {
	f := o.Fmt21c.String()
	if f != "" {
		return fmt.Sprintf("0x%x: sget-wide %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: sget-wide", o.pos)
}

type OpSgetObject struct {
	opBase
	Fmt21c
}

func (o OpSgetObject) Code() OpCode {
	return OpCodeSgetObject
}

func (o OpSgetObject) Fmt() Fmt {
	return o.Fmt21c
}

func (o OpSgetObject) String() string {
	f := o.Fmt21c.String()
	if f != "" {
		return fmt.Sprintf("0x%x: sget-object %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: sget-object", o.pos)
}

type OpSgetBoolean struct {
	opBase
	Fmt21c
}

func (o OpSgetBoolean) Code() OpCode {
	return OpCodeSgetBoolean
}

func (o OpSgetBoolean) Fmt() Fmt {
	return o.Fmt21c
}

func (o OpSgetBoolean) String() string {
	f := o.Fmt21c.String()
	if f != "" {
		return fmt.Sprintf("0x%x: sget-boolean %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: sget-boolean", o.pos)
}

type OpSgetByte struct {
	opBase
	Fmt21c
}

func (o OpSgetByte) Code() OpCode {
	return OpCodeSgetByte
}

func (o OpSgetByte) Fmt() Fmt {
	return o.Fmt21c
}

func (o OpSgetByte) String() string {
	f := o.Fmt21c.String()
	if f != "" {
		return fmt.Sprintf("0x%x: sget-byte %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: sget-byte", o.pos)
}

type OpSgetChar struct {
	opBase
	Fmt21c
}

func (o OpSgetChar) Code() OpCode {
	return OpCodeSgetChar
}

func (o OpSgetChar) Fmt() Fmt {
	return o.Fmt21c
}

func (o OpSgetChar) String() string {
	f := o.Fmt21c.String()
	if f != "" {
		return fmt.Sprintf("0x%x: sget-char %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: sget-char", o.pos)
}

type OpSgetShort struct {
	opBase
	Fmt21c
}

func (o OpSgetShort) Code() OpCode {
	return OpCodeSgetShort
}

func (o OpSgetShort) Fmt() Fmt {
	return o.Fmt21c
}

func (o OpSgetShort) String() string {
	f := o.Fmt21c.String()
	if f != "" {
		return fmt.Sprintf("0x%x: sget-short %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: sget-short", o.pos)
}

type OpSput struct {
	opBase
	Fmt21c
}

func (o OpSput) Code() OpCode {
	return OpCodeSput
}

func (o OpSput) Fmt() Fmt {
	return o.Fmt21c
}

func (o OpSput) String() string {
	f := o.Fmt21c.String()
	if f != "" {
		return fmt.Sprintf("0x%x: sput %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: sput", o.pos)
}

type OpSputWide struct {
	opBase
	Fmt21c
}

func (o OpSputWide) Code() OpCode {
	return OpCodeSputWide
}

func (o OpSputWide) Fmt() Fmt {
	return o.Fmt21c
}

func (o OpSputWide) String() string {
	f := o.Fmt21c.String()
	if f != "" {
		return fmt.Sprintf("0x%x: sput-wide %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: sput-wide", o.pos)
}

type OpSputObject struct {
	opBase
	Fmt21c
}

func (o OpSputObject) Code() OpCode {
	return OpCodeSputObject
}

func (o OpSputObject) Fmt() Fmt {
	return o.Fmt21c
}

func (o OpSputObject) String() string {
	f := o.Fmt21c.String()
	if f != "" {
		return fmt.Sprintf("0x%x: sput-object %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: sput-object", o.pos)
}

type OpSputBoolean struct {
	opBase
	Fmt21c
}

func (o OpSputBoolean) Code() OpCode {
	return OpCodeSputBoolean
}

func (o OpSputBoolean) Fmt() Fmt {
	return o.Fmt21c
}

func (o OpSputBoolean) String() string {
	f := o.Fmt21c.String()
	if f != "" {
		return fmt.Sprintf("0x%x: sput-boolean %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: sput-boolean", o.pos)
}

type OpSputByte struct {
	opBase
	Fmt21c
}

func (o OpSputByte) Code() OpCode {
	return OpCodeSputByte
}

func (o OpSputByte) Fmt() Fmt {
	return o.Fmt21c
}

func (o OpSputByte) String() string {
	f := o.Fmt21c.String()
	if f != "" {
		return fmt.Sprintf("0x%x: sput-byte %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: sput-byte", o.pos)
}

type OpSputChar struct {
	opBase
	Fmt21c
}

func (o OpSputChar) Code() OpCode {
	return OpCodeSputChar
}

func (o OpSputChar) Fmt() Fmt {
	return o.Fmt21c
}

func (o OpSputChar) String() string {
	f := o.Fmt21c.String()
	if f != "" {
		return fmt.Sprintf("0x%x: sput-char %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: sput-char", o.pos)
}

type OpSputShort struct {
	opBase
	Fmt21c
}

func (o OpSputShort) Code() OpCode {
	return OpCodeSputShort
}

func (o OpSputShort) Fmt() Fmt {
	return o.Fmt21c
}

func (o OpSputShort) String() string {
	f := o.Fmt21c.String()
	if f != "" {
		return fmt.Sprintf("0x%x: sput-short %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: sput-short", o.pos)
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

type OpInvokeVirtualRange struct {
	opBase
	Fmt3rc
}

func (o OpInvokeVirtualRange) Code() OpCode {
	return OpCodeInvokeVirtualRange
}

func (o OpInvokeVirtualRange) Fmt() Fmt {
	return o.Fmt3rc
}

func (o OpInvokeVirtualRange) String() string {
	f := o.Fmt3rc.String()
	if f != "" {
		return fmt.Sprintf("0x%x: invoke-virtual/range %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: invoke-virtual/range", o.pos)
}

type OpInvokeSuperRange struct {
	opBase
	Fmt3rc
}

func (o OpInvokeSuperRange) Code() OpCode {
	return OpCodeInvokeSuperRange
}

func (o OpInvokeSuperRange) Fmt() Fmt {
	return o.Fmt3rc
}

func (o OpInvokeSuperRange) String() string {
	f := o.Fmt3rc.String()
	if f != "" {
		return fmt.Sprintf("0x%x: invoke-super/range %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: invoke-super/range", o.pos)
}

type OpInvokeDirectRange struct {
	opBase
	Fmt3rc
}

func (o OpInvokeDirectRange) Code() OpCode {
	return OpCodeInvokeDirectRange
}

func (o OpInvokeDirectRange) Fmt() Fmt {
	return o.Fmt3rc
}

func (o OpInvokeDirectRange) String() string {
	f := o.Fmt3rc.String()
	if f != "" {
		return fmt.Sprintf("0x%x: invoke-direct/range %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: invoke-direct/range", o.pos)
}

type OpInvokeStaticRange struct {
	opBase
	Fmt3rc
}

func (o OpInvokeStaticRange) Code() OpCode {
	return OpCodeInvokeStaticRange
}

func (o OpInvokeStaticRange) Fmt() Fmt {
	return o.Fmt3rc
}

func (o OpInvokeStaticRange) String() string {
	f := o.Fmt3rc.String()
	if f != "" {
		return fmt.Sprintf("0x%x: invoke-static/range %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: invoke-static/range", o.pos)
}

type OpInvokeInterfaceRange struct {
	opBase
	Fmt3rc
}

func (o OpInvokeInterfaceRange) Code() OpCode {
	return OpCodeInvokeInterfaceRange
}

func (o OpInvokeInterfaceRange) Fmt() Fmt {
	return o.Fmt3rc
}

func (o OpInvokeInterfaceRange) String() string {
	f := o.Fmt3rc.String()
	if f != "" {
		return fmt.Sprintf("0x%x: invoke-interface/range %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: invoke-interface/range", o.pos)
}

type OpAddIntLit8 struct {
	opBase
	Fmt22b
}

func (o OpAddIntLit8) Code() OpCode {
	return OpCodeAddIntLit8
}

func (o OpAddIntLit8) Fmt() Fmt {
	return o.Fmt22b
}

func (o OpAddIntLit8) String() string {
	f := o.Fmt22b.String()
	if f != "" {
		return fmt.Sprintf("0x%x: add-int/lit8 %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: add-int/lit8", o.pos)
}

type OpRsubIntLit8 struct {
	opBase
	Fmt22b
}

func (o OpRsubIntLit8) Code() OpCode {
	return OpCodeRsubIntLit8
}

func (o OpRsubIntLit8) Fmt() Fmt {
	return o.Fmt22b
}

func (o OpRsubIntLit8) String() string {
	f := o.Fmt22b.String()
	if f != "" {
		return fmt.Sprintf("0x%x: rsub-int/lit8 %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: rsub-int/lit8", o.pos)
}

type OpMulIntLit8 struct {
	opBase
	Fmt22b
}

func (o OpMulIntLit8) Code() OpCode {
	return OpCodeMulIntLit8
}

func (o OpMulIntLit8) Fmt() Fmt {
	return o.Fmt22b
}

func (o OpMulIntLit8) String() string {
	f := o.Fmt22b.String()
	if f != "" {
		return fmt.Sprintf("0x%x: mul-int/lit8 %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: mul-int/lit8", o.pos)
}

type OpDivIntLit8 struct {
	opBase
	Fmt22b
}

func (o OpDivIntLit8) Code() OpCode {
	return OpCodeDivIntLit8
}

func (o OpDivIntLit8) Fmt() Fmt {
	return o.Fmt22b
}

func (o OpDivIntLit8) String() string {
	f := o.Fmt22b.String()
	if f != "" {
		return fmt.Sprintf("0x%x: div-int/lit8 %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: div-int/lit8", o.pos)
}

type OpRemIntLit8 struct {
	opBase
	Fmt22b
}

func (o OpRemIntLit8) Code() OpCode {
	return OpCodeRemIntLit8
}

func (o OpRemIntLit8) Fmt() Fmt {
	return o.Fmt22b
}

func (o OpRemIntLit8) String() string {
	f := o.Fmt22b.String()
	if f != "" {
		return fmt.Sprintf("0x%x: rem-int/lit8 %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: rem-int/lit8", o.pos)
}

type OpAndIntLit8 struct {
	opBase
	Fmt22b
}

func (o OpAndIntLit8) Code() OpCode {
	return OpCodeAndIntLit8
}

func (o OpAndIntLit8) Fmt() Fmt {
	return o.Fmt22b
}

func (o OpAndIntLit8) String() string {
	f := o.Fmt22b.String()
	if f != "" {
		return fmt.Sprintf("0x%x: and-int/lit8 %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: and-int/lit8", o.pos)
}

type OpOrIntLit8 struct {
	opBase
	Fmt22b
}

func (o OpOrIntLit8) Code() OpCode {
	return OpCodeOrIntLit8
}

func (o OpOrIntLit8) Fmt() Fmt {
	return o.Fmt22b
}

func (o OpOrIntLit8) String() string {
	f := o.Fmt22b.String()
	if f != "" {
		return fmt.Sprintf("0x%x: or-int/lit8 %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: or-int/lit8", o.pos)
}

type OpXorIntLit8 struct {
	opBase
	Fmt22b
}

func (o OpXorIntLit8) Code() OpCode {
	return OpCodeXorIntLit8
}

func (o OpXorIntLit8) Fmt() Fmt {
	return o.Fmt22b
}

func (o OpXorIntLit8) String() string {
	f := o.Fmt22b.String()
	if f != "" {
		return fmt.Sprintf("0x%x: xor-int/lit8 %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: xor-int/lit8", o.pos)
}

type OpShlIntLit8 struct {
	opBase
	Fmt22b
}

func (o OpShlIntLit8) Code() OpCode {
	return OpCodeShlIntLit8
}

func (o OpShlIntLit8) Fmt() Fmt {
	return o.Fmt22b
}

func (o OpShlIntLit8) String() string {
	f := o.Fmt22b.String()
	if f != "" {
		return fmt.Sprintf("0x%x: shl-int/lit8 %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: shl-int/lit8", o.pos)
}

type OpShrIntLit8 struct {
	opBase
	Fmt22b
}

func (o OpShrIntLit8) Code() OpCode {
	return OpCodeShrIntLit8
}

func (o OpShrIntLit8) Fmt() Fmt {
	return o.Fmt22b
}

func (o OpShrIntLit8) String() string {
	f := o.Fmt22b.String()
	if f != "" {
		return fmt.Sprintf("0x%x: shr-int/lit8 %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: shr-int/lit8", o.pos)
}

type OpUshrIntLit8 struct {
	opBase
	Fmt22b
}

func (o OpUshrIntLit8) Code() OpCode {
	return OpCodeUshrIntLit8
}

func (o OpUshrIntLit8) Fmt() Fmt {
	return o.Fmt22b
}

func (o OpUshrIntLit8) String() string {
	f := o.Fmt22b.String()
	if f != "" {
		return fmt.Sprintf("0x%x: ushr-int/lit8 %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: ushr-int/lit8", o.pos)
}
