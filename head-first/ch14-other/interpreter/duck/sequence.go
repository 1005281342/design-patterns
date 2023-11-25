package duck

type Sequence struct {
	exp1 Expression
	exp2 Expression
}

func NewSequence(exp1 Expression, exp2 Expression) *Sequence {
	return &Sequence{exp1: exp1, exp2: exp2}
}

func (s *Sequence) Interpret(ctx *Context) string {
	return s.exp1.Interpret(ctx) + "; " + s.exp2.Interpret(ctx)
}
