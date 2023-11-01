package dextk

import (
	"errors"
	"fmt"
	"log"
	"math"
)

var ErrBadOp = errors.New("malformed op")

type OpNode interface {
	RawOp() Op

	String() string
}

type NopOpNode struct {
	Raw Op
}

func (n NopOpNode) RawOp() Op {
	return n.Raw
}

func (n NopOpNode) String() string {
	return n.Raw.Code().String()
}

type MoveOpNode struct {
	Raw Op
	Src uint16
	Dst uint16
}

func (n MoveOpNode) RawOp() Op {
	return n.Raw
}

func (n MoveOpNode) String() string {
	return fmt.Sprintf("%v src=%v dst=%v", n.Raw.Code(), n.Src, n.Dst)
}

type MoveResultOpNode struct {
	Raw Op
	Dst uint8
}

func (n MoveResultOpNode) RawOp() Op {
	return n.Raw
}

func (n MoveResultOpNode) String() string {
	return fmt.Sprintf("%v dst=%v", n.Raw.Code(), n.Dst)
}

type MoveExceptionOpNode struct {
	Raw Op
	Dst uint8
}

func (n MoveExceptionOpNode) RawOp() Op {
	return n.Raw
}

func (n MoveExceptionOpNode) String() string {
	return fmt.Sprintf("%v dst=%v", n.Raw.Code(), n.Dst)
}

type ReturnOpNode struct {
	Raw Op
	// -1 = void
	Value int16
}

func (n ReturnOpNode) RawOp() Op {
	return n.Raw
}

func (n ReturnOpNode) String() string {
	return fmt.Sprintf("%v value=%v", n.Raw.Code(), n.Value)
}

type ConstOpNode struct {
	Raw   Op
	Dst   uint8
	Value any
}

func (n ConstOpNode) RawOp() Op {
	return n.Raw
}

func (n ConstOpNode) String() string {
	return fmt.Sprintf("%v dst=%v value='%v'", n.Raw.Code(), n.Dst, n.Value)
}

type MonitorOpNode struct {
	Raw   Op
	Ref   uint8
	Enter bool
}

func (n MonitorOpNode) RawOp() Op {
	return n.Raw
}

func (n MonitorOpNode) String() string {
	return fmt.Sprintf("%v ref=%v enter=%v", n.Raw.Code(), n.Ref, n.Enter)
}

type CheckCastOpNode struct {
	Raw  Op
	Ref  uint8
	Type TypeDescriptor
}

func (n CheckCastOpNode) RawOp() Op {
	return n.Raw
}

func (n CheckCastOpNode) String() string {
	return fmt.Sprintf("%v ref=%v type=%v", n.Raw.Code(), n.Ref, n.Type)
}

type NewInstanceOpNode struct {
	Raw  Op
	Dst  uint8
	Type string
}

func (n NewInstanceOpNode) RawOp() Op {
	return n.Raw
}

func (n NewInstanceOpNode) String() string {
	return fmt.Sprintf("%v dst=%v type=%v", n.Raw.Code(), n.Dst, n.Type)
}

type GotoOpNode struct {
	Raw Op
	Tgt int
}

func (n GotoOpNode) RawOp() Op {
	return n.Raw
}

func (n GotoOpNode) String() string {
	return fmt.Sprintf("%v tgt=%v", n.Raw.Code(), n.Tgt)
}

type IfTestOpNode struct {
	Raw Op
	A   uint8
	// -1 = const 0
	B   int16
	Tgt int
}

func (n IfTestOpNode) RawOp() Op {
	return n.Raw
}

func (n IfTestOpNode) String() string {
	return fmt.Sprintf("%v a=%v b=%v tgt=%v", n.Raw.Code(), n.A, n.B, n.Tgt)
}

type IGetOpNode struct {
	Raw   Op
	Dst   uint8
	Obj   uint8
	Field FieldRef
}

func (n IGetOpNode) RawOp() Op {
	return n.Raw
}

func (n IGetOpNode) String() string {
	return fmt.Sprintf("%v obj=%v field=%v dst=%v", n.Raw.Code(), n.Obj, n.Field, n.Dst)
}

type IPutOpNode struct {
	Raw   Op
	Src   uint8
	Obj   uint8
	Field FieldRef
}

func (n IPutOpNode) RawOp() Op {
	return n.Raw
}

func (n IPutOpNode) String() string {
	return fmt.Sprintf("%v obj=%v field=%v src=%v", n.Raw.Code(), n.Obj, n.Field, n.Src)
}

type SGetOpNode struct {
	Raw   Op
	Dst   uint8
	Field FieldRef
}

func (n SGetOpNode) RawOp() Op {
	return n.Raw
}

