package model

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
