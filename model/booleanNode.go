package model

import "fmt"

type BooleanNode struct {
	Value bool
}

func (n BooleanNode) Type() NodeType {
	return BooleanType
}

func (n *BooleanNode) MarshalJSON() ([]byte, error) {
	if n.Value {
		return []byte("true"), nil
	}
	return []byte("false"), nil
}

func (n *BooleanNode) Set(v any) error {
	val, ok := v.(bool)
	if !ok {
		return fmt.Errorf("%v is not boolean", v)
	}
	n.Value = val
	return nil
}
