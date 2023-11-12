package receiver

import "log"

type LivingRoomStereo struct {
	volume int
}

func NewLivingRoomStereo() *LivingRoomStereo {
	return &LivingRoomStereo{}
}

func (l *LivingRoomStereo) On() {
	log.Println("Living room stereo is on")
}

func (l *LivingRoomStereo) Off() {
	log.Println("Living room stereo is off")
}

func (l *LivingRoomStereo) SetCD() {
	log.Println("Living room stereo is set for CD input")
}

func (l *LivingRoomStereo) SetDVD() {
	log.Println("Living room stereo is set for DVD input")
}

func (l *LivingRoomStereo) SetRadio() {
	log.Println("Living room stereo is set for Radio input")
}

func (l *LivingRoomStereo) SetVolume(value int) {
	l.volume = value
	log.Printf("Living room stereo volume set to %d\n", l.volume)
}

func (l *LivingRoomStereo) GetVolume() int {
	return l.volume
}
