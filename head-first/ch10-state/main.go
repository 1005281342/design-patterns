package main

import (
	"fmt"

	"github.com/1005281342/design-patterns/head-first/ch10-state/gumballmachine"
)

func main() {
	// 创建糖果机并初始化状态
	gumballMachine := gumballmachine.NewGumballMachine(5)

	// 模拟操作糖果机
	gumballMachine.InsertQuarter()
	gumballMachine.TurnCrank()
	fmt.Println()

	gumballMachine.InsertQuarter()
	gumballMachine.EjectQuarter()
	fmt.Println()

	gumballMachine.InsertQuarter()
	gumballMachine.TurnCrank()
	gumballMachine.TurnCrank()
	fmt.Println()

	gumballMachine.InsertQuarter()
	gumballMachine.TurnCrank()
	gumballMachine.EjectQuarter()

	//output:
	//你投入了 25 分钱
	//转动曲柄...
	//糖果已发放
	//
	//你投入了 25 分钱
	//25 分钱已退回
	//
	//你投入了 25 分钱
	//转动曲柄...
	//糖果已发放
	//请先投入 25 分钱
	//请先投入 25 分钱
	//
	//你投入了 25 分钱
	//转动曲柄...
	//糖果已发放
	//你没有投入 25 分钱
}
