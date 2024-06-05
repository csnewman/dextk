package dextk

import (
	"errors"
	"fmt"
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

type InstanceOfOpNode struct {
	Raw  Op
	Src  uint8
	Dst  uint8
	Type TypeDescriptor
}

func (n InstanceOfOpNode) RawOp() Op {
	return n.Raw
}

func (n InstanceOfOpNode) String() string {
	return fmt.Sprintf("%v src=%v dst=%v type=%v", n.Raw.Code(), n.Src, n.Dst, n.Type)
}

type ArrayLenOpNode struct {
	Raw Op
	Ref uint8
	Dst uint8
}

func (n ArrayLenOpNode) RawOp() Op {
	return n.Raw
}

func (n ArrayLenOpNode) String() string {
	return fmt.Sprintf("%v ref=%v dst=%v", n.Raw.Code(), n.Ref, n.Dst)
}

type NewInstanceOpNode struct {
	Raw  Op
	Dst  uint8
	Type String
}

func (n NewInstanceOpNode) RawOp() Op {
	return n.Raw
}

func (n NewInstanceOpNode) String() string {
	return fmt.Sprintf("%v dst=%v type=%v", n.Raw.Code(), n.Dst, n.Type)
}

type NewArrayOpNode struct {
	Raw     Op
	Dst     uint8
	SizeReg uint8
	Type    TypeDescriptor
}

func (n NewArrayOpNode) RawOp() Op {
	return n.Raw
}

func (n NewArrayOpNode) String() string {
	return fmt.Sprintf("%v dst=%v size=%v type=%v", n.Raw.Code(), n.Dst, n.SizeReg, n.Type)
}

type FilledNewArrayOpNode struct {
	Raw       Op
	Type      TypeDescriptor
	ValueRegs []uint16
}

func (n FilledNewArrayOpNode) RawOp() Op {
	return n.Raw
}

func (n FilledNewArrayOpNode) String() string {
	return fmt.Sprintf("%v type=%v values=%v", n.Raw.Code(), n.Type, n.ValueRegs)
}

type FillArrayDataOpNode struct {
	Raw Op
	Ref uint8
	FmtFillArrayDataPayload
}

func (n FillArrayDataOpNode) RawOp() Op {
	return n.Raw
}

func (n FillArrayDataOpNode) String() string {
	return fmt.Sprintf("%v ref=%v %v", n.Raw.Code(), n.Ref, n.FmtFillArrayDataPayload)
}

type ThrowOpNode struct {
	Raw Op
	Reg uint8
}

func (n ThrowOpNode) RawOp() Op {
	return n.Raw
}

func (n ThrowOpNode) String() string {
	return fmt.Sprintf("%v reg=%v", n.Raw.Code(), n.Reg)
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

type PackedSwitchOpNode struct {
	Raw      Op
	Reg      uint8
	FirstKey int32
	Targets  []int
}

func (n PackedSwitchOpNode) RawOp() Op {
	return n.Raw
}

func (n PackedSwitchOpNode) String() string {
	return fmt.Sprintf("%v first=%v tgts=%v", n.Raw.Code(), n.FirstKey, n.Targets)
}

type SparseSwitchOpNode struct {
	Raw     Op
	Reg     uint8
	Keys    []int32
	Targets []int
}

func (n SparseSwitchOpNode) RawOp() Op {
	return n.Raw
}

func (n SparseSwitchOpNode) String() string {
	return fmt.Sprintf("%v keys=%v tgts=%v", n.Raw.Code(), n.Keys, n.Targets)
}

type CmpOpNode struct {
	Raw  Op
	SrcA uint8
	SrcB uint8
	Dst  uint8
}

func (n CmpOpNode) RawOp() Op {
	return n.Raw
}

func (n CmpOpNode) String() string {
	return fmt.Sprintf("%v a=%v b=%v tgt=%v", n.Raw.Code(), n.SrcA, n.SrcB, n.Dst)
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

type AGetOpNode struct {
	Raw   Op
	Dst   uint8
	Array uint8
	Index uint8
}

func (n AGetOpNode) RawOp() Op {
	return n.Raw
}

func (n AGetOpNode) String() string {
	return fmt.Sprintf("%v array=%v index=%v dst=%v", n.Raw.Code(), n.Array, n.Index, n.Dst)
}

type APutOpNode struct {
	Raw   Op
	Src   uint8
	Array uint8
	Index uint8
}

func (n APutOpNode) RawOp() Op {
	return n.Raw
}

func (n APutOpNode) String() string {
	return fmt.Sprintf("%v array=%v index=%v src=%v", n.Raw.Code(), n.Array, n.Index, n.Src)
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
	Args   []uint16
}

func (n InvokeOpNode) RawOp() Op {
	return n.Raw
}

func (n InvokeOpNode) String() string {
	return fmt.Sprintf("%v method=%v args=%v", n.Raw.Code(), n.Method, n.Args)
}

type UnaryOpNode struct {
	Raw Op
	Src uint8
	Dst uint8
}

func (n UnaryOpNode) RawOp() Op {
	return n.Raw
}

func (n UnaryOpNode) String() string {
	return fmt.Sprintf("%v src=%v dst=%v", n.Raw.Code(), n.Src, n.Dst)
}

type BinaryOpNode struct {
	Raw  Op
	SrcA uint8
	SrcB uint8
	Dst  uint8
}

func (n BinaryOpNode) RawOp() Op {
	return n.Raw
}

func (n BinaryOpNode) String() string {
	return fmt.Sprintf("%v srca=%v srcb=%v dst=%v", n.Raw.Code(), n.SrcA, n.SrcB, n.Dst)
}

type BinaryLiteralOpNode struct {
	Raw    Op
	SrcA   uint8
	ValueB int16
	Dst    uint8
}

func (n BinaryLiteralOpNode) RawOp() Op {
	return n.Raw
}

func (n BinaryLiteralOpNode) String() string {
	return fmt.Sprintf("%v srca=%v valueb=%v dst=%v", n.Raw.Code(), n.SrcA, n.ValueB, n.Dst)
}

type UnknownOpNode struct {
	Raw Op
}

func (n UnknownOpNode) RawOp() Op {
	return n.Raw
}

func (n UnknownOpNode) String() string {
	return fmt.Sprintf("%v (unknown)", n.Raw.Code())
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
		return res, err
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
		case OpNop, OpPseudoPackedSwitchPayload, OpPseudoSparseSwitchPayload, OpPseudoFillArrayDataPayload:
			parsed = NopOpNode{
				Raw: ot,
			}
		case OpMove, OpMoveWide, OpMoveObject:
			data := (ot.Fmt()).(Fmt12x)

			parsed = MoveOpNode{
				Raw: ot,
				Src: uint16(data.B),
				Dst: uint16(data.A),
			}
		case OpMoveFrom16, OpMoveWideFrom16, OpMoveObjectFrom16:
			data := (ot.Fmt()).(Fmt22x)

			parsed = MoveOpNode{
				Raw: ot,
				Src: data.B,
				Dst: uint16(data.A),
			}
		case OpMove16, OpMoveWide16, OpMoveObject16:
			data := (ot.Fmt()).(Fmt32x)

			parsed = MoveOpNode{
				Raw: ot,
				Src: data.B,
				Dst: data.A,
			}
		case OpMoveResult, OpMoveResultWide, OpMoveResultObject:
			data := (ot.Fmt()).(Fmt11x)

			parsed = MoveResultOpNode{
				Raw: ot,
				Dst: data.A,
			}
		case OpMoveException:
			data := (ot.Fmt()).(Fmt11x)

			parsed = MoveExceptionOpNode{
				Raw: ot,
				Dst: data.A,
			}
		case OpReturnVoid:
			parsed = ReturnOpNode{
				Raw:   ot,
				Value: -1,
			}
		case OpReturn, OpReturnWide, OpReturnObject:
			data := (ot.Fmt()).(Fmt11x)

			parsed = ReturnOpNode{
				Raw:   ot,
				Value: int16(data.A),
			}
		case OpConst4:
			data := (ot.Fmt()).(Fmt11n)

			val := uint32(data.B)
			if (val & 0b1000) != 0 {
				val |= math.MaxUint32 >> 4 << 4
			}

			conv := int32(val)

			parsed = ConstOpNode{
				Raw:   ot,
				Dst:   data.A,
				Value: conv,
			}
		case OpConst16:
			data := (ot.Fmt()).(Fmt21s)

			val := uint32(data.B)
			if (val & (1 << 15)) != 0 {
				val |= math.MaxUint32 >> 16 << 16
			}

			conv := int32(val)

			parsed = ConstOpNode{
				Raw:   ot,
				Dst:   data.A,
				Value: conv,
			}
		case OpConst:
			data := (ot.Fmt()).(Fmt31i)
			conv := int32(data.B)

			parsed = ConstOpNode{
				Raw:   ot,
				Dst:   data.A,
				Value: conv,
			}
		case OpConstHigh16:
			data := (ot.Fmt()).(Fmt21h)
			conv := int32(uint32(data.B) << 16)

			parsed = ConstOpNode{
				Raw:   ot,
				Dst:   data.A,
				Value: conv,
			}
		case OpConstWide16:
			data := (ot.Fmt()).(Fmt21s)

			val := uint64(data.B)
			if (val & (1 << 15)) != 0 {
				val |= math.MaxUint64 >> 16 << 16
			}

			conv := int64(val)

			parsed = ConstOpNode{
				Raw:   ot,
				Dst:   data.A,
				Value: conv,
			}
		case OpConstWide32:
			data := (ot.Fmt()).(Fmt31i)

			val := uint64(data.B)
			if (val & (1 << 31)) != 0 {
				val |= math.MaxUint64 >> 32 << 32
			}

			conv := int64(val)

			parsed = ConstOpNode{
				Raw:   ot,
				Dst:   data.A,
				Value: conv,
			}
		case OpConstWide:
			data := (ot.Fmt()).(Fmt51l)
			conv := int64(data.B)

			parsed = ConstOpNode{
				Raw:   ot,
				Dst:   data.A,
				Value: conv,
			}
		case OpConstWideHigh16:
			data := (ot.Fmt()).(Fmt21h)
			conv := int64(uint64(data.B) << 48)

			parsed = ConstOpNode{
				Raw:   ot,
				Dst:   data.A,
				Value: conv,
			}
		case OpConstString:
			data := (ot.Fmt()).(Fmt21c)
			val, err := r.ReadString(uint32(data.B))
			if err != nil {
				return res, err
			}

			parsed = ConstOpNode{
				Raw:   ot,
				Dst:   data.A,
				Value: val,
			}
		case OpConstStringJumbo:
			data := (ot.Fmt()).(Fmt31c)
			val, err := r.ReadString(data.B)
			if err != nil {
				return res, err
			}

			parsed = ConstOpNode{
				Raw:   ot,
				Dst:   data.A,
				Value: val,
			}
		case OpConstClass:
			data := (ot.Fmt()).(Fmt21c)
			val, err := r.ReadTypeAndParse(uint32(data.B))
			if err != nil {
				return res, err
			}

			parsed = ConstOpNode{
				Raw:   ot,
				Dst:   data.A,
				Value: val,
			}
		case OpMonitorEnter:
			data := (ot.Fmt()).(Fmt11x)
			parsed = MonitorOpNode{
				Raw:   ot,
				Ref:   data.A,
				Enter: true,
			}
		case OpMonitorExit:
			data := (ot.Fmt()).(Fmt11x)
			parsed = MonitorOpNode{
				Raw:   ot,
				Ref:   data.A,
				Enter: false,
			}
		case OpCheckCast:
			data := (ot.Fmt()).(Fmt21c)
			val, err := r.ReadTypeAndParse(uint32(data.B))
			if err != nil {
				return res, err
			}

			parsed = CheckCastOpNode{
				Raw:  ot,
				Ref:  data.A,
				Type: val,
			}
		case OpInstanceOf:
			data := (ot.Fmt()).(Fmt22c)
			val, err := r.ReadTypeAndParse(uint32(data.C))
			if err != nil {
				return res, err
			}

			parsed = InstanceOfOpNode{
				Raw:  ot,
				Dst:  data.A,
				Src:  data.B,
				Type: val,
			}
		case OpArrayLength:
			data := (ot.Fmt()).(Fmt12x)

			parsed = ArrayLenOpNode{
				Raw: ot,
				Dst: data.A,
				Ref: data.B,
			}
		case OpNewInstance:
			data := (ot.Fmt()).(Fmt21c)

			parsedDesc, err := r.ReadTypeAndParse(uint32(data.B))
			if err != nil {
				return res, fmt.Errorf("bad class type: %w", err)
			}

			if !parsedDesc.IsClass() || parsedDesc.IsArray() {
				return res, fmt.Errorf("%w: invalid descriptor", ErrBadOp)
			}

			parsed = NewInstanceOpNode{
				Raw:  ot,
				Dst:  data.A,
				Type: parsedDesc.ClassName,
			}
		case OpNewArray:
			data := (ot.Fmt()).(Fmt22c)

			parsedDesc, err := r.ReadTypeAndParse(uint32(data.C))
			if err != nil {
				return res, fmt.Errorf("bad class type: %w", err)
			}

			if !parsedDesc.IsArray() {
				return res, fmt.Errorf("%w: %v: invalid descriptor", ErrBadOp, ot)
			}

			parsed = NewArrayOpNode{
				Raw:     ot,
				Dst:     data.A,
				SizeReg: data.B,
				Type:    parsedDesc,
			}
		case OpFilledNewArray:
			data := (ot.Fmt()).(Fmt35c)

			parsedDesc, err := r.ReadTypeAndParse(uint32(data.B))
			if err != nil {
				return res, fmt.Errorf("bad class type: %w", err)
			}

			if !parsedDesc.IsArray() {
				return res, fmt.Errorf("%w: %v: invalid descriptor", ErrBadOp, ot)
			}

			values := make([]uint16, data.A)

			if data.A >= 1 {
				values[0] = uint16(data.C)
			}

			if data.A >= 2 {
				values[1] = uint16(data.D)
			}

			if data.A >= 3 {
				values[2] = uint16(data.E)
			}

			if data.A >= 4 {
				values[3] = uint16(data.F)
			}

			if data.A >= 5 {
				values[4] = uint16(data.G)
			}

			parsed = FilledNewArrayOpNode{
				Raw:       ot,
				Type:      parsedDesc,
				ValueRegs: values,
			}
		case OpFilledNewArrayRange:
			data := (ot.Fmt()).(Fmt3rc)

			parsedDesc, err := r.ReadTypeAndParse(uint32(data.B))
			if err != nil {
				return res, fmt.Errorf("bad class type: %w", err)
			}

			if !parsedDesc.IsArray() {
				return res, fmt.Errorf("%w: %v: invalid descriptor", ErrBadOp, ot)
			}

			values := make([]uint16, data.A)

			for i := uint8(0); i < data.A; i++ {
				values[i] = data.C + uint16(i)
			}

			parsed = FilledNewArrayOpNode{
				Raw:       ot,
				Type:      parsedDesc,
				ValueRegs: values,
			}
		case OpFillArrayData:
			data := (ot.Fmt()).(Fmt31t)

			tgt := pos + int(int32(data.B))
			if tgt < 0 || tgt >= len(idMap) {
				return res, fmt.Errorf("%w: invalid data offset %v", ErrBadOp, tgt)
			}

			oldPos := or.Pos()
			or.Seek(tgt)

			rtop, err := or.Read()
			if err != nil {
				return res, fmt.Errorf("%w: invalid data %w", ErrBadOp, err)
			}

			tableOp, ok := (rtop).(OpPseudoFillArrayDataPayload)
			if !ok {
				return res, fmt.Errorf("%w: invalid data %v", ErrBadOp, tableOp)
			}

			or.Seek(oldPos)

			parsed = FillArrayDataOpNode{
				Raw:                     ot,
				Ref:                     data.A,
				FmtFillArrayDataPayload: tableOp.FmtFillArrayDataPayload,
			}
		case OpThrow:
			data := (ot.Fmt()).(Fmt11x)

			parsed = ThrowOpNode{
				Raw: ot,
				Reg: data.A,
			}
		case OpGoto:
			data := (ot.Fmt()).(Fmt10t)

			tgt := pos + int(int8(data.A))
			if tgt < 0 || tgt >= len(idMap) {
				return res, fmt.Errorf("%w: invalid goto jump offset %v", ErrBadOp, tgt)
			}

			id := idMap[tgt]
			if id == 0 && tgt != 0 {
				return res, fmt.Errorf("%w: invalid goto jump offset %v", ErrBadOp, tgt)
			}

			parsed = GotoOpNode{
				Raw: ot,
				Tgt: int(id),
			}
		case OpGoto16:
			data := (ot.Fmt()).(Fmt20t)

			tgt := pos + int(int16(data.A))
			if tgt < 0 || tgt >= len(idMap) {
				return res, fmt.Errorf("%w: invalid goto16 jump offset %v", ErrBadOp, tgt)
			}

			id := idMap[tgt]
			if id == 0 && tgt != 0 {
				return res, fmt.Errorf("%w: invalid goto16 jump offset %v", ErrBadOp, tgt)
			}

			parsed = GotoOpNode{
				Raw: ot,
				Tgt: int(id),
			}
		case OpGoto32:
			data := (ot.Fmt()).(Fmt30t)

			tgt := pos + int(int32(data.A))
			if tgt < 0 || tgt >= len(idMap) {
				return res, fmt.Errorf("%w: invalid goto32 jump offset %v", ErrBadOp, tgt)
			}

			id := idMap[tgt]
			if id == 0 && tgt != 0 {
				return res, fmt.Errorf("%w: invalid goto32 jump offset %v", ErrBadOp, tgt)
			}

			parsed = GotoOpNode{
				Raw: ot,
				Tgt: int(id),
			}
		case OpPackedSwitch:
			data := (ot.Fmt()).(Fmt31t)

			tgt := pos + int(int32(data.B))
			if tgt < 0 || tgt >= len(idMap) {
				return res, fmt.Errorf("%w: invalid table offset %v", ErrBadOp, tgt)
			}

			oldPos := or.Pos()
			or.Seek(tgt)

			rtop, err := or.Read()
			if err != nil {
				return res, fmt.Errorf("%w: invalid table %w", ErrBadOp, err)
			}

			tableOp, ok := (rtop).(OpPseudoPackedSwitchPayload)
			if !ok {
				return res, fmt.Errorf("%w: invalid table %v", ErrBadOp, tableOp)
			}

			or.Seek(oldPos)

			tgts := make([]int, len(tableOp.Targets))

			for i, target := range tableOp.Targets {
				tgts[i] = pos + int(target)
			}

			parsed = PackedSwitchOpNode{
				Raw:      ot,
				Reg:      data.A,
				FirstKey: tableOp.FirstKey,
				Targets:  tgts,
			}
		case OpSparseSwitch:
			data := (ot.Fmt()).(Fmt31t)

			tgt := pos + int(int32(data.B))
			if tgt < 0 || tgt >= len(idMap) {
				return res, fmt.Errorf("%w: invalid table offset %v", ErrBadOp, tgt)
			}

			oldPos := or.Pos()
			or.Seek(tgt)

			rtop, err := or.Read()
			if err != nil {
				return res, fmt.Errorf("%w: invalid table %w", ErrBadOp, err)
			}

			tableOp, ok := (rtop).(OpPseudoSparseSwitchPayload)
			if !ok {
				return res, fmt.Errorf("%w: invalid table %v", ErrBadOp, tableOp)
			}

			or.Seek(oldPos)

			tgts := make([]int, len(tableOp.Targets))

			for i, target := range tableOp.Targets {
				tgts[i] = pos + int(target)
			}

			parsed = SparseSwitchOpNode{
				Raw:     ot,
				Reg:     data.A,
				Keys:    tableOp.Keys,
				Targets: tgts,
			}
		case OpCmplFloat, OpCmpgFloat, OpCmplDouble, OpCmpgDouble, OpCmpLong:
			data := (ot.Fmt()).(Fmt23x)

			parsed = CmpOpNode{
				Raw:  ot,
				SrcA: data.B,
				SrcB: data.C,
				Dst:  data.A,
			}
		case OpIfEq, OpIfNe, OpIfLt, OpIfGe, OpIfGt, OpIfLe:
			data := (ot.Fmt()).(Fmt22t)

			tgt := pos + int(int16(data.C))
			if tgt < 0 || tgt >= len(idMap) {
				return res, fmt.Errorf("%w: invalid if jump offset %v", ErrBadOp, tgt)
			}

			id := idMap[tgt]
			if id == 0 && tgt != 0 {
				return res, fmt.Errorf("%w: invalid if jump offset %v", ErrBadOp, tgt)
			}

			parsed = IfTestOpNode{
				Raw: ot,
				A:   data.A,
				B:   int16(data.B),
				Tgt: int(id),
			}
		case OpIfEqz, OpIfNez, OpIfLtz, OpIfGez, OpIfGtz, OpIfLez:
			data := (ot.Fmt()).(Fmt21t)

			tgt := pos + int(int16(data.B))
			if tgt < 0 || tgt >= len(idMap) {
				return res, fmt.Errorf("%w: invalid ifz jump offset %v", ErrBadOp, tgt)
			}

			id := idMap[tgt]
			if id == 0 && tgt != 0 {
				return res, fmt.Errorf("%w: invalid ifz jump offset %v", ErrBadOp, tgt)
			}

			parsed = IfTestOpNode{
				Raw: ot,
				A:   data.A,
				B:   -1,
				Tgt: int(id),
			}
		case OpAget, OpAgetWide, OpAgetObject, OpAgetBoolean, OpAgetByte, OpAgetChar, OpAgetShort:
			data := (ot.Fmt()).(Fmt23x)

			parsed = AGetOpNode{
				Raw:   ot,
				Dst:   data.A,
				Array: data.B,
				Index: data.C,
			}
		case OpAput, OpAputWide, OpAputObject, OpAputBoolean, OpAputByte, OpAputChar, OpAputShort:
			data := (ot.Fmt()).(Fmt23x)

			parsed = APutOpNode{
				Raw:   ot,
				Src:   data.A,
				Array: data.B,
				Index: data.C,
			}
		case OpIget, OpIgetWide, OpIgetObject, OpIgetBoolean, OpIgetByte, OpIgetChar, OpIgetShort:
			data := (ot.Fmt()).(Fmt22c)

			f, err := r.ReadFieldAndParse(uint32(data.C))
			if err != nil {
				return res, err
			}

			parsed = IGetOpNode{
				Raw:   ot,
				Dst:   data.A,
				Obj:   data.B,
				Field: f,
			}
		case OpIput, OpIputWide, OpIputObject, OpIputBoolean, OpIputByte, OpIputChar, OpIputShort:
			data := (ot.Fmt()).(Fmt22c)

			f, err := r.ReadFieldAndParse(uint32(data.C))
			if err != nil {
				return res, err
			}

			parsed = IPutOpNode{
				Raw:   ot,
				Src:   data.A,
				Obj:   data.B,
				Field: f,
			}
		case OpSget, OpSgetWide, OpSgetObject, OpSgetBoolean, OpSgetByte, OpSgetChar, OpSgetShort:
			data := (ot.Fmt()).(Fmt21c)

			f, err := r.ReadFieldAndParse(uint32(data.B))
			if err != nil {
				return res, err
			}

			parsed = SGetOpNode{
				Raw:   ot,
				Dst:   data.A,
				Field: f,
			}
		case OpSput, OpSputWide, OpSputObject, OpSputBoolean, OpSputByte, OpSputChar, OpSputShort:
			data := (ot.Fmt()).(Fmt21c)

			f, err := r.ReadFieldAndParse(uint32(data.B))
			if err != nil {
				return res, err
			}

			parsed = SPutOpNode{
				Raw:   ot,
				Src:   data.A,
				Field: f,
			}
		case OpInvokeVirtual, OpInvokeSuper, OpInvokeDirect, OpInvokeStatic, OpInvokeInterface:
			data := (ot.Fmt()).(Fmt35c)

			m, err := r.ReadMethodAndParse(uint32(data.B))
			if err != nil {
				return res, err
			}

			args := make([]uint16, data.A)

			if data.A >= 1 {
				args[0] = uint16(data.C)
			}

			if data.A >= 2 {
				args[1] = uint16(data.D)
			}

			if data.A >= 3 {
				args[2] = uint16(data.E)
			}

			if data.A >= 4 {
				args[3] = uint16(data.F)
			}

			if data.A >= 5 {
				args[4] = uint16(data.G)
			}

			parsed = InvokeOpNode{
				Raw:    ot,
				Method: m,
				Args:   args,
			}
		case OpInvokeVirtualRange, OpInvokeSuperRange, OpInvokeDirectRange, OpInvokeStaticRange,
			OpInvokeInterfaceRange:
			data := (ot.Fmt()).(Fmt3rc)

			m, err := r.ReadMethodAndParse(uint32(data.B))
			if err != nil {
				return res, err
			}

			args := make([]uint16, data.A)

			for i := uint8(0); i < data.A; i++ {
				args[i] = data.C + uint16(i)
			}

			parsed = InvokeOpNode{
				Raw:    ot,
				Method: m,
				Args:   args,
			}
		case OpNegInt, OpNotInt, OpNegLong, OpNotLong, OpNegFloat, OpNegDouble, OpIntToLong, OpIntToFloat,
			OpIntToDouble, OpLongToInt, OpLongToFloat, OpLongToDouble, OpFloatToInt, OpFloatToLong,
			OpFloatToDouble, OpDoubleToInt, OpDoubleToLong, OpDoubleToFloat, OpIntToByte, OpIntToChar,
			OpIntToShort:
			data := (ot.Fmt()).(Fmt12x)

			parsed = UnaryOpNode{
				Raw: ot,
				Dst: data.A,
				Src: data.B,
			}
		case OpAddInt, OpSubInt, OpMulInt, OpDivInt, OpRemInt, OpAndInt, OpOrInt, OpXorInt, OpShlInt,
			OpShrInt, OpUshrInt, OpAddLong, OpSubLong, OpMulLong, OpDivLong, OpRemLong, OpAndLong, OpOrLong,
			OpXorLong, OpShlLong, OpShrLong, OpUshrLong, OpAddFloat, OpSubFloat, OpMulFloat, OpDivFloat,
			OpRemFloat, OpAddDouble, OpSubDouble, OpMulDouble, OpDivDouble, OpRemDouble:
			data := (ot.Fmt()).(Fmt23x)

			parsed = BinaryOpNode{
				Raw:  ot,
				Dst:  data.A,
				SrcA: data.B,
				SrcB: data.C,
			}
		case OpAddInt2Addr, OpSubInt2Addr, OpMulInt2Addr, OpDivInt2Addr, OpRemInt2Addr, OpAndInt2Addr,
			OpOrInt2Addr, OpXorInt2Addr, OpShlInt2Addr, OpShrInt2Addr, OpUshrInt2Addr, OpAddLong2Addr,
			OpSubLong2Addr, OpMulLong2Addr, OpDivLong2Addr, OpRemLong2Addr, OpAndLong2Addr, OpOrLong2Addr,
			OpXorLong2Addr, OpShlLong2Addr, OpShrLong2Addr, OpUshrLong2Addr, OpAddFloat2Addr, OpSubFloat2Addr,
			OpMulFloat2Addr, OpDivFloat2Addr, OpRemFloat2Addr, OpAddDouble2Addr, OpSubDouble2Addr,
			OpMulDouble2Addr, OpDivDouble2Addr, OpRemDouble2Addr:
			data := (ot.Fmt()).(Fmt12x)

			parsed = BinaryOpNode{
				Raw:  ot,
				Dst:  data.A,
				SrcA: data.A,
				SrcB: data.B,
			}
		case OpAddIntLit16, OpRsubIntLit16, OpMulIntLit16, OpDivIntLit16, OpRemIntLit16, OpAndIntLit16,
			OpOrIntLit16, OpXorIntLit16:
			data := (ot.Fmt()).(Fmt22s)

			parsed = BinaryLiteralOpNode{
				Raw:    ot,
				Dst:    data.A,
				SrcA:   data.B,
				ValueB: int16(data.C),
			}
		case OpAddIntLit8, OpRsubIntLit8, OpMulIntLit8, OpDivIntLit8, OpRemIntLit8, OpAndIntLit8, OpOrIntLit8,
			OpXorIntLit8, OpShlIntLit8, OpShrIntLit8, OpUshrIntLit8:
			data := (ot.Fmt()).(Fmt22b)

			parsed = BinaryLiteralOpNode{
				Raw:    ot,
				Dst:    data.A,
				SrcA:   data.B,
				ValueB: int16(int8(data.C)),
			}
		// TODO: invoke-polymorphic, invoke-polymorphic/range, invoke-custom, invoke-custom/range, const-method-handle,
		//       const-method-type
		default:
			parsed = UnknownOpNode{
				Raw: ot,
			}
		}

		res.Ops[idPos] = parsed
		idPos++
	}

	return res, nil
}
