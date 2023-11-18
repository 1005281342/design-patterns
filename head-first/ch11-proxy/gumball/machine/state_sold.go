package machine

import "fmt"

// StateSold 售出糖果状态
type StateSold struct {
	machine *GumballMachine
}

const msgPleaseWait = "请等待，正在发放糖果"

func (s *StateSold) String() string {
	return "sold"
}

func (s *StateSold) InsertQuarter() {
	fmt.Println(msgPleaseWait)
}

func (s *StateSold) EjectQuarter() {
	fmt.Println("对不起，已经转动曲柄，无法退回 25 分钱")
}

func (s *StateSold) TurnCrank() {
	fmt.Println(msgPleaseWait)
}

func (s *StateSold) Dispense() {
	s.machine.count--
	fmt.Println("糖果已发放")

	if s.machine.count == 0 {
		fmt.Println("糖果已售罄")
		s.machine.state = s.machine.soldOutState
	} else {
		s.machine.state = s.machine.noQuarterState
	}
}
