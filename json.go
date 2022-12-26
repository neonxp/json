package json

import (
	"fmt"

	"go.neonxp.dev/json/internal/lexer"
)

type JSON struct {
	Factory NodeFactory
}

func (j *JSON) Unmarshal(input string) (Node, error) {
	lex := lexer.NewLexer(input)
	go lex.Run(lexer.InitJson)
	return j.parse(lex.Output)
}

func (j *JSON) MustUnmarshal(input string) Node {
	n, err := j.Unmarshal(input)
	if err != nil {
		panic(err)
	}
	return n
}

func (j *JSON) Marshal(n Node) string {
	return n.String()
}

func (j *JSON) Node(value any) (Node, error) {
	switch value := value.(type) {
	case string:
		n, err := j.Factory(StringType)
		if err != nil {
			return nil, err
		}
		n.(StringNode).SetString(value)
		return n, nil
	case float64:
		n, err := j.Factory(NumberType)
		if err != nil {
			return nil, err
		}
		n.(NumberNode).SetNumber(value)
		return n, nil
	case int:
		n, err := j.Factory(NumberType)
		if err != nil {
			return nil, err
		}
		n.(NumberNode).SetNumber(float64(value))
		return n, nil
	case bool:
		n, err := j.Factory(BooleanType)
		if err != nil {
			return nil, err
		}
		n.(BooleanNode).SetBool(value)
		return n, nil
	case nil:
		return j.Factory(NullType)
	case map[string]Node:
		n, err := j.Factory(ObjectType)
		if err != nil {
			return nil, err
		}
		on := n.(ObjectNode)
		for k, v := range value {
			on.SetKeyValue(k, v)
		}
		return on, nil
	case []Node:
		n, err := j.Factory(ArrayType)
		if err != nil {
			return nil, err
		}
		an := n.(ArrayNode)
		for _, v := range value {
			an.Append(v)
		}
		return an, nil
	default:
		return nil, fmt.Errorf("invalid type %t", value)
	}
}

func (j *JSON) MustNode(value any) Node {
	n, err := j.Node(value)
	if err != nil {
		panic(err)
	}
	return n
}

func New(factory NodeFactory) *JSON {
	return &JSON{
		Factory: factory,
	}
}
