package dextk

import (
	"fmt"
	"io"
)

type Fmt interface {
	internalFmt()

	Size() int

	String() string
}

type Fmt10x struct {
}

func (f Fmt10x) internalFmt() {}

func (f Fmt10x) Size() int {
	return 1
}

func (f Fmt10x) String() string {
	return ""
}

func (r *OpReader) readFmt10x() (Fmt10x, error) {
	var res Fmt10x

	if r.pos+1 > len(r.ops) {
		return res, io.EOF
	}

	r.pos += 1

	return res, nil
}

func fmt10xSize(_ *OpReader) (int, error) {
	return 1, nil
}

type Fmt11 struct {
	A uint8
}

func (f Fmt11) internalFmt() {}

func (f Fmt11) Size() int {
	return 1
}

func (f Fmt11) String() string {
	return fmt.Sprintf("a=%v", f.A)
}

type Fmt11x struct{ Fmt11 }
type Fmt10t struct{ Fmt11 }

func (r *OpReader) readFmt11() (Fmt11, error) {
	var res Fmt11

	if r.pos+1 > len(r.ops) {
		return res, io.EOF
	}

	res.A = uint8((r.ops[r.pos] >> 8) & 0xFF)

	r.pos += 1

	return res, nil
}

func (r *OpReader) readFmt11x() (Fmt11x, error) {
	f, e := r.readFmt11()
	return Fmt11x{f}, e
}

func fmt11xSize(_ *OpReader) (int, error) {
	return 1, nil
}

func (r *OpReader) readFmt10t() (Fmt10t, error) {
	f, e := r.readFmt11()
	return Fmt10t{f}, e
}

func fmt10tSize(_ *OpReader) (int, error) {
	return 1, nil
}

type Fmt12 struct {
	A uint8
	B uint8
}

type Fmt12x struct{ Fmt12 }
type Fmt11n struct{ Fmt12 }

func (f Fmt12) internalFmt() {}

func (f Fmt12) Size() int {
	return 1
}

func (f Fmt12) String() string {
	return fmt.Sprintf("a=%v, b=%v", f.A, f.B)
}

func (r *OpReader) readFmt12() (Fmt12, error) {
	var res Fmt12

	if r.pos+1 > len(r.ops) {
		return res, io.EOF
	}

	lower := r.ops[r.pos]
	res.A = uint8((lower >> 8) & 0xF)
	res.B = uint8((lower >> 12) & 0xF)

	r.pos += 1

	return res, nil
}

func (r *OpReader) readFmt12x() (Fmt12x, error) {
	f, e := r.readFmt12()
	return Fmt12x{f}, e
}

func fmt12xSize(_ *OpReader) (int, error) {
	return 1, nil
}

func (r *OpReader) readFmt11n() (Fmt11n, error) {
	f, e := r.readFmt12()
	return Fmt11n{f}, e
}

func fmt11nSize(_ *OpReader) (int, error) {
	return 1, nil
}

type Fmt20t struct {
	A uint16
}

func (f Fmt20t) internalFmt() {}

func (f Fmt20t) Size() int {
	return 2
}

func (f Fmt20t) String() string {
	return fmt.Sprintf("a=%v", f.A)
}

func (r *OpReader) readFmt20t() (Fmt20t, error) {
	var res Fmt20t

	if r.pos+2 > len(r.ops) {
		return res, io.EOF
	}

	res.A = r.ops[r.pos+1]

	r.pos += 2

	return res, nil
}

func fmt20tSize(_ *OpReader) (int, error) {
	return 2, nil
}

type Fmt22 struct {
	A uint8
	B uint16
}

type Fmt22x struct{ Fmt22 }
type Fmt21t struct{ Fmt22 }
type Fmt21h struct{ Fmt22 }
type Fmt21c struct{ Fmt22 }
type Fmt21s struct{ Fmt22 }

func (f Fmt22) internalFmt() {}

func (f Fmt22) Size() int {
	return 2
}

func (f Fmt22) String() string {
	return fmt.Sprintf("a=%v, b=%v", f.A, f.B)
}

func (r *OpReader) readFmt22() (Fmt22, error) {
	var res Fmt22

	if r.pos+2 > len(r.ops) {
		return res, io.EOF
	}

	res.A = uint8((r.ops[r.pos] >> 8) & 0xFF)
	res.B = r.ops[r.pos+1]

	r.pos += 2

	return res, nil
}

