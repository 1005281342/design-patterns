package dough

type ThinDough struct{}

func NewThinDough() *ThinDough {
	return &ThinDough{}
}

func (t *ThinDough) Name() string {
	return "薄饼面团"
}
