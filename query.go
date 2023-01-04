package json

import (
	"fmt"
	"strconv"
	"strings"
)

func Query(parent Node, path []string) (Node, error) {
	if len(path) == 0 {
		return parent, nil
	}
	head, rest := path[0], path[1:]
	switch parent := parent.(type) {
	case ObjectNode:
		next, ok := parent.Get(head)
		if !ok {
			return nil, fmt.Errorf("key %s not found at object %v", head, parent)
		}
		return Query(next, rest)
	case ArrayNode:
		stringIdx := strings.Trim(head, "[]")
		idx, err := strconv.Atoi(stringIdx)
		if err != nil {
			return nil, fmt.Errorf("key %s is invalid index: %w", stringIdx, err)
		}
		if idx >= parent.Len() {
			return nil, fmt.Errorf("index %d is out of range (len=%d)", idx, parent.Len())
		}
		next := parent.Index(idx)
		return Query(next, rest)
	default:
		return nil, fmt.Errorf("can't get key=%s from node type = %t", head, parent)
	}
}

func MustQuery(parent Node, path []string) Node {
	n, err := Query(parent, path)
	if err != nil {
		panic(err)
	}
	return n
}