func (n SGetOpNode) String() string {
	return fmt.Sprintf("%v field=%v dst=%v", n.Raw.Code(), n.Field, n.Dst)
}

type SPutOpNode struct {
	Raw   Op
	Src   uint8
	Field FieldRef
}

func (n SPutOpNode) RawOp() Op {
	return n.Raw
}

func (n SPutOpNode) String() string {
	return fmt.Sprintf("%v field=%v src=%v", n.Raw.Code(), n.Field, n.Src)
}

type InvokeOpNode struct {
	Raw    Op
	Method MethodRef
	Args   []uint8
}

func (n InvokeOpNode) RawOp() Op {
	return n.Raw
}

func (n InvokeOpNode) String() string {
	return fmt.Sprintf("%v method=%v args=%v", n.Raw.Code(), n.Method, n.Args)
}

type CodeNode struct {
	RegisterCount uint16
	IncomingCount uint16
	OutgoingCount uint16
	Ops           []OpNode
	// TODO: DebugInfoOff, Tries, Handlers
}

func (r *Reader) ReadCodeAndParse(off uint32) (CodeNode, error) {
	var res CodeNode

	c, err := r.ReadCode(off)
	if err != nil {
		return res, nil
	}

	res.RegisterCount = c.RegisterCount
	res.IncomingCount = c.IncomingCount
	res.OutgoingCount = c.OutgoingCount

	or := NewOpReader(c.Insns)

	// Perform an initial pass to discover instruction positions
	idMap := make([]uint, len(c.Insns))
	idPos := uint(0)

	for or.HasMore() {
		idMap[or.pos] = idPos

		err := or.Skip()
		if err != nil {
			return res, err
		}

		idPos++
	}

	// Second pass to parse instructions
	res.Ops = make([]OpNode, idPos)
	idPos = 0

	or.Seek(0)

	for or.HasMore() {
		pos := or.pos

		o, err := or.Read()
		if err != nil {
			return res, err
		}

		var parsed OpNode

		switch ot := o.(type) {
		case *OpNop:
			parsed = &NopOpNode{
				Raw: ot,
			}
		case *OpMove, *OpMoveWide, *OpMoveObject:
			data := (ot.Fmt()).(Fmt12x)

			parsed = &MoveOpNode{
				Raw: ot,
				Src: uint16(data.B),
				Dst: uint16(data.A),
			}
		case *OpMoveFrom16, *OpMoveWideFrom16, *OpMoveObjectFrom16:
			data := (ot.Fmt()).(Fmt22x)

			parsed = &MoveOpNode{
				Raw: ot,
				Src: data.B,
				Dst: uint16(data.A),
			}
		case *OpMove16, *OpMoveWide16, *OpMoveObject16:
			data := (ot.Fmt()).(Fmt32x)

			parsed = &MoveOpNode{
				Raw: ot,
				Src: data.B,
				Dst: data.A,
			}
		case *OpMoveResult, *OpMoveResultWide, *OpMoveResultObject:
			data := (ot.Fmt()).(Fmt11x)

			parsed = &MoveResultOpNode{
				Raw: ot,
				Dst: data.A,
			}
		case *OpMoveException:
			data := (ot.Fmt()).(Fmt11x)

			parsed = &MoveExceptionOpNode{
				Raw: ot,
				Dst: data.A,
			}
		case *OpReturnVoid:
			parsed = &ReturnOpNode{
				Raw:   ot,
				Value: -1,
			}
		case *OpReturn, *OpReturnWide, *OpReturnObject:
			data := (ot.Fmt()).(Fmt11x)

			parsed = &ReturnOpNode{
				Raw:   ot,
				Value: int16(data.A),
			}
		case *OpConst4:
			data := (ot.Fmt()).(Fmt11n)

			val := uint32(data.B)
			if (val & 0b1000) != 0 {
				val |= math.MaxUint32 >> 4 << 4
			}

			conv := int32(val)

			parsed = &ConstOpNode{
				Raw:   ot,
				Dst:   data.A,
				Value: conv,
			}
		case *OpConst16:
			data := (ot.Fmt()).(Fmt21s)

			val := uint32(data.B)
			if (val & (1 << 15)) != 0 {
				val |= math.MaxUint32 >> 16 << 16
			}

			conv := int32(val)

			parsed = &ConstOpNode{
				Raw:   ot,
				Dst:   data.A,
				Value: conv,
			}
		case *OpConst:
			data := (ot.Fmt()).(Fmt31i)
			conv := int32(data.B)

			parsed = &ConstOpNode{
				Raw:   ot,
				Dst:   data.A,
				Value: conv,
			}
		case *OpConstHigh16:
			data := (ot.Fmt()).(Fmt21h)
			conv := int32(uint32(data.B) << 16)

			parsed = &ConstOpNode{
				Raw:   ot,
				Dst:   data.A,
				Value: conv,
			}
		case *OpConstWide16:
			data := (ot.Fmt()).(Fmt21s)

			val := uint64(data.B)
			if (val & (1 << 15)) != 0 {
				val |= math.MaxUint64 >> 16 << 16
			}

			conv := int64(val)

			parsed = &ConstOpNode{
				Raw:   ot,
				Dst:   data.A,
				Value: conv,
			}
		case *OpConstWide32:
			data := (ot.Fmt()).(Fmt31i)

			val := uint64(data.B)
			if (val & (1 << 31)) != 0 {
				val |= math.MaxUint64 >> 32 << 32
			}

			conv := int64(val)

			parsed = &ConstOpNode{
				Raw:   ot,
				Dst:   data.A,
				Value: conv,
			}
		case *OpConstWide:
			data := (ot.Fmt()).(Fmt51l)
			conv := int64(data.B)

			parsed = &ConstOpNode{
				Raw:   ot,
				Dst:   data.A,
				Value: conv,
			}
		case *OpConstWideHigh16:
			data := (ot.Fmt()).(Fmt21h)
			conv := int64(uint64(data.B) << 48)

			parsed = &ConstOpNode{
				Raw:   ot,
				Dst:   data.A,
				Value: conv,
			}
		case *OpConstString:
			data := (ot.Fmt()).(Fmt21c)
			val, err := r.ReadString(uint32(data.B))
			if err != nil {
				return res, err
			}

			parsed = &ConstOpNode{
				Raw:   ot,
				Dst:   data.A,
				Value: val,
			}
		case *OpConstStringJumbo:
			data := (ot.Fmt()).(Fmt31c)
			val, err := r.ReadString(data.B)
			if err != nil {
				return res, err
			}

			parsed = &ConstOpNode{
				Raw:   ot,
				Dst:   data.A,
				Value: val,
			}
		case *OpConstClass:
			data := (ot.Fmt()).(Fmt21c)
			val, err := r.ReadTypeAndParse(uint32(data.B))
			if err != nil {
				return res, err
			}

			parsed = &ConstOpNode{
				Raw:   ot,
				Dst:   data.A,
				Value: val,
			}
		case *OpMonitorEnter:
			data := (ot.Fmt()).(Fmt11x)
			parsed = &MonitorOpNode{
				Raw:   ot,
				Ref:   data.A,
				Enter: true,
			}
		case *OpMonitorExit:
			data := (ot.Fmt()).(Fmt11x)
			parsed = &MonitorOpNode{
				Raw:   ot,
				Ref:   data.A,
				Enter: false,
			}
		case *OpCheckCast:
			data := (ot.Fmt()).(Fmt21c)
			val, err := r.ReadTypeAndParse(uint32(data.B))
			if err != nil {
				return res, err
			}

			parsed = &CheckCastOpNode{
				Raw:  ot,
				Ref:  data.A,
				Type: val,
			}
		case *OpNewInstance:
			data := (ot.Fmt()).(Fmt21c)

			parsedDesc, err := r.ReadTypeAndParse(uint32(data.B))
			if err != nil {
				return res, fmt.Errorf("bad class type: %w", err)
			}

			if !parsedDesc.IsClass() || parsedDesc.IsArray() {
				return res, fmt.Errorf("%w: invalid descriptor", ErrBadOp)
			}

			parsed = &NewInstanceOpNode{
				Raw:  ot,
				Dst:  data.A,
				Type: parsedDesc.ClassName,
			}
		case *OpGoto:
			data := (ot.Fmt()).(Fmt10t)

			tgt := pos + int(int8(data.A))
			if tgt < 0 || tgt >= len(idMap) {
				return res, fmt.Errorf("%w: invalid jump offset %v", ErrBadOp, tgt)
			}

			id := idMap[tgt]
			if id == 0 && tgt != 0 {
				return res, fmt.Errorf("%w: invalid jump offset %v", ErrBadOp, tgt)
			}

			parsed = &GotoOpNode{
				Raw: ot,
				Tgt: int(id),
			}
		case *OpGoto16:
			data := (ot.Fmt()).(Fmt20t)

			tgt := pos + int(int16(data.A))
			if tgt < 0 || tgt >= len(idMap) {
				return res, fmt.Errorf("%w: invalid jump offset %v", ErrBadOp, tgt)
			}

			id := idMap[tgt]
			if id == 0 && tgt != 0 {
				return res, fmt.Errorf("%w: invalid jump offset %v", ErrBadOp, tgt)
			}

			parsed = &GotoOpNode{
				Raw: ot,
				Tgt: int(id),
			}
		case *OpGoto32:
			data := (ot.Fmt()).(Fmt30t)

			tgt := pos + int(int32(data.A))
			if tgt < 0 || tgt >= len(idMap) {
				return res, fmt.Errorf("%w: invalid jump offset %v", ErrBadOp, tgt)
			}

			id := idMap[tgt]
			if id == 0 && tgt != 0 {
				return res, fmt.Errorf("%w: invalid jump offset %v", ErrBadOp, tgt)
			}

			parsed = &GotoOpNode{
				Raw: ot,
				Tgt: int(id),
			}
		case *OpIfEq, *OpIfNe, *OpIfLt, *OpIfGe, *OpIfGt, *OpIfLe:
			data := (ot.Fmt()).(Fmt22t)

			tgt := pos + int(int16(data.C))
			if tgt < 0 || tgt >= len(idMap) {
				return res, fmt.Errorf("%w: invalid jump offset %v", ErrBadOp, tgt)
			}

			id := idMap[tgt]
			if id == 0 && tgt != 0 {
				return res, fmt.Errorf("%w: invalid jump offset %v", ErrBadOp, tgt)
			}

			parsed = &IfTestOpNode{
				Raw: ot,
				A:   data.A,
				B:   int16(data.B),
				Tgt: int(id),
			}
		case *OpIfEqz, *OpIfNez, *OpIfLtz, *OpIfGez, *OpIfGtz, *OpIfLez:
			data := (ot.Fmt()).(Fmt21t)

			tgt := pos + int(int16(data.B))
			if tgt < 0 || tgt >= len(idMap) {
				return res, fmt.Errorf("%w: invalid jump offset %v", ErrBadOp, tgt)
			}

			id := idMap[tgt]
			if id == 0 && tgt != 0 {
				return res, fmt.Errorf("%w: invalid jump offset %v", ErrBadOp, tgt)
			}

			parsed = &IfTestOpNode{
				Raw: ot,
				A:   data.A,
				B:   -1,
				Tgt: int(id),
			}
		case *OpIget, *OpIgetWide, *OpIgetObject, *OpIgetBoolean, *OpIgetByte, *OpIgetChar, *OpIgetShort:
			data := (ot.Fmt()).(Fmt22c)

			f, err := r.ReadFieldAndParse(uint32(data.C))
			if err != nil {
				return res, err
			}

			parsed = &IGetOpNode{
				Raw:   ot,
				Dst:   data.A,
				Obj:   data.B,
				Field: f,
			}
		case *OpIput, *OpIputWide, *OpIputObject, *OpIputBoolean, *OpIputByte, *OpIputChar, *OpIputShort:
			data := (ot.Fmt()).(Fmt22c)

			f, err := r.ReadFieldAndParse(uint32(data.C))
			if err != nil {
				return res, err
			}

			parsed = &IPutOpNode{
				Raw:   ot,
				Src:   data.A,
				Obj:   data.B,
				Field: f,
			}
		case *OpSget, *OpSgetWide, *OpSgetObject, *OpSgetBoolean, *OpSgetByte, *OpSgetChar, *OpSgetShort:
			data := (ot.Fmt()).(Fmt21c)

			f, err := r.ReadFieldAndParse(uint32(data.B))
			if err != nil {
				return res, err
			}

			parsed = &SGetOpNode{
				Raw:   ot,
				Dst:   data.A,
				Field: f,
			}
		case *OpSput, *OpSputWide, *OpSputObject, *OpSputBoolean, *OpSputByte, *OpSputChar, *OpSputShort:
			data := (ot.Fmt()).(Fmt21c)

			f, err := r.ReadFieldAndParse(uint32(data.B))
			if err != nil {
				return res, err
			}

			parsed = &SPutOpNode{
				Raw:   ot,
				Src:   data.A,
				Field: f,
			}
		case *OpInvokeVirtual, *OpInvokeSuper, *OpInvokeDirect, *OpInvokeStatic, *OpInvokeInterface:
			data := (ot.Fmt()).(Fmt35c)

			m, err := r.ReadMethodAndParse(uint32(data.B))
			if err != nil {
				return res, err
			}

			args := make([]uint8, data.A)

			if data.A >= 1 {
				args[0] = data.C
			}

			if data.A >= 2 {
				args[1] = data.D
			}

			if data.A >= 3 {
				args[2] = data.E
			}

			if data.A >= 4 {
				args[3] = data.F
			}

			if data.A >= 5 {
				args[4] = data.G
			}

			parsed = &InvokeOpNode{
				Raw:    ot,
				Method: m,
				Args:   args,
			}

		default:
			log.Panic(o)
		}

		res.Ops[idPos] = parsed
		idPos++
	}

	return res, nil
}
