package component

type Decaf struct {
	*coffer
}

func NewDecaf() *Decaf {
	return &Decaf{
		coffer: &coffer{
			desc: "低咖啡因",
			size: SizeTall,
		},
	}
}

func (h *Decaf) Cost() float32 {
	return 1.05
}
