package model

import (
	stdJSON "encoding/json"
	"reflect"
	"testing"
)

func TestNode_MarshalJSON(t *testing.T) {
	type fields struct {
		node Node
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			name: "empty",
			fields: fields{
				node: NewNode(nil),
			},
			want: []byte(`null`),
		},
		{
			name: "string",
			fields: fields{
				node: NewNode("this is a string"),
			},
			want: []byte(`"this is a string"`),
		},
		{
			name: "int",
			fields: fields{
				node: NewNode(123),
			},
			want: []byte(`123`),
		},
		{
			name: "float",
			fields: fields{
				node: NewNode(123.321),
			},
			want: []byte(`123.321`),
		},
		{
			name: "booleant",
			fields: fields{
				node: NewNode(true),
			},
			want: []byte(`true`),
		},
		{
			name: "booleanf",
			fields: fields{
				node: NewNode(false),
			},
			want: []byte(`false`),
		},
		{
			name: "complex",
			fields: fields{
				node: NewNode(
					NodeObjectValue{
						"string key": NewNode("string value"),
						"number key": NewNode(1337),
						"float key":  NewNode(123.3),
						"object key": NewNode(NodeObjectValue{
							"ab": NewNode("cd"),
						}),
						"array key": NewNode(NodeArrayValue{
							NewNode(1), NewNode(2), NewNode("three"),
						}),
						"boolean key": NewNode(true),
						"null key":    NewNode(nil),
					},
				),
			},
			want: []byte(
				`{"string key": "string value", "number key": 1337, "float key": 123.3, "object key": {"ab": "cd"}, "array key": [1, 2, "three"], "boolean key": true, "null key": null}`,
			),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				gotObj  any
				wantObj any
			)

			got, err := tt.fields.node.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("Node.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			err = stdJSON.Unmarshal(got, &gotObj) // TODO use own unmarshaller
			if err != nil {
				t.Errorf("Generated invalid json = %s, error = %v", got, err)
			}
			_ = stdJSON.Unmarshal(tt.want, &wantObj) // I belive, test is correct
			if !reflect.DeepEqual(gotObj, wantObj) {
				t.Errorf("Node.MarshalJSON() = %s, want %s", got, tt.want)
			}
		})
	}
}
