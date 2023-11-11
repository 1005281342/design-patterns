package pizzastore

import "github.com/1005281342/design-patterns/head-first/ch4-factory/factory-method/pizza"

type NYPizzaStore struct {
	*pizzaStore
}

func NewNYPizzaStore() *NYPizzaStore {
	return &NYPizzaStore{
		pizzaStore: &pizzaStore{
			createPizza: pizza.CreateNYStyleFactory,
		},
	}
}
