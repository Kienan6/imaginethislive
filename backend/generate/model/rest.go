package model

import (
	"io"
	"reflect"
	"text/template"
)

type RestField struct {
	Name string
	Type string
}

type TemplateParams struct {
	Prefix      string
	PackageName string
	Imports     []string
	Types       map[string][]RestField
}

func ProcessRestModelConversion(t reflect.Type, fields map[string][]RestField, key string) {
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		_, ok := field.Tag.Lookup("rest")
		if !ok {
			continue
		}
		fields[key] = append(fields[key], RestField{
			Name: field.Name,
			Type: field.Type.String(),
		})
	}
}

func ToTemplate(params TemplateParams, f io.Writer) {
	temp := `package {{ .PackageName }}

import (
{{- range $key, $import := .Imports }}
	{{ $import | printf "%q"  }}
{{- end }}
)
{{- range $key, $value := .Types }}
type {{ $.Prefix }}{{ $key }} struct {
	{{- range $index, $field := $value }}
	{{ $field.Name }} {{ $field.Type }}
	{{- end }}
}

{{- end }}
	`
	t, err := template.New("rest").Parse(temp)
	if err != nil {
		panic(err)
	}

	err = t.Execute(f, params)
	if err != nil {
		panic(err)
	}

}
