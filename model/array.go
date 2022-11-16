package model

import "fmt"

// Index returns node by index from array
func (n *Node) Index(idx int) (*Node, error) {
	arrlen := len(n.arrayValue)
	if idx >= arrlen {
		return nil, fmt.Errorf("index %d out of range (len=%d)", idx, arrlen)
	}
	return n.arrayValue[idx], nil
}

// SetIndex sets node to array by index
func (n *Node) SetIndex(idx int, value *Node) error {
	arrlen := len(n.arrayValue)
	if idx >= arrlen {
		return fmt.Errorf("index %d out of range (len=%d)", idx, arrlen)
	}
	n.arrayValue[idx] = value
	return nil
}

// Each applies callback to each element of array
func (n *Node) Each(cb func(idx int, value *Node) error) error {
	for i, v := range n.arrayValue {
		if err := cb(i, v); err != nil {
			return err
		}
	}
	return nil
}
