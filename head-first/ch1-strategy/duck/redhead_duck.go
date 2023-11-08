package duck

import (
	"log"

	"github.com/1005281342/design-patterns/head-first/ch1-strategy/behavior"
)

type RedHeadDuck struct {
	*baseDuck
}

func NewRedHeadDuck(quackBehavior behavior.QuackBehavior, flyBehavior behavior.FlyBehavior) *RedHeadDuck {
	return &RedHeadDuck{
		baseDuck: newbaseDuck(quackBehavior, flyBehavior),
	}
}

func (d *RedHeadDuck) Display() {
	log.Println("看起来像红头鸭")
}
