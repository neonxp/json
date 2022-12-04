package model

type NodeType string

const (
	StringType  NodeType = "string"
	NumberType  NodeType = "number"
	ObjectType  NodeType = "object"
	ArrayType   NodeType = "array"
	BooleanType NodeType = "boolean"
	NullType    NodeType = "null"
)

type NodeObjectValue map[string]Node

func (n NodeObjectValue) Set(k string, v any) error {
	n[k] = NewNode(v)
	return nil
}
