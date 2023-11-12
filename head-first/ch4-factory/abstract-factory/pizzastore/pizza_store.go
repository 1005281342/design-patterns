package pizzastore

import "github.com/1005281342/design-patterns/head-first/ch4-factory/abstract-factory/pizza"

type PizzaStore interface {
	OrderPizza(menu string) (pizza.Pizza, error)
}

type pizzaStore struct {
	createPizza func(menu string) (pizza.Pizza, error)
}

func (p *pizzaStore) OrderPizza(menu string) (pizza.Pizza, error) {
	var piz, err = p.createPizza(menu)
	if err != nil {
		return nil, err
	}

	piz.Prepare()
	piz.Bake()
	piz.Cut()
	piz.Box()
	return piz, nil
}
