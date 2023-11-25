package duck

type CommandFly struct{}

func (c *CommandFly) Interpret(ctx *Context) string {
	return "Flying"
}
