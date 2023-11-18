package machine

type GumballMachine struct {
	count    int
	location string
	state    State

	noQuarterState  State
	hasQuarterState State
	soldState       State
	soldOutState    State
}

func NewGumballMachine(count int, location string) *GumballMachine {
	var gm = &GumballMachine{
		count:    count,
		location: location,
	}
	gm.soldOutState = &StateSoldOut{machine: gm}
	gm.noQuarterState = &StateNoQuarter{machine: gm}
	gm.hasQuarterState = &StateHasQuarter{machine: gm}
	gm.soldState = &StateSold{machine: gm}

	gm.state = gm.noQuarterState
	if gm.count <= 0 {
		gm.state = gm.soldOutState
	}

	return gm
}

func (g *GumballMachine) InsertQuarter() {
	g.state.InsertQuarter()
}

func (g *GumballMachine) TurnCrank() {
	g.state.TurnCrank()
	g.state.Dispense()
}

func (g *GumballMachine) EjectQuarter() {
	g.state.EjectQuarter()
}

func (g *GumballMachine) GetCount() int {
	return g.count
}

func (g *GumballMachine) GetLocation() string {
	return g.location
}

func (g *GumballMachine) GetState() string {
	return g.state.String()
}
