package component

type Espresso struct {
	*coffer
}

func NewEspresso() *Espresso {
	return &Espresso{
		coffer: &coffer{
			desc: "浓缩咖啡",
			size: SizeTall,
		},
	}
}

func (h *Espresso) Cost() float32 {
	return 1.99
}
