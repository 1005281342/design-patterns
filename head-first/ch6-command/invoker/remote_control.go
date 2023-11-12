package invoker

import (
	"fmt"
	"log"

	"github.com/1005281342/design-patterns/head-first/ch6-command/command"
)

type Slot int

func (s Slot) Int() int {
	return int(s)
}

const (
	Slot0 Slot = iota
	Slot1
	Slot2
	Slot3
	Slot4
	Slot5
	Slot6
	slots = 7
)

type RemoteControl struct {
	onCommand   [slots]command.Command
	offCommand  [slots]command.Command
	undoCommand command.Command
}

func NewRemoteControl() *RemoteControl {
	var (
		noCommand     = command.NewNoCommand()
		remoteControl = &RemoteControl{
			undoCommand: noCommand,
		}
	)
	for i := 0; i < slots; i++ {
		remoteControl.onCommand[i] = noCommand
		remoteControl.offCommand[i] = noCommand
	}
	return remoteControl
}

func (r *RemoteControl) SetOnCommand(slot Slot, cmd command.Command) {
	log.Printf("Set on command button:%d\n", slot)
	r.onCommand[slot.Int()] = cmd
}

func (r *RemoteControl) SetOffCommand(slot Slot, cmd command.Command) {
	log.Printf("Set off command button:%d\n", slot)
	r.offCommand[slot.Int()] = cmd
}

func (r *RemoteControl) OnButtonWasPushed(slot Slot) {
	fmt.Printf("Push on command button:%d\n", slot)
	var cmd = r.onCommand[slot.Int()]
	cmd.Execute()
	r.undoCommand = cmd
}

func (r *RemoteControl) OffButtonWasPushed(slot Slot) {
	fmt.Printf("Push off command button:%d\n", slot)
	var cmd = r.offCommand[slot.Int()]
	cmd.Execute()
	r.undoCommand = cmd
}

func (r *RemoteControl) UndoButtonWasPushed() {
	r.undoCommand.Undo()
}
