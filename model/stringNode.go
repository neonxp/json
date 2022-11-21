package model

import "fmt"

type StringNode struct {
	Value string
}

func (n StringNode) Type() NodeType {
	return StringType
}

func (n *StringNode) MarshalJSON() ([]byte, error) {
	return []byte(`"` + n.Value + `"`), nil
}

func (n *StringNode) Set(v any) error {
	val, ok := v.(string)
	if !ok {
		return fmt.Errorf("%v is not string", v)
	}
	n.Value = val
	return nil
}
