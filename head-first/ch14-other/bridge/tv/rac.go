package tv

import "fmt"

type RAC struct{}

func NewRAC() *RAC {
	return &RAC{}
}

func (r *RAC) On() {
	fmt.Println("RAC TV on")
}

func (r *RAC) Off() {
	fmt.Println("RAC TV off")
}

func (r *RAC) TuneChannel(index uint) {
	fmt.Printf("RAC TV tune channel to %d\n", index)
}
