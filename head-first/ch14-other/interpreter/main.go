package main

import "github.com/1005281342/design-patterns/head-first/ch14-other/interpreter/duck"

func main() {
	var (
		rightExp      = &duck.CommandRight{}
		quackExp      = &duck.CommandQuack{}
		sequenceExp   = duck.NewSequence(rightExp, quackExp)
		repetitionExp = duck.NewRepetition("variable", sequenceExp)
		ctx           = duck.NewContext()
	)

	duck.InterpretExpression(ctx, rightExp)
	duck.InterpretExpression(ctx, sequenceExp)
	duck.InterpretExpression(ctx, repetitionExp)
}
