package menu

type Menu interface {
	CrateIterator() Iterator
}

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

func (i *Item) GetVegetarian() bool {
	return i.vegetarian
}

func (i *Item) GetPrice() float32 {
	return i.price
}
