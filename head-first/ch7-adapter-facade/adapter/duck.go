package adapter

import "log"

type Duck interface {
	Quack()
	Fly()
}

var _ Duck = (*MallardDuck)(nil)

type MallardDuck struct{}

func NewMallardDuck() *MallardDuck {
	return &MallardDuck{}
}

func (m *MallardDuck) Quack() {
	log.Println("Quack")
}
func (m *MallardDuck) Fly() {
	log.Println("I'm flying")
}
