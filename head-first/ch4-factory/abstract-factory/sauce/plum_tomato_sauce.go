package sauce

type PlumTomatoSauce struct{}

func NewPlumTomatoSauce() *PlumTomatoSauce {
	return &PlumTomatoSauce{}
}

func (p *PlumTomatoSauce) Name() string {
	return "李子番茄酱"
}
