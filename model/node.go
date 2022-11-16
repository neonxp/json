package model

import (
	"bytes"
	"fmt"
	"strconv"
)

// Node of JSON tree
type Node struct {
	Type         NodeType
	stringValue  string
	numberValue  float64
	objectValue  NodeObjectValue
	arrayValue   NodeArrayValue
	booleanValue bool
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
		return n.stringValue
	case NumberNode:
		return n.numberValue
	case ObjectNode:
		return n.objectValue
	case ArrayNode:
		return n.arrayValue
	case BooleanNode:
		return n.booleanValue
	default:
		return nil
	}
}

// SetValue to node
func (n *Node) SetValue(value any) {
	switch value := value.(type) {
	case string:
		n.Type = StringNode
		n.stringValue = value
	case float64:
		n.Type = NumberNode
		n.numberValue = value
	case int:
		n.Type = NumberNode
		n.numberValue = float64(value)
	case NodeObjectValue:
		n.Type = ObjectNode
		n.objectValue = value
	case NodeArrayValue:
		n.Type = ArrayNode
		n.arrayValue = value
	case bool:
		n.Type = BooleanNode
		n.booleanValue = value
	default:
		n.Type = NullNode
	}
}

// MarshalJSON to []byte
func (n *Node) MarshalJSON() ([]byte, error) {
	switch n.Type {
	case StringNode:
		return []byte(`"` + n.stringValue + `"`), nil
	case NumberNode:
		return []byte(strconv.FormatFloat(n.numberValue, 'g', -1, 64)), nil
	case ObjectNode:
		result := make([][]byte, 0, len(n.objectValue))
		for k, v := range n.objectValue {
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
		result := make([][]byte, 0, len(n.arrayValue))
		for _, v := range n.arrayValue {
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
		if n.booleanValue {
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
		for k, v := range node.objectValue {
			n.objectValue[k] = v
		}
	case ArrayNode:
		n.arrayValue = append(n.arrayValue, node.arrayValue...)
	default:
		return fmt.Errorf("merge not implemented for type %s", n.Type)
	}
	return nil
}

// Len returns length of object or array nodes
func (n *Node) Len() (int, error) {
	switch n.Type {
	case ObjectNode:
		return len(n.objectValue), nil
	case ArrayNode:
		return len(n.arrayValue), nil
	default:
		return 0, fmt.Errorf("merge not implemented for type %s", n.Type)
	}
}
