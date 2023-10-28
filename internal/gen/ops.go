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
	// TODO: 0x18
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
	// TODO: missing
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
	// TODO: missing
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
	// TODO: missing
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
	// TODO: missing
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
}
