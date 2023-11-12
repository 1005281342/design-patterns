package command

type NoCommand struct{}

func NewNoCommand() *NoCommand {
	return &NoCommand{}
}

func (n *NoCommand) Execute() {}

func (n *NoCommand) Undo() {}
