package condiment

import "github.com/1005281342/design-patterns/head-first/ch3-decorator/component"

type Mocha struct {
	component.Beverage
}

func AddMocha(beverage component.Beverage) *Mocha {
	return &Mocha{Beverage: beverage}
}

func (m *Mocha) Cost() float32 {
	var cost = m.Beverage.Cost()
	switch m.Beverage.GetSize() {
	case component.SizeTall:
		return cost + 0.2
	case component.SizeGrand:
		return cost + 0.25
	case component.SizeVenti:
		return cost + 0.3
	}
	return cost
}

func (m *Mocha) GetDesc() string {
	return m.Beverage.GetDesc() + ", 摩卡"
}
