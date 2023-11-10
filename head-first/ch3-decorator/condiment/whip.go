package condiment

import "github.com/1005281342/design-patterns/head-first/ch3-decorator/component"

type Whip struct {
	component.Beverage
}

func AddWhip(beverage component.Beverage) *Whip {
	return &Whip{Beverage: beverage}
}

func (m *Whip) Cost() float32 {
	var cost = m.Beverage.Cost()
	switch m.Beverage.GetSize() {
	case component.SizeTall:
		return cost + 0.1
	case component.SizeGrand:
		return cost + 0.15
	case component.SizeVenti:
		return cost + 0.2
	}
	return cost
}

func (m *Whip) GetDesc() string {
	return m.Beverage.GetDesc() + ", 奶泡"
}
