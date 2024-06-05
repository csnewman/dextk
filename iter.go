package dextk

import (
	"errors"
	"fmt"
)

var (
	ErrIterEnd  = errors.New("iter end reached")
	ErrBadClass = errors.New("malformed class def")
)

type StringIter struct {
	r   *Reader
	pos uint32
}

type StringNode struct {
	Id    uint32
	Value String
}

func (r *Reader) StringIter() *StringIter {
	return &StringIter{
		r:   r,
		pos: 0,
	}
}

func (i *StringIter) Seek(pos uint32) {
	i.pos = pos
}

func (i *StringIter) HasNext() bool {
	return i.pos < i.r.StringIDCount
}

func (i *StringIter) Next() (StringNode, error) {
	var res StringNode

	res.Id = i.pos

	if i.pos >= i.r.StringIDCount {
		return res, ErrIterEnd
	}

	value, err := i.r.ReadString(i.pos)
	if err != nil {
		return res, err
	}

	res.Value = value
	i.pos++

	return res, nil
}

type ClassIter struct {
	r   *Reader
	pos uint32
}

type ClassNode struct {
	Id          uint32
	Name        String
	AccessFlags uint32
	SuperClass  String
	Interfaces  []String
	SourceFile  String

	StaticFields   []FieldNode
	InstanceFields []FieldNode
	DirectMethods  []MethodNode
	VirtualMethods []MethodNode
}

type FieldNode struct {
	Id          uint32
	AccessFlags uint32
	Type        TypeDescriptor
	Name        String
}

type MethodNode struct {
	Id          uint32
	AccessFlags uint32
	Name        String
	Shorty      String
	ReturnType  TypeDescriptor
	Params      []TypeDescriptor
	CodeOff     uint32
}

func (r *Reader) ClassIter() *ClassIter {
	return &ClassIter{
		r:   r,
		pos: 0,
	}
}

func (i *ClassIter) Seek(pos uint32) {
	i.pos = pos
}

func (i *ClassIter) HasNext() bool {
	return i.pos < i.r.ClassDefCount
}

func (i *ClassIter) Next() (ClassNode, error) {
	if i.pos >= i.r.ClassDefCount {
		return ClassNode{}, ErrIterEnd
	}

	res, err := i.r.ReadClassAndParse(i.pos)
	if err != nil {
		return res, err
	}

	i.pos++

	return res, nil
}

func (r *Reader) ReadClassAndParse(id uint32) (ClassNode, error) {
	var res ClassNode

	if id >= r.ClassDefCount {
		return res, ErrInvalidClassDefID
	}

	res.Id = id

	def, err := r.ReadClassDef(id)
	if err != nil {
		return res, fmt.Errorf("%w: %w", ErrBadClass, err)
	}

	res.AccessFlags = def.AccessFlags

	// Read class name
	{
		parsedDesc, err := r.ReadTypeAndParse(def.ClassTypeID)
		if err != nil {
			return res, fmt.Errorf("%w: bad class type: %w", ErrBadClass, err)
		}

		if !parsedDesc.IsClass() || parsedDesc.IsArray() {
			return res, fmt.Errorf("%w: invalid descriptor", ErrBadClass)
		}

		res.Name = parsedDesc.ClassName
	}

	// Parse superclass
	if def.SuperclassTypeID != NoIndex {
		parsedDesc, err := r.ReadTypeAndParse(def.SuperclassTypeID)
		if err != nil {
			return res, fmt.Errorf("%w: bad superclass type: %w", ErrBadClass, err)
		}

		if !parsedDesc.IsClass() || parsedDesc.IsArray() {
			return res, fmt.Errorf("%w: invalid superclass descriptor", ErrBadClass)
		}

		res.SuperClass = parsedDesc.ClassName
	}

	// Parse interfaces
	if def.InterfacesTypeListOff != 0 {
		list, err := r.ReadTypeList(def.InterfacesTypeListOff)
		if err != nil {
			return res, fmt.Errorf("%w: bad interface list: %w", ErrBadClass, err)
		}

		res.Interfaces = make([]String, len(list.TypeIds))

		for p, id := range list.TypeIds {
			parsedDesc, err := r.ReadTypeAndParse(uint32(id))
			if err != nil {
				return res, fmt.Errorf("%w: bad iface type: %w", ErrBadClass, err)
			}

			if !parsedDesc.IsClass() || parsedDesc.IsArray() {
				return res, fmt.Errorf("%w: invalid iface descriptor", ErrBadClass)
			}

			res.Interfaces[p] = parsedDesc.ClassName
		}
	}

	// Parse source file
	if def.SourceFileStringID != NoIndex {
		res.SourceFile, err = r.ReadString(def.SourceFileStringID)
		if err != nil {
			return res, fmt.Errorf("%w: bad src file: %w", ErrBadClass, err)
		}
	}

	// Parse annotations
	if def.AnnotationsDirectoryOff != 0 {
		// TODO: Parse annotations
	}

	// Parse data
	if def.ClassDataOff != 0 {
		data, err := r.ReadClassData(def.ClassDataOff)
		if err != nil {
			return res, fmt.Errorf("%w: bad data: %w", ErrBadClass, err)
		}

		res.StaticFields, err = r.parseFields(def.ClassTypeID, data.StaticFields)
		if err != nil {
			return res, err
		}

		res.InstanceFields, err = r.parseFields(def.ClassTypeID, data.InstanceFields)
		if err != nil {
			return res, err
		}

		res.DirectMethods, err = r.parseMethods(def.ClassTypeID, data.DirectMethods)
		if err != nil {
			return res, err
		}

		res.VirtualMethods, err = r.parseMethods(def.ClassTypeID, data.VirtualMethods)
		if err != nil {
			return res, err
		}
	}

	return res, nil
}

