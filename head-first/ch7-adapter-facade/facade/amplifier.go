package facade

import "log"

type IAmplifier interface {
	On()
	Off()
	SetStreamingPlayer(player IStreamingPlayer)
	SetVolume(value int)
	SetSurroundSound()
}

type Amplifier struct {
	streamingPlayer IStreamingPlayer
}

func NewAmplifier() *Amplifier {
	return &Amplifier{}
}

func (a *Amplifier) On() {
	log.Println("Amplifier on")
}

func (a *Amplifier) Off() {
	log.Println("Amplifier off")
}

func (a *Amplifier) SetStreamingPlayer(player IStreamingPlayer) {
	log.Println("Amplifier setting Streaming player to Streaming Player")
}

func (a *Amplifier) SetVolume(value int) {
	log.Printf("Amplifier set volume to %d\n", value)
}

func (a *Amplifier) SetSurroundSound() {
	log.Println("Amplifier surround sound on (5 speakers, 1 subwoofer)")
}
