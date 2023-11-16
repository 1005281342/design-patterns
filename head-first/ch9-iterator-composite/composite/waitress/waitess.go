package waitress

import "github.com/1005281342/design-patterns/head-first/ch9-iterator-composite/composite/menu"

type Waitress struct {
	menus menu.Component
}

func NewWaitress(menus menu.Component) *Waitress {
	return &Waitress{menus: menus}
}

func (w *Waitress) Print() {
	w.menus.Print()
}
