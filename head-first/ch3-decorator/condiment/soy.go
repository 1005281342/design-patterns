package condiment

import "github.com/1005281342/design-patterns/head-first/ch3-decorator/component"

type Soy struct {
	component.Beverage
}

func AddSoy(beverage component.Beverage) *Soy {
	return &Soy{Beverage: beverage}
}

func (m *Soy) Cost() float32 {
	var cost = m.Beverage.Cost()
	switch m.Beverage.GetSize() {
	case component.SizeTall:
		return cost + 0.15
	case component.SizeGrand:
		return cost + 0.2
	case component.SizeVenti:
		return cost + 0.25
	}
	return cost
}

func (m *Soy) GetDesc() string {
	return m.Beverage.GetDesc() + ", 豆奶"
}
