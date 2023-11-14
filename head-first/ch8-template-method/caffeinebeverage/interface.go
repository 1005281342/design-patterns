package caffeinebeverage

type ICaffeineBeverage interface {
	BoilWater()
	Brew()
	PourInCup()
	AddCondiments()
}

type CaffeineBeverage struct {
	beverage ICaffeineBeverage
}

func New(beverage ICaffeineBeverage) *CaffeineBeverage {
	return &CaffeineBeverage{beverage: beverage}
}

func (c *CaffeineBeverage) PrepareRecipe() {
	c.beverage.BoilWater()
	c.beverage.Brew()
	c.beverage.PourInCup()
	c.beverage.AddCondiments()
}
