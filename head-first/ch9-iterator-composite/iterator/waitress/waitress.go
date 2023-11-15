package waitress

import (
	"log"

	"github.com/1005281342/design-patterns/head-first/ch9-iterator-composite/iterator/menu"
)

type Waitress struct {
	pancakeHouseMenu *menu.PancakeHouseMenu
	dinerMenu        *menu.DinerMenu
}

func NewWaitress(pancakeHouseMenu *menu.PancakeHouseMenu, dinerMenu *menu.DinerMenu) *Waitress {
	return &Waitress{pancakeHouseMenu: pancakeHouseMenu, dinerMenu: dinerMenu}
}

func (w *Waitress) PrintMenu() {
	var pancakeHouseMenuIterator = menu.NewPancakeHouseMenuIterator(w.pancakeHouseMenu.GetItems())
	w.printMenu(pancakeHouseMenuIterator)

	var dinerMenu = menu.NewDinerMenuIterator(w.dinerMenu.GetItems())
	w.printMenu(dinerMenu)
}

func (w *Waitress) printMenu(iterator menu.Iterator) {
	for iterator.HasNext() {
		var item = iterator.Next()
		log.Printf("%s, %f -- %s\n", item.GetName(), item.GetPrice(), item.GetDesc())
	}
}
