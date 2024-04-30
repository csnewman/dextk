package java

import (
	"fmt"
	"strings"
)

type Writer struct {
	buffer *strings.Builder
	indent int
}

func NewWriter() *Writer {
	return &Writer{
		buffer: &strings.Builder{},
	}
}

func (w *Writer) Indent(d int) {
	w.indent += d
}

func (w *Writer) WriteLinef(format string, args ...any) {
	w.WriteLine(fmt.Sprintf(format, args...))
}

func (w *Writer) WriteLine(text string) {
	for i := 0; i < w.indent; i++ {
		w.buffer.WriteString("    ")
	}

	w.buffer.WriteString(text)
	w.buffer.WriteRune('\n')
}

func (w *Writer) EmptyLine() {
	w.buffer.WriteRune('\n')
}

func (w *Writer) String() string {
	return w.buffer.String()
}

func (w *Writer) Reset() {
	w.buffer = &strings.Builder{}
}
