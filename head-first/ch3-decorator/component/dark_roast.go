package component

type DarkRoast struct {
	*coffer
}

func NewDarkRoast() *DarkRoast {
	return &DarkRoast{
		coffer: &coffer{
			desc: "深度烘焙",
			size: SizeTall,
		},
	}
}

func (h *DarkRoast) Cost() float32 {
	return 0.99
}
