# JSON parsing library

This library is an marshaler/unmarshaler for JSON in a tree of nodes. Also allows you to make queries over these trees.

## Library interface

```go
package json // import "go.neonxp.dev/json"

// Marshal Node tree to []byte
func Marshal(node *model.Node) ([]byte, error)

// Unmarshal data to Node tree
func Unmarshal(data []byte) (*model.Node, error)

// Query returns node by query string (dot notation)
func Query(json string, query string) (*model.Node, error)

// QueryArray returns node by array query
func QueryArray(json string, query []string) (*model.Node, error)
```

## Node methods

```go
package model // import "go.neonxp.dev/json/model"

// Node of JSON tree
type Node struct {
    Type NodeType
}

// NewNode creates new node from value
func NewNode(value any) *Node

// Get node from object by key
func (n *Node) Get(key string) (*Node, error)

// Index returns node by index from array
func (n *Node) Index(idx int) (*Node, error)

// Set node to object by key
func (n *Node) Set(key string, value *Node) error

// SetIndex sets node to array by index
func (n *Node) SetIndex(idx int, value *Node) error

// SetValue to node
func (n *Node) SetValue(value any)

// Map callback to each key value pair of object
func (n *Node) Map(cb func(key string, value *Node) (*Node, error)) error

// Each applies callback to each element of array
func (n *Node) Each(cb func(idx int, value *Node) error) error

// Query returns node by array query
func (n *Node) Query(query []string) (*Node, error)

// Value returns value of node
func (n *Node) Value() any

// MarshalJSON to []byte
func (n *Node) MarshalJSON() ([]byte, error)
```
