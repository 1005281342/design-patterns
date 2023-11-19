package adapter

import (
	"github.com/1005281342/design-patterns/head-first/ch12-compound/quackable/goose"
	"github.com/1005281342/design-patterns/head-first/ch12-compound/quackable/observable"
)

type GooseAdapter struct {
	goose *goose.Goose

	observable.QuackObservable
}

func NewGooseAdapter(goose *goose.Goose) *GooseAdapter {
	var g = &GooseAdapter{goose: goose}
	g.QuackObservable = observable.NewObservable(g)
	return g
}

func (g *GooseAdapter) Quack() {
	g.goose.Honk()
}

func (g *GooseAdapter) String() string {
	return "GooseAdapter"
}
