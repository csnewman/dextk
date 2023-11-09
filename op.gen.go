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
	OpCodeConstWide OpCode = 0x18
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
	OpCodeFilledNewArray OpCode = 0x24
	OpCodeFilledNewArrayRange OpCode = 0x25
	OpCodeFillArrayData OpCode = 0x26
	OpCodeThrow OpCode = 0x27
	OpCodeGoto OpCode = 0x28
	OpCodeGoto16 OpCode = 0x29
	OpCodeGoto32 OpCode = 0x2a
	OpCodePackedSwitch OpCode = 0x2b
	OpCodeSparseSwitch OpCode = 0x2c
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
	OpCodeNegInt OpCode = 0x7b
	OpCodeNotInt OpCode = 0x7c
	OpCodeNegLong OpCode = 0x7d
	OpCodeNotLong OpCode = 0x7e
	OpCodeNegFloat OpCode = 0x7f
	OpCodeNegDouble OpCode = 0x80
	OpCodeIntToLong OpCode = 0x81
	OpCodeIntToFloat OpCode = 0x82
	OpCodeIntToDouble OpCode = 0x83
	OpCodeLongToInt OpCode = 0x84
	OpCodeLongToFloat OpCode = 0x85
	OpCodeLongToDouble OpCode = 0x86
	OpCodeFloatToInt OpCode = 0x87
	OpCodeFloatToLong OpCode = 0x88
	OpCodeFloatToDouble OpCode = 0x89
	OpCodeDoubleToInt OpCode = 0x8a
	OpCodeDoubleToLong OpCode = 0x8b
	OpCodeDoubleToFloat OpCode = 0x8c
	OpCodeIntToByte OpCode = 0x8d
	OpCodeIntToChar OpCode = 0x8e
	OpCodeIntToShort OpCode = 0x8f
	OpCodeAddInt OpCode = 0x90
	OpCodeSubInt OpCode = 0x91
	OpCodeMulInt OpCode = 0x92
	OpCodeDivInt OpCode = 0x93
	OpCodeRemInt OpCode = 0x94
	OpCodeAndInt OpCode = 0x95
	OpCodeOrInt OpCode = 0x96
	OpCodeXorInt OpCode = 0x97
	OpCodeShlInt OpCode = 0x98
	OpCodeShrInt OpCode = 0x99
	OpCodeUshrInt OpCode = 0x9a
	OpCodeAddLong OpCode = 0x9b
	OpCodeSubLong OpCode = 0x9c
	OpCodeMulLong OpCode = 0x9d
	OpCodeDivLong OpCode = 0x9e
	OpCodeRemLong OpCode = 0x9f
	OpCodeAndLong OpCode = 0xa0
	OpCodeOrLong OpCode = 0xa1
	OpCodeXorLong OpCode = 0xa2
	OpCodeShlLong OpCode = 0xa3
	OpCodeShrLong OpCode = 0xa4
	OpCodeUshrLong OpCode = 0xa5
	OpCodeAddFloat OpCode = 0xa6
	OpCodeSubFloat OpCode = 0xa7
	OpCodeMulFloat OpCode = 0xa8
	OpCodeDivFloat OpCode = 0xa9
	OpCodeRemFloat OpCode = 0xaa
	OpCodeAddDouble OpCode = 0xab
	OpCodeSubDouble OpCode = 0xac
	OpCodeMulDouble OpCode = 0xad
	OpCodeDivDouble OpCode = 0xae
	OpCodeRemDouble OpCode = 0xaf
	OpCodeAddInt2Addr OpCode = 0xb0
	OpCodeSubInt2Addr OpCode = 0xb1
	OpCodeMulInt2Addr OpCode = 0xb2
	OpCodeDivInt2Addr OpCode = 0xb3
	OpCodeRemInt2Addr OpCode = 0xb4
	OpCodeAndInt2Addr OpCode = 0xb5
	OpCodeOrInt2Addr OpCode = 0xb6
	OpCodeXorInt2Addr OpCode = 0xb7
	OpCodeShlInt2Addr OpCode = 0xb8
	OpCodeShrInt2Addr OpCode = 0xb9
	OpCodeUshrInt2Addr OpCode = 0xba
	OpCodeAddLong2Addr OpCode = 0xbb
	OpCodeSubLong2Addr OpCode = 0xbc
	OpCodeMulLong2Addr OpCode = 0xbd
	OpCodeDivLong2Addr OpCode = 0xbe
	OpCodeRemLong2Addr OpCode = 0xbf
	OpCodeAndLong2Addr OpCode = 0xc0
	OpCodeOrLong2Addr OpCode = 0xc1
	OpCodeXorLong2Addr OpCode = 0xc2
	OpCodeShlLong2Addr OpCode = 0xc3
	OpCodeShrLong2Addr OpCode = 0xc4
	OpCodeUshrLong2Addr OpCode = 0xc5
	OpCodeAddFloat2Addr OpCode = 0xc6
	OpCodeSubFloat2Addr OpCode = 0xc7
	OpCodeMulFloat2Addr OpCode = 0xc8
	OpCodeDivFloat2Addr OpCode = 0xc9
	OpCodeRemFloat2Addr OpCode = 0xca
	OpCodeAddDouble2Addr OpCode = 0xcb
	OpCodeSubDouble2Addr OpCode = 0xcc
	OpCodeMulDouble2Addr OpCode = 0xcd
	OpCodeDivDouble2Addr OpCode = 0xce
	OpCodeRemDouble2Addr OpCode = 0xcf
	OpCodeAddIntLit16 OpCode = 0xd0
	OpCodeRsubIntLit16 OpCode = 0xd1
	OpCodeMulIntLit16 OpCode = 0xd2
	OpCodeDivIntLit16 OpCode = 0xd3
	OpCodeRemIntLit16 OpCode = 0xd4
	OpCodeAndIntLit16 OpCode = 0xd5
	OpCodeOrIntLit16 OpCode = 0xd6
	OpCodeXorIntLit16 OpCode = 0xd7
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
	OpCodeInvokePolymorphic OpCode = 0xfa
	OpCodeInvokePolymorphicRange OpCode = 0xfb
	OpCodeInvokeCustom OpCode = 0xfc
	OpCodeInvokeCustomRange OpCode = 0xfd
	OpCodeConstMethodHandle OpCode = 0xfe
	OpCodeConstMethodType OpCode = 0xff
)

var opConfigsExtra = map[OpCode]opConfig{}

