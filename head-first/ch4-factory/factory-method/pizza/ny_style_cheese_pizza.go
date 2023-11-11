package pizza

type NYStyleCheesePizza struct {
	*pizza
}

func NewNYStyleCheesePizza() *NYStyleCheesePizza {
	return &NYStyleCheesePizza{
		pizza: &pizza{
			name:     "NY Style Sauce and Cheese Pizza",
			dough:    "Thin Crust Dough",
			sauce:    "Marinara Sauce",
			toppings: []string{"Grated Reggiano Cheese"},
		},
	}
}
