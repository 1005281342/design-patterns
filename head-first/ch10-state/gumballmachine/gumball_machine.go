package gumballmachine

type GumballMachine struct {
	count int
	state State

	noQuarterState  State
	hasQuarterState State
	soldState       State
	soldOutState    State
}

func NewGumballMachine(count int) *GumballMachine {
	var gm = &GumballMachine{count: count}
	gm.soldOutState = &StateSoldOut{machine: gm}
	gm.noQuarterState = &StateNoQuarter{machine: gm}
	gm.hasQuarterState = &StateHasQuarter{machine: gm}
	gm.soldState = &StateSold{machine: gm}
	gm.state = gm.noQuarterState
	return gm
}

func (g *GumballMachine) InsertQuarter() {
	g.state.InsertQuarter()
}

func (g *GumballMachine) TurnCrank() {
	g.state.TurnCrank()
}

func (g *GumballMachine) EjectQuarter() {
	g.state.EjectQuarter()
}
