package model

type NodeType string

const (
	StringNode  NodeType = "string"
	NumberNode  NodeType = "number"
	ObjectNode  NodeType = "object"
	ArrayNode   NodeType = "array"
	BooleanNode NodeType = "boolean"
	NullNode    NodeType = "null"
)

type NodeObjectValue map[string]*Node

func (n NodeObjectValue) Set(k string, v any) {
	n[k] = NewNode(v)
}

type NodeArrayValue []*Node
