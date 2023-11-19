package duck

import (
	"fmt"

	"github.com/1005281342/design-patterns/head-first/ch12-compound/quackable/observable"
)

type MallardDuck struct {
	observable.QuackObservable
}

func NewMallardDuck() *MallardDuck {
	var m = &MallardDuck{}
	m.QuackObservable = observable.NewObservable(m)
	return m
}

func (m *MallardDuck) Quack() {
	fmt.Println(quack)

	m.Notify()
}

func (m *MallardDuck) String() string {
	return "MallardDuck"
}
