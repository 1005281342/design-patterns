package duck

import (
	"fmt"

	"github.com/1005281342/design-patterns/head-first/ch12-compound/quackable/observable"
)

type RubberDuck struct {
	observable.QuackObservable
}

func NewRubberDuck() *RubberDuck {
	var r = &RubberDuck{}
	r.QuackObservable = observable.NewObservable(r)
	return r
}

func (r *RubberDuck) Quack() {
	fmt.Println("Squeak")

	r.Notify()
}

func (r *RubberDuck) String() string {
	return "RubberDuck"
}
