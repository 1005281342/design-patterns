package command

import "github.com/1005281342/design-patterns/head-first/ch6-command/receiver"

type LightOffCommand struct {
	light receiver.Light
}

func NewLightOffCommand(light receiver.Light) *LightOffCommand {
	return &LightOffCommand{light: light}
}

func (l *LightOffCommand) Execute() {
	l.light.Off()
}

func (l *LightOffCommand) Undo() {
	l.light.On()
}
