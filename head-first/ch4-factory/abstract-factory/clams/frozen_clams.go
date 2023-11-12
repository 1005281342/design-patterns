package clams

type FrozenClams struct{}

func NewFrozenClams() *FrozenClams {
	return &FrozenClams{}
}

func (f *FrozenClams) Name() string {
	return "冷冻蛤蜊"
}
