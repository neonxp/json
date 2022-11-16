package model

import (
	"fmt"
	"strconv"
)

// Query returns node by array query
func (n *Node) Query(query []string) (*Node, error) {
	if len(query) == 0 {
		return n, nil
	}
	head, rest := query[0], query[1:]
	switch n.Type {
	case ArrayNode:
		idx, err := strconv.Atoi(head)
		if err != nil {
			return nil, fmt.Errorf("index must be a number, got %s", head)
		}
		next, err := n.Index(idx)
		if err != nil {
			return nil, err
		}
		return next.Query(rest)
	case ObjectNode:
		next, err := n.Get(head)
		if err != nil {
			return nil, err
		}
		return next.Query(rest)
	}
	return nil, fmt.Errorf("can't get %s from node type %s", head, n.Type)
}
