package model

import (
	"fmt"
	"strconv"
)

type NumberNode struct {
	Value float64
}

func (n NumberNode) Type() NodeType {
	return NumberType
}

func (n *NumberNode) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatFloat(n.Value, 'g', -1, 64)), nil
}

func (n *NumberNode) Set(v any) error {
	switch v := v.(type) {
	case float64:
		n.Value = v
	case int:
		n.Value = float64(v)
	}
	return fmt.Errorf("%v is not number", v)
}
