package duck

import (
	"log"

	"github.com/1005281342/design-patterns/head-first/ch1-strategy/behavior"
)

type DecoyDuck struct {
	*baseDuck
}

func NewDecoyDuck(quackBehavior behavior.QuackBehavior, flyBehavior behavior.FlyBehavior) *DecoyDuck {
	return &DecoyDuck{
		baseDuck: newbaseDuck(quackBehavior, flyBehavior),
	}
}

func (d *DecoyDuck) Display() {
	log.Println("看起来像诱饵鸭")
}
