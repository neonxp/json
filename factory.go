package json

type NodeFactory func(typ NodeType) (Node, error)

type Node interface {
	String() string
}

type ObjectNode interface {
	Node
	SetKeyValue(k string, v Node)
	GetByKey(k string) (Node, bool)
}

type ArrayNode interface {
	Node
	Append(v Node)
	Index(i int) Node
	Len() int
}

type StringNode interface {
	Node
	SetString(v string)
	GetString() string
}

type NumberNode interface {
	Node
	SetNumber(v float64)
	GetNumber() float64
}

type BooleanNode interface {
	Node
	SetBool(v bool)
	GetBool() bool
}

type NullNode interface {
	Node
}

type AcceptParent interface {
	SetParent(n Node)
}
