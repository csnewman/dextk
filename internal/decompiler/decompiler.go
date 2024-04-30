package decompiler

import (
	"strconv"
	"strings"

	"github.com/csnewman/dextk"
	"github.com/csnewman/dextk/internal/java"
)

func ProcessClass(reader *dextk.Reader, node dextk.ClassNode) string {
	d := &Decompiler{
		reader: reader,
		in:     node,
	}

	d.process()

	w := java.NewWriter()

	d.out.Write(w)

	out := w.String()

	return out
}

type Decompiler struct {
	reader *dextk.Reader
	in     dextk.ClassNode
	out    *java.Class
}

func (d *Decompiler) process() {
	name, pkg := SplitType(d.in.Name)

	d.out = &java.Class{
		Name:    name,
		Package: pkg,
	}

	flags := NewFlags(d.in.AccessFlags)

	if flags.Take(dextk.AccPublic) {
		d.out.Access = java.AccessPublic
	}

	if flags.Take(dextk.AccFinal) {
		d.out.Final = true
	}

	if flags.Take(dextk.AccSuper) {
		panic("todo super")
	}

	if flags.Take(dextk.AccAbstract) {
		panic("todo abstract")
	}

	if flags.Take(dextk.AccSynthetic) {
		d.out.Synthetic = true
	}

	if flags.Take(dextk.AccInterface) {
		panic("todo interface")
	}

	if flags.Take(dextk.AccAnnotation) {
		panic("todo annotation")
	}

	if flags.Take(dextk.AccEnum) {
		panic("todo enum")
	}

	if flags.Remaining() != 0 {
		panic("Flags unparsed: " + strconv.FormatInt(int64(flags.Remaining()), 2))
	}

	for _, field := range d.in.StaticFields {
		d.processField(field, true)
	}

	for _, field := range d.in.InstanceFields {
		d.processField(field, false)
	}
}

func (d *Decompiler) processField(src dextk.FieldNode, static bool) {
	out := &java.Field{
		Name:   src.Name,
		Static: static,
		Type:   ConvertType(src.Type),
	}

	flags := NewFlags(src.AccessFlags)

	if flags.Take(dextk.AccPrivate) {
		out.Access = java.AccessPrivate
	}

	if flags.Take(dextk.AccProtected) {
		out.Access = java.AccessProtected
	}

	if flags.Take(dextk.AccPublic) {
		out.Access = java.AccessPublic
	}

	if flags.Take(dextk.AccFinal) {
		out.Final = true
	}

	if flags.Take(dextk.AccStatic) {
		panic("todo static")
	}

	if flags.Take(dextk.AccVolatile) {
		panic("todo volatile")
	}

	if flags.Take(dextk.AccTransient) {
		panic("todo transient")
	}

	if flags.Take(dextk.AccSynthetic) {
		out.Synthetic = true
	}

	if flags.Take(dextk.AccEnum) {
		panic("todo enum")
	}

	if flags.Remaining() != 0 {
		panic("Flags unparsed: " + strconv.FormatInt(int64(flags.Remaining()), 2))
	}

	d.out.AddField(out)
}

func ConvertType(td dextk.TypeDescriptor) *java.Type {
	switch td.Type {
	case 'L':
		name, pkg := SplitType(td.ClassName)

		return &java.Type{
			Name:        name,
			Package:     pkg,
			ArrayLength: td.ArrayLength,
		}

	default:
		panic("Unknown type: " + strconv.Itoa(int(td.Type)) + " " + string(rune(td.Type)))
	}
}

func SplitType(full string) (string, string) {
	parts := strings.Split(full, "/")

	name := parts[len(parts)-1]
	pkg := strings.Join(parts[:len(parts)-1], ".")

	return name, pkg
}
