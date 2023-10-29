package dextk

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
)

var (
	ErrInvalidHeader           = errors.New("invalid header")
	ErrInvalidStringID         = errors.New("invalid string id")
	ErrInvalidTypeID           = errors.New("invalid type id")
	ErrInvalidProtoID          = errors.New("invalid proto id")
	ErrInvalidFieldID          = errors.New("invalid field id")
	ErrInvalidMethodID         = errors.New("invalid method id")
	ErrInvalidClassDefID       = errors.New("invalid class def id")
	ErrInvalidTryHandlerOffset = errors.New("invalid try handler offset")
)

const (
	EndianConst        = uint32(0x12345678)
	ReverseEndianConst = uint32(0x78563412)
	NoIndex            = uint32(0xFFFFFFFF)
)

type Reader struct {
	file          io.ReaderAt
	Version       string
	ReverseEndian bool

	mapOff        uint32
	StringIDCount uint32
	stringIDOff   uint32
	TypeIDCount   uint32
	typeIDOff     uint32
	ProtoIDCount  uint32
	protoIDOff    uint32
	FieldIDCount  uint32
	fieldIDOff    uint32
	MethodIDCount uint32
	methodIDOff   uint32
	ClassDefCount uint32
	classDefOff   uint32
}

func Read(file io.ReaderAt) (*Reader, error) {
	header := make([]byte, 112)

	_, err := file.ReadAt(header[0:40], 0)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrInvalidHeader, err)
	}

	// Validate file magic
	if header[0] != 'd' || header[1] != 'e' || header[2] != 'x' || header[3] != '\n' || header[7] != 0 {
		return nil, fmt.Errorf("%w: invalid magic", ErrInvalidHeader)
	}

	ver := string(header[4:7])

	// TODO: checksum, signature

	headerSize := binary.LittleEndian.Uint32(header[36:40])
	if headerSize < 112 {
		return nil, fmt.Errorf("%w: header too small %d", ErrInvalidHeader, headerSize)
	}

	_, err = file.ReadAt(header[40:112], 40)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrInvalidHeader, err)
	}

	endian := binary.LittleEndian.Uint32(header[40:44])
	reverseEndian := endian == ReverseEndianConst

	r := &Reader{
		file:          file,
		Version:       ver,
		ReverseEndian: reverseEndian,
	}

	r.mapOff = r.toUint(header[52:56])
	r.StringIDCount = r.toUint(header[56:60])
	r.stringIDOff = r.toUint(header[60:64])
	r.TypeIDCount = r.toUint(header[64:68])
	r.typeIDOff = r.toUint(header[68:72])
	r.ProtoIDCount = r.toUint(header[72:76])
	r.protoIDOff = r.toUint(header[76:80])
	r.FieldIDCount = r.toUint(header[80:84])
	r.fieldIDOff = r.toUint(header[84:88])
	r.MethodIDCount = r.toUint(header[88:92])
	r.methodIDOff = r.toUint(header[92:96])
	r.ClassDefCount = r.toUint(header[96:100])
	r.classDefOff = r.toUint(header[100:104])

	return r, nil
}

func (r *Reader) ReadString(id uint32) (string, error) {
	if id >= r.StringIDCount {
		return "", ErrInvalidStringID
	}

	idPos := r.stringIDOff + (id * 4)

	strPos, err := r.readUint(idPos)
	if err != nil {
		return "", err
	}

	strSize, n, err := r.readUleb128(strPos)
	if err != nil {
		return "", err
	}

	data := make([]byte, strSize)

	_, err = r.file.ReadAt(data, int64(strPos+n))
	if err != nil {
		return "", err
	}

	return string(data), nil
}

type Type struct {
	DescriptorStringID uint32
}

func (r *Reader) ReadType(id uint32) (Type, error) {
	var (
		res Type
		err error
	)

	if id >= r.TypeIDCount {
		return res, ErrInvalidTypeID
	}

	entryPos := r.typeIDOff + (id * 4)

	res.DescriptorStringID, err = r.readUint(entryPos)
	if err != nil {
		return res, err
	}

	return res, nil
}

