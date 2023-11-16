package menu

import (
	"log"
)

type Item struct {
	name       string
	desc       string
	vegetarian bool
	price      float32
}

func NewItem(name string, desc string, vegetarian bool, price float32) *Item {
	return &Item{
		name:       name,
		desc:       desc,
		vegetarian: vegetarian,
		price:      price,
	}
}

func (i *Item) GetName() string {
	return i.name
}

func (i *Item) GetDesc() string {
	return i.desc
}

func (i *Item) IsVegetarian() bool {
	return i.vegetarian
}

func (i *Item) GetPrice() float32 {
	return i.price
}

func (i *Item) Print() {
	log.Printf("Item: %+v\n", i)
}

func (i *Item) Add(_ Component) {}

func (i *Item) Remove(_ Component) {}

func (i *Item) GetChild(_ int) Component {
	return nil
}
