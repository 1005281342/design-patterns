package sample_pizza_factory

type CheesePizza struct {
	*pizza
}

func NewCheesePizza() *CheesePizza {
	return &CheesePizza{
		pizza: &pizza{},
	}
}
