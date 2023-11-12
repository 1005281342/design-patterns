package pizzastore

import (
	"github.com/1005281342/design-patterns/head-first/ch4-factory/abstract-factory/errs"
	"github.com/1005281342/design-patterns/head-first/ch4-factory/abstract-factory/pizza"
)

type ChicagoPizzaStore struct {
	*pizzaStore
}

func NewChicagoPizzaStore() *ChicagoPizzaStore {
	var store = &ChicagoPizzaStore{
		pizzaStore: &pizzaStore{},
	}
	store.createPizza = store.CreatePizza
	return store
}

func (n *ChicagoPizzaStore) CreatePizza(menu string) (pizza.Pizza, error) {
	switch menu {
	case pizza.TypeCheese.String():
		return pizza.NewChicagoStyleCheesePizza(), nil
	}
	return nil, errs.NotFound
}
