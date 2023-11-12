package command

import "github.com/1005281342/design-patterns/head-first/ch6-command/receiver"

type StereoOnWithCDCommand struct {
	stereo     receiver.Stereo
	prevVolume int
}

func NewStereoOnWithCDCommand(stereo receiver.Stereo) *StereoOnWithCDCommand {
	return &StereoOnWithCDCommand{stereo: stereo}
}

func (s *StereoOnWithCDCommand) Execute() {
	s.stereo.On()
	s.stereo.SetCD()

	s.prevVolume = s.stereo.GetVolume()
	s.stereo.SetVolume(11)
}

func (s *StereoOnWithCDCommand) Undo() {
	s.stereo.SetVolume(s.prevVolume)

	s.stereo.Off()
}
