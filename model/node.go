package model

import (
	"bytes"
	"fmt"
	"strconv"
)

// Node of JSON tree
type Node struct {
	Type         NodeType
	Meta         NodeObjectValue
	StringValue  string
	NumberValue  float64
	ObjectValue  NodeObjectValue
	ArrayValue   NodeArrayValue
	BooleanValue bool
}

// NewNode creates new node from value
func NewNode(value any) *Node {
	n := new(Node)
	n.SetValue(value)
	return n
}

// Value returns value of node
func (n *Node) Value() any {
	switch n.Type {
	case StringNode:
		return n.StringValue
	case NumberNode:
		return n.NumberValue
	case ObjectNode:
		return n.ObjectValue
	case ArrayNode:
		return n.ArrayValue
	case BooleanNode:
		return n.BooleanValue
	default:
		return nil
	}
}

// SetValue to node
func (n *Node) SetValue(value any) {
	switch value := value.(type) {
	case string:
		n.Type = StringNode
		n.StringValue = value
	case float64:
		n.Type = NumberNode
		n.NumberValue = value
	case int:
		n.Type = NumberNode
		n.NumberValue = float64(value)
	case NodeObjectValue:
		n.Type = ObjectNode
		meta, hasMeta := value["@"]
		if hasMeta {
			n.Meta = meta.ObjectValue
			delete(value, "@")
		}
		n.ObjectValue = value
	case NodeArrayValue:
		n.Type = ArrayNode
		n.ArrayValue = value
	case bool:
		n.Type = BooleanNode
		n.BooleanValue = value
	default:
		n.Type = NullNode
	}
}

// MarshalJSON to []byte
func (n *Node) MarshalJSON() ([]byte, error) {
	switch n.Type {
	case StringNode:
		return []byte(`"` + n.StringValue + `"`), nil
	case NumberNode:
		return []byte(strconv.FormatFloat(n.NumberValue, 'g', -1, 64)), nil
	case ObjectNode:
		result := make([][]byte, 0, len(n.ObjectValue))
		for k, v := range n.ObjectValue {
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
	case ArrayNode:
		result := make([][]byte, 0, len(n.ArrayValue))
		for _, v := range n.ArrayValue {
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
	case BooleanNode:
		if n.BooleanValue {
			return []byte("true"), nil
		}
		return []byte("false"), nil
	default:
		return []byte("null"), nil
	}
}

// Merge two object or array nodes
func (n *Node) Merge(node *Node) error {
	if n.Type != node.Type {
		return fmt.Errorf("can't merge nodes of different types")
	}
	switch n.Type {
	case ObjectNode:
		for k, v := range node.ObjectValue {
			n.ObjectValue[k] = v
		}
	case ArrayNode:
		n.ArrayValue = append(n.ArrayValue, node.ArrayValue...)
	default:
		return fmt.Errorf("merge not implemented for type %s", n.Type)
	}
	return nil
}

// Len returns length of object or array nodes
func (n *Node) Len() (int, error) {
	switch n.Type {
	case ObjectNode:
		return len(n.ObjectValue), nil
	case ArrayNode:
		return len(n.ArrayValue), nil
	default:
		return 0, fmt.Errorf("merge not implemented for type %s", n.Type)
	}
}

// Meta represents node metadata
type Meta map[string]any
