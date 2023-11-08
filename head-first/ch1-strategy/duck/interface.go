package duck

import "github.com/1005281342/design-patterns/head-first/ch1-strategy/behavior"

type Duck interface {
	Display()
	PerformQuack()
	SetQuackBehavior(quackBehavior behavior.QuackBehavior)
	PerformFly()
	SetFlyBehavior(flyBehavior behavior.FlyBehavior)

	// Swim() 暂时认为所有鸭子都会游泳
}
