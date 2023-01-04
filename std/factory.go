package std

import (
	"fmt"
	"strconv"
	"strings"

	"go.neonxp.dev/json"
)

type Factory struct{}

func (f *Factory) Produce(typ json.NodeType) (json.Node, error) {
	switch typ {
	case json.ObjectType:
		return ObjectNode{}, nil
	case json.ArrayType:
		return &ArrayNode{}, nil
	case json.StringType:
		return &StringNode{}, nil
	case json.NumberType:
		return &NumberNode{}, nil
	case json.BooleanType:
		return &BooleanNode{}, nil
	case json.NullType:
		return NullNode{}, nil
	}
	return nil, fmt.Errorf("unknown type: %s", typ)
}

func (f *Factory) Fill(n json.Node, value any) {
	switch n := n.(type) {
	case *ObjectNode:
		for k, v := range value.(map[string]json.Node) {
			n.Set(k, v)
		}
	case *ArrayNode:
		for _, v := range value.([]json.Node) {
			n.Append(v)
		}
	case *StringNode:
		n.Value = value.(string)
	case *NumberNode:
		n.Value = value.(float64)
	case *BooleanNode:
		n.Value = value.(bool)
	}
}

type ObjectNode map[string]json.Node

func (o ObjectNode) Set(k string, v json.Node) {
	o[k] = v
}

func (o ObjectNode) Get(k string) (json.Node, bool) {
	v, ok := o[k]
	return v, ok
}

func (o ObjectNode) ToJSON() string {
	res := make([]string, 0, len(o))
	for k, n := range o {
		res = append(res, fmt.Sprintf(`"%s":%s`, k, n.ToJSON()))
	}
	return fmt.Sprintf(`{%s}`, strings.Join(res, ","))
}

type ArrayNode []json.Node

func (o *ArrayNode) Append(v json.Node) {
	na := append(*o, v)
	*o = na
}

func (o *ArrayNode) Index(i int) json.Node {
	return (*o)[i]
}

func (o *ArrayNode) Len() int {
	return len(*o)
}

func (o *ArrayNode) ToJSON() string {
	res := make([]string, 0, len(*o))
	for _, v := range *o {
		res = append(res, v.ToJSON())
	}
	return fmt.Sprintf(`[%s]`, strings.Join(res, ","))
}

type StringNode struct {
	Value string
}

func (o *StringNode) ToJSON() string {
	return `"` + o.Value + `"`
}

type NumberNode struct {
	Value float64
}

func (o *NumberNode) ToJSON() string {
	return strconv.FormatFloat(float64(o.Value), 'g', 15, 64)
}

type BooleanNode struct {
	Value bool
}

func (o BooleanNode) ToJSON() string {
	if o.Value {
		return "true"
	}
	return "false"
}

type NullNode struct{}

func (o NullNode) ToJSON() string {
	return "null"
}
