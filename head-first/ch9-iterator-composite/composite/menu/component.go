package menu

type Component interface {
	GetName() string
	GetDesc() string
	GetPrice() float32
	IsVegetarian() bool

	Print()
	Add(component Component)
	Remove(component Component)
	GetChild(index int) Component
}
