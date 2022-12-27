package json

import (
	"fmt"
	"strconv"
	"strings"

	"go.neonxp.dev/json/internal/lexer"
)

func (j *JSON) parse(ch chan lexer.Lexem) (Node, error) {
	prefix := <-ch
	return j.createChild(nil, prefix, ch)
}

func (j *JSON) createChild(parent Node, l lexer.Lexem, ch chan lexer.Lexem) (Node, error) {
	switch l.Type {
	case lexer.LString:
		c, err := j.Factory(StringType)
		if err != nil {
			return nil, err
		}
		if c, ok := c.(AcceptParent); ok {
			c.SetParent(parent)
		}
		child := c.(StringNode)
		child.SetString(strings.Trim(l.Value, `"`))
		return child, nil
	case lexer.LNumber:
		num, err := strconv.ParseFloat(l.Value, 64)
		if err != nil {
			return nil, err
		}
		c, err := j.Factory(NumberType)
		if err != nil {
			return nil, err
		}
		if c, ok := c.(AcceptParent); ok {
			c.SetParent(parent)
		}
		child := c.(NumberNode)
		child.SetNumber(num)
		return child, nil
	case lexer.LBoolean:
		b := strings.ToLower(l.Value) == "true"
		c, err := j.Factory(BooleanType)
		if err != nil {
			return nil, err
		}
		if c, ok := c.(AcceptParent); ok {
			c.SetParent(parent)
		}
		child := c.(BooleanNode)
		child.SetBool(b)
		return child, nil
	case lexer.LObjectStart:
		child, err := j.parseObject(parent, ch)
		if err != nil {
			return nil, err
		}
		return child, nil
	case lexer.LArrayStart:
		child, err := j.parseArray(parent, ch)
		if err != nil {
			return nil, err
		}
		return child, nil
	case lexer.LNull:
		c, err := j.Factory(NullType)
		if err != nil {
			return nil, err
		}
		if c, ok := c.(AcceptParent); ok {
			c.SetParent(parent)
		}
		return c.(NullNode), nil
	default:
		return nil, fmt.Errorf("ivalid token: '%s' type=%s", l.Value, l.Type.String())
	}
}

func (j *JSON) parseObject(parent Node, ch chan lexer.Lexem) (ObjectNode, error) {
	c, err := j.Factory(ObjectType)
	if err != nil {
		return nil, err
	}
	if c, ok := c.(AcceptParent); ok {
		c.SetParent(parent)
	}
	n := c.(ObjectNode)
	nextKey := ""
	for l := range ch {
		switch l.Type {
		case lexer.LObjectKey:
			nextKey = strings.Trim(l.Value, `"`)
		case lexer.LObjectEnd:
			return n, nil
		case lexer.LObjectValue:
			continue
		default:
			child, err := j.createChild(n, l, ch)
			if err != nil {
				return nil, err
			}
			n.SetKeyValue(nextKey, child)
		}
	}
	return nil, fmt.Errorf("unexpected end of object")
}

func (j *JSON) parseArray(parent Node, ch chan lexer.Lexem) (ArrayNode, error) {
	c, err := j.Factory(ArrayType)
	if err != nil {
		return nil, err
	}
	if c, ok := c.(AcceptParent); ok {
		c.SetParent(parent)
	}
	n := c.(ArrayNode)
	for l := range ch {
		switch l.Type {
		case lexer.LArrayEnd:
			return n, nil
		default:
			child, err := j.createChild(n, l, ch)
			if err != nil {
				return nil, err
			}
			n.Append(child)
		}
	}
	return nil, fmt.Errorf("unexpected end of object")
}
