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
	node, ok := n.ObjectValue[key]
	if !ok {
		keys := make([]string, 0, len(n.ObjectValue))
		for k := range n.ObjectValue {
			keys = append(keys, k)
		}
		return nil, fmt.Errorf("field '%s' does not exist in object (keys %s)", key, strings.Join(keys, ", "))
	}
	return node, nil
}

// Set node to object by key
func (n *Node) Set(key string, value Node) error {
	if n.Type != ObjectNode {
		return fmt.Errorf("node must be object, got %s", n.Type)
	}
	n.ObjectValue[key] = &value
	return nil
}

// Map callback to each key value pair of object
func (n *Node) Map(cb func(key string, value *Node) (*Node, error)) error {
	for k, v := range n.ObjectValue {
		newNode, err := cb(k, v)
		if err != nil {
			return err
		}
		n.ObjectValue[k] = newNode
	}
	return nil
}

// Remove by key from object
func (n *Node) Remove(key string) error {
	if n.Type != ObjectNode {
		return fmt.Errorf("node must be object, got %s", n.Type)
	}
	delete(n.ObjectValue, key)
	return nil
}
