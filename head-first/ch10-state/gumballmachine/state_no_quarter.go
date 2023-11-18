package gumballmachine

import "fmt"

// StateNoQuarter 没有 25 分钱状态
type StateNoQuarter struct {
	machine *GumballMachine
}

const msgPleaseInsertQuarter = "请先投入 25 分钱"

func (n *StateNoQuarter) InsertQuarter() {
	fmt.Println("你投入了 25 分钱")
	n.machine.state = n.machine.hasQuarterState
}

func (n *StateNoQuarter) EjectQuarter() {
	fmt.Println("你没有投入 25 分钱")
}

func (n *StateNoQuarter) TurnCrank() {
	fmt.Println(msgPleaseInsertQuarter)
}

func (n *StateNoQuarter) Dispense() {
	fmt.Println(msgPleaseInsertQuarter)
}
