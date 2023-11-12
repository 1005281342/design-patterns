package pizzastore

import (
	"github.com/1005281342/design-patterns/head-first/ch4-factory/abstract-factory/errs"
	"github.com/1005281342/design-patterns/head-first/ch4-factory/abstract-factory/pizza"
)

type NYPizzaStore struct {
	*pizzaStore
}

func NewNYPizzaStore() *NYPizzaStore {
	var store = &NYPizzaStore{
		pizzaStore: &pizzaStore{},
	}
	store.createPizza = store.CreatePizza
	return store
}

func (n *NYPizzaStore) CreatePizza(menu string) (pizza.Pizza, error) {
	switch menu {
	case pizza.TypeCheese.String():
		return pizza.NewNYStyleCheesePizza(), nil
	}
	return nil, errs.NotFound
}
