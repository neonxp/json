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

Other methods: https://pkg.go.dev/go.neonxp.dev/json