package receiver

import "log"

type KitchenLight struct{}

func NewKitchenLight() *KitchenLight {
	return &KitchenLight{}
}

func (k *KitchenLight) On() {
	log.Println("Kitchen light is on")
}

func (k *KitchenLight) Off() {
	log.Println("Kitchen light is off")
}
