package facade

import "log"

type IPopcornPopper interface {
	On()
	Pop()
	Off()
}

type PopcornPopper struct{}

func NewPopper() *PopcornPopper {
	return &PopcornPopper{}
}

func (p *PopcornPopper) On() {
	log.Println("Popcorn Popper on")
}

func (p *PopcornPopper) Off() {
	log.Println("Popcorn Popper off")
}

func (p *PopcornPopper) Pop() {
	log.Println("Popcorn Popper popping popcorn!")
}