func (r *OpReader) readFmt22x() (Fmt22x, error) {
	f, e := r.readFmt22()
	return Fmt22x{f}, e
}

func fmt22xSize(_ *OpReader) (int, error) {
	return 2, nil
}

func (r *OpReader) readFmt21t() (Fmt21t, error) {
	f, e := r.readFmt22()
	return Fmt21t{f}, e
}

func fmt21tSize(_ *OpReader) (int, error) {
	return 2, nil
}

func (r *OpReader) readFmt21h() (Fmt21h, error) {
	f, e := r.readFmt22()
	return Fmt21h{f}, e
}

func fmt21hSize(_ *OpReader) (int, error) {
	return 2, nil
}

func (r *OpReader) readFmt21c() (Fmt21c, error) {
	f, e := r.readFmt22()
	return Fmt21c{f}, e
}

func fmt21cSize(_ *OpReader) (int, error) {
	return 2, nil
}

func (r *OpReader) readFmt21s() (Fmt21s, error) {
	f, e := r.readFmt22()
	return Fmt21s{f}, e
}

func fmt21sSize(_ *OpReader) (int, error) {
	return 2, nil
}

type Fmt231 struct {
	A uint8
	B uint8
	C uint8
}

type Fmt23x struct{ Fmt231 }
type Fmt22b struct{ Fmt231 }

func (f Fmt231) internalFmt() {}

func (f Fmt231) Size() int {
	return 2
}

func (f Fmt231) String() string {
	return fmt.Sprintf("a=%v, b=%v, c=%v", f.A, f.B, f.C)
}

func (r *OpReader) readFmt231() (Fmt231, error) {
	var res Fmt231

	if r.pos+2 > len(r.ops) {
		return res, io.EOF
	}

	res.A = uint8((r.ops[r.pos] >> 8) & 0xFF)

	upper := r.ops[r.pos+1]
	res.B = uint8(upper & 0xFF)
	res.C = uint8((upper >> 8) & 0xFF)

	r.pos += 2

	return res, nil
}

func (r *OpReader) readFmt23x() (Fmt23x, error) {
	f, e := r.readFmt231()
	return Fmt23x{f}, e
}

func fmt23xSize(_ *OpReader) (int, error) {
	return 2, nil
}

func (r *OpReader) readFmt22b() (Fmt22b, error) {
	f, e := r.readFmt231()
	return Fmt22b{f}, e
}

func fmt22bSize(_ *OpReader) (int, error) {
	return 2, nil
}

type Fmt232 struct {
	A uint8
	B uint8
	C uint16
}

type Fmt22t struct{ Fmt232 }
type Fmt22s struct{ Fmt232 }
type Fmt22c struct{ Fmt232 }
type Fmt22cs struct{ Fmt232 }

func (f Fmt232) internalFmt() {}

func (f Fmt232) Size() int {
	return 2
}

func (f Fmt232) String() string {
	return fmt.Sprintf("a=%v, b=%v, c=%v", f.A, f.B, f.C)
}

func (r *OpReader) readFmt232() (Fmt232, error) {
	var res Fmt232

	if r.pos+2 > len(r.ops) {
		return res, io.EOF
	}

	lower := r.ops[r.pos]
	res.A = uint8((lower >> 8) & 0xF)
	res.B = uint8((lower >> 12) & 0xF)
	res.C = r.ops[r.pos+1]

	r.pos += 2

	return res, nil
}

func (r *OpReader) readFmt22t() (Fmt22t, error) {
	f, e := r.readFmt232()
	return Fmt22t{f}, e
}

func fmt22tSize(_ *OpReader) (int, error) {
	return 2, nil
}

func (r *OpReader) readFmt22s() (Fmt22s, error) {
	f, e := r.readFmt232()
	return Fmt22s{f}, e
}

func fmt22sSize(_ *OpReader) (int, error) {
	return 2, nil
}

func (r *OpReader) readFmt22c() (Fmt22c, error) {
	f, e := r.readFmt232()
	return Fmt22c{f}, e
}

func fmt22cSize(_ *OpReader) (int, error) {
	return 2, nil
}

func (r *OpReader) readFmt22cs() (Fmt22cs, error) {
	f, e := r.readFmt232()
	return Fmt22cs{f}, e
}

func fmt22csSize(_ *OpReader) (int, error) {
	return 2, nil
}

type Fmt30t struct {
	A uint32
}

func (f Fmt30t) internalFmt() {}

