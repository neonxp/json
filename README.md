# JSON parsing library

Библиотека для разбора JSON в дерево объектов. Так же позволяет выполнять поисковые запросы над ними.

## Использование

```go
import "go.neonxp.dev/json"

jsonString := `{
    "string key": "string value",
    "number key": 123.321,
    "bool key": true,
    "object": {
        "one": "two",
        "object 2": {
            "three": "four"
        }
    },
    "array": [
        "one",
        2,
        true,
        null,
        {
            "five": "six"
        }
    ]
}`

j := json.New(std.Factory) // в качестве фабрики можно передавать имплементацию интерфейса NodeFactory
rootNode, err := j.Unmarshal(jsonString)

// Запрос по получившемуся дереву узлов
found := json.MustQuery(rootNode, []string{ "array", "4", "five" }) // == six
```

В результате `rootNode` будет содержать:

```go
std.ObjectNode{
    "string key": &std.StringNode{ "string value" },
    "number key": &std.NumberNode{ 123.321 },
    "bool key":   &std.BoolNode{ true },
    "object": std.ObjectNode{
        "one": &std.StringNode{ "two" },
        "object 2": std.ObjectNode{
            "three": &std.StringNode{ "four" },
        },
    },
    "array": &std.ArrayNode{
        &std.StringNode{ "one" },
        &std.NumberNode{ 2 },
        &std.BoolNode{ true },
        &std.NullNode{},
        std.ObjectNode{
            "five": &std.StringNode{ "six" },
        },
    },
},
```

## Своя фабрика

```
type Factory interface {
    // Непосредственно фабрика возвращающая заготовки нужного типа
	Produce(typ NodeType) (Node, error)
    // Заполнение заготовки определенного типа соответствующим ей значением
	Fill(n Node, value any)
}

type Node interface {
    // Возвращает json представление узла
	ToJSON() string
}

type ObjectNode interface {
	Node
	Set(k string, v Node)
	Get(k string) (Node, bool)
}

type ArrayNode interface {
	Node
	Append(v Node)
	Index(i int) Node
	Len() int
}

type StringNode interface {
	Node
}

type NumberNode interface {
	Node
}

type BooleanNode interface {
	Node
}

type NullNode interface {
	Node
}

// Если узел имплементирует этот интерфейс то вызывается метод Parent передающий родительский узел
type AcceptParent interface {
	SetParent(n Node)
}
```

[Пример реализации самой простой фабрики и типов](/std/factory.go)