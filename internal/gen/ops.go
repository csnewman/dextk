package main

type Op struct {
	Code int
	Name string
	Fmt  string
}

var Ops = []Op{
	{
		Code: 0x00,
		Name: "nop",
		Fmt:  "10x",
	},
	{
		Code: 0x01,
		Name: "move",
		Fmt:  "12x",
	},
	{
		Code: 0x02,
		Name: "move/from16",
		Fmt:  "22x",
	},
	{
		Code: 0x03,
		Name: "move/16",
		Fmt:  "32x",
	},
	{
		Code: 0x04,
		Name: "move-wide",
		Fmt:  "12x",
	},
	{
		Code: 0x05,
		Name: "move-wide/from16",
		Fmt:  "22x",
	},
	{
		Code: 0x06,
		Name: "move-wide/16",
		Fmt:  "32x",
	},
	{
		Code: 0x07,
		Name: "move-object",
		Fmt:  "12x",
	},
	{
		Code: 0x08,
		Name: "move-object/from16",
		Fmt:  "22x",
	},
	{
		Code: 0x09,
		Name: "move-object/16",
		Fmt:  "32x",
	},
	{
		Code: 0x0a,
		Name: "move-result",
		Fmt:  "11x",
	},
	{
		Code: 0x0b,
		Name: "move-result-wide",
		Fmt:  "11x",
	},
	{
		Code: 0x0c,
		Name: "move-result-object",
		Fmt:  "11x",
	},
	{
		Code: 0x0d,
		Name: "move-exception",
		Fmt:  "11x",
	},
	{
		Code: 0x0e,
		Name: "return-void",
		Fmt:  "10x",
	},
	{
		Code: 0x0f,
		Name: "return",
		Fmt:  "11x",
	},
	{
		Code: 0x10,
		Name: "return-wide",
		Fmt:  "11x",
	},
	{
		Code: 0x11,
		Name: "return-object",
		Fmt:  "11x",
	},
	{
		Code: 0x12,
		Name: "const/4",
		Fmt:  "11n",
	},
	{
		Code: 0x13,
		Name: "const/16",
		Fmt:  "21s",
	},
	{
		Code: 0x14,
		Name: "const",
		Fmt:  "31i",
	},
	{
		Code: 0x15,
		Name: "const/high16",
		Fmt:  "21h",
	},
	{
		Code: 0x16,
		Name: "const-wide/16",
		Fmt:  "21s",
	},
	{
		Code: 0x17,
		Name: "const-wide/32",
		Fmt:  "31i",
	},
	{
		Code: 0x18,
		Name: "const-wide",
		Fmt:  "51l",
	},
	{
		Code: 0x19,
		Name: "const-wide/high16",
		Fmt:  "21h",
	},
	{
		Code: 0x1a,
		Name: "const-string",
		Fmt:  "21c",
	},
	{
		Code: 0x1b,
		Name: "const-string/jumbo",
		Fmt:  "31c",
	},
	{
		Code: 0x1c,
		Name: "const-class",
		Fmt:  "21c",
	},
	{
		Code: 0x1d,
		Name: "monitor-enter",
		Fmt:  "11x",
	},
	{
		Code: 0x1e,
		Name: "monitor-exit",
		Fmt:  "11x",
	},
	{
		Code: 0x1f,
		Name: "check-cast",
		Fmt:  "21c",
	},
	{
		Code: 0x20,
		Name: "instance-of",
		Fmt:  "22c",
	},
	{
		Code: 0x21,
		Name: "array-length",
		Fmt:  "12x",
	},
	{
		Code: 0x22,
		Name: "new-instance",
		Fmt:  "21c",
	},
	{
		Code: 0x23,
		Name: "new-array",
		Fmt:  "22c",
	},
	{
		Code: 0x24,
		Name: "filled-new-array",
		Fmt:  "35c",
	},
	{
		Code: 0x25,
		Name: "filled-new-array/range",
		Fmt:  "3rc",
	},
	{
		Code: 0x26,
		Name: "fill-array-data",
		Fmt:  "31t",
	},
	{
		Code: 0x27,
		Name: "throw",
		Fmt:  "11x",
	},
	{
		Code: 0x28,
		Name: "goto",
		Fmt:  "10t",
	},
	{
		Code: 0x29,
		Name: "goto/16",
		Fmt:  "20t",
	},
	{
		Code: 0x2a,
		Name: "goto/32",
		Fmt:  "30t",
	},
	{
		Code: 0x2b,
		Name: "packed-switch",
		Fmt:  "31t",
	},
	{
		Code: 0x2c,
		Name: "sparse-switch",
		Fmt:  "31t",
	},
	{
		Code: 0x2d,
		Name: "cmpl-float",
		Fmt:  "23x",
	},
	{
		Code: 0x2e,
		Name: "cmpg-float",
		Fmt:  "23x",
	},
	{
		Code: 0x2f,
		Name: "cmpl-double",
		Fmt:  "23x",
	},
	{
		Code: 0x30,
		Name: "cmpg-double",
		Fmt:  "23x",
	},
	{
		Code: 0x31,
		Name: "cmp-long",
		Fmt:  "23x",
	},
	{
		Code: 0x32,
		Name: "if-eq",
		Fmt:  "22t",
	},
	{
		Code: 0x33,
		Name: "if-ne",
		Fmt:  "22t",
	},
	{
		Code: 0x34,
		Name: "if-lt",
		Fmt:  "22t",
	},
	{
		Code: 0x35,
		Name: "if-ge",
		Fmt:  "22t",
	},
	{
		Code: 0x36,
		Name: "if-gt",
		Fmt:  "22t",
	},
	{
		Code: 0x37,
		Name: "if-le",
		Fmt:  "22t",
	},
	{
		Code: 0x38,
		Name: "if-eqz",
		Fmt:  "21t",
	},
	{
		Code: 0x39,
		Name: "if-nez",
		Fmt:  "21t",
	},
	{
		Code: 0x3a,
		Name: "if-ltz",
		Fmt:  "21t",
	},
	{
		Code: 0x3b,
		Name: "if-gez",
		Fmt:  "21t",
	},
	{
		Code: 0x3c,
		Name: "if-gtz",
		Fmt:  "21t",
	},
	{
		Code: 0x3d,
		Name: "if-lez",
		Fmt:  "21t",
	},
	// 0x3e-0x43 unused
	{
		Code: 0x44,
		Name: "aget",
		Fmt:  "23x",
	},
	{
		Code: 0x45,
		Name: "aget-wide",
		Fmt:  "23x",
	},
	{
		Code: 0x46,
		Name: "aget-object",
		Fmt:  "23x",
	},
	{
		Code: 0x47,
		Name: "aget-boolean",
		Fmt:  "23x",
	},
	{
		Code: 0x48,
		Name: "aget-byte",
		Fmt:  "23x",
	},
	{
		Code: 0x49,
		Name: "aget-char",
		Fmt:  "23x",
	},
	{
		Code: 0x4a,
		Name: "aget-short",
		Fmt:  "23x",
	},
	{
		Code: 0x4b,
		Name: "aput",
		Fmt:  "23x",
	},
	{
		Code: 0x4c,
		Name: "aput-wide",
		Fmt:  "23x",
	},
	{
		Code: 0x4d,
		Name: "aput-object",
		Fmt:  "23x",
	},
	{
		Code: 0x4e,
		Name: "aput-boolean",
		Fmt:  "23x",
	},
	{
		Code: 0x4f,
		Name: "aput-byte",
		Fmt:  "23x",
	},
	{
		Code: 0x50,
		Name: "aput-char",
		Fmt:  "23x",
	},
	{
		Code: 0x51,
		Name: "aput-short",
		Fmt:  "23x",
	},
	{
		Code: 0x52,
		Name: "iget",
		Fmt:  "22c",
	},
	{
		Code: 0x53,
		Name: "iget-wide",
		Fmt:  "22c",
	},
	{
		Code: 0x54,
		Name: "iget-object",
		Fmt:  "22c",
	},
	{
		Code: 0x55,
		Name: "iget-boolean",
		Fmt:  "22c",
	},
	{
		Code: 0x56,
		Name: "iget-byte",
		Fmt:  "22c",
	},
	{
		Code: 0x57,
		Name: "iget-char",
		Fmt:  "22c",
	},
	{
		Code: 0x58,
		Name: "iget-short",
		Fmt:  "22c",
	},
	{
		Code: 0x59,
		Name: "iput",
		Fmt:  "22c",
	},
	{
		Code: 0x5a,
		Name: "iput-wide",
		Fmt:  "22c",
	},
	{
		Code: 0x5b,
		Name: "iput-object",
		Fmt:  "22c",
	},
	{
		Code: 0x5c,
		Name: "iput-boolean",
		Fmt:  "22c",
	},
	{
		Code: 0x5d,
		Name: "iput-byte",
		Fmt:  "22c",
	},
	{
		Code: 0x5e,
		Name: "iput-char",
		Fmt:  "22c",
	},
	{
		Code: 0x5f,
		Name: "iput-short",
		Fmt:  "22c",
	},
	{
		Code: 0x60,
		Name: "sget",
		Fmt:  "21c",
	},
	{
		Code: 0x61,
		Name: "sget-wide",
		Fmt:  "21c",
	},
	{
		Code: 0x62,
		Name: "sget-object",
		Fmt:  "21c",
	},
	{
		Code: 0x63,
		Name: "sget-boolean",
		Fmt:  "21c",
	},
	{
		Code: 0x64,
		Name: "sget-byte",
		Fmt:  "21c",
	},
	{
		Code: 0x65,
		Name: "sget-char",
		Fmt:  "21c",
	},
	{
		Code: 0x66,
		Name: "sget-short",
		Fmt:  "21c",
	},
	{
		Code: 0x67,
		Name: "sput",
		Fmt:  "21c",
	},
	{
		Code: 0x68,
		Name: "sput-wide",
		Fmt:  "21c",
	},
	{
		Code: 0x69,
		Name: "sput-object",
		Fmt:  "21c",
	},
	{
		Code: 0x6a,
		Name: "sput-boolean",
		Fmt:  "21c",
	},
	{
		Code: 0x6b,
		Name: "sput-byte",
		Fmt:  "21c",
	},
	{
		Code: 0x6c,
		Name: "sput-char",
		Fmt:  "21c",
	},
	{
		Code: 0x6d,
		Name: "sput-short",
		Fmt:  "21c",
	},
	{
		Code: 0x6e,
		Name: "invoke-virtual",
		Fmt:  "35c",
	},
	{
		Code: 0x6f,
		Name: "invoke-super",
		Fmt:  "35c",
	},
	{
		Code: 0x70,
		Name: "invoke-direct",
		Fmt:  "35c",
	},
	{
		Code: 0x71,
		Name: "invoke-static",
		Fmt:  "35c",
	},
	{
		Code: 0x72,
		Name: "invoke-interface",
		Fmt:  "35c",
	},
	// 0x73 unused
	{
		Code: 0x74,
		Name: "invoke-virtual/range",
		Fmt:  "3rc",
	},
	{
		Code: 0x75,
		Name: "invoke-super/range",
		Fmt:  "3rc",
	},
	{
		Code: 0x76,
		Name: "invoke-direct/range",
		Fmt:  "3rc",
	},
	{
		Code: 0x77,
		Name: "invoke-static/range",
		Fmt:  "3rc",
	},
	{
		Code: 0x78,
		Name: "invoke-interface/range",
		Fmt:  "3rc",
	},
	// 0x79-0x7a unused
	{
		Code: 0x7b,
		Name: "neg-int",
		Fmt:  "12x",
	},
	{
		Code: 0x7c,
		Name: "not-int",
		Fmt:  "12x",
	},
	{
		Code: 0x7d,
		Name: "neg-long",
		Fmt:  "12x",
	},
	{
		Code: 0x7e,
		Name: "not-long",
		Fmt:  "12x",
	},
	{
		Code: 0x7f,
		Name: "neg-float",
		Fmt:  "12x",
	},
	{
		Code: 0x80,
		Name: "neg-double",
		Fmt:  "12x",
	},
	{
		Code: 0x81,
		Name: "int-to-long",
		Fmt:  "12x",
	},
	{
		Code: 0x82,
		Name: "int-to-float",
		Fmt:  "12x",
	},
	{
		Code: 0x83,
		Name: "int-to-double",
		Fmt:  "12x",
	},
	{
		Code: 0x84,
		Name: "long-to-int",
		Fmt:  "12x",
	},
	{
		Code: 0x85,
		Name: "long-to-float",
		Fmt:  "12x",
	},
	{
		Code: 0x86,
		Name: "long-to-double",
		Fmt:  "12x",
	},
	{
		Code: 0x87,
		Name: "float-to-int",
		Fmt:  "12x",
	},
	{
		Code: 0x88,
		Name: "float-to-long",
		Fmt:  "12x",
	},
	{
		Code: 0x89,
		Name: "float-to-double",
		Fmt:  "12x",
	},
	{
		Code: 0x8a,
		Name: "double-to-int",
		Fmt:  "12x",
	},
	{
		Code: 0x8b,
		Name: "double-to-long",
		Fmt:  "12x",
	},
	{
		Code: 0x8c,
		Name: "double-to-float",
		Fmt:  "12x",
	},
	{
		Code: 0x8d,
		Name: "int-to-byte",
		Fmt:  "12x",
	},
	{
		Code: 0x8e,
		Name: "int-to-char",
		Fmt:  "12x",
	},
	{
		Code: 0x8f,
		Name: "int-to-short",
		Fmt:  "12x",
	},
	{
		Code: 0x90,
		Name: "add-int",
		Fmt:  "23x",
	},
	{
		Code: 0x91,
		Name: "sub-int",
		Fmt:  "23x",
	},
	{
		Code: 0x92,
		Name: "mul-int",
		Fmt:  "23x",
	},
	{
		Code: 0x93,
		Name: "div-int",
		Fmt:  "23x",
	},
	{
		Code: 0x94,
		Name: "rem-int",
		Fmt:  "23x",
	},
	{
		Code: 0x95,
		Name: "and-int",
		Fmt:  "23x",
	},
	{
		Code: 0x96,
		Name: "or-int",
		Fmt:  "23x",
	},
	{
		Code: 0x97,
		Name: "xor-int",
		Fmt:  "23x",
	},
	{
		Code: 0x98,
		Name: "shl-int",
		Fmt:  "23x",
	},
	{
		Code: 0x99,
		Name: "shr-int",
		Fmt:  "23x",
	},
	{
		Code: 0x9a,
		Name: "ushr-int",
		Fmt:  "23x",
	},
	{
		Code: 0x9b,
		Name: "add-long",
		Fmt:  "23x",
	},
	{
		Code: 0x9c,
		Name: "sub-long",
		Fmt:  "23x",
	},
	{
		Code: 0x9d,
		Name: "mul-long",
		Fmt:  "23x",
	},
	{
		Code: 0x9e,
		Name: "div-long",
		Fmt:  "23x",
	},
	{
		Code: 0x9f,
		Name: "rem-long",
		Fmt:  "23x",
	},
	{
		Code: 0xa0,
		Name: "and-long",
		Fmt:  "23x",
	},
	{
		Code: 0xa1,
		Name: "or-long",
		Fmt:  "23x",
	},
	{
		Code: 0xa2,
		Name: "xor-long",
		Fmt:  "23x",
	},
	{
		Code: 0xa3,
		Name: "shl-long",
		Fmt:  "23x",
	},
	{
		Code: 0xa4,
		Name: "shr-long",
		Fmt:  "23x",
	},
	{
		Code: 0xa5,
		Name: "ushr-long",
		Fmt:  "23x",
	},
	{
		Code: 0xa6,
		Name: "add-float",
		Fmt:  "23x",
	},
	{
		Code: 0xa7,
		Name: "sub-float",
		Fmt:  "23x",
	},
	{
		Code: 0xa8,
		Name: "mul-float",
		Fmt:  "23x",
	},
	{
		Code: 0xa9,
		Name: "div-float",
		Fmt:  "23x",
	},
	{
		Code: 0xaa,
		Name: "rem-float",
		Fmt:  "23x",
	},
	{
		Code: 0xab,
		Name: "add-double",
		Fmt:  "23x",
	},
	{
		Code: 0xac,
		Name: "sub-double",
		Fmt:  "23x",
	},
	{
		Code: 0xad,
		Name: "mul-double",
		Fmt:  "23x",
	},
	{
		Code: 0xae,
		Name: "div-double",
		Fmt:  "23x",
	},
	{
		Code: 0xaf,
		Name: "rem-double",
		Fmt:  "23x",
	},
	{
		Code: 0xb0,
		Name: "add-int/2addr",
		Fmt:  "12x",
	},
	{
		Code: 0xb1,
		Name: "sub-int/2addr",
		Fmt:  "12x",
	},
	{
		Code: 0xb2,
		Name: "mul-int/2addr",
		Fmt:  "12x",
	},
	{
		Code: 0xb3,
		Name: "div-int/2addr",
		Fmt:  "12x",
	},
	{
		Code: 0xb4,
		Name: "rem-int/2addr",
		Fmt:  "12x",
	},
	{
		Code: 0xb5,
		Name: "and-int/2addr",
		Fmt:  "12x",
	},
	{
		Code: 0xb6,
		Name: "or-int/2addr",
		Fmt:  "12x",
	},
	{
		Code: 0xb7,
		Name: "xor-int/2addr",
		Fmt:  "12x",
	},
	{
		Code: 0xb8,
		Name: "shl-int/2addr",
		Fmt:  "12x",
	},
	{
		Code: 0xb9,
		Name: "shr-int/2addr",
		Fmt:  "12x",
	},
	{
		Code: 0xba,
		Name: "ushr-int/2addr",
		Fmt:  "12x",
	},
	{
		Code: 0xbb,
		Name: "add-long/2addr",
		Fmt:  "12x",
	},
	{
		Code: 0xbc,
		Name: "sub-long/2addr",
		Fmt:  "12x",
	},
	{
		Code: 0xbd,
		Name: "mul-long/2addr",
		Fmt:  "12x",
	},
	{
		Code: 0xbe,
		Name: "div-long/2addr",
		Fmt:  "12x",
	},
	{
		Code: 0xbf,
		Name: "rem-long/2addr",
		Fmt:  "12x",
	},
	{
		Code: 0xc0,
		Name: "and-long/2addr",
		Fmt:  "12x",
	},
	{
		Code: 0xc1,
		Name: "or-long/2addr",
		Fmt:  "12x",
	},
	{
		Code: 0xc2,
		Name: "xor-long/2addr",
		Fmt:  "12x",
	},
	{
		Code: 0xc3,
		Name: "shl-long/2addr",
		Fmt:  "12x",
	},
	{
		Code: 0xc4,
		Name: "shr-long/2addr",
		Fmt:  "12x",
	},
	{
		Code: 0xc5,
		Name: "ushr-long/2addr",
		Fmt:  "12x",
	},
	{
		Code: 0xc6,
		Name: "add-float/2addr",
		Fmt:  "12x",
	},
	{
		Code: 0xc7,
		Name: "sub-float/2addr",
		Fmt:  "12x",
	},
	{
		Code: 0xc8,
		Name: "mul-float/2addr",
		Fmt:  "12x",
	},
	{
		Code: 0xc9,
		Name: "div-float/2addr",
		Fmt:  "12x",
	},
	{
		Code: 0xca,
		Name: "rem-float/2addr",
		Fmt:  "12x",
	},
	{
		Code: 0xcb,
		Name: "add-double/2addr",
		Fmt:  "12x",
	},
	{
		Code: 0xcc,
		Name: "sub-double/2addr",
		Fmt:  "12x",
	},
	{
		Code: 0xcd,
		Name: "mul-double/2addr",
		Fmt:  "12x",
	},
	{
		Code: 0xce,
		Name: "div-double/2addr",
		Fmt:  "12x",
	},
	{
		Code: 0xcf,
		Name: "rem-double/2addr",
		Fmt:  "12x",
	},
	{
		Code: 0xd0,
		Name: "add-int/lit16",
		Fmt:  "22s",
	},
	{
		Code: 0xd1,
		Name: "rsub-int/lit16",
		Fmt:  "22s",
	},
	{
		Code: 0xd2,
		Name: "mul-int/lit16",
		Fmt:  "22s",
	},
	{
		Code: 0xd3,
		Name: "div-int/lit16",
		Fmt:  "22s",
	},
	{
		Code: 0xd4,
		Name: "rem-int/lit16",
		Fmt:  "22s",
	},
	{
		Code: 0xd5,
		Name: "and-int/lit16",
		Fmt:  "22s",
	},
	{
		Code: 0xd6,
		Name: "or-int/lit16",
		Fmt:  "22s",
	},
	{
		Code: 0xd7,
		Name: "xor-int/lit16",
		Fmt:  "22s",
	},
	{
		Code: 0xd8,
		Name: "add-int/lit8",
		Fmt:  "22b",
	},
	{
		Code: 0xd9,
		Name: "rsub-int/lit8",
		Fmt:  "22b",
	},
	{
		Code: 0xda,
		Name: "mul-int/lit8",
		Fmt:  "22b",
	},
	{
		Code: 0xdb,
		Name: "div-int/lit8",
		Fmt:  "22b",
	},
	{
		Code: 0xdc,
		Name: "rem-int/lit8",
		Fmt:  "22b",
	},
	{
		Code: 0xdd,
		Name: "and-int/lit8",
		Fmt:  "22b",
	},
	{
		Code: 0xde,
		Name: "or-int/lit8",
		Fmt:  "22b",
	},
	{
		Code: 0xdf,
		Name: "xor-int/lit8",
		Fmt:  "22b",
	},
	{
		Code: 0xe0,
		Name: "shl-int/lit8",
		Fmt:  "22b",
	},
	{
		Code: 0xe1,
		Name: "shr-int/lit8",
		Fmt:  "22b",
	},
	{
		Code: 0xe2,
		Name: "ushr-int/lit8",
		Fmt:  "22b",
	},
	// 0xe3-0xf9 unused
	{
		Code: 0xfa,
		Name: "invoke-polymorphic",
		Fmt:  "45cc",
	},
	{
		Code: 0xfb,
		Name: "invoke-polymorphic/range",
		Fmt:  "4rcc",
	},
	{
		Code: 0xfc,
		Name: "invoke-custom",
		Fmt:  "35c",
	},
	{
		Code: 0xfd,
		Name: "invoke-custom/range",
		Fmt:  "3rc",
	},
	{
		Code: 0xfe,
		Name: "const-method-handle",
		Fmt:  "21c",
	},
	{
		Code: 0xff,
		Name: "const-method-type",
		Fmt:  "21c",
	},
}
