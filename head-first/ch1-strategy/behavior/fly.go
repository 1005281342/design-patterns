package behavior

import "log"

type FlyBehavior interface {
	Fly()
}

type FlyWithWings struct{}

func (f *FlyWithWings) Fly() {
	log.Println("俺会飞！")
}

type FlyNoWay struct{}

func (f *FlyNoWay) Fly() {
	log.Println("俺不会飞！")
}

type FlyRocketPower struct{}

func (f *FlyRocketPower) Fly() {
	log.Println("俺能飞上天和太阳肩并肩，还和火箭一样快！")
}
