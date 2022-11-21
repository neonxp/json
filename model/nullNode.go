package model

type NullNode struct{}

func (n NullNode) Type() NodeType {
	return NullType
}

func (n NullNode) MarshalJSON() ([]byte, error) {
	return []byte("null"), nil
}
