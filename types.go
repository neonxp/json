package json

type NodeType string

const (
	StringType  NodeType = "string"
	NumberType  NodeType = "number"
	ObjectType  NodeType = "object"
	ArrayType   NodeType = "array"
	BooleanType NodeType = "boolean"
	NullType    NodeType = "null"
)
