package goose

import "fmt"

type Goose struct{}

func NewGoose() *Goose {
	return &Goose{}
}

func (g *Goose) Honk() {
	fmt.Println("Honk")
}
