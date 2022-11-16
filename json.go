package json

import (
	"strings"

	"go.neonxp.dev/json/model"
	"go.neonxp.dev/json/parser"
)

// Marshal Node tree to []byte
func Marshal(node *model.Node) ([]byte, error) {
	return node.MarshalJSON()
}

// Unmarshal data to Node tree
func Unmarshal(data []byte) (*model.Node, error) {
	return parser.Parse(string(data))
}

// Query returns node by query string (dot notation)
func Query(json string, query string) (*model.Node, error) {
	n, err := parser.Parse(json)
	if err != nil {
		return nil, err
	}
	return n.Query(strings.Split(query, "."))
}

// QueryArray returns node by array query
func QueryArray(json string, query []string) (*model.Node, error) {
	n, err := parser.Parse(json)
	if err != nil {
		return nil, err
	}
	return n.Query(query)
}
