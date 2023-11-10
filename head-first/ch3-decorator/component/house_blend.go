package component

type HouseBlend struct {
	*coffer
}

func NewHouseBlend() *HouseBlend {
	return &HouseBlend{
		coffer: &coffer{
			desc: "家常综合",
			size: SizeTall,
		},
	}
}

func (h *HouseBlend) Cost() float32 {
	return 0.89
}
