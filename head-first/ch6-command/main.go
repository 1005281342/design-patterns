package main

import (
	"github.com/1005281342/design-patterns/head-first/ch6-command/command"
	"github.com/1005281342/design-patterns/head-first/ch6-command/invoker"
	"github.com/1005281342/design-patterns/head-first/ch6-command/receiver"
)

func main() {
	var (
		// 调用者
		remoteCtl = invoker.NewRemoteControl()

		// 接收者
		kitchenLight     = receiver.NewKitchenLight()
		livingRoomLight  = receiver.NewLivingRoomLight()
		livingRoomStereo = receiver.NewLivingRoomStereo()

		// 命令
		kitchenLightOnCommand     = command.NewLightOnCommand(kitchenLight)
		kitchenLightOffCommand    = command.NewLightOffCommand(kitchenLight)
		livingRoomLightOnCommand  = command.NewLightOnCommand(livingRoomLight)
		livingRoomLightOffCommand = command.NewLightOffCommand(livingRoomLight)
		livingRoomOnStereoWithCD  = command.NewStereoOnWithCDCommand(livingRoomStereo)
		livingRoomOffStereoWithCD = command.NewStereoOffWithCDCommand(livingRoomStereo)
	)

	// 命令加载至遥控器槽
	remoteCtl.SetOnCommand(invoker.Slot0, kitchenLightOnCommand)
	remoteCtl.SetOffCommand(invoker.Slot0, kitchenLightOffCommand)
	remoteCtl.SetOnCommand(invoker.Slot1, livingRoomLightOnCommand)
	remoteCtl.SetOffCommand(invoker.Slot1, livingRoomLightOffCommand)
	remoteCtl.SetOnCommand(invoker.Slot2, livingRoomOnStereoWithCD)
	remoteCtl.SetOffCommand(invoker.Slot2, livingRoomOffStereoWithCD)

	// 使用遥控器
	remoteCtl.OnButtonWasPushed(invoker.Slot0)
	remoteCtl.OffButtonWasPushed(invoker.Slot1)
	remoteCtl.UndoButtonWasPushed()
}
