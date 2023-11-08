package duck

import (
	"log"

	"github.com/1005281342/design-patterns/head-first/ch1-strategy/behavior"
)

type MallardDuck struct {
	*baseDuck
}

func NewMallardDuck(quackBehavior behavior.QuackBehavior, flyBehavior behavior.FlyBehavior) *MallardDuck {
	return &MallardDuck{
		baseDuck: newbaseDuck(quackBehavior, flyBehavior),
	}
}

func (d *MallardDuck) Display() {
	log.Println("看起来像绿头鸭")
}
