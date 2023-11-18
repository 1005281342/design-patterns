package main

import (
	"github.com/1005281342/design-patterns/head-first/ch11-proxy/gumball/machine"
	"github.com/1005281342/design-patterns/head-first/ch11-proxy/gumball/monitor"
)

func main() {
	var (
		szGumballMachine = machine.NewGumballMachine(5, "SZ")
		szGumballMonitor = monitor.NewGumballMonitor(szGumballMachine)
	)
	szGumballMonitor.Report()
}
