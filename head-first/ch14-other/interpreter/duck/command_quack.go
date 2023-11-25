package duck

type CommandQuack struct{}

func (c *CommandQuack) Interpret(ctx *Context) string {
	return "Quacking"
}
