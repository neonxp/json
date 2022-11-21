package json

import (
	"reflect"
	"testing"

	"go.neonxp.dev/json/model"
)

func TestQuery(t *testing.T) {
	type args struct {
		json  string
		query string
	}
	tests := []struct {
		name    string
		args    args
		want    model.Node
		wantErr bool
	}{
		{
			name: "Complex",
			args: args{
				json: `{
					"key1": "value1",
					"key2": [
						"item 1",
						"item 2",
						"item 3",
						"item 4",
						"item 5",
						"item 6",
						{
							"status": "invalid"
						},
						{
							"status": "valid",
							"embededArray": [
								"not target",
								"not target",
								"not target",
								"target",
								"not target",
								"not target",
							]
						}
					]
				}`,
				query: "key2.7.embededArray.3",
			},
			want:    model.NewNode("target"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Query(tt.args.json, tt.args.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("Query() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Query() = %v, want %v", got, tt.want)
			}
		})
	}
}
