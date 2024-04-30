package java

type Field struct {
	Name      string
	Access    Access
	Final     bool
	Static    bool
	Synthetic bool
	Type      *Type
}

func (f *Field) Write(w *Writer, c *Class) {
	acc := f.Access
	if acc != "" {
		acc += " "
	}

	if f.Static {
		acc += "static "
	}

	if f.Final {
		acc += "final "
	}

	ty := c.GenType(f.Type)

	var cmt string

	if f.Synthetic {
		cmt = "/* synthetic */"
	}

	if cmt != "" {
		cmt = " "
	}

	w.WriteLinef("%v%v %v;%v", acc, ty, f.Name, cmt)
}
