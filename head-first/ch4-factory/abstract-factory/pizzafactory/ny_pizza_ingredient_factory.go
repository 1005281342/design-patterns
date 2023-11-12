package pizzafactory

import (
	"github.com/1005281342/design-patterns/head-first/ch4-factory/abstract-factory/cheese"
	"github.com/1005281342/design-patterns/head-first/ch4-factory/abstract-factory/clams"
	"github.com/1005281342/design-patterns/head-first/ch4-factory/abstract-factory/dough"
	"github.com/1005281342/design-patterns/head-first/ch4-factory/abstract-factory/sauce"
)

type NYPizzaIngredientFactory struct{}

func (n *NYPizzaIngredientFactory) CreateDough() dough.Dough {
	return dough.NewThinDough()
}

func (n *NYPizzaIngredientFactory) CreateSauce() sauce.Sauce {
	return sauce.NewMarinaraSauce()
}

func (n *NYPizzaIngredientFactory) CreateCheese() cheese.Cheese {
	return cheese.NewReggianoCheese()
}

func (n *NYPizzaIngredientFactory) CreateClam() clams.Clams {
	return clams.NewFreshClams()
}