func (r *Reader) parseFields(classTypeId uint32, input []EncodedField) ([]FieldNode, error) {
	res := make([]FieldNode, len(input))

	for i, f := range input {
		fp, err := r.ReadField(f.FieldID)
		if err != nil {
			return res, fmt.Errorf("%w: bad field: %w", ErrBadClass, err)
		}

		if classTypeId != uint32(fp.ClassTypeID) {
			return res, fmt.Errorf("%w: bad field: mismatch class id", ErrBadClass)
		}

		name, err := r.ReadString(fp.NameStringID)
		if err != nil {
			return res, fmt.Errorf("%w: bad field name: %w", ErrBadClass, err)
		}

		def, err := r.ReadTypeAndParse(uint32(fp.TypeID))
		if err != nil {
			return res, fmt.Errorf("%w: bad field type: %w", ErrBadClass, err)
		}

		res[i] = FieldNode{
			Id:          f.FieldID,
			AccessFlags: f.AccessFlags,
			Type:        def,
			Name:        name,
		}
	}

	return res, nil
}

func (r *Reader) parseMethods(classTypeId uint32, input []EncodedMethod) ([]MethodNode, error) {
	res := make([]MethodNode, len(input))

	for i, m := range input {
		mp, err := r.ReadMethod(m.MethodID)
		if err != nil {
			return res, fmt.Errorf("%w: bad method: %w", ErrBadClass, err)
		}

		if classTypeId != uint32(mp.ClassTypeID) {
			return res, fmt.Errorf("%w: bad method: mismatch class id", ErrBadClass)
		}

		name, err := r.ReadString(mp.NameStringID)
		if err != nil {
			return res, fmt.Errorf("%w: bad method name: %w", ErrBadClass, err)
		}

		proto, err := r.ReadProto(uint32(mp.ProtoID))
		if err != nil {
			return res, fmt.Errorf("%w: bad method proto: %w", ErrBadClass, err)
		}

		shorty, err := r.ReadString(proto.ShortyStringID)
		if err != nil {
			return res, fmt.Errorf("%w: bad method shorty: %w", ErrBadClass, err)
		}

		retDef, err := r.ReadTypeAndParse(proto.ReturnTypeID)
		if err != nil {
			return res, fmt.Errorf("%w: bad method type: %w", ErrBadClass, err)
		}

		res[i] = MethodNode{
			Id:          m.MethodID,
			AccessFlags: m.AccessFlags,
			Name:        name,
			Shorty:      shorty,
			ReturnType:  retDef,
			CodeOff:     m.CodeOff,
		}

		if proto.ParametersTypeListOff != 0 {
			plist, err := r.ReadTypeList(proto.ParametersTypeListOff)
			if err != nil {
				return res, fmt.Errorf("%w: bad method params: %w", ErrBadClass, err)
			}

			res[i].Params = make([]TypeDescriptor, len(plist.TypeIds))

			for j, tid := range plist.TypeIds {
				res[i].Params[j], err = r.ReadTypeAndParse(uint32(tid))
				if err != nil {
					return res, fmt.Errorf("%w: bad method paramm: %w", ErrBadClass, err)
				}
			}
		}
	}

	return res, nil
}
