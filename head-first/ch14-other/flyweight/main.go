package main

import "github.com/1005281342/design-patterns/head-first/ch14-other/flyweight/forest"

func main() {
	f := forest.NewForest()

	f.PlantTree(10, 20, 5, "Pine", "Green")
	f.PlantTree(30, 40, 8, "Maple", "Red")
	f.PlantTree(50, 60, 6, "Pine", "Green")

	f.Display()
}
