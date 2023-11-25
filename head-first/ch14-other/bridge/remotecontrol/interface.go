package remotecontrol

type RemoteControl interface {
	On()
	Off()
	SetChannel(uint)
}
