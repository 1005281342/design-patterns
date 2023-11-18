package gumballmachine

import "fmt"

// StateSoldOut 售罄状态
type StateSoldOut struct {
	machine *GumballMachine
}

const msgSoldOut = "对不起，糖果已售罄"

func (s *StateSoldOut) InsertQuarter() {
	fmt.Println(msgSoldOut)
}

func (s *StateSoldOut) EjectQuarter() {
	fmt.Println(msgSoldOut)
}

func (s *StateSoldOut) TurnCrank() {
	fmt.Println(msgSoldOut)
}

func (s *StateSoldOut) Dispense() {
	fmt.Println(msgSoldOut)
}
