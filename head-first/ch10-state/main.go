package main

import "github.com/1005281342/design-patterns/head-first/ch10-state/gumballmachine"

func main() {
	// 创建糖果机并初始化状态
	gumballMachine := gumballmachine.NewGumballMachine(5)

	// 模拟操作糖果机
	gumballMachine.InsertQuarter()
	gumballMachine.TurnCrank()
	gumballMachine.InsertQuarter()
	gumballMachine.EjectQuarter()
	gumballMachine.InsertQuarter()
	gumballMachine.TurnCrank()
	gumballMachine.InsertQuarter()
	gumballMachine.TurnCrank()
	//output:
	//你投入了 25 分钱
	//转动曲柄...
	//请等待，正在发放糖果
	//对不起，已经转动曲柄，无法退回 25 分钱
	//请等待，正在发放糖果
	//请等待，正在发放糖果
	//请等待，正在发放糖果
	//请等待，正在发放糖果
}
