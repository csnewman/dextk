package dextk

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrEmptyTypeDesc = errors.New("empty type desc")
	ErrBadTypeDesc   = errors.New("malformed type desc")
)

func (r *Reader) ReadTypeAndParse(id uint32) (TypeDescriptor, error) {
	var res TypeDescriptor

	typeDef, err := r.ReadType(id)
	if err != nil {
		return res, err
	}

	typeDesc, err := r.ReadString(typeDef.DescriptorStringID)
	if err != nil {
		return res, err
	}

	return ParseTypeDescriptor(typeDesc)
}

type TypeDescriptor struct {
	Type        uint8
	ArrayLength int
	ClassName   String
}

func ParseTypeDescriptor(value String) (TypeDescriptor, error) {
	var res TypeDescriptor

	l := len(value.Raw)

	if l == 0 {
		return res, ErrEmptyTypeDesc
	}

	for value.Raw[res.ArrayLength] == '[' {
		res.ArrayLength++

		// Ensure there is a next character
		if res.ArrayLength == l {
			return res, fmt.Errorf("%w: %v", ErrBadTypeDesc, value)
		}
	}

	res.Type = uint8(value.Raw[res.ArrayLength])

	// Check if a string
	if res.Type != 'L' {
		// Only a single character should appear if not a string
		if res.ArrayLength+1 != l {
			return res, fmt.Errorf("%w: %v", ErrBadTypeDesc, value)
		}

		return res, nil
	}

	if value.Raw[l-1] != ';' {
		return res, fmt.Errorf("%w: %v", ErrBadTypeDesc, value)
	}

	if l-2-res.ArrayLength <= 0 {
		return res, fmt.Errorf("%w: %v", ErrBadTypeDesc, value)
	}

	res.ClassName = StringFromUTF16(value.Raw[1+res.ArrayLength : l-1])

	return res, nil
}

func (d TypeDescriptor) String() string {
	ap := strings.Repeat("[", d.ArrayLength)

	if d.Type == 'L' {
		return fmt.Sprintf("%vL%s;", ap, d.ClassName)
	}

	return fmt.Sprintf("%v%c", ap, d.Type)
}

func (d TypeDescriptor) Base() TypeDescriptor {
	return TypeDescriptor{
		Type:        d.Type,
		ArrayLength: 0,
		ClassName:   d.ClassName,
	}
}

func (d TypeDescriptor) IsArray() bool {
	return d.ArrayLength != 0
}

func (d TypeDescriptor) IsClass() bool {
	return d.Type == 'L'
}
