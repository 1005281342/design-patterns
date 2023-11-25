package main

import (
	"github.com/1005281342/design-patterns/head-first/ch14-other/bridge/remotecontrol"
	"github.com/1005281342/design-patterns/head-first/ch14-other/bridge/tv"
)

func main() {
	var sonyRemoteControl = remotecontrol.NewConcreteRemote(tv.NewSony())
	sonyRemoteControl.On()
	defer sonyRemoteControl.Off()

	sonyRemoteControl.SetChannel(100)
	sonyRemoteControl.SetChannel(21)
	sonyRemoteControl.NextChannel()
	sonyRemoteControl.SetChannel(23)
	sonyRemoteControl.PrevChannel()
}
