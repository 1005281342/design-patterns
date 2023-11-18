package monitor

import (
	"fmt"
	"github.com/1005281342/design-patterns/head-first/ch11-proxy/gumball/machine"
)

type GumballMonitor struct {
	gumballMachine *machine.GumballMachine
}

func NewGumballMonitor(gumballMachine *machine.GumballMachine) *GumballMonitor {
	return &GumballMonitor{gumballMachine: gumballMachine}
}

func (g *GumballMonitor) Report() {
	fmt.Printf("GumballMachine machine %s\n", g.gumballMachine.GetLocation())
	fmt.Printf("Current inventory %d\n", g.gumballMachine.GetCount())
	fmt.Printf("Current state %s\n", g.gumballMachine.GetState())
}
