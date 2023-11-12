package sauce

type MarinaraSauce struct{}

func NewMarinaraSauce() *MarinaraSauce {
	return &MarinaraSauce{}
}

func (m *MarinaraSauce) Name() string {
	return "意式番茄酱"
}