type Proto struct {
	ShortyStringID        uint32
	ReturnTypeID          uint32
	ParametersTypeListOff uint32
}

func (r *Reader) ReadProto(id uint32) (Proto, error) {
	var (
		res Proto
		err error
	)

	if id >= r.ProtoIDCount {
		return res, ErrInvalidProtoID
	}

	entryPos := r.protoIDOff + (id * 12)

	res.ShortyStringID, err = r.readUint(entryPos)
	if err != nil {
		return res, err
	}

	res.ReturnTypeID, err = r.readUint(entryPos + 4)
	if err != nil {
		return res, err
	}

	res.ParametersTypeListOff, err = r.readUint(entryPos + 8)
	if err != nil {
		return res, err
	}

	return res, nil
}

type Field struct {
	ClassTypeID  uint16
	TypeID       uint16
	NameStringID uint32
}

func (r *Reader) ReadField(id uint32) (Field, error) {
	var (
		res Field
		err error
	)

	if id >= r.FieldIDCount {
		return res, ErrInvalidFieldID
	}

	entryPos := r.fieldIDOff + (id * 8)

	res.ClassTypeID, err = r.readUshort(entryPos)
	if err != nil {
		return res, err
	}

	res.TypeID, err = r.readUshort(entryPos + 2)
	if err != nil {
		return res, err
	}

	res.NameStringID, err = r.readUint(entryPos + 4)
	if err != nil {
		return res, err
	}

	return res, nil
}

type Method struct {
	ClassTypeID  uint16
	ProtoID      uint16
	NameStringID uint32
}

func (r *Reader) ReadMethod(id uint32) (Method, error) {
	var (
		res Method
		err error
	)

	if id >= r.MethodIDCount {
		return res, ErrInvalidMethodID
	}

	entryPos := r.methodIDOff + (id * 8)

	res.ClassTypeID, err = r.readUshort(entryPos)
	if err != nil {
		return res, err
	}

	res.ProtoID, err = r.readUshort(entryPos + 2)
	if err != nil {
		return res, err
	}

	res.NameStringID, err = r.readUint(entryPos + 4)
	if err != nil {
		return res, err
	}

	return res, nil
}

type ClassDef struct {
	ClassTypeID             uint32
	AccessFlags             uint32
	SuperclassTypeID        uint32
	InterfacesTypeListOff   uint32
	SourceFileStringID      uint32
	AnnotationsDirectoryOff uint32
	ClassDataOff            uint32
	StaticValuesOff         uint32
}

func (r *Reader) ReadClassDef(id uint32) (ClassDef, error) {
	var (
		res ClassDef
		err error
	)

	if id >= r.ClassDefCount {
		return res, ErrInvalidClassDefID
	}

	entryPos := r.classDefOff + (id * 32)

	res.ClassTypeID, err = r.readUint(entryPos)
	if err != nil {
		return res, err
	}

	res.AccessFlags, err = r.readUint(entryPos + 4)
	if err != nil {
		return res, err
	}

	res.SuperclassTypeID, err = r.readUint(entryPos + 8)
	if err != nil {
		return res, err
	}

	res.InterfacesTypeListOff, err = r.readUint(entryPos + 12)
	if err != nil {
		return res, err
	}

	res.SourceFileStringID, err = r.readUint(entryPos + 16)
	if err != nil {
		return res, err
	}

	res.AnnotationsDirectoryOff, err = r.readUint(entryPos + 20)
	if err != nil {
		return res, err
	}

	res.ClassDataOff, err = r.readUint(entryPos + 24)
	if err != nil {
		return res, err
	}

	res.StaticValuesOff, err = r.readUint(entryPos + 28)
	if err != nil {
		return res, err
	}

	return res, nil
}

