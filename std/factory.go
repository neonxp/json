package std

import (
	"fmt"
	"strconv"
	"strings"

	"go.neonxp.dev/json"
)

func Factory(typ json.NodeType) (json.Node, error) {
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

type ObjectNode map[string]json.Node

func (o ObjectNode) SetKeyValue(k string, v json.Node) {
	o[k] = v
}

func (o ObjectNode) GetByKey(k string) (json.Node, bool) {
	v, ok := o[k]
	return v, ok
}

func (o ObjectNode) String() string {
	res := make([]string, 0, len(o))
	for k, n := range o {
		res = append(res, fmt.Sprintf(`"%s":%s`, k, n.String()))
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

func (o *ArrayNode) String() string {
	res := make([]string, 0, len(*o))
	for _, v := range *o {
		res = append(res, v.String())
	}
	return fmt.Sprintf(`[%s]`, strings.Join(res, ","))
}

type StringNode struct {
	Value string
}

func (o *StringNode) SetString(v string) {
	o.Value = v
}

func (o *StringNode) GetString() string {
	return o.Value
}

func (o *StringNode) String() string {
	return `"` + o.Value + `"`
}

type NumberNode struct {
	Value float64
}

func (o *NumberNode) SetNumber(v float64) {
	o.Value = v
}

func (o *NumberNode) GetNumber() float64 {
	return o.Value
}

func (o *NumberNode) String() string {
	return strconv.FormatFloat(float64(o.Value), 'g', 15, 64)
}

type BooleanNode struct {
	Value bool
}

func (o *BooleanNode) SetBool(v bool) {
	o.Value = v
}

func (o *BooleanNode) GetBool() bool {
	return o.Value
}

func (o BooleanNode) String() string {
	if o.Value {
		return "true"
	}
	return "false"
}

type NullNode struct{}

func (o NullNode) String() string {
	return "null"
}
