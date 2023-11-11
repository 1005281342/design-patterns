package sample_pizza_factory

type PepperoniPizza struct {
	*pizza
}

func NewPepperoniPizza() *PepperoniPizza {
	return &PepperoniPizza{
		pizza: &pizza{},
	}
}
