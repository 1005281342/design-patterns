package duck

import (
	"log"

	"github.com/1005281342/design-patterns/head-first/ch1-strategy/behavior"
)

type baseDuck struct {
	quackBehavior behavior.QuackBehavior
	flyBehavior   behavior.FlyBehavior
}

func newbaseDuck(quackBehavior behavior.QuackBehavior, flyBehavior behavior.FlyBehavior) *baseDuck {
	return &baseDuck{
		quackBehavior: quackBehavior,
		flyBehavior:   flyBehavior,
	}
}

func (d *baseDuck) Display() {
	log.Println("就是一只普通的鸭子")
}

func (d *baseDuck) PerformQuack() {
	d.quackBehavior.Quack()
}

func (d *baseDuck) SetQuackBehavior(quackBehavior behavior.QuackBehavior) {
	d.quackBehavior = quackBehavior
}

func (d *baseDuck) PerformFly() {
	d.flyBehavior.Fly()
}

func (d *baseDuck) SetFlyBehavior(flyBehavior behavior.FlyBehavior) {
	d.flyBehavior = flyBehavior
}
