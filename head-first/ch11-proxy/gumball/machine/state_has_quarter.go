package machine

import "fmt"

// StateHasQuarter 有 25 分钱状态
type StateHasQuarter struct {
	machine *GumballMachine
}

func (h *StateHasQuarter) String() string {
	return "has quarter"
}

func (h *StateHasQuarter) InsertQuarter() {
	fmt.Println("你已经投入了 25 分钱，无需再投")
}

func (h *StateHasQuarter) EjectQuarter() {
	fmt.Println("25 分钱已退回")
	h.machine.state = h.machine.noQuarterState
}

func (h *StateHasQuarter) TurnCrank() {
	fmt.Println("转动曲柄...")
	h.machine.state = h.machine.soldState
}

func (h *StateHasQuarter) Dispense() {
	fmt.Println("请先转动曲柄")
}
