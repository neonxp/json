package model

// Compare current node with another node
func (n *Node) Compare(op Operand, node *Node) bool {
	switch op {
	case OpEq:
		return n.Value() == node.Value()
	case OpNeq:
		return n.Value() != node.Value()
	case OpLess:
		return less(n, node)
	case OpGt:
		return less(node, n)
	case OpLessEq:
		return less(n, node) || n.Value() == node.Value()
	case OpGtEq:
		return less(node, n) || n.Value() == node.Value()
	case OpIn:
		if n.Type != ArrayNode {
			return false
		}
		for _, v := range n.arrayValue {
			if v.Value() == node.Value() {
				return true
			}
		}
	}
	return false
}

func less(n1 *Node, n2 *Node) bool {
	if n1.Type != n2.Type {
		return false
	}
	switch n1.Type {
	case NumberNode:
		return n1.numberValue < n2.numberValue
	case StringNode:
		return n1.stringValue < n2.stringValue
	default:
		return false
	}
}

type Operand int

const (
	OpEq Operand = iota
	OpNeq
	OpLess
	OpLessEq
	OpGt
	OpGtEq
	OpIn
)
