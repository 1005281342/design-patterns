package sample_pizza_factory

type SimplePizzaFactory struct{}

func (s *SimplePizzaFactory) CreatePizza(typ Type) Pizza {
	switch typ {
	case TypeCheese:
		return NewCheesePizza()
	case TypeVeggie:
		return NewVeggiePizza()
	case TypeClam:
		return NewClamPizza()
	case TypePepperoni:
		return NewPepperoniPizza()
	}

	return &pizza{}
}
