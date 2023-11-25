package duck

type CommandRight struct{}

func (c *CommandRight) Interpret(ctx *Context) string {
	return "Moving right"
}
