package pdfext

import (
	"reflect"
	"testing"
)

type TestA struct {
	Name    string
	Account string
	Age     string
}

func TestToTableData(t *testing.T) {
	type args struct {
		data interface{}
	}
	tests := []struct {
		name string
		args args
		want []map[string]string
	}{
		// TODO: Add test cases.
		{"case1",
			args{[]TestA{{"tesat", "21234", "13"}, {"3241", "g4211", "63"}}},
			[]map[string]string{
				{
					"Name":    "tesat",
					"Account": "21234",
					"Age":     "13",
				},
				{
					"Name":    "3241",
					"Account": "g4211",
					"Age":     "63",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToTableData(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToTableData() = %v, want %v", got, tt.want)
			}
		})
	}
}
