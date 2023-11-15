package menu

type PancakeHouseMenu struct {
	items []*Item
}

func NewPancakeHouseMenu() *PancakeHouseMenu {
	var p = &PancakeHouseMenu{}
	p.AddItem("K&B's Pancake Breakfast", "Pancakes with scrambled eggs and toast", true, 2.99)
	p.AddItem("Regular Pancake Breakfast", "Pancakes with fried eggs, sausage", false, 2.99)
	p.AddItem("Blueberry Pancakes", "Pancakes made with fresh blueberries", true, 3.49)
	p.AddItem("Waffles", "Waffles with your choice of blueberries or strawberries", true, 3.59)
	return p
}

func (p *PancakeHouseMenu) AddItem(name string, desc string, vegetarian bool, price float32) {
	p.items = append(p.items, NewItem(name, desc, vegetarian, price))
}

func (p *PancakeHouseMenu) GetItems() []*Item {
	return p.items
}

type PancakeHouseMenuIterator struct {
	items    []*Item
	position int
}

func NewPancakeHouseMenuIterator(items []*Item) *PancakeHouseMenuIterator {
	return &PancakeHouseMenuIterator{items: items}
}

func (p *PancakeHouseMenuIterator) HasNext() bool {
	return len(p.items) > p.position
}

func (p *PancakeHouseMenuIterator) Next() *Item {
	if !p.HasNext() {
		return nil
	}

	p.position++
	return p.items[p.position-1]
}
