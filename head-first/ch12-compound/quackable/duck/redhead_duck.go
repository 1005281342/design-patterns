package duck

import (
	"fmt"

	"github.com/1005281342/design-patterns/head-first/ch12-compound/quackable/observable"
)

type RedheadDuck struct {
	observable.QuackObservable
}

func NewRedheadDuck() *RedheadDuck {
	var r = &RedheadDuck{}
	r.QuackObservable = observable.NewObservable(r)
	return r
}

func (r *RedheadDuck) Quack() {
	fmt.Println(quack)

	r.Notify()
}

func (r *RedheadDuck) String() string {
	return "RedheadDuck"
}
