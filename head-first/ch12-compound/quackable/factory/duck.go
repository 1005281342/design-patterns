package factory

import (
	"github.com/1005281342/design-patterns/head-first/ch12-compound/quackable"
	"github.com/1005281342/design-patterns/head-first/ch12-compound/quackable/decorator"
	"github.com/1005281342/design-patterns/head-first/ch12-compound/quackable/duck"
)

type AbstractDuckFactory interface {
	CreateMallardDuck() quackable.Quackable
	CreateRedheadDuck() quackable.Quackable
	CreateDuckCall() quackable.Quackable
	CreateRubberDuck() quackable.Quackable
}

type DuckFactory struct{}

func NewDuckFactory() *DuckFactory {
	return &DuckFactory{}
}

func (d *DuckFactory) CreateMallardDuck() quackable.Quackable {
	return duck.NewMallardDuck()
}

func (d *DuckFactory) CreateRedheadDuck() quackable.Quackable {
	return duck.NewRedheadDuck()
}

func (d *DuckFactory) CreateDuckCall() quackable.Quackable {
	return duck.NewDuckCall()
}

func (d *DuckFactory) CreateRubberDuck() quackable.Quackable {
	return duck.NewRubberDuck()
}

type CountingDuckFactory struct{}

func NewCountingDuckFactory() *CountingDuckFactory {
	return &CountingDuckFactory{}
}

func (c *CountingDuckFactory) CreateMallardDuck() quackable.Quackable {
	return decorator.NewQuackCounter(duck.NewMallardDuck())
}

func (c *CountingDuckFactory) CreateRedheadDuck() quackable.Quackable {
	return decorator.NewQuackCounter(duck.NewRedheadDuck())
}

func (c *CountingDuckFactory) CreateDuckCall() quackable.Quackable {
	return decorator.NewQuackCounter(duck.NewDuckCall())
}

func (c *CountingDuckFactory) CreateRubberDuck() quackable.Quackable {
	return decorator.NewQuackCounter(duck.NewRedheadDuck())
}
