package model

import "strconv"

type NumberNode struct {
	Value float64
}

func (n NumberNode) Type() NodeType {
	return NumberType
}

func (n *NumberNode) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatFloat(n.Value, 'g', -1, 64)), nil
}
