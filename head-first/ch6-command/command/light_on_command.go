package command

import "github.com/1005281342/design-patterns/head-first/ch6-command/receiver"

type LightOnCommand struct {
	light receiver.Light
}

func NewLightOnCommand(light receiver.Light) *LightOnCommand {
	return &LightOnCommand{light: light}
}

func (l *LightOnCommand) Execute() {
	l.light.On()
}

func (l *LightOnCommand) Undo() {
	l.light.Off()
}
