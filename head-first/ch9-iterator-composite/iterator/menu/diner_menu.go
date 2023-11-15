package menu

import "log"

const dinerMenuMaxItems = 6

type DinerMenu struct {
	items         [dinerMenuMaxItems]*Item
	numberOffItem int
}

func NewDinerMenu() *DinerMenu {
	var d = &DinerMenu{}
	d.AddItem("Vegetarian BLT", "(Fakin') Bacon with lettuce & tomato on whole wheat", true, 2.99)
	d.AddItem("BLT", "Bacon with lettuce & tomato on whole wheat", false, 2.99)
	d.AddItem("Soup of the day", "Soup of the day, with a side of potato salad", false, 3.29)
	d.AddItem("Hotdog", "A hot dog, with sauerkraut, relish, onions, topped with cheese", false, 2.99)
	return d
}

func (d *DinerMenu) AddItem(name string, desc string, vegetarian bool, price float32) {
	if d.numberOffItem >= dinerMenuMaxItems {
		log.Println("Sorry, menu is full! Can't add item to menu")
		return
	}

	d.items[d.numberOffItem] = NewItem(name, desc, vegetarian, price)
	d.numberOffItem++
}

func (d *DinerMenu) GetItems() [dinerMenuMaxItems]*Item {
	return d.items
}

type DinerMenuIterator struct {
	items    [dinerMenuMaxItems]*Item
	position int
}

func NewDinerMenuIterator(items [dinerMenuMaxItems]*Item) *DinerMenuIterator {
	return &DinerMenuIterator{items: items}
}

func (d *DinerMenuIterator) HasNext() bool {
	if d.position+1 >= dinerMenuMaxItems {
		return false
	}

	if d.items[d.position+1] == nil {
		return false
	}

	return true
}

func (d *DinerMenuIterator) Next() *Item {
	if !d.HasNext() {
		return nil
	}

	d.position++
	return d.items[d.position-1]
}
