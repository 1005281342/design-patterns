package receiver

type Stereo interface {
	On()
	Off()
	SetCD()
	SetDVD()
	SetRadio()
	SetVolume(int)
	GetVolume() int
}
