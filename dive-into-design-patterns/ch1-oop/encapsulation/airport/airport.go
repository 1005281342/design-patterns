package airport

import "github.com/1005281342/design-patterns/dive-into-design-patterns/ch1-oop/encapsulation/flyingtransport"

type Airport struct{}

func NewAirport() *Airport {
	return &Airport{}
}

func (a *Airport) Accept(vehicle flyingtransport.FlyingTransport) {}
