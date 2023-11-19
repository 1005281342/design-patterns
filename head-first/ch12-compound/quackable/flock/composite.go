package flock

import (
	"github.com/1005281342/design-patterns/head-first/ch12-compound/quackable"
	"github.com/1005281342/design-patterns/head-first/ch12-compound/quackable/observable"
)

type Flock struct {
	quackers []quackable.Quackable

	observable.QuackObservable
}

func NewFlock() *Flock {
	var f = &Flock{}
	f.QuackObservable = observable.NewObservable(f)
	return f
}

func (f *Flock) Add(quacker quackable.Quackable) {
	f.quackers = append(f.quackers, quacker)
}

func (f *Flock) Quack() {
	for _, quacker := range f.quackers {
		quacker.Quack()
	}
}

func (f *Flock) Register(observer observable.Observer) {
	for _, quacker := range f.quackers {
		quacker.Register(observer)
	}
}
