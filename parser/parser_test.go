package parser

import (
	"reflect"
	"testing"

	"go.neonxp.dev/json/model"
)

func TestParse(t *testing.T) {
	type args struct {
		json string
	}
	tests := []struct {
		name    string
		args    args
		want    *model.Node
		wantErr bool
	}{
		{
			name: "complex",
			args: args{
				json: `{
					"string key": "string value",
					"number key": 1337,
					"float key": 123.3,
					"object key": {
						"ab": "cd"
					},
					"array key": [
						1, 
						2, 
						"three"
					],
					"null key":null,
					"boolean key":true
					}`,
			},
			want: model.NewNode(
				model.NodeObjectValue{
					"string key": model.NewNode("string value"),
					"number key": model.NewNode(1337),
					"float key":  model.NewNode(123.3),
					"object key": model.NewNode(model.NodeObjectValue{
						"ab": model.NewNode("cd"),
					}),
					"array key": model.NewNode(model.NodeArrayValue{
						model.NewNode(1),
						model.NewNode(2),
						model.NewNode("three"),
					}),
					"null key":    model.NewNode(nil),
					"boolean key": model.NewNode(true),
				},
			),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args.json)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