type EncodedField struct {
	FieldID     uint32
	AccessFlags uint32
}

type EncodedMethod struct {
	MethodID    uint32
	AccessFlags uint32
	CodeOff     uint32
}

type ClassData struct {
	StaticFields   []EncodedField
	InstanceFields []EncodedField
	DirectMethods  []EncodedMethod
	VirtualMethods []EncodedMethod
}

func (r *Reader) ReadClassData(off uint32) (ClassData, error) {
	if off == 0 {
		panic("invalid class data offset")
	}

	var (
		res ClassData
		err error
	)

	staticCount, n, err := r.readUleb128(off)
	if err != nil {
		return res, err
	}

	res.StaticFields = make([]EncodedField, staticCount)
	off += n

	instanceCount, n, err := r.readUleb128(off)
	if err != nil {
		return res, err
	}

	res.InstanceFields = make([]EncodedField, instanceCount)
	off += n

	directCount, n, err := r.readUleb128(off)
	if err != nil {
		return res, err
	}

	res.DirectMethods = make([]EncodedMethod, directCount)
	off += n

	virtualCount, n, err := r.readUleb128(off)
	if err != nil {
		return res, err
	}

	res.VirtualMethods = make([]EncodedMethod, virtualCount)
	off += n

	lastFieldId := uint32(0)

	for i := uint32(0); i < staticCount; i++ {
		idOff, n, err := r.readUleb128(off)
		if err != nil {
			return res, err
		}

		lastFieldId += idOff
		off += n

		flags, n, err := r.readUleb128(off)
		if err != nil {
			return res, err
		}

		off += n

		res.StaticFields[i].FieldID = lastFieldId
		res.StaticFields[i].AccessFlags = flags
	}

	lastFieldId = uint32(0)

	for i := uint32(0); i < instanceCount; i++ {
		idOff, n, err := r.readUleb128(off)
		if err != nil {
			return res, err
		}

		lastFieldId += idOff
		off += n

		flags, n, err := r.readUleb128(off)
		if err != nil {
			return res, err
		}

		off += n

		res.InstanceFields[i].FieldID = lastFieldId
		res.InstanceFields[i].AccessFlags = flags
	}

	lastMethodId := uint32(0)

	for i := uint32(0); i < directCount; i++ {
		idOff, n, err := r.readUleb128(off)
		if err != nil {
			return res, err
		}

		lastMethodId += idOff
		off += n

		flags, n, err := r.readUleb128(off)
		if err != nil {
			return res, err
		}

		off += n

		code, n, err := r.readUleb128(off)
		if err != nil {
			return res, err
		}

		off += n

		res.DirectMethods[i].MethodID = lastMethodId
		res.DirectMethods[i].AccessFlags = flags
		res.DirectMethods[i].CodeOff = code
	}

	lastMethodId = uint32(0)

	for i := uint32(0); i < virtualCount; i++ {
		idOff, n, err := r.readUleb128(off)
		if err != nil {
			return res, err
		}

		lastMethodId += idOff
		off += n

		flags, n, err := r.readUleb128(off)
		if err != nil {
			return res, err
		}

		off += n

		code, n, err := r.readUleb128(off)
		if err != nil {
			return res, err
		}

		off += n

		res.VirtualMethods[i].MethodID = lastMethodId
		res.VirtualMethods[i].AccessFlags = flags
		res.VirtualMethods[i].CodeOff = code
	}

	return res, nil
}

type TypeList struct {
	TypeIds []uint16
}

func (r *Reader) ReadTypeList(off uint32) (TypeList, error) {
	if off == 0 {
		panic("invalid type list offset")
	}

	var (
		res TypeList
		err error
	)

	size, err := r.readUint(off)
	if err != nil {
		return res, err
	}

	off += 4
	res.TypeIds = make([]uint16, size)

	for i := uint32(0); i < size; i++ {
		id, err := r.readUshort(off)
		if err != nil {
			return res, err
		}

		off += 2

		res.TypeIds[i] = id
	}

	return res, nil
}

