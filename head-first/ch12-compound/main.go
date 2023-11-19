package main

import (
	"fmt"

	"github.com/1005281342/design-patterns/head-first/ch12-compound/quackable/adapter"
	"github.com/1005281342/design-patterns/head-first/ch12-compound/quackable/decorator"
	"github.com/1005281342/design-patterns/head-first/ch12-compound/quackable/factory"
	"github.com/1005281342/design-patterns/head-first/ch12-compound/quackable/flock"
	"github.com/1005281342/design-patterns/head-first/ch12-compound/quackable/goose"
	"github.com/1005281342/design-patterns/head-first/ch12-compound/quackable/observable"
)

func main() {
	exampleQuacked()
	fmt.Println("----------")
	exampleQuackologist()
}

func exampleQuacked() {
	var duckFactory = factory.NewCountingDuckFactory()

	var flockDucks = flock.NewFlock()
	flockDucks.Add(duckFactory.CreateMallardDuck())
	flockDucks.Add(duckFactory.CreateRedheadDuck())
	flockDucks.Add(duckFactory.CreateDuckCall())
	flockDucks.Add(duckFactory.CreateRubberDuck())
	flockDucks.Add(adapter.NewGooseAdapter(goose.NewGoose()))

	var flockMallardDucks = flock.NewFlock()
	flockMallardDucks.Add(duckFactory.CreateMallardDuck())
	flockMallardDucks.Add(duckFactory.CreateMallardDuck())
	flockMallardDucks.Add(duckFactory.CreateMallardDuck())
	flockMallardDucks.Add(duckFactory.CreateMallardDuck())

	flockDucks.Add(flockMallardDucks)

	flockDucks.Quack()

	fmt.Printf("The ducks quacked %d times\n", decorator.GetQuacks())
}

func exampleQuackologist() {
	var duckFactory = factory.NewDuckFactory()

	var flockDucks = flock.NewFlock()
	flockDucks.Add(duckFactory.CreateMallardDuck())
	flockDucks.Add(duckFactory.CreateRedheadDuck())
	flockDucks.Add(duckFactory.CreateDuckCall())
	flockDucks.Add(duckFactory.CreateRubberDuck())
	flockDucks.Add(adapter.NewGooseAdapter(goose.NewGoose()))

	var flockMallardDucks = flock.NewFlock()
	flockMallardDucks.Add(duckFactory.CreateMallardDuck())
	flockMallardDucks.Add(duckFactory.CreateMallardDuck())
	flockMallardDucks.Add(duckFactory.CreateMallardDuck())
	flockMallardDucks.Add(duckFactory.CreateMallardDuck())

	flockDucks.Add(flockMallardDucks)

	var quackologist = observable.NewQuackologist()
	// 设置 quackologist 为观察者
	flockDucks.Register(quackologist)

	flockDucks.Quack()
}