var opConfigs = [256]opConfig{
	OpCodeNop: {
		Name: "nop",
		Size: fmt10xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt10x()
			if err != nil {
				return nil, err
			}

			return OpNop {
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

			return OpMove {
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

			return OpMoveFrom16 {
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

			return OpMove16 {
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

			return OpMoveWide {
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

			return OpMoveWideFrom16 {
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

			return OpMoveWide16 {
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

			return OpMoveObject {
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

			return OpMoveObjectFrom16 {
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

			return OpMoveObject16 {
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

			return OpMoveResult {
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

			return OpMoveResultWide {
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

			return OpMoveResultObject {
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

			return OpMoveException {
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

			return OpReturnVoid {
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

			return OpReturn {
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

			return OpReturnWide {
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

			return OpReturnObject {
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

			return OpConst4 {
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

			return OpConst16 {
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

			return OpConst {
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

			return OpConstHigh16 {
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

			return OpConstWide16 {
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

			return OpConstWide32 {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeConstWide: {
		Name: "const-wide",
		Size: fmt51lSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt51l()
			if err != nil {
				return nil, err
			}

			return OpConstWide {
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

			return OpConstWideHigh16 {
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

			return OpConstString {
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

			return OpConstStringJumbo {
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

			return OpConstClass {
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

			return OpMonitorEnter {
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

			return OpMonitorExit {
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

			return OpCheckCast {
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

			return OpInstanceOf {
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

			return OpArrayLength {
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

			return OpNewInstance {
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

			return OpNewArray {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeFilledNewArray: {
		Name: "filled-new-array",
		Size: fmt35cSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt35c()
			if err != nil {
				return nil, err
			}

			return OpFilledNewArray {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeFilledNewArrayRange: {
		Name: "filled-new-array/range",
		Size: fmt3rcSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt3rc()
			if err != nil {
				return nil, err
			}

			return OpFilledNewArrayRange {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeFillArrayData: {
		Name: "fill-array-data",
		Size: fmt31tSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt31t()
			if err != nil {
				return nil, err
			}

			return OpFillArrayData {
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

			return OpThrow {
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

			return OpGoto {
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

			return OpGoto16 {
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

			return OpGoto32 {
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

			return OpPackedSwitch {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeSparseSwitch: {
		Name: "sparse-switch",
		Size: fmt31tSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt31t()
			if err != nil {
				return nil, err
			}

			return OpSparseSwitch {
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

			return OpCmplFloat {
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

			return OpCmpgFloat {
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

			return OpCmplDouble {
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

			return OpCmpgDouble {
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

			return OpCmpLong {
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

			return OpIfEq {
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

			return OpIfNe {
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

			return OpIfLt {
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

			return OpIfGe {
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

			return OpIfGt {
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

			return OpIfLe {
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

			return OpIfEqz {
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

			return OpIfNez {
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

			return OpIfLtz {
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

			return OpIfGez {
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

			return OpIfGtz {
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

			return OpIfLez {
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

			return OpAget {
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

			return OpAgetWide {
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

			return OpAgetObject {
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

			return OpAgetBoolean {
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

			return OpAgetByte {
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

			return OpAgetChar {
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

			return OpAgetShort {
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

			return OpAput {
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

			return OpAputWide {
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

			return OpAputObject {
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

			return OpAputBoolean {
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

			return OpAputByte {
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

			return OpAputChar {
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

			return OpAputShort {
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

			return OpIget {
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

			return OpIgetWide {
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

			return OpIgetObject {
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

			return OpIgetBoolean {
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

			return OpIgetByte {
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

			return OpIgetChar {
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

			return OpIgetShort {
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

			return OpIput {
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

			return OpIputWide {
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

			return OpIputObject {
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

			return OpIputBoolean {
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

			return OpIputByte {
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

			return OpIputChar {
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

			return OpIputShort {
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

			return OpSget {
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

			return OpSgetWide {
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

			return OpSgetObject {
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

			return OpSgetBoolean {
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

			return OpSgetByte {
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

			return OpSgetChar {
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

			return OpSgetShort {
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

			return OpSput {
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

			return OpSputWide {
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

			return OpSputObject {
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

			return OpSputBoolean {
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

			return OpSputByte {
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

			return OpSputChar {
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

			return OpSputShort {
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

			return OpInvokeVirtual {
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

			return OpInvokeSuper {
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

			return OpInvokeDirect {
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

			return OpInvokeStatic {
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

			return OpInvokeInterface {
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

			return OpInvokeVirtualRange {
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

			return OpInvokeSuperRange {
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

			return OpInvokeDirectRange {
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

			return OpInvokeStaticRange {
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

			return OpInvokeInterfaceRange {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeNegInt: {
		Name: "neg-int",
		Size: fmt12xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt12x()
			if err != nil {
				return nil, err
			}

			return OpNegInt {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeNotInt: {
		Name: "not-int",
		Size: fmt12xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt12x()
			if err != nil {
				return nil, err
			}

			return OpNotInt {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeNegLong: {
		Name: "neg-long",
		Size: fmt12xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt12x()
			if err != nil {
				return nil, err
			}

			return OpNegLong {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeNotLong: {
		Name: "not-long",
		Size: fmt12xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt12x()
			if err != nil {
				return nil, err
			}

			return OpNotLong {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeNegFloat: {
		Name: "neg-float",
		Size: fmt12xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt12x()
			if err != nil {
				return nil, err
			}

			return OpNegFloat {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeNegDouble: {
		Name: "neg-double",
		Size: fmt12xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt12x()
			if err != nil {
				return nil, err
			}

			return OpNegDouble {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeIntToLong: {
		Name: "int-to-long",
		Size: fmt12xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt12x()
			if err != nil {
				return nil, err
			}

			return OpIntToLong {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeIntToFloat: {
		Name: "int-to-float",
		Size: fmt12xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt12x()
			if err != nil {
				return nil, err
			}

			return OpIntToFloat {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeIntToDouble: {
		Name: "int-to-double",
		Size: fmt12xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt12x()
			if err != nil {
				return nil, err
			}

			return OpIntToDouble {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeLongToInt: {
		Name: "long-to-int",
		Size: fmt12xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt12x()
			if err != nil {
				return nil, err
			}

			return OpLongToInt {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeLongToFloat: {
		Name: "long-to-float",
		Size: fmt12xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt12x()
			if err != nil {
				return nil, err
			}

			return OpLongToFloat {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeLongToDouble: {
		Name: "long-to-double",
		Size: fmt12xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt12x()
			if err != nil {
				return nil, err
			}

			return OpLongToDouble {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeFloatToInt: {
		Name: "float-to-int",
		Size: fmt12xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt12x()
			if err != nil {
				return nil, err
			}

			return OpFloatToInt {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeFloatToLong: {
		Name: "float-to-long",
		Size: fmt12xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt12x()
			if err != nil {
				return nil, err
			}

			return OpFloatToLong {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeFloatToDouble: {
		Name: "float-to-double",
		Size: fmt12xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt12x()
			if err != nil {
				return nil, err
			}

			return OpFloatToDouble {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeDoubleToInt: {
		Name: "double-to-int",
		Size: fmt12xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt12x()
			if err != nil {
				return nil, err
			}

			return OpDoubleToInt {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeDoubleToLong: {
		Name: "double-to-long",
		Size: fmt12xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt12x()
			if err != nil {
				return nil, err
			}

			return OpDoubleToLong {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeDoubleToFloat: {
		Name: "double-to-float",
		Size: fmt12xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt12x()
			if err != nil {
				return nil, err
			}

			return OpDoubleToFloat {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeIntToByte: {
		Name: "int-to-byte",
		Size: fmt12xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt12x()
			if err != nil {
				return nil, err
			}

			return OpIntToByte {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeIntToChar: {
		Name: "int-to-char",
		Size: fmt12xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt12x()
			if err != nil {
				return nil, err
			}

			return OpIntToChar {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeIntToShort: {
		Name: "int-to-short",
		Size: fmt12xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt12x()
			if err != nil {
				return nil, err
			}

			return OpIntToShort {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeAddInt: {
		Name: "add-int",
		Size: fmt23xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt23x()
			if err != nil {
				return nil, err
			}

			return OpAddInt {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeSubInt: {
		Name: "sub-int",
		Size: fmt23xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt23x()
			if err != nil {
				return nil, err
			}

			return OpSubInt {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeMulInt: {
		Name: "mul-int",
		Size: fmt23xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt23x()
			if err != nil {
				return nil, err
			}

			return OpMulInt {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeDivInt: {
		Name: "div-int",
		Size: fmt23xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt23x()
			if err != nil {
				return nil, err
			}

			return OpDivInt {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeRemInt: {
		Name: "rem-int",
		Size: fmt23xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt23x()
			if err != nil {
				return nil, err
			}

			return OpRemInt {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeAndInt: {
		Name: "and-int",
		Size: fmt23xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt23x()
			if err != nil {
				return nil, err
			}

			return OpAndInt {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeOrInt: {
		Name: "or-int",
		Size: fmt23xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt23x()
			if err != nil {
				return nil, err
			}

			return OpOrInt {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeXorInt: {
		Name: "xor-int",
		Size: fmt23xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt23x()
			if err != nil {
				return nil, err
			}

			return OpXorInt {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeShlInt: {
		Name: "shl-int",
		Size: fmt23xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt23x()
			if err != nil {
				return nil, err
			}

			return OpShlInt {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeShrInt: {
		Name: "shr-int",
		Size: fmt23xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt23x()
			if err != nil {
				return nil, err
			}

			return OpShrInt {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeUshrInt: {
		Name: "ushr-int",
		Size: fmt23xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt23x()
			if err != nil {
				return nil, err
			}

			return OpUshrInt {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeAddLong: {
		Name: "add-long",
		Size: fmt23xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt23x()
			if err != nil {
				return nil, err
			}

			return OpAddLong {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeSubLong: {
		Name: "sub-long",
		Size: fmt23xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt23x()
			if err != nil {
				return nil, err
			}

			return OpSubLong {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeMulLong: {
		Name: "mul-long",
		Size: fmt23xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt23x()
			if err != nil {
				return nil, err
			}

			return OpMulLong {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeDivLong: {
		Name: "div-long",
		Size: fmt23xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt23x()
			if err != nil {
				return nil, err
			}

			return OpDivLong {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeRemLong: {
		Name: "rem-long",
		Size: fmt23xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt23x()
			if err != nil {
				return nil, err
			}

			return OpRemLong {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeAndLong: {
		Name: "and-long",
		Size: fmt23xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt23x()
			if err != nil {
				return nil, err
			}

			return OpAndLong {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeOrLong: {
		Name: "or-long",
		Size: fmt23xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt23x()
			if err != nil {
				return nil, err
			}

			return OpOrLong {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeXorLong: {
		Name: "xor-long",
		Size: fmt23xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt23x()
			if err != nil {
				return nil, err
			}

			return OpXorLong {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeShlLong: {
		Name: "shl-long",
		Size: fmt23xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt23x()
			if err != nil {
				return nil, err
			}

			return OpShlLong {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeShrLong: {
		Name: "shr-long",
		Size: fmt23xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt23x()
			if err != nil {
				return nil, err
			}

			return OpShrLong {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeUshrLong: {
		Name: "ushr-long",
		Size: fmt23xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt23x()
			if err != nil {
				return nil, err
			}

			return OpUshrLong {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeAddFloat: {
		Name: "add-float",
		Size: fmt23xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt23x()
			if err != nil {
				return nil, err
			}

			return OpAddFloat {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeSubFloat: {
		Name: "sub-float",
		Size: fmt23xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt23x()
			if err != nil {
				return nil, err
			}

			return OpSubFloat {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeMulFloat: {
		Name: "mul-float",
		Size: fmt23xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt23x()
			if err != nil {
				return nil, err
			}

			return OpMulFloat {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeDivFloat: {
		Name: "div-float",
		Size: fmt23xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt23x()
			if err != nil {
				return nil, err
			}

			return OpDivFloat {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeRemFloat: {
		Name: "rem-float",
		Size: fmt23xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt23x()
			if err != nil {
				return nil, err
			}

			return OpRemFloat {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeAddDouble: {
		Name: "add-double",
		Size: fmt23xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt23x()
			if err != nil {
				return nil, err
			}

			return OpAddDouble {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeSubDouble: {
		Name: "sub-double",
		Size: fmt23xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt23x()
			if err != nil {
				return nil, err
			}

			return OpSubDouble {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeMulDouble: {
		Name: "mul-double",
		Size: fmt23xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt23x()
			if err != nil {
				return nil, err
			}

			return OpMulDouble {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeDivDouble: {
		Name: "div-double",
		Size: fmt23xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt23x()
			if err != nil {
				return nil, err
			}

			return OpDivDouble {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeRemDouble: {
		Name: "rem-double",
		Size: fmt23xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt23x()
			if err != nil {
				return nil, err
			}

			return OpRemDouble {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeAddInt2Addr: {
		Name: "add-int/2addr",
		Size: fmt12xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt12x()
			if err != nil {
				return nil, err
			}

			return OpAddInt2Addr {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeSubInt2Addr: {
		Name: "sub-int/2addr",
		Size: fmt12xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt12x()
			if err != nil {
				return nil, err
			}

			return OpSubInt2Addr {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeMulInt2Addr: {
		Name: "mul-int/2addr",
		Size: fmt12xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt12x()
			if err != nil {
				return nil, err
			}

			return OpMulInt2Addr {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeDivInt2Addr: {
		Name: "div-int/2addr",
		Size: fmt12xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt12x()
			if err != nil {
				return nil, err
			}

			return OpDivInt2Addr {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeRemInt2Addr: {
		Name: "rem-int/2addr",
		Size: fmt12xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt12x()
			if err != nil {
				return nil, err
			}

			return OpRemInt2Addr {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeAndInt2Addr: {
		Name: "and-int/2addr",
		Size: fmt12xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt12x()
			if err != nil {
				return nil, err
			}

			return OpAndInt2Addr {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeOrInt2Addr: {
		Name: "or-int/2addr",
		Size: fmt12xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt12x()
			if err != nil {
				return nil, err
			}

			return OpOrInt2Addr {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeXorInt2Addr: {
		Name: "xor-int/2addr",
		Size: fmt12xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt12x()
			if err != nil {
				return nil, err
			}

			return OpXorInt2Addr {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeShlInt2Addr: {
		Name: "shl-int/2addr",
		Size: fmt12xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt12x()
			if err != nil {
				return nil, err
			}

			return OpShlInt2Addr {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeShrInt2Addr: {
		Name: "shr-int/2addr",
		Size: fmt12xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt12x()
			if err != nil {
				return nil, err
			}

			return OpShrInt2Addr {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeUshrInt2Addr: {
		Name: "ushr-int/2addr",
		Size: fmt12xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt12x()
			if err != nil {
				return nil, err
			}

			return OpUshrInt2Addr {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeAddLong2Addr: {
		Name: "add-long/2addr",
		Size: fmt12xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt12x()
			if err != nil {
				return nil, err
			}

			return OpAddLong2Addr {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeSubLong2Addr: {
		Name: "sub-long/2addr",
		Size: fmt12xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt12x()
			if err != nil {
				return nil, err
			}

			return OpSubLong2Addr {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeMulLong2Addr: {
		Name: "mul-long/2addr",
		Size: fmt12xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt12x()
			if err != nil {
				return nil, err
			}

			return OpMulLong2Addr {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeDivLong2Addr: {
		Name: "div-long/2addr",
		Size: fmt12xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt12x()
			if err != nil {
				return nil, err
			}

			return OpDivLong2Addr {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeRemLong2Addr: {
		Name: "rem-long/2addr",
		Size: fmt12xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt12x()
			if err != nil {
				return nil, err
			}

			return OpRemLong2Addr {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeAndLong2Addr: {
		Name: "and-long/2addr",
		Size: fmt12xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt12x()
			if err != nil {
				return nil, err
			}

			return OpAndLong2Addr {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeOrLong2Addr: {
		Name: "or-long/2addr",
		Size: fmt12xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt12x()
			if err != nil {
				return nil, err
			}

			return OpOrLong2Addr {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeXorLong2Addr: {
		Name: "xor-long/2addr",
		Size: fmt12xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt12x()
			if err != nil {
				return nil, err
			}

			return OpXorLong2Addr {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeShlLong2Addr: {
		Name: "shl-long/2addr",
		Size: fmt12xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt12x()
			if err != nil {
				return nil, err
			}

			return OpShlLong2Addr {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeShrLong2Addr: {
		Name: "shr-long/2addr",
		Size: fmt12xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt12x()
			if err != nil {
				return nil, err
			}

			return OpShrLong2Addr {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeUshrLong2Addr: {
		Name: "ushr-long/2addr",
		Size: fmt12xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt12x()
			if err != nil {
				return nil, err
			}

			return OpUshrLong2Addr {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeAddFloat2Addr: {
		Name: "add-float/2addr",
		Size: fmt12xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt12x()
			if err != nil {
				return nil, err
			}

			return OpAddFloat2Addr {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeSubFloat2Addr: {
		Name: "sub-float/2addr",
		Size: fmt12xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt12x()
			if err != nil {
				return nil, err
			}

			return OpSubFloat2Addr {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeMulFloat2Addr: {
		Name: "mul-float/2addr",
		Size: fmt12xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt12x()
			if err != nil {
				return nil, err
			}

			return OpMulFloat2Addr {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeDivFloat2Addr: {
		Name: "div-float/2addr",
		Size: fmt12xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt12x()
			if err != nil {
				return nil, err
			}

			return OpDivFloat2Addr {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeRemFloat2Addr: {
		Name: "rem-float/2addr",
		Size: fmt12xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt12x()
			if err != nil {
				return nil, err
			}

			return OpRemFloat2Addr {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeAddDouble2Addr: {
		Name: "add-double/2addr",
		Size: fmt12xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt12x()
			if err != nil {
				return nil, err
			}

			return OpAddDouble2Addr {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeSubDouble2Addr: {
		Name: "sub-double/2addr",
		Size: fmt12xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt12x()
			if err != nil {
				return nil, err
			}

			return OpSubDouble2Addr {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeMulDouble2Addr: {
		Name: "mul-double/2addr",
		Size: fmt12xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt12x()
			if err != nil {
				return nil, err
			}

			return OpMulDouble2Addr {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeDivDouble2Addr: {
		Name: "div-double/2addr",
		Size: fmt12xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt12x()
			if err != nil {
				return nil, err
			}

			return OpDivDouble2Addr {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeRemDouble2Addr: {
		Name: "rem-double/2addr",
		Size: fmt12xSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt12x()
			if err != nil {
				return nil, err
			}

			return OpRemDouble2Addr {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeAddIntLit16: {
		Name: "add-int/lit16",
		Size: fmt22sSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt22s()
			if err != nil {
				return nil, err
			}

			return OpAddIntLit16 {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeRsubIntLit16: {
		Name: "rsub-int/lit16",
		Size: fmt22sSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt22s()
			if err != nil {
				return nil, err
			}

			return OpRsubIntLit16 {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeMulIntLit16: {
		Name: "mul-int/lit16",
		Size: fmt22sSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt22s()
			if err != nil {
				return nil, err
			}

			return OpMulIntLit16 {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeDivIntLit16: {
		Name: "div-int/lit16",
		Size: fmt22sSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt22s()
			if err != nil {
				return nil, err
			}

			return OpDivIntLit16 {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeRemIntLit16: {
		Name: "rem-int/lit16",
		Size: fmt22sSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt22s()
			if err != nil {
				return nil, err
			}

			return OpRemIntLit16 {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeAndIntLit16: {
		Name: "and-int/lit16",
		Size: fmt22sSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt22s()
			if err != nil {
				return nil, err
			}

			return OpAndIntLit16 {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeOrIntLit16: {
		Name: "or-int/lit16",
		Size: fmt22sSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt22s()
			if err != nil {
				return nil, err
			}

			return OpOrIntLit16 {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeXorIntLit16: {
		Name: "xor-int/lit16",
		Size: fmt22sSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt22s()
			if err != nil {
				return nil, err
			}

			return OpXorIntLit16 {
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

			return OpAddIntLit8 {
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

			return OpRsubIntLit8 {
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

			return OpMulIntLit8 {
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

			return OpDivIntLit8 {
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

			return OpRemIntLit8 {
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

			return OpAndIntLit8 {
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

			return OpOrIntLit8 {
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

			return OpXorIntLit8 {
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

			return OpShlIntLit8 {
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

			return OpShrIntLit8 {
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

			return OpUshrIntLit8 {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeInvokePolymorphic: {
		Name: "invoke-polymorphic",
		Size: fmt45ccSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt45cc()
			if err != nil {
				return nil, err
			}

			return OpInvokePolymorphic {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeInvokePolymorphicRange: {
		Name: "invoke-polymorphic/range",
		Size: fmt4rccSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt4rcc()
			if err != nil {
				return nil, err
			}

			return OpInvokePolymorphicRange {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeInvokeCustom: {
		Name: "invoke-custom",
		Size: fmt35cSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt35c()
			if err != nil {
				return nil, err
			}

			return OpInvokeCustom {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeInvokeCustomRange: {
		Name: "invoke-custom/range",
		Size: fmt3rcSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt3rc()
			if err != nil {
				return nil, err
			}

			return OpInvokeCustomRange {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeConstMethodHandle: {
		Name: "const-method-handle",
		Size: fmt21cSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt21c()
			if err != nil {
				return nil, err
			}

			return OpConstMethodHandle {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
	OpCodeConstMethodType: {
		Name: "const-method-type",
		Size: fmt21cSize,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt21c()
			if err != nil {
				return nil, err
			}

			return OpConstMethodType {
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

type OpConstWide struct {
	opBase
	Fmt51l
}

func (o OpConstWide) Code() OpCode {
	return OpCodeConstWide
}

func (o OpConstWide) Fmt() Fmt {
	return o.Fmt51l
}

func (o OpConstWide) String() string {
	f := o.Fmt51l.String()
	if f != "" {
		return fmt.Sprintf("0x%x: const-wide %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: const-wide", o.pos)
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

type OpFilledNewArray struct {
	opBase
	Fmt35c
}

func (o OpFilledNewArray) Code() OpCode {
	return OpCodeFilledNewArray
}

func (o OpFilledNewArray) Fmt() Fmt {
	return o.Fmt35c
}

func (o OpFilledNewArray) String() string {
	f := o.Fmt35c.String()
	if f != "" {
		return fmt.Sprintf("0x%x: filled-new-array %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: filled-new-array", o.pos)
}

type OpFilledNewArrayRange struct {
	opBase
	Fmt3rc
}

func (o OpFilledNewArrayRange) Code() OpCode {
	return OpCodeFilledNewArrayRange
}

func (o OpFilledNewArrayRange) Fmt() Fmt {
	return o.Fmt3rc
}

func (o OpFilledNewArrayRange) String() string {
	f := o.Fmt3rc.String()
	if f != "" {
		return fmt.Sprintf("0x%x: filled-new-array/range %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: filled-new-array/range", o.pos)
}

type OpFillArrayData struct {
	opBase
	Fmt31t
}

func (o OpFillArrayData) Code() OpCode {
	return OpCodeFillArrayData
}

func (o OpFillArrayData) Fmt() Fmt {
	return o.Fmt31t
}

func (o OpFillArrayData) String() string {
	f := o.Fmt31t.String()
	if f != "" {
		return fmt.Sprintf("0x%x: fill-array-data %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: fill-array-data", o.pos)
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

type OpSparseSwitch struct {
	opBase
	Fmt31t
}

func (o OpSparseSwitch) Code() OpCode {
	return OpCodeSparseSwitch
}

func (o OpSparseSwitch) Fmt() Fmt {
	return o.Fmt31t
}

func (o OpSparseSwitch) String() string {
	f := o.Fmt31t.String()
	if f != "" {
		return fmt.Sprintf("0x%x: sparse-switch %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: sparse-switch", o.pos)
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

type OpNegInt struct {
	opBase
	Fmt12x
}

func (o OpNegInt) Code() OpCode {
	return OpCodeNegInt
}

func (o OpNegInt) Fmt() Fmt {
	return o.Fmt12x
}

func (o OpNegInt) String() string {
	f := o.Fmt12x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: neg-int %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: neg-int", o.pos)
}

type OpNotInt struct {
	opBase
	Fmt12x
}

func (o OpNotInt) Code() OpCode {
	return OpCodeNotInt
}

func (o OpNotInt) Fmt() Fmt {
	return o.Fmt12x
}

func (o OpNotInt) String() string {
	f := o.Fmt12x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: not-int %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: not-int", o.pos)
}

type OpNegLong struct {
	opBase
	Fmt12x
}

func (o OpNegLong) Code() OpCode {
	return OpCodeNegLong
}

func (o OpNegLong) Fmt() Fmt {
	return o.Fmt12x
}

func (o OpNegLong) String() string {
	f := o.Fmt12x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: neg-long %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: neg-long", o.pos)
}

type OpNotLong struct {
	opBase
	Fmt12x
}

func (o OpNotLong) Code() OpCode {
	return OpCodeNotLong
}

func (o OpNotLong) Fmt() Fmt {
	return o.Fmt12x
}

func (o OpNotLong) String() string {
	f := o.Fmt12x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: not-long %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: not-long", o.pos)
}

type OpNegFloat struct {
	opBase
	Fmt12x
}

func (o OpNegFloat) Code() OpCode {
	return OpCodeNegFloat
}

func (o OpNegFloat) Fmt() Fmt {
	return o.Fmt12x
}

func (o OpNegFloat) String() string {
	f := o.Fmt12x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: neg-float %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: neg-float", o.pos)
}

type OpNegDouble struct {
	opBase
	Fmt12x
}

func (o OpNegDouble) Code() OpCode {
	return OpCodeNegDouble
}

func (o OpNegDouble) Fmt() Fmt {
	return o.Fmt12x
}

func (o OpNegDouble) String() string {
	f := o.Fmt12x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: neg-double %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: neg-double", o.pos)
}

type OpIntToLong struct {
	opBase
	Fmt12x
}

func (o OpIntToLong) Code() OpCode {
	return OpCodeIntToLong
}

func (o OpIntToLong) Fmt() Fmt {
	return o.Fmt12x
}

func (o OpIntToLong) String() string {
	f := o.Fmt12x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: int-to-long %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: int-to-long", o.pos)
}

type OpIntToFloat struct {
	opBase
	Fmt12x
}

func (o OpIntToFloat) Code() OpCode {
	return OpCodeIntToFloat
}

func (o OpIntToFloat) Fmt() Fmt {
	return o.Fmt12x
}

func (o OpIntToFloat) String() string {
	f := o.Fmt12x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: int-to-float %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: int-to-float", o.pos)
}

type OpIntToDouble struct {
	opBase
	Fmt12x
}

func (o OpIntToDouble) Code() OpCode {
	return OpCodeIntToDouble
}

func (o OpIntToDouble) Fmt() Fmt {
	return o.Fmt12x
}

func (o OpIntToDouble) String() string {
	f := o.Fmt12x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: int-to-double %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: int-to-double", o.pos)
}

type OpLongToInt struct {
	opBase
	Fmt12x
}

func (o OpLongToInt) Code() OpCode {
	return OpCodeLongToInt
}

func (o OpLongToInt) Fmt() Fmt {
	return o.Fmt12x
}

func (o OpLongToInt) String() string {
	f := o.Fmt12x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: long-to-int %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: long-to-int", o.pos)
}

type OpLongToFloat struct {
	opBase
	Fmt12x
}

func (o OpLongToFloat) Code() OpCode {
	return OpCodeLongToFloat
}

func (o OpLongToFloat) Fmt() Fmt {
	return o.Fmt12x
}

func (o OpLongToFloat) String() string {
	f := o.Fmt12x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: long-to-float %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: long-to-float", o.pos)
}

type OpLongToDouble struct {
	opBase
	Fmt12x
}

func (o OpLongToDouble) Code() OpCode {
	return OpCodeLongToDouble
}

func (o OpLongToDouble) Fmt() Fmt {
	return o.Fmt12x
}

func (o OpLongToDouble) String() string {
	f := o.Fmt12x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: long-to-double %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: long-to-double", o.pos)
}

type OpFloatToInt struct {
	opBase
	Fmt12x
}

func (o OpFloatToInt) Code() OpCode {
	return OpCodeFloatToInt
}

func (o OpFloatToInt) Fmt() Fmt {
	return o.Fmt12x
}

func (o OpFloatToInt) String() string {
	f := o.Fmt12x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: float-to-int %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: float-to-int", o.pos)
}

type OpFloatToLong struct {
	opBase
	Fmt12x
}

func (o OpFloatToLong) Code() OpCode {
	return OpCodeFloatToLong
}

func (o OpFloatToLong) Fmt() Fmt {
	return o.Fmt12x
}

func (o OpFloatToLong) String() string {
	f := o.Fmt12x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: float-to-long %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: float-to-long", o.pos)
}

type OpFloatToDouble struct {
	opBase
	Fmt12x
}

func (o OpFloatToDouble) Code() OpCode {
	return OpCodeFloatToDouble
}

func (o OpFloatToDouble) Fmt() Fmt {
	return o.Fmt12x
}

func (o OpFloatToDouble) String() string {
	f := o.Fmt12x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: float-to-double %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: float-to-double", o.pos)
}

type OpDoubleToInt struct {
	opBase
	Fmt12x
}

func (o OpDoubleToInt) Code() OpCode {
	return OpCodeDoubleToInt
}

func (o OpDoubleToInt) Fmt() Fmt {
	return o.Fmt12x
}

func (o OpDoubleToInt) String() string {
	f := o.Fmt12x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: double-to-int %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: double-to-int", o.pos)
}

type OpDoubleToLong struct {
	opBase
	Fmt12x
}

func (o OpDoubleToLong) Code() OpCode {
	return OpCodeDoubleToLong
}

func (o OpDoubleToLong) Fmt() Fmt {
	return o.Fmt12x
}

func (o OpDoubleToLong) String() string {
	f := o.Fmt12x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: double-to-long %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: double-to-long", o.pos)
}

type OpDoubleToFloat struct {
	opBase
	Fmt12x
}

func (o OpDoubleToFloat) Code() OpCode {
	return OpCodeDoubleToFloat
}

func (o OpDoubleToFloat) Fmt() Fmt {
	return o.Fmt12x
}

func (o OpDoubleToFloat) String() string {
	f := o.Fmt12x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: double-to-float %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: double-to-float", o.pos)
}

type OpIntToByte struct {
	opBase
	Fmt12x
}

func (o OpIntToByte) Code() OpCode {
	return OpCodeIntToByte
}

func (o OpIntToByte) Fmt() Fmt {
	return o.Fmt12x
}

func (o OpIntToByte) String() string {
	f := o.Fmt12x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: int-to-byte %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: int-to-byte", o.pos)
}

type OpIntToChar struct {
	opBase
	Fmt12x
}

func (o OpIntToChar) Code() OpCode {
	return OpCodeIntToChar
}

func (o OpIntToChar) Fmt() Fmt {
	return o.Fmt12x
}

func (o OpIntToChar) String() string {
	f := o.Fmt12x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: int-to-char %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: int-to-char", o.pos)
}

type OpIntToShort struct {
	opBase
	Fmt12x
}

func (o OpIntToShort) Code() OpCode {
	return OpCodeIntToShort
}

func (o OpIntToShort) Fmt() Fmt {
	return o.Fmt12x
}

func (o OpIntToShort) String() string {
	f := o.Fmt12x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: int-to-short %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: int-to-short", o.pos)
}

type OpAddInt struct {
	opBase
	Fmt23x
}

func (o OpAddInt) Code() OpCode {
	return OpCodeAddInt
}

func (o OpAddInt) Fmt() Fmt {
	return o.Fmt23x
}

func (o OpAddInt) String() string {
	f := o.Fmt23x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: add-int %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: add-int", o.pos)
}

type OpSubInt struct {
	opBase
	Fmt23x
}

func (o OpSubInt) Code() OpCode {
	return OpCodeSubInt
}

func (o OpSubInt) Fmt() Fmt {
	return o.Fmt23x
}

func (o OpSubInt) String() string {
	f := o.Fmt23x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: sub-int %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: sub-int", o.pos)
}

type OpMulInt struct {
	opBase
	Fmt23x
}

func (o OpMulInt) Code() OpCode {
	return OpCodeMulInt
}

func (o OpMulInt) Fmt() Fmt {
	return o.Fmt23x
}

func (o OpMulInt) String() string {
	f := o.Fmt23x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: mul-int %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: mul-int", o.pos)
}

type OpDivInt struct {
	opBase
	Fmt23x
}

func (o OpDivInt) Code() OpCode {
	return OpCodeDivInt
}

func (o OpDivInt) Fmt() Fmt {
	return o.Fmt23x
}

func (o OpDivInt) String() string {
	f := o.Fmt23x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: div-int %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: div-int", o.pos)
}

type OpRemInt struct {
	opBase
	Fmt23x
}

func (o OpRemInt) Code() OpCode {
	return OpCodeRemInt
}

func (o OpRemInt) Fmt() Fmt {
	return o.Fmt23x
}

func (o OpRemInt) String() string {
	f := o.Fmt23x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: rem-int %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: rem-int", o.pos)
}

type OpAndInt struct {
	opBase
	Fmt23x
}

func (o OpAndInt) Code() OpCode {
	return OpCodeAndInt
}

func (o OpAndInt) Fmt() Fmt {
	return o.Fmt23x
}

func (o OpAndInt) String() string {
	f := o.Fmt23x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: and-int %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: and-int", o.pos)
}

type OpOrInt struct {
	opBase
	Fmt23x
}

func (o OpOrInt) Code() OpCode {
	return OpCodeOrInt
}

func (o OpOrInt) Fmt() Fmt {
	return o.Fmt23x
}

func (o OpOrInt) String() string {
	f := o.Fmt23x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: or-int %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: or-int", o.pos)
}

type OpXorInt struct {
	opBase
	Fmt23x
}

func (o OpXorInt) Code() OpCode {
	return OpCodeXorInt
}

func (o OpXorInt) Fmt() Fmt {
	return o.Fmt23x
}

func (o OpXorInt) String() string {
	f := o.Fmt23x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: xor-int %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: xor-int", o.pos)
}

type OpShlInt struct {
	opBase
	Fmt23x
}

func (o OpShlInt) Code() OpCode {
	return OpCodeShlInt
}

func (o OpShlInt) Fmt() Fmt {
	return o.Fmt23x
}

func (o OpShlInt) String() string {
	f := o.Fmt23x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: shl-int %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: shl-int", o.pos)
}

type OpShrInt struct {
	opBase
	Fmt23x
}

func (o OpShrInt) Code() OpCode {
	return OpCodeShrInt
}

func (o OpShrInt) Fmt() Fmt {
	return o.Fmt23x
}

func (o OpShrInt) String() string {
	f := o.Fmt23x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: shr-int %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: shr-int", o.pos)
}

type OpUshrInt struct {
	opBase
	Fmt23x
}

func (o OpUshrInt) Code() OpCode {
	return OpCodeUshrInt
}

func (o OpUshrInt) Fmt() Fmt {
	return o.Fmt23x
}

func (o OpUshrInt) String() string {
	f := o.Fmt23x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: ushr-int %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: ushr-int", o.pos)
}

type OpAddLong struct {
	opBase
	Fmt23x
}

func (o OpAddLong) Code() OpCode {
	return OpCodeAddLong
}

func (o OpAddLong) Fmt() Fmt {
	return o.Fmt23x
}

func (o OpAddLong) String() string {
	f := o.Fmt23x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: add-long %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: add-long", o.pos)
}

type OpSubLong struct {
	opBase
	Fmt23x
}

func (o OpSubLong) Code() OpCode {
	return OpCodeSubLong
}

func (o OpSubLong) Fmt() Fmt {
	return o.Fmt23x
}

func (o OpSubLong) String() string {
	f := o.Fmt23x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: sub-long %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: sub-long", o.pos)
}

type OpMulLong struct {
	opBase
	Fmt23x
}

func (o OpMulLong) Code() OpCode {
	return OpCodeMulLong
}

func (o OpMulLong) Fmt() Fmt {
	return o.Fmt23x
}

func (o OpMulLong) String() string {
	f := o.Fmt23x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: mul-long %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: mul-long", o.pos)
}

type OpDivLong struct {
	opBase
	Fmt23x
}

func (o OpDivLong) Code() OpCode {
	return OpCodeDivLong
}

func (o OpDivLong) Fmt() Fmt {
	return o.Fmt23x
}

func (o OpDivLong) String() string {
	f := o.Fmt23x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: div-long %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: div-long", o.pos)
}

type OpRemLong struct {
	opBase
	Fmt23x
}

func (o OpRemLong) Code() OpCode {
	return OpCodeRemLong
}

func (o OpRemLong) Fmt() Fmt {
	return o.Fmt23x
}

func (o OpRemLong) String() string {
	f := o.Fmt23x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: rem-long %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: rem-long", o.pos)
}

type OpAndLong struct {
	opBase
	Fmt23x
}

func (o OpAndLong) Code() OpCode {
	return OpCodeAndLong
}

func (o OpAndLong) Fmt() Fmt {
	return o.Fmt23x
}

func (o OpAndLong) String() string {
	f := o.Fmt23x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: and-long %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: and-long", o.pos)
}

type OpOrLong struct {
	opBase
	Fmt23x
}

func (o OpOrLong) Code() OpCode {
	return OpCodeOrLong
}

func (o OpOrLong) Fmt() Fmt {
	return o.Fmt23x
}

func (o OpOrLong) String() string {
	f := o.Fmt23x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: or-long %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: or-long", o.pos)
}

type OpXorLong struct {
	opBase
	Fmt23x
}

func (o OpXorLong) Code() OpCode {
	return OpCodeXorLong
}

func (o OpXorLong) Fmt() Fmt {
	return o.Fmt23x
}

func (o OpXorLong) String() string {
	f := o.Fmt23x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: xor-long %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: xor-long", o.pos)
}

type OpShlLong struct {
	opBase
	Fmt23x
}

func (o OpShlLong) Code() OpCode {
	return OpCodeShlLong
}

func (o OpShlLong) Fmt() Fmt {
	return o.Fmt23x
}

func (o OpShlLong) String() string {
	f := o.Fmt23x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: shl-long %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: shl-long", o.pos)
}

type OpShrLong struct {
	opBase
	Fmt23x
}

func (o OpShrLong) Code() OpCode {
	return OpCodeShrLong
}

func (o OpShrLong) Fmt() Fmt {
	return o.Fmt23x
}

func (o OpShrLong) String() string {
	f := o.Fmt23x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: shr-long %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: shr-long", o.pos)
}

type OpUshrLong struct {
	opBase
	Fmt23x
}

func (o OpUshrLong) Code() OpCode {
	return OpCodeUshrLong
}

func (o OpUshrLong) Fmt() Fmt {
	return o.Fmt23x
}

func (o OpUshrLong) String() string {
	f := o.Fmt23x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: ushr-long %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: ushr-long", o.pos)
}

type OpAddFloat struct {
	opBase
	Fmt23x
}

func (o OpAddFloat) Code() OpCode {
	return OpCodeAddFloat
}

func (o OpAddFloat) Fmt() Fmt {
	return o.Fmt23x
}

func (o OpAddFloat) String() string {
	f := o.Fmt23x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: add-float %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: add-float", o.pos)
}

type OpSubFloat struct {
	opBase
	Fmt23x
}

func (o OpSubFloat) Code() OpCode {
	return OpCodeSubFloat
}

func (o OpSubFloat) Fmt() Fmt {
	return o.Fmt23x
}

func (o OpSubFloat) String() string {
	f := o.Fmt23x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: sub-float %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: sub-float", o.pos)
}

type OpMulFloat struct {
	opBase
	Fmt23x
}

func (o OpMulFloat) Code() OpCode {
	return OpCodeMulFloat
}

func (o OpMulFloat) Fmt() Fmt {
	return o.Fmt23x
}

func (o OpMulFloat) String() string {
	f := o.Fmt23x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: mul-float %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: mul-float", o.pos)
}

type OpDivFloat struct {
	opBase
	Fmt23x
}

func (o OpDivFloat) Code() OpCode {
	return OpCodeDivFloat
}

func (o OpDivFloat) Fmt() Fmt {
	return o.Fmt23x
}

func (o OpDivFloat) String() string {
	f := o.Fmt23x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: div-float %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: div-float", o.pos)
}

type OpRemFloat struct {
	opBase
	Fmt23x
}

func (o OpRemFloat) Code() OpCode {
	return OpCodeRemFloat
}

func (o OpRemFloat) Fmt() Fmt {
	return o.Fmt23x
}

func (o OpRemFloat) String() string {
	f := o.Fmt23x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: rem-float %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: rem-float", o.pos)
}

type OpAddDouble struct {
	opBase
	Fmt23x
}

func (o OpAddDouble) Code() OpCode {
	return OpCodeAddDouble
}

func (o OpAddDouble) Fmt() Fmt {
	return o.Fmt23x
}

func (o OpAddDouble) String() string {
	f := o.Fmt23x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: add-double %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: add-double", o.pos)
}

type OpSubDouble struct {
	opBase
	Fmt23x
}

func (o OpSubDouble) Code() OpCode {
	return OpCodeSubDouble
}

func (o OpSubDouble) Fmt() Fmt {
	return o.Fmt23x
}

func (o OpSubDouble) String() string {
	f := o.Fmt23x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: sub-double %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: sub-double", o.pos)
}

type OpMulDouble struct {
	opBase
	Fmt23x
}

func (o OpMulDouble) Code() OpCode {
	return OpCodeMulDouble
}

func (o OpMulDouble) Fmt() Fmt {
	return o.Fmt23x
}

func (o OpMulDouble) String() string {
	f := o.Fmt23x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: mul-double %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: mul-double", o.pos)
}

type OpDivDouble struct {
	opBase
	Fmt23x
}

func (o OpDivDouble) Code() OpCode {
	return OpCodeDivDouble
}

func (o OpDivDouble) Fmt() Fmt {
	return o.Fmt23x
}

func (o OpDivDouble) String() string {
	f := o.Fmt23x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: div-double %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: div-double", o.pos)
}

type OpRemDouble struct {
	opBase
	Fmt23x
}

func (o OpRemDouble) Code() OpCode {
	return OpCodeRemDouble
}

func (o OpRemDouble) Fmt() Fmt {
	return o.Fmt23x
}

func (o OpRemDouble) String() string {
	f := o.Fmt23x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: rem-double %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: rem-double", o.pos)
}

type OpAddInt2Addr struct {
	opBase
	Fmt12x
}

func (o OpAddInt2Addr) Code() OpCode {
	return OpCodeAddInt2Addr
}

func (o OpAddInt2Addr) Fmt() Fmt {
	return o.Fmt12x
}

func (o OpAddInt2Addr) String() string {
	f := o.Fmt12x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: add-int/2addr %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: add-int/2addr", o.pos)
}

type OpSubInt2Addr struct {
	opBase
	Fmt12x
}

func (o OpSubInt2Addr) Code() OpCode {
	return OpCodeSubInt2Addr
}

func (o OpSubInt2Addr) Fmt() Fmt {
	return o.Fmt12x
}

func (o OpSubInt2Addr) String() string {
	f := o.Fmt12x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: sub-int/2addr %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: sub-int/2addr", o.pos)
}

type OpMulInt2Addr struct {
	opBase
	Fmt12x
}

func (o OpMulInt2Addr) Code() OpCode {
	return OpCodeMulInt2Addr
}

func (o OpMulInt2Addr) Fmt() Fmt {
	return o.Fmt12x
}

func (o OpMulInt2Addr) String() string {
	f := o.Fmt12x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: mul-int/2addr %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: mul-int/2addr", o.pos)
}

type OpDivInt2Addr struct {
	opBase
	Fmt12x
}

func (o OpDivInt2Addr) Code() OpCode {
	return OpCodeDivInt2Addr
}

func (o OpDivInt2Addr) Fmt() Fmt {
	return o.Fmt12x
}

func (o OpDivInt2Addr) String() string {
	f := o.Fmt12x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: div-int/2addr %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: div-int/2addr", o.pos)
}

type OpRemInt2Addr struct {
	opBase
	Fmt12x
}

func (o OpRemInt2Addr) Code() OpCode {
	return OpCodeRemInt2Addr
}

func (o OpRemInt2Addr) Fmt() Fmt {
	return o.Fmt12x
}

func (o OpRemInt2Addr) String() string {
	f := o.Fmt12x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: rem-int/2addr %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: rem-int/2addr", o.pos)
}

type OpAndInt2Addr struct {
	opBase
	Fmt12x
}

func (o OpAndInt2Addr) Code() OpCode {
	return OpCodeAndInt2Addr
}

func (o OpAndInt2Addr) Fmt() Fmt {
	return o.Fmt12x
}

func (o OpAndInt2Addr) String() string {
	f := o.Fmt12x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: and-int/2addr %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: and-int/2addr", o.pos)
}

type OpOrInt2Addr struct {
	opBase
	Fmt12x
}

func (o OpOrInt2Addr) Code() OpCode {
	return OpCodeOrInt2Addr
}

func (o OpOrInt2Addr) Fmt() Fmt {
	return o.Fmt12x
}

func (o OpOrInt2Addr) String() string {
	f := o.Fmt12x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: or-int/2addr %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: or-int/2addr", o.pos)
}

type OpXorInt2Addr struct {
	opBase
	Fmt12x
}

func (o OpXorInt2Addr) Code() OpCode {
	return OpCodeXorInt2Addr
}

func (o OpXorInt2Addr) Fmt() Fmt {
	return o.Fmt12x
}

func (o OpXorInt2Addr) String() string {
	f := o.Fmt12x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: xor-int/2addr %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: xor-int/2addr", o.pos)
}

type OpShlInt2Addr struct {
	opBase
	Fmt12x
}

func (o OpShlInt2Addr) Code() OpCode {
	return OpCodeShlInt2Addr
}

func (o OpShlInt2Addr) Fmt() Fmt {
	return o.Fmt12x
}

func (o OpShlInt2Addr) String() string {
	f := o.Fmt12x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: shl-int/2addr %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: shl-int/2addr", o.pos)
}

type OpShrInt2Addr struct {
	opBase
	Fmt12x
}

func (o OpShrInt2Addr) Code() OpCode {
	return OpCodeShrInt2Addr
}

func (o OpShrInt2Addr) Fmt() Fmt {
	return o.Fmt12x
}

func (o OpShrInt2Addr) String() string {
	f := o.Fmt12x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: shr-int/2addr %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: shr-int/2addr", o.pos)
}

type OpUshrInt2Addr struct {
	opBase
	Fmt12x
}

func (o OpUshrInt2Addr) Code() OpCode {
	return OpCodeUshrInt2Addr
}

func (o OpUshrInt2Addr) Fmt() Fmt {
	return o.Fmt12x
}

func (o OpUshrInt2Addr) String() string {
	f := o.Fmt12x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: ushr-int/2addr %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: ushr-int/2addr", o.pos)
}

type OpAddLong2Addr struct {
	opBase
	Fmt12x
}

func (o OpAddLong2Addr) Code() OpCode {
	return OpCodeAddLong2Addr
}

func (o OpAddLong2Addr) Fmt() Fmt {
	return o.Fmt12x
}

func (o OpAddLong2Addr) String() string {
	f := o.Fmt12x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: add-long/2addr %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: add-long/2addr", o.pos)
}

type OpSubLong2Addr struct {
	opBase
	Fmt12x
}

func (o OpSubLong2Addr) Code() OpCode {
	return OpCodeSubLong2Addr
}

func (o OpSubLong2Addr) Fmt() Fmt {
	return o.Fmt12x
}

func (o OpSubLong2Addr) String() string {
	f := o.Fmt12x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: sub-long/2addr %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: sub-long/2addr", o.pos)
}

type OpMulLong2Addr struct {
	opBase
	Fmt12x
}

func (o OpMulLong2Addr) Code() OpCode {
	return OpCodeMulLong2Addr
}

func (o OpMulLong2Addr) Fmt() Fmt {
	return o.Fmt12x
}

func (o OpMulLong2Addr) String() string {
	f := o.Fmt12x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: mul-long/2addr %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: mul-long/2addr", o.pos)
}

type OpDivLong2Addr struct {
	opBase
	Fmt12x
}

func (o OpDivLong2Addr) Code() OpCode {
	return OpCodeDivLong2Addr
}

func (o OpDivLong2Addr) Fmt() Fmt {
	return o.Fmt12x
}

func (o OpDivLong2Addr) String() string {
	f := o.Fmt12x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: div-long/2addr %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: div-long/2addr", o.pos)
}

type OpRemLong2Addr struct {
	opBase
	Fmt12x
}

func (o OpRemLong2Addr) Code() OpCode {
	return OpCodeRemLong2Addr
}

func (o OpRemLong2Addr) Fmt() Fmt {
	return o.Fmt12x
}

func (o OpRemLong2Addr) String() string {
	f := o.Fmt12x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: rem-long/2addr %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: rem-long/2addr", o.pos)
}

type OpAndLong2Addr struct {
	opBase
	Fmt12x
}

func (o OpAndLong2Addr) Code() OpCode {
	return OpCodeAndLong2Addr
}

func (o OpAndLong2Addr) Fmt() Fmt {
	return o.Fmt12x
}

func (o OpAndLong2Addr) String() string {
	f := o.Fmt12x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: and-long/2addr %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: and-long/2addr", o.pos)
}

type OpOrLong2Addr struct {
	opBase
	Fmt12x
}

func (o OpOrLong2Addr) Code() OpCode {
	return OpCodeOrLong2Addr
}

func (o OpOrLong2Addr) Fmt() Fmt {
	return o.Fmt12x
}

func (o OpOrLong2Addr) String() string {
	f := o.Fmt12x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: or-long/2addr %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: or-long/2addr", o.pos)
}

type OpXorLong2Addr struct {
	opBase
	Fmt12x
}

func (o OpXorLong2Addr) Code() OpCode {
	return OpCodeXorLong2Addr
}

func (o OpXorLong2Addr) Fmt() Fmt {
	return o.Fmt12x
}

func (o OpXorLong2Addr) String() string {
	f := o.Fmt12x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: xor-long/2addr %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: xor-long/2addr", o.pos)
}

type OpShlLong2Addr struct {
	opBase
	Fmt12x
}

func (o OpShlLong2Addr) Code() OpCode {
	return OpCodeShlLong2Addr
}

func (o OpShlLong2Addr) Fmt() Fmt {
	return o.Fmt12x
}

func (o OpShlLong2Addr) String() string {
	f := o.Fmt12x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: shl-long/2addr %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: shl-long/2addr", o.pos)
}

type OpShrLong2Addr struct {
	opBase
	Fmt12x
}

func (o OpShrLong2Addr) Code() OpCode {
	return OpCodeShrLong2Addr
}

func (o OpShrLong2Addr) Fmt() Fmt {
	return o.Fmt12x
}

func (o OpShrLong2Addr) String() string {
	f := o.Fmt12x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: shr-long/2addr %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: shr-long/2addr", o.pos)
}

type OpUshrLong2Addr struct {
	opBase
	Fmt12x
}

func (o OpUshrLong2Addr) Code() OpCode {
	return OpCodeUshrLong2Addr
}

func (o OpUshrLong2Addr) Fmt() Fmt {
	return o.Fmt12x
}

func (o OpUshrLong2Addr) String() string {
	f := o.Fmt12x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: ushr-long/2addr %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: ushr-long/2addr", o.pos)
}

type OpAddFloat2Addr struct {
	opBase
	Fmt12x
}

func (o OpAddFloat2Addr) Code() OpCode {
	return OpCodeAddFloat2Addr
}

func (o OpAddFloat2Addr) Fmt() Fmt {
	return o.Fmt12x
}

func (o OpAddFloat2Addr) String() string {
	f := o.Fmt12x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: add-float/2addr %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: add-float/2addr", o.pos)
}

type OpSubFloat2Addr struct {
	opBase
	Fmt12x
}

func (o OpSubFloat2Addr) Code() OpCode {
	return OpCodeSubFloat2Addr
}

func (o OpSubFloat2Addr) Fmt() Fmt {
	return o.Fmt12x
}

func (o OpSubFloat2Addr) String() string {
	f := o.Fmt12x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: sub-float/2addr %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: sub-float/2addr", o.pos)
}

type OpMulFloat2Addr struct {
	opBase
	Fmt12x
}

func (o OpMulFloat2Addr) Code() OpCode {
	return OpCodeMulFloat2Addr
}

func (o OpMulFloat2Addr) Fmt() Fmt {
	return o.Fmt12x
}

func (o OpMulFloat2Addr) String() string {
	f := o.Fmt12x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: mul-float/2addr %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: mul-float/2addr", o.pos)
}

type OpDivFloat2Addr struct {
	opBase
	Fmt12x
}

func (o OpDivFloat2Addr) Code() OpCode {
	return OpCodeDivFloat2Addr
}

func (o OpDivFloat2Addr) Fmt() Fmt {
	return o.Fmt12x
}

func (o OpDivFloat2Addr) String() string {
	f := o.Fmt12x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: div-float/2addr %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: div-float/2addr", o.pos)
}

type OpRemFloat2Addr struct {
	opBase
	Fmt12x
}

func (o OpRemFloat2Addr) Code() OpCode {
	return OpCodeRemFloat2Addr
}

func (o OpRemFloat2Addr) Fmt() Fmt {
	return o.Fmt12x
}

func (o OpRemFloat2Addr) String() string {
	f := o.Fmt12x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: rem-float/2addr %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: rem-float/2addr", o.pos)
}

type OpAddDouble2Addr struct {
	opBase
	Fmt12x
}

func (o OpAddDouble2Addr) Code() OpCode {
	return OpCodeAddDouble2Addr
}

func (o OpAddDouble2Addr) Fmt() Fmt {
	return o.Fmt12x
}

func (o OpAddDouble2Addr) String() string {
	f := o.Fmt12x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: add-double/2addr %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: add-double/2addr", o.pos)
}

type OpSubDouble2Addr struct {
	opBase
	Fmt12x
}

func (o OpSubDouble2Addr) Code() OpCode {
	return OpCodeSubDouble2Addr
}

func (o OpSubDouble2Addr) Fmt() Fmt {
	return o.Fmt12x
}

func (o OpSubDouble2Addr) String() string {
	f := o.Fmt12x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: sub-double/2addr %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: sub-double/2addr", o.pos)
}

type OpMulDouble2Addr struct {
	opBase
	Fmt12x
}

func (o OpMulDouble2Addr) Code() OpCode {
	return OpCodeMulDouble2Addr
}

func (o OpMulDouble2Addr) Fmt() Fmt {
	return o.Fmt12x
}

func (o OpMulDouble2Addr) String() string {
	f := o.Fmt12x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: mul-double/2addr %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: mul-double/2addr", o.pos)
}

type OpDivDouble2Addr struct {
	opBase
	Fmt12x
}

func (o OpDivDouble2Addr) Code() OpCode {
	return OpCodeDivDouble2Addr
}

func (o OpDivDouble2Addr) Fmt() Fmt {
	return o.Fmt12x
}

func (o OpDivDouble2Addr) String() string {
	f := o.Fmt12x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: div-double/2addr %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: div-double/2addr", o.pos)
}

type OpRemDouble2Addr struct {
	opBase
	Fmt12x
}

func (o OpRemDouble2Addr) Code() OpCode {
	return OpCodeRemDouble2Addr
}

func (o OpRemDouble2Addr) Fmt() Fmt {
	return o.Fmt12x
}

func (o OpRemDouble2Addr) String() string {
	f := o.Fmt12x.String()
	if f != "" {
		return fmt.Sprintf("0x%x: rem-double/2addr %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: rem-double/2addr", o.pos)
}

type OpAddIntLit16 struct {
	opBase
	Fmt22s
}

func (o OpAddIntLit16) Code() OpCode {
	return OpCodeAddIntLit16
}

func (o OpAddIntLit16) Fmt() Fmt {
	return o.Fmt22s
}

func (o OpAddIntLit16) String() string {
	f := o.Fmt22s.String()
	if f != "" {
		return fmt.Sprintf("0x%x: add-int/lit16 %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: add-int/lit16", o.pos)
}

type OpRsubIntLit16 struct {
	opBase
	Fmt22s
}

func (o OpRsubIntLit16) Code() OpCode {
	return OpCodeRsubIntLit16
}

func (o OpRsubIntLit16) Fmt() Fmt {
	return o.Fmt22s
}

func (o OpRsubIntLit16) String() string {
	f := o.Fmt22s.String()
	if f != "" {
		return fmt.Sprintf("0x%x: rsub-int/lit16 %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: rsub-int/lit16", o.pos)
}

type OpMulIntLit16 struct {
	opBase
	Fmt22s
}

func (o OpMulIntLit16) Code() OpCode {
	return OpCodeMulIntLit16
}

func (o OpMulIntLit16) Fmt() Fmt {
	return o.Fmt22s
}

func (o OpMulIntLit16) String() string {
	f := o.Fmt22s.String()
	if f != "" {
		return fmt.Sprintf("0x%x: mul-int/lit16 %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: mul-int/lit16", o.pos)
}

type OpDivIntLit16 struct {
	opBase
	Fmt22s
}

func (o OpDivIntLit16) Code() OpCode {
	return OpCodeDivIntLit16
}

func (o OpDivIntLit16) Fmt() Fmt {
	return o.Fmt22s
}

func (o OpDivIntLit16) String() string {
	f := o.Fmt22s.String()
	if f != "" {
		return fmt.Sprintf("0x%x: div-int/lit16 %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: div-int/lit16", o.pos)
}

type OpRemIntLit16 struct {
	opBase
	Fmt22s
}

func (o OpRemIntLit16) Code() OpCode {
	return OpCodeRemIntLit16
}

func (o OpRemIntLit16) Fmt() Fmt {
	return o.Fmt22s
}

func (o OpRemIntLit16) String() string {
	f := o.Fmt22s.String()
	if f != "" {
		return fmt.Sprintf("0x%x: rem-int/lit16 %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: rem-int/lit16", o.pos)
}

type OpAndIntLit16 struct {
	opBase
	Fmt22s
}

func (o OpAndIntLit16) Code() OpCode {
	return OpCodeAndIntLit16
}

func (o OpAndIntLit16) Fmt() Fmt {
	return o.Fmt22s
}

func (o OpAndIntLit16) String() string {
	f := o.Fmt22s.String()
	if f != "" {
		return fmt.Sprintf("0x%x: and-int/lit16 %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: and-int/lit16", o.pos)
}

type OpOrIntLit16 struct {
	opBase
	Fmt22s
}

func (o OpOrIntLit16) Code() OpCode {
	return OpCodeOrIntLit16
}

func (o OpOrIntLit16) Fmt() Fmt {
	return o.Fmt22s
}

func (o OpOrIntLit16) String() string {
	f := o.Fmt22s.String()
	if f != "" {
		return fmt.Sprintf("0x%x: or-int/lit16 %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: or-int/lit16", o.pos)
}

type OpXorIntLit16 struct {
	opBase
	Fmt22s
}

func (o OpXorIntLit16) Code() OpCode {
	return OpCodeXorIntLit16
}

func (o OpXorIntLit16) Fmt() Fmt {
	return o.Fmt22s
}

func (o OpXorIntLit16) String() string {
	f := o.Fmt22s.String()
	if f != "" {
		return fmt.Sprintf("0x%x: xor-int/lit16 %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: xor-int/lit16", o.pos)
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

type OpInvokePolymorphic struct {
	opBase
	Fmt45cc
}

func (o OpInvokePolymorphic) Code() OpCode {
	return OpCodeInvokePolymorphic
}

func (o OpInvokePolymorphic) Fmt() Fmt {
	return o.Fmt45cc
}

func (o OpInvokePolymorphic) String() string {
	f := o.Fmt45cc.String()
	if f != "" {
		return fmt.Sprintf("0x%x: invoke-polymorphic %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: invoke-polymorphic", o.pos)
}

type OpInvokePolymorphicRange struct {
	opBase
	Fmt4rcc
}

func (o OpInvokePolymorphicRange) Code() OpCode {
	return OpCodeInvokePolymorphicRange
}

func (o OpInvokePolymorphicRange) Fmt() Fmt {
	return o.Fmt4rcc
}

func (o OpInvokePolymorphicRange) String() string {
	f := o.Fmt4rcc.String()
	if f != "" {
		return fmt.Sprintf("0x%x: invoke-polymorphic/range %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: invoke-polymorphic/range", o.pos)
}

type OpInvokeCustom struct {
	opBase
	Fmt35c
}

func (o OpInvokeCustom) Code() OpCode {
	return OpCodeInvokeCustom
}

func (o OpInvokeCustom) Fmt() Fmt {
	return o.Fmt35c
}

func (o OpInvokeCustom) String() string {
	f := o.Fmt35c.String()
	if f != "" {
		return fmt.Sprintf("0x%x: invoke-custom %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: invoke-custom", o.pos)
}

type OpInvokeCustomRange struct {
	opBase
	Fmt3rc
}

func (o OpInvokeCustomRange) Code() OpCode {
	return OpCodeInvokeCustomRange
}

func (o OpInvokeCustomRange) Fmt() Fmt {
	return o.Fmt3rc
}

func (o OpInvokeCustomRange) String() string {
	f := o.Fmt3rc.String()
	if f != "" {
		return fmt.Sprintf("0x%x: invoke-custom/range %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: invoke-custom/range", o.pos)
}

type OpConstMethodHandle struct {
	opBase
	Fmt21c
}

func (o OpConstMethodHandle) Code() OpCode {
	return OpCodeConstMethodHandle
}

func (o OpConstMethodHandle) Fmt() Fmt {
	return o.Fmt21c
}

func (o OpConstMethodHandle) String() string {
	f := o.Fmt21c.String()
	if f != "" {
		return fmt.Sprintf("0x%x: const-method-handle %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: const-method-handle", o.pos)
}

type OpConstMethodType struct {
	opBase
	Fmt21c
}

func (o OpConstMethodType) Code() OpCode {
	return OpCodeConstMethodType
}

func (o OpConstMethodType) Fmt() Fmt {
	return o.Fmt21c
}

func (o OpConstMethodType) String() string {
	f := o.Fmt21c.String()
	if f != "" {
		return fmt.Sprintf("0x%x: const-method-type %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: const-method-type", o.pos)
}
