package tv

import "fmt"

type Sony struct{}

func NewSony() *Sony {
	return &Sony{}
}

func (s *Sony) On() {
	fmt.Println("Sony TV on")
}

func (s *Sony) Off() {
	fmt.Println("Sony TV off")
}

func (s *Sony) TuneChannel(index uint) {
	fmt.Printf("Sony TV tune channel to %d\n", index)
}
