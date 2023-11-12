package receiver

import "log"

type LivingRoomLight struct{}

func NewLivingRoomLight() *LivingRoomLight {
	return &LivingRoomLight{}
}

func (k *LivingRoomLight) On() {
	log.Println("Living room light is on")
}

func (k *LivingRoomLight) Off() {
	log.Println("Living room light is off")
}
