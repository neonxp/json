package json_test

import (
	"reflect"
	"testing"

	"go.neonxp.dev/json"
	"go.neonxp.dev/json/std"
)

func TestJSON_Unmarshal(t *testing.T) {
	j := &json.JSON{
		Factory: &std.Factory{},
	}
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		want    json.Node
		wantErr bool
	}{
		{
			name: "object with strings",
			args: args{
				input: `{
					"hello": "world",
				}`,
			},
			want: std.ObjectNode{
				"hello": &std.StringNode{Value: "world"},
			},
			wantErr: false,
		},
		{
			name: "complex object",
			args: args{
				input: `{
					"string key": "string value",
					"number key": 123.321,
					"bool key": true,
					"object": {
						"one": "two",
						"object 2": {
							"three": "four"
						}
					},
					"array": [
						"one",
						2,
						true,
						null,
						{
							"five": "six"
						}
					]
				}`,
			},
			want: std.ObjectNode{
				"string key": j.MustNode("string value"),
				"number key": j.MustNode(123.321),
				"bool key":   j.MustNode(true),
				"object": std.ObjectNode{
					"one": j.MustNode("two"),
					"object 2": std.ObjectNode{
						"three": j.MustNode("four"),
					},
				},
				"array": &std.ArrayNode{
					j.MustNode("one"),
					j.MustNode(2),
					j.MustNode(true),
					j.MustNode(nil),
					std.ObjectNode{
						"five": j.MustNode("six"),
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := j.Unmarshal(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("JSON.Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JSON.Unmarshal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJSON_Marshal(t *testing.T) {
	j := &json.JSON{
		Factory: &std.Factory{},
	}
	type args struct {
		n json.Node
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "complex object",
			args: args{
				n: std.ObjectNode{
					"string key": j.MustNode("string value"),
					"number key": j.MustNode(123.321),
					"bool key":   j.MustNode(true),
					"object": std.ObjectNode{
						"one": j.MustNode("two"),
						"object 2": std.ObjectNode{
							"three": j.MustNode("four"),
						},
					},
					"array": &std.ArrayNode{
						j.MustNode("one"),
						j.MustNode(2),
						j.MustNode(true),
						j.MustNode(nil),
						std.ObjectNode{
							"five": j.MustNode("six"),
						},
					},
				},
			},
			want: `{"string key":"string value","number key":123.321,"bool key":true,"object":{"one":"two","object 2":{"three":"four"}},"array":["one",2,true,null,{"five":"six"}]}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := j.Marshal(tt.args.n); len(got) != len(tt.want) {
				t.Errorf("JSON.Marshal() = %v, want %v", got, tt.want)
			}
		})
	}
}
