package model

import (
	"bytes"
	"fmt"
)

type ObjectNode struct {
	Value map[string]Node
}

func (n ObjectNode) Type() NodeType {
	return ObjectType
}

func (n *ObjectNode) MarshalJSON() ([]byte, error) {
	result := make([][]byte, 0, len(n.Value))
	for k, v := range n.Value {
		b, err := v.MarshalJSON()
		if err != nil {
			return nil, err
		}
		result = append(result, []byte(fmt.Sprintf("\"%s\": %s", k, b)))
	}
	return bytes.Join(
		[][]byte{
			[]byte("{"),
			bytes.Join(result, []byte(", ")),
			[]byte("}"),
		}, []byte("")), nil
}

func (n *ObjectNode) Set(k string, v any) {
	n.Value[k] = NewNode(v)
}

func (n *ObjectNode) Get(k string) (Node, error) {
	child, ok := n.Value[k]
	if !ok {
		return nil, fmt.Errorf("field %s not found", k)
	}
	return child, nil
}

func (n *ObjectNode) Merge(n2 *ObjectNode) {
	for k, v := range n2.Value {
		n.Value[k] = v
	}
}

func (n *ObjectNode) Len() int {
	return len(n.Value)
}
