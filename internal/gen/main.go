package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
)

func main() {
	t, err := template.New("tmpl").Parse(opsTemplateSrc)
	if err != nil {
		panic(err)
	}

	f, err := os.Create("op.gen.go")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	d := &Data{}

	for _, op := range Ops {
		d.Ops = append(d.Ops, GenOp{
			DisplayName: op.Name,
			PascalName:  strcase.ToCamel(strings.ReplaceAll(op.Name, "/", "_")),
			CodeConst:   fmt.Sprintf("%x", op.Code),
			Fmt:         op.Fmt,
		})
	}

	if err := t.Execute(f, d); err != nil {
		panic(err)
	}
}

type Data struct {
	Ops []GenOp
}

type GenOp struct {
	DisplayName string
	PascalName  string
	CodeConst   string
	Fmt         string
}

var opsTemplateSrc = `{{$top := . -}}
package dextk

import "fmt"

const (
{{- range $o := $top.Ops}}
	OpCode{{$o.PascalName}} OpCode = 0x{{$o.CodeConst}}
{{- end}}
)

var opConfigs = map[OpCode]opConfig{
{{- range $o := $top.Ops}}
	OpCode{{$o.PascalName}}: {
		Name: "{{$o.DisplayName}}",
		Size: fmt{{$o.Fmt}}Size,
		Reader: func(r *OpReader) (Op, error) {
			pos := r.pos

			f, err := r.readFmt{{$o.Fmt}}()
			if err != nil {
				return nil, err
			}

			return Op{{$o.PascalName}} {
				opBase{pos: pos},
				f,
			}, nil
		},
	},
{{- end}}
}
{{range $o := $top.Ops}}
type Op{{$o.PascalName}} struct {
	opBase
	Fmt{{$o.Fmt}}
}

func (o Op{{$o.PascalName}}) Code() OpCode {
	return OpCode{{$o.PascalName}}
}

func (o Op{{$o.PascalName}}) Fmt() Fmt {
	return o.Fmt{{$o.Fmt}}
}

func (o Op{{$o.PascalName}}) String() string {
	f := o.Fmt{{$o.Fmt}}.String()
	if f != "" {
		return fmt.Sprintf("0x%x: {{$o.DisplayName}} %v", o.pos, f)
	}

	return fmt.Sprintf("0x%x: {{$o.DisplayName}}", o.pos)
}
{{end}}`
