package cheese

type MozzarellaCheese struct{}

func NewMozzarellaCheese() *MozzarellaCheese {
	return &MozzarellaCheese{}
}

func (m *MozzarellaCheese) Name() string {
	return "Mozzarella 芝士"
}
