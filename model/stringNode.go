package model

type StringNode struct {
	Value string
}

func (n StringNode) Type() NodeType {
	return StringType
}

func (n *StringNode) MarshalJSON() ([]byte, error) {
	return []byte(`"` + n.Value + `"`), nil
}
