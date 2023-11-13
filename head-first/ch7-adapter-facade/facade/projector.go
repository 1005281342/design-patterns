package facade

import "log"

type IProjector interface {
	On()
	Off()
	WideScreenMode()
}

type Projector struct{}

func NewProjector() *Projector {
	return &Projector{}
}

func (p *Projector) On() {
	log.Println("Projector on")
}

func (p *Projector) Off() {
	log.Println("Projector off")
}

func (p *Projector) WideScreenMode() {
	log.Println("Projector in widescreen mode (16x9 aspect ratio)")
}
