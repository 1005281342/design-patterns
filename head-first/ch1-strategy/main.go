package main

import (
	"github.com/1005281342/design-patterns/head-first/ch1-strategy/behavior"
	"github.com/1005281342/design-patterns/head-first/ch1-strategy/duck"
)

func main() {
	// 创建一只会嘎嘎叫且会飞的绿头鸭
	var mallardDuck = duck.NewMallardDuck(&behavior.Quack{}, &behavior.FlyWithWings{})
	mallardDuck.Display()
	mallardDuck.PerformQuack()
	mallardDuck.PerformFly()

	// 改变 quack 行为
	mallardDuck.SetQuackBehavior(&behavior.QuackMute{})
	mallardDuck.PerformQuack()

	// 改变 fly 行为
	mallardDuck.SetFlyBehavior(&behavior.FlyRocketPower{})
	mallardDuck.PerformFly()

	//2023/11/09 00:09:16 看起来像绿头鸭
	//2023/11/09 00:09:16 嘎嘎嘎！
	//2023/11/09 00:09:16 俺会飞！
	//2023/11/09 00:09:16 俺能飞上天和太阳肩并肩，还和火箭一样快！
}
