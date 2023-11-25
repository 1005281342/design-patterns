package duck

import (
	"math/rand"
	"strings"
)

type Repetition struct {
	variable string
	exp      Expression
}

func NewRepetition(variable string, exp Expression) *Repetition {
	return &Repetition{variable: variable, exp: exp}
}

func (r *Repetition) Interpret(ctx *Context) string {
	const sep = " __ "
	result := ""
	for {
		// Interpret the expression until the variable condition is met
		if checkVariableCondition(ctx, r.variable) {
			result += r.exp.Interpret(ctx) + sep
		} else {
			break
		}
	}
	return strings.TrimRight(result, sep)
}

func checkVariableCondition(ctx *Context, variable string) bool {
	// Implement logic to check variable condition (e.g., whether variable is true or false)
	// Return true or false based on the condition

	return rand.Intn(5) > 1 // example
}
