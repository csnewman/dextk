package java

import "strings"

type Class struct {
	Name      string
	Package   string
	Access    Access
	Final     bool
	Synthetic bool
	Fields    []*Field
	Methods   []*Method
}

func (c *Class) AddField(field *Field) {
	c.Fields = append(c.Fields, field)
}

func (c *Class) Write(w *Writer) {
	if c.Package != "" {
		w.WriteLinef("package %v;", c.Package)
		w.EmptyLine()
	}

	acc := c.Access
	if acc != "" {
		acc += " "
	}

	if c.Final {
		acc += "final "
	}

	if c.Synthetic {
		w.WriteLine("/* synthetic */")
	}

	w.WriteLinef("%vclass %v {", acc, c.Name)
	w.Indent(1)

	for _, field := range c.Fields {
		field.Write(w, c)
	}

	if len(c.Fields) > 0 && len(c.Methods) > 0 {
		w.EmptyLine()
	}

	w.Indent(-1)
	w.WriteLine("}")
}

func (c *Class) GenType(t *Type) string {
	sb := &strings.Builder{}

	if t.Package != "" {
		sb.WriteString(t.Package)
		sb.WriteString(".")
	}

	sb.WriteString(t.Name)

	for i := 0; i < t.ArrayLength; i++ {
		sb.WriteString("[]")
	}

	return sb.String()
}
