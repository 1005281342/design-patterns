package duck

import "fmt"

type Expression interface {
	Interpret(ctx *Context) string
}

func InterpretExpression(ctx *Context, exp Expression) {
	fmt.Println("Interpreted expression:", exp.Interpret(ctx))
}
