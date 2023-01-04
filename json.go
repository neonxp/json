package json

import (
	"fmt"

	"go.neonxp.dev/json/internal/lexer"
)

type JSON struct {
	Factory Factory
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
	return n.ToJSON()
}

func (j *JSON) Node(value any) (Node, error) {
	switch value := value.(type) {
	case string:
		n, err := j.Factory.Produce(StringType)
		if err != nil {
			return nil, err
		}
		j.Factory.Fill(n, value)
		return n, nil
	case float64:
		n, err := j.Factory.Produce(NumberType)
		if err != nil {
			return nil, err
		}
		j.Factory.Fill(n, value)
		return n, nil
	case int:
		n, err := j.Factory.Produce(NumberType)
		if err != nil {
			return nil, err
		}
		j.Factory.Fill(n, float64(value))
		return n, nil
	case bool:
		n, err := j.Factory.Produce(BooleanType)
		if err != nil {
			return nil, err
		}
		j.Factory.Fill(n, value)
		return n, nil
	case nil:
		return j.Factory.Produce(NullType)
	case map[string]Node:
		n, err := j.Factory.Produce(ObjectType)
		if err != nil {
			return nil, err
		}
		j.Factory.Fill(n, value)
		return n, nil
	case []Node:
		n, err := j.Factory.Produce(ArrayType)
		if err != nil {
			return nil, err
		}
		j.Factory.Fill(n, value)
		return n, nil
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

func New(factory Factory) *JSON {
	return &JSON{
		Factory: factory,
	}
}
