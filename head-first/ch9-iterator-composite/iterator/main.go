package main

import (
	"github.com/1005281342/design-patterns/head-first/ch9-iterator-composite/iterator/menu"
	"github.com/1005281342/design-patterns/head-first/ch9-iterator-composite/iterator/waitress"
)

func main() {
	var w = waitress.NewWaitress(menu.NewPancakeHouseMenu(), menu.NewDinerMenu())
	w.PrintMenu()
}
