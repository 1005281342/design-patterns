package pizza

import "log"

type ChicagoStyleCheesePizza struct {
	*pizza
}

func NewChicagoStyleCheesePizza() *ChicagoStyleCheesePizza {
	return &ChicagoStyleCheesePizza{pizza: &pizza{
		name:     "Chicago Style Deep Dish Cheese Pizza",
		dough:    "Extra Thick Crust Dough",
		sauce:    "Plum Tomato Sauce",
		toppings: []string{"Shredded Mozzarella Cheese"},
	}}
}

func (c *ChicagoStyleCheesePizza) Cut() {
	log.Println("Cutting the Pizza into square slices")
}
