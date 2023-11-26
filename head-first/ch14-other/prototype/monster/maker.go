package monster

import "math/rand"

var registry = [...]Monster{&WellKnown{}, &DynamicPlayerGenerated{}}

type Maker struct{}

func (m *Maker) MakeRandomMonster() Monster {
	var idx = rand.Intn(len(registry))
	return registry[idx].clone()
}
