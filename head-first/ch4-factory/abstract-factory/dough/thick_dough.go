package dough

type ThickDough struct{}

func NewThickDough() *ThickDough {
	return &ThickDough{}
}

func (t *ThickDough) Name() string {
	return "厚饼面团"
}
