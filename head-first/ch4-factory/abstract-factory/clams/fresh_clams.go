package clams

type FreshClams struct{}

func NewFreshClams() *FreshClams {
	return &FreshClams{}
}

func (f *FreshClams) Name() string {
	return "新鲜蛤蜊"
}
