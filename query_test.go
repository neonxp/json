package json_test

import (
	"reflect"
	"testing"

	"go.neonxp.dev/json"
	"go.neonxp.dev/json/std"
)

func TestMustQuery(t *testing.T) {
	jsonString := `{
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
	}`
	type args struct {
		parent json.Node
		path   []string
	}
	tests := []struct {
		name string
		args args
		want json.Node
	}{
		{
			name: "find in object",
			args: args{
				parent: json.New(std.Factory).MustUnmarshal(jsonString),
				path:   []string{"object", "object 2", "three"},
			},
			want: &std.StringNode{Value: "four"},
		},
		{
			name: "find in array",
			args: args{
				parent: json.New(std.Factory).MustUnmarshal(jsonString),
				path:   []string{"array", "[4]", "five"},
			},
			want: &std.StringNode{Value: "six"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := json.MustQuery(tt.args.parent, tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MustQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}
