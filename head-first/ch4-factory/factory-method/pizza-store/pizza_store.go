package pizzastore

import (
	"github.com/1005281342/design-patterns/head-first/ch4-factory/factory-method/pizza"
)

type PizzaStore interface {
	OrderPizza(menu string) (pizza.Pizza, error)
}

type pizzaStore struct {
	createPizza pizza.CreatePizzaFactory
}

func (p *pizzaStore) OrderPizza(menu string) (pizza.Pizza, error) {
	var piz, err = p.createPizza(menu)
	if err != nil {
		return piz, err
	}

	piz.Prepare()
	piz.Bake()
	piz.Cut()
	piz.Box()
	return piz, nil
}
