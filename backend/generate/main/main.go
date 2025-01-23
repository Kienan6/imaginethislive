package main

import (
	"itl/generate/model"
	"itl/model/domain"
	"os"
	"reflect"
)

//go:generate go run ./main.go

// Output TODO: more generic - package name, paths, prefix
type Output struct {
	fileName string
	imports  []string
	structs  map[string]any
}

func main() {

	packageName := "rest"

	output := []Output{
		{
			fileName: "comment.go",
			imports: []string{
				"github.com/google/uuid",
				"time",
			},
			structs: map[string]any{
				"Comment": domain.Comment{},
			},
		},
	}

	for _, o := range output {
		path := "../../model/" + packageName + "/" + o.fileName
		f, err := os.Create(path)
		if err != nil {
			panic(err)
		}
		for key, val := range o.structs {

			fields := make(map[string][]model.RestField)

			model.ProcessRestModelConversion(reflect.TypeOf(val), fields, key)
			params := model.TemplateParams{
				Prefix:      "",
				Imports:     o.imports,
				PackageName: packageName,
				Types:       fields,
			}

			model.ToTemplate(params, f)
			err = f.Close()
			if err != nil {
				panic(err)
			}
		}
	}

}
