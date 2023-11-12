package command

import "github.com/1005281342/design-patterns/head-first/ch6-command/receiver"

type StereoOffWithCDCommand struct {
	stereo     receiver.Stereo
	prevVolume int
}

func NewStereoOffWithCDCommand(stereo receiver.Stereo) *StereoOffWithCDCommand {
	return &StereoOffWithCDCommand{stereo: stereo}
}

func (s *StereoOffWithCDCommand) Execute() {
	s.prevVolume = s.stereo.GetVolume()
	s.stereo.Off()
}

func (s *StereoOffWithCDCommand) Undo() {
	s.stereo.On()
	s.stereo.SetVolume(s.prevVolume)
}
