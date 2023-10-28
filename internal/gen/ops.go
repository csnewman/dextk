package main

type Op struct {
	Code int
	Name string
	Fmt  string
}

var Ops = []Op{
	{
		Code: 0x01,
		Name: "nop",
		Fmt:  "10x",
	},
	{
		Code: 0x11,
		Name: "return-object",
		Fmt:  "11x",
	},
	{
		Code: 0x22,
		Name: "new-instance",
		Fmt:  "21c",
	},
	{
		Code: 0x70,
		Name: "invoke-direct",
		Fmt:  "35c",
	},
}
