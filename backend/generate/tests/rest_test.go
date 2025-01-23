package tests

import (
	"github.com/stretchr/testify/assert"
	"itl/generate/model"
	"reflect"
	"testing"
)

type before struct {
	included string `rest:"include"`
	excluded string
}

type test struct {
	key    string
	model  interface{}
	fields map[string][]model.RestField
}

func TestRestModelConversion(t *testing.T) {

	var tests = []test{
		{
			key:   "before",
			model: before{},
			fields: map[string][]model.RestField{
				"before": {
					model.RestField{
						Name: "included",
						Type: "string",
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			fields := make(map[string][]model.RestField)
			model.ProcessRestModelConversion(reflect.TypeOf(test.model), fields, test.key)

			assert.Equal(t, test.fields, fields)
		})
	}

}
