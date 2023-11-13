package facade

import "log"

type IScreen interface {
	Up()
	Down()
}

type Screen struct{}

func NewScreen() *Screen {
	return &Screen{}
}

func (s *Screen) Up() {
	log.Println("Theater Screen going up")
}

func (s *Screen) Down() {
	log.Println("Theater Screen going down")
}
