package sample_pizza_factory

type ClamPizza struct {
	*pizza
}

func NewClamPizza() *ClamPizza {
	return &ClamPizza{
		pizza: &pizza{},
	}
}
