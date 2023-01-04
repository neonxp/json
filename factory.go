package json

type Factory interface {
	Produce(typ NodeType) (Node, error)
	Fill(n Node, value any)
}

type Node interface {
	ToJSON() string
}

type ObjectNode interface {
	Node
	Set(k string, v Node)
	Get(k string) (Node, bool)
}

type ArrayNode interface {
	Node
	Append(v Node)
	Index(i int) Node
	Len() int
}

type StringNode interface {
	Node
}

type NumberNode interface {
	Node
}

type BooleanNode interface {
	Node
}

type NullNode interface {
	Node
}

type AcceptParent interface {
	SetParent(n Node)
}
