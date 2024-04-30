package decompiler

type Flags struct {
	value uint32
}

func NewFlags(value uint32) *Flags {
	return &Flags{value}
}

func (f *Flags) Take(bits uint32) bool {
	if f.value&bits == bits {
		f.value &= ^bits

		return true
	}

	return false
}

func (f *Flags) Remaining() uint32 {
	return f.value
}
