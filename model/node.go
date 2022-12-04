package model

// Node of JSON tree
type Node interface {
	Type() NodeType
	MarshalJSON() ([]byte, error)
	Set(v any) error
}

// NewNode creates new node from value
func NewNode(value any) Node {
	if value, ok := value.(Node); ok {
		return value
	}
	switch value := value.(type) {
	case string:
		return &StringNode{
			Value: value,
		}
	case float64:
		return &NumberNode{
			Value: value,
		}
	case int:
		return &NumberNode{
			Value: float64(value),
		}
	case NodeObjectValue:
		return &ObjectNode{
			Value: value,
			Meta:  make(map[string]any),
		}
	case NodeArrayValue:
		return &ArrayNode{
			Value: value,
		}
	case bool:
		return &BooleanNode{
			Value: value,
		}
	default:
		return NullNode{}
	}
}