func (f Fmt30t) Size() int {
	return 2
}

func (f Fmt30t) String() string {
	return fmt.Sprintf("a=%v", f.A)
}

func (r *OpReader) readFmt30t() (Fmt30t, error) {
	var res Fmt30t

	if r.pos+3 > len(r.ops) {
		return res, io.EOF
	}

	res.A = (uint32(r.ops[r.pos+2]) << 16) | uint32(r.ops[r.pos+1])

	r.pos += 3

	return res, nil
}

func fmt30tSize(_ *OpReader) (int, error) {
	return 3, nil
}

type Fmt32x struct {
	A uint16
	B uint16
}

func (f Fmt32x) internalFmt() {}

func (f Fmt32x) Size() int {
	return 3
}

func (f Fmt32x) String() string {
	return fmt.Sprintf("a=%v, b=%v", f.A, f.B)
}

func (r *OpReader) readFmt32x() (Fmt32x, error) {
	var res Fmt32x

	if r.pos+3 > len(r.ops) {
		return res, io.EOF
	}

	res.A = r.ops[r.pos+1]
	res.B = r.ops[r.pos+2]

	r.pos += 3

	return res, nil
}

func fmt32xSize(_ *OpReader) (int, error) {
	return 3, nil
}

type Fmt35c struct {
	A uint8
	B uint16
	C uint8
	D uint8
	E uint8
	F uint8
	G uint8
}

func (f Fmt35c) internalFmt() {}

func (f Fmt35c) Size() int {
	return 3
}

func (f Fmt35c) String() string {
	return fmt.Sprintf("a=%v, b=%v, c=%v, d=%v, e=%v, f=%v, g=%v", f.A, f.B, f.C, f.D, f.E, f.F, f.G)
}

func (r *OpReader) readFmt35c() (Fmt35c, error) {
	var res Fmt35c

	if r.pos+3 > len(r.ops) {
		return res, io.EOF
	}

	lower := r.ops[r.pos]
	res.A = uint8((lower >> 12) & 0xF)
	res.G = uint8((lower >> 8) & 0xF)

	res.B = r.ops[r.pos+1]

	upper := r.ops[r.pos+2]
	res.C = uint8(upper & 0xF)
	res.D = uint8((upper >> 4) & 0xF)
	res.E = uint8((upper >> 8) & 0xF)
	res.F = uint8((upper >> 12) & 0xF)

	r.pos += 3

	return res, nil
}

func fmt35cSize(_ *OpReader) (int, error) {
	return 3, nil
}

type Fmt31 struct {
	A uint8
	B uint32
}

func (f Fmt31) internalFmt() {}

func (f Fmt31) Size() int {
	return 3
}

func (f Fmt31) String() string {
	return fmt.Sprintf("a=%v, b=%v", f.A, f.B)
}

type Fmt31i struct{ Fmt31 }
type Fmt31t struct{ Fmt31 }
type Fmt31c struct{ Fmt31 }

func (r *OpReader) readFmt31() (Fmt31, error) {
	var res Fmt31

	if r.pos+3 > len(r.ops) {
		return res, io.EOF
	}

	lower := r.ops[r.pos]
	res.A = uint8((lower >> 8) & 0xFF)
	res.B = (uint32(r.ops[r.pos+2]) << 16) | uint32(r.ops[r.pos+1])

	r.pos += 3

	return res, nil
}

func (r *OpReader) readFmt31i() (Fmt31i, error) {
	f, e := r.readFmt31()
	return Fmt31i{f}, e
}

func fmt31iSize(_ *OpReader) (int, error) {
	return 3, nil
}

func (r *OpReader) readFmt31t() (Fmt31t, error) {
	f, e := r.readFmt31()
	return Fmt31t{f}, e
}

func fmt31tSize(_ *OpReader) (int, error) {
	return 3, nil
}

func (r *OpReader) readFmt31c() (Fmt31c, error) {
	f, e := r.readFmt31()
	return Fmt31c{f}, e
}

func fmt31cSize(_ *OpReader) (int, error) {
	return 3, nil
}

type Fmt3r struct {
	A uint8
	B uint16
	C uint16
}

func (f Fmt3r) internalFmt() {}

func (f Fmt3r) Size() int {
	return 3
}

func (f Fmt3r) String() string {
	return fmt.Sprintf("a=%v, b=%v, c=%v", f.A, f.B, f.C)
}

type Fmt3rc struct{ Fmt3r }
type Fmt3rms struct{ Fmt3r }
type Fmt3rmi struct{ Fmt3r }

