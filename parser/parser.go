package parser

import (
	"fmt"
	"strconv"
	"strings"

	"go.neonxp.dev/json/model"
)

func Parse(json string) (model.Node, error) {
	l := newLexer(json)
	go l.Run(initJson)
	n, err := parse(l.Output)
	if err != nil {
		return nil, err
	}
	return model.NewNode(n), nil
}

func parse(ch chan lexem) (any, error) {
	prefix := <-ch
	switch prefix.Type {
	case lObjectStart:
		return parseObject(ch)
	case lArrayStart:
		return parseArray(ch)
	case lString:
		return strings.Trim(prefix.Value, `"`), nil
	case lNumber:
		num, err := strconv.ParseFloat(prefix.Value, 64)
		if err != nil {
			return nil, err
		}
		return num, nil
	case lBoolean:
		if strings.ToLower(prefix.Value) == "true" {
			return true, nil
		}
		return false, nil
	case lNull:
		return nil, nil
	}
	return nil, fmt.Errorf("ivalid token: '%s' type=%s", prefix.Value, prefix.Type.String())
}

func parseObject(ch chan lexem) (model.NodeObjectValue, error) {
	m := model.NodeObjectValue{}
	nextKey := ""
	for l := range ch {
		switch l.Type {
		case lObjectKey:
			nextKey = strings.Trim(l.Value, `"`)
		case lString:
			m.Set(nextKey, strings.Trim(l.Value, `"`))
		case lNumber:
			num, err := strconv.ParseFloat(l.Value, 64)
			if err != nil {
				return nil, err
			}
			m.Set(nextKey, num)
		case lBoolean:
			if strings.ToLower(l.Value) == "true" {
				m.Set(nextKey, true)
				continue
			}
			m.Set(nextKey, false)
		case lNull:
			m.Set(nextKey, nil)
		case lObjectStart:
			obj, err := parseObject(ch)
			if err != nil {
				return nil, err
			}
			m.Set(nextKey, obj)
		case lArrayStart:
			arr, err := parseArray(ch)
			if err != nil {
				return nil, err
			}
			m.Set(nextKey, arr)
		case lObjectEnd:
			return m, nil
		}
	}
	return nil, fmt.Errorf("unexpected end of object")
}

func parseArray(ch chan lexem) (model.NodeArrayValue, error) {
	m := model.NodeArrayValue{}
	for l := range ch {
		switch l.Type {
		case lString:
			m = append(m, model.NewNode(strings.Trim(l.Value, `"`)))
		case lNumber:
			num, err := strconv.ParseFloat(l.Value, 64)
			if err != nil {
				return nil, err
			}
			m = append(m, model.NewNode(num))
		case lBoolean:
			if strings.ToLower(l.Value) == "true" {
				m = append(m, model.NewNode(true))
				continue
			}
			m = append(m, model.NewNode(false))
		case lNull:
			m = append(m, model.NewNode(nil))
		case lObjectStart:
			obj, err := parseObject(ch)
			if err != nil {
				return nil, err
			}
			m = append(m, model.NewNode(obj))
		case lArrayStart:
			arr, err := parseArray(ch)
			if err != nil {
				return nil, err
			}
			m = append(m, model.NewNode(arr))
		case lArrayEnd:
			return m, nil
		}
	}
	return nil, fmt.Errorf("unexpected end of object")
}
