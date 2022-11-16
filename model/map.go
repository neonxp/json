package model

import (
	"fmt"
	"strings"
)

// Get node from object by key
func (n *Node) Get(key string) (*Node, error) {
	if n.Type != ObjectNode {
		return nil, fmt.Errorf("node must be object, got %s", n.Type)
	}
	node, ok := n.objectValue[key]
	if !ok {
		keys := make([]string, 0, len(n.objectValue))
		for k := range n.objectValue {
			keys = append(keys, k)
		}
		return nil, fmt.Errorf("field '%s' does not exist in object (keys %s)", key, strings.Join(keys, ", "))
	}
	return node, nil
}

// Set node to object by key
func (n *Node) Set(key string, value *Node) error {
	if n.Type != ObjectNode {
		return fmt.Errorf("node must be object, got %s", n.Type)
	}
	n.objectValue[key] = value
	return nil
}

// Map callback to each key value pair of object
func (n *Node) Map(cb func(key string, value *Node) (*Node, error)) error {
	for k, v := range n.objectValue {
		newNode, err := cb(k, v)
		if err != nil {
			return err
		}
		n.objectValue[k] = newNode
	}
	return nil
}