func (r *OpReader) readFmt3r() (Fmt3r, error) {
	var res Fmt3r

	if r.pos+3 > len(r.ops) {
		return res, io.EOF
	}

	res.A = uint8((r.ops[r.pos] >> 8) & 0xFF)
	res.B = r.ops[r.pos+1]
	res.C = r.ops[r.pos+2]

	r.pos += 3

	return res, nil
}

func (r *OpReader) readFmt3rc() (Fmt3rc, error) {
	f, e := r.readFmt3r()
	return Fmt3rc{f}, e
}

func fmt3rcSize(_ *OpReader) (int, error) {
	return 3, nil
}

func (r *OpReader) readFmt3rms() (Fmt3rms, error) {
	f, e := r.readFmt3r()
	return Fmt3rms{f}, e
}

func fmt3rmsSize(_ *OpReader) (int, error) {
	return 3, nil
}

func (r *OpReader) readFmt3rmi() (Fmt3rmi, error) {
	f, e := r.readFmt3r()
	return Fmt3rmi{f}, e
}

func fmt3rmiSize(_ *OpReader) (int, error) {
	return 3, nil
}

type Fmt45cc struct {
	A uint8
	B uint16
	C uint8
	D uint8
	E uint8
	F uint8
	G uint8
	H uint16
}

func (f Fmt45cc) internalFmt() {}

func (f Fmt45cc) Size() int {
	return 4
}

func (f Fmt45cc) String() string {
	return fmt.Sprintf("a=%v, b=%v, c=%v, d=%v, e=%v, f=%v, g=%v, h=%v", f.A, f.B, f.C, f.D, f.E, f.F, f.G, f.H)
}

func (r *OpReader) readFmt45cc() (Fmt45cc, error) {
	var res Fmt45cc

	if r.pos+4 > len(r.ops) {
		return res, io.EOF
	}

	lower := r.ops[r.pos]
	res.A = uint8((lower >> 12) & 0xF)
	res.G = uint8((lower >> 8) & 0xF)

	res.B = r.ops[r.pos+1]

	upper := r.ops[r.pos+2]
	res.C = uint8(upper & 0xF)
	res.D = uint8((upper >> 4) & 0xF)
	res.E = uint8((upper >> 8) & 0xF)
	res.F = uint8((upper >> 12) & 0xF)

	res.H = r.ops[r.pos+3]

	r.pos += 4

	return res, nil
}

func fmt45ccSize(_ *OpReader) (int, error) {
	return 4, nil
}

type Fmt4rcc struct {
	A uint8
	B uint16
	C uint16
	H uint16
}

func (f Fmt4rcc) internalFmt() {}

func (f Fmt4rcc) Size() int {
	return 4
}

func (f Fmt4rcc) String() string {
	return fmt.Sprintf("a=%v, b=%v, c=%v, h=%v", f.A, f.B, f.C, f.H)
}

func (r *OpReader) readFmt4rcc() (Fmt4rcc, error) {
	var res Fmt4rcc

	if r.pos+4 > len(r.ops) {
		return res, io.EOF
	}

	lower := r.ops[r.pos]
	res.A = uint8((lower >> 8) & 0xFF)

	res.B = r.ops[r.pos+1]
	res.C = r.ops[r.pos+1]
	res.H = r.ops[r.pos+3]

	r.pos += 4

	return res, nil
}

func fmt4rccSize(_ *OpReader) (int, error) {
	return 4, nil
}

type Fmt51l struct {
	A uint8
	B uint64
}

func (f Fmt51l) internalFmt() {}

func (f Fmt51l) Size() int {
	return 5
}

func (f Fmt51l) String() string {
	return fmt.Sprintf("a=%v, b=%v", f.A, f.B)
}

func (r *OpReader) readFmt51l() (Fmt51l, error) {
	var res Fmt51l

	if r.pos+5 > len(r.ops) {
		return res, io.EOF
	}

	lower := r.ops[r.pos]
	res.A = uint8((lower >> 8) & 0xFF)
	res.B = (uint64(r.ops[r.pos+4]) << 48) | (uint64(r.ops[r.pos+3]) << 32) |
		(uint64(r.ops[r.pos+2]) << 16) | uint64(r.ops[r.pos+1])

	r.pos += 5

	return res, nil
}

func fmt51lSize(_ *OpReader) (int, error) {
	return 5, nil
}
