package duck

import (
	"log"

	"github.com/1005281342/design-patterns/head-first/ch1-strategy/behavior"
)

type RubberDuck struct {
	*baseDuck
}

func NewRubberDuck(quackBehavior behavior.QuackBehavior, flyBehavior behavior.FlyBehavior) *RubberDuck {
	return &RubberDuck{
		baseDuck: newbaseDuck(quackBehavior, flyBehavior),
	}
}

func (d *RubberDuck) Display() {
	log.Println("看起来像橡皮鸭")
}
