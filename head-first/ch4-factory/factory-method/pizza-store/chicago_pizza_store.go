package pizzastore

import "github.com/1005281342/design-patterns/head-first/ch4-factory/factory-method/pizza"

type ChicagoPizzaStore struct {
	*pizzaStore
}

func NewChicagoPizzaStore() *ChicagoPizzaStore {
	return &ChicagoPizzaStore{
		pizzaStore: &pizzaStore{
			createPizza: pizza.CreateChicagoStyleFactory,
		},
	}
}
