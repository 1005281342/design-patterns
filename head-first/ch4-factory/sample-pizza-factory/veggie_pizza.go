package sample_pizza_factory

type VeggiePizza struct {
	*pizza
}

func NewVeggiePizza() *VeggiePizza {
	return &VeggiePizza{
		pizza: &pizza{},
	}
}