type Code struct {
	RegisterCount uint16
	IncomingCount uint16
	OutgoingCount uint16
	DebugInfoOff  uint32
	Insns         []uint16
	Tries         []Try
	Handlers      []CatchHandler
}

type Try struct {
	StartAddr uint32
	InsnCount uint16
	HandlerId uint16
}

type CatchHandler struct {
	Handlers     []HandlerPair
	HasCatchAll  bool
	CatchAllAddr uint32
}

type HandlerPair struct {
	TypeID uint32
	Addr   uint32
}

func (r *Reader) ReadCode(off uint32) (Code, error) {
	if off == 0 {
		panic("invalid code offset")
	}

	var (
		res Code
		err error
	)

	res.RegisterCount, err = r.readUshort(off)
	if err != nil {
		return res, err
	}

	res.IncomingCount, err = r.readUshort(off + 2)
	if err != nil {
		return res, err
	}

	res.OutgoingCount, err = r.readUshort(off + 4)
	if err != nil {
		return res, err
	}

	triesCount, err := r.readUshort(off + 6)
	if err != nil {
		return res, err
	}

	res.DebugInfoOff, err = r.readUint(off + 8)
	if err != nil {
		return res, err
	}

	insCount, err := r.readUint(off + 12)
	if err != nil {
		return res, err
	}

	off += 16

	res.Insns = make([]uint16, insCount)

	for i := uint32(0); i < insCount; i++ {
		res.Insns[i], err = r.readUshort(off)
		if err != nil {
			return res, err
		}

		off += 2
	}

	if triesCount == 0 {
		return res, err
	}

	// Padding
	if triesCount != 0 && insCount%2 == 1 {
		off += 2
	}

	// Skip ahead and read handlers
	handlerMap := make(map[uint16]uint16)
	handOff := off + (uint32(triesCount) * 8)
	handOffStart := handOff

	handlerCount, n, err := r.readUleb128(handOff)
	if err != nil {
		return res, err
	}

	handOff += n
	res.Handlers = make([]CatchHandler, handlerCount)

	for i := uint16(0); i < uint16(handlerCount); i++ {
		handlerMap[uint16(handOff-handOffStart)] = i

		size, n, err := r.readSleb128(handOff)
		if err != nil {
			return res, err
		}

		handOff += n

		if size <= 0 {
			res.Handlers[i].HasCatchAll = true
			size = -size
		}

		res.Handlers[i].Handlers = make([]HandlerPair, size)

		for j := uint16(0); j < uint16(size); j++ {
			typeIdx, n, err := r.readUleb128(handOff)
			if err != nil {
				return res, err
			}
			handOff += n

			addr, n, err := r.readUleb128(handOff)
			if err != nil {
				return res, err
			}
			handOff += n

			res.Handlers[i].Handlers[j] = HandlerPair{
				TypeID: typeIdx,
				Addr:   addr,
			}
		}

		if res.Handlers[i].HasCatchAll {
			res.Handlers[i].CatchAllAddr, n, err = r.readUleb128(handOff)
			if err != nil {
				return res, err
			}
			handOff += n
		}
	}

	// Jump back and read try blocks
	res.Tries = make([]Try, triesCount)

	for i := uint16(0); i < triesCount; i++ {
		res.Tries[i].StartAddr, err = r.readUint(off)
		if err != nil {
			return res, err
		}

		res.Tries[i].InsnCount, err = r.readUshort(off + 4)
		if err != nil {
			return res, err
		}

		handlerOff, err := r.readUshort(off + 6)
		if err != nil {
			return res, err
		}

		handlerId, found := handlerMap[handlerOff]

		if !found {
			return res, ErrInvalidTryHandlerOffset
		}

		res.Tries[i].HandlerId = handlerId

		off += 8
	}

	return res, nil
}
