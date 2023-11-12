package pizzafactory

import (
	"github.com/1005281342/design-patterns/head-first/ch4-factory/abstract-factory/cheese"
	"github.com/1005281342/design-patterns/head-first/ch4-factory/abstract-factory/clams"
	"github.com/1005281342/design-patterns/head-first/ch4-factory/abstract-factory/dough"
	"github.com/1005281342/design-patterns/head-first/ch4-factory/abstract-factory/sauce"
)

type PizzaIngredientFactory interface {
	CreateDough() dough.Dough
	CreateSauce() sauce.Sauce
	CreateCheese() cheese.Cheese
	CreateClam() clams.Clams
}
