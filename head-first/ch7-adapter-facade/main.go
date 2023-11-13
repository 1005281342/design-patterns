package main

import (
	"github.com/1005281342/design-patterns/head-first/ch7-adapter-facade/adapter"
	"github.com/1005281342/design-patterns/head-first/ch7-adapter-facade/facade"
)

func main() {
	testAdapter()

	testFacade()
}

func testAdapter() {
	var (
		mallardDuck   = adapter.NewMallardDuck()
		wildTurkey    = adapter.NewWildTurkey()
		turkeyAdapter = adapter.NewTurkeyAdapter(wildTurkey)
	)

	// 火鸡
	wildTurkey.Gobble()
	wildTurkey.Fly()

	// 测试绿头鸭
	testDuck(mallardDuck)
	// 测试适配成鸭子的火鸡
	testDuck(turkeyAdapter)
}

func testDuck(duck adapter.Duck) {
	duck.Quack()
	duck.Fly()
}

func testFacade() {
	var homeTheater = facade.NewHomeTheaterFacade(
		facade.NewAmplifier(),
		facade.NewStreamingPlayer(),
		facade.NewPopper(),
		facade.NewProjector(),
		facade.NewTheaterLights(),
		facade.NewScreen(),
	)
	homeTheater.WatchMovie("Raiders of the Lost Ark")
	homeTheater.EndMovie()
}
