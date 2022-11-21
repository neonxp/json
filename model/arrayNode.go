package model

import (
	"bytes"
	"fmt"
)

type ArrayNode struct {
	Value NodeArrayValue
}

func (n ArrayNode) Type() NodeType {
	return ArrayType
}

func (n *ArrayNode) MarshalJSON() ([]byte, error) {
	result := make([][]byte, 0, len(n.Value))
	for _, v := range n.Value {
		b, err := v.MarshalJSON()
		if err != nil {
			return nil, err
		}
		result = append(result, b)
	}
	return bytes.Join(
		[][]byte{
			[]byte("["),
			bytes.Join(result, []byte(", ")),
			[]byte("]"),
		}, []byte("")), nil
}

func (n *ArrayNode) Set(v any) error {
	val, ok := v.(NodeArrayValue)
	if !ok {
		return fmt.Errorf("%v is not array", v)
	}
	n.Value = val
	return nil
}

func (n *ArrayNode) Index(idx int) (Node, error) {
	if len(n.Value) <= idx {
		return nil, fmt.Errorf("index %d out of range [0...%d]", idx, len(n.Value)-1)
	}
	return n.Value[idx], nil
}

func (n *ArrayNode) Merge(n2 *ArrayNode) {
	n.Value = append(n.Value, n2.Value...)
}

func (n *ArrayNode) Len() int {
	return len(n.Value)
}

type NodeArrayValue []Node
