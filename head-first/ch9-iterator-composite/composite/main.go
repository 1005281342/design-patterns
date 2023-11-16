package main

import (
	"github.com/1005281342/design-patterns/head-first/ch9-iterator-composite/composite/menu"
	"github.com/1005281342/design-patterns/head-first/ch9-iterator-composite/composite/waitress"
)

func main() {
	var (
		pancakeHouseMenu = menu.NewMenu("PANCAKE HOUSE MENU", "Breakfast")
		dinerMenu        = menu.NewMenu("DINER MENU", "Lunch")
		cafeMenu         = menu.NewMenu("CAFE MENU", "Dinner")
		menus            = menu.NewMenu("ALL MENUS", "All menus combined")
	)
	menus.Add(pancakeHouseMenu)
	menus.Add(dinerMenu)
	menus.Add(cafeMenu)

	dinerMenu.Add(menu.NewItem(
		"Pasta",
		"Spaghetti with Marinara Sauce, and a slice of sourdough bread",
		true,
		3.89,
	))
	dinerMenu.Add(menu.NewItem(
		"Apple pie",
		"Apple pie with a flakey crust, topped with vanilla ice cream",
		true,
		1.59,
	))

	var w = waitress.NewWaitress(menus)
	w.Print()
}
