package adapter

type TurkeyAdapter struct {
	turkey Turkey
}

// 使用火鸡充当鸭子
var _ Duck = (*TurkeyAdapter)(nil)

func NewTurkeyAdapter(turkey Turkey) *TurkeyAdapter {
	return &TurkeyAdapter{turkey: turkey}
}

func (t *TurkeyAdapter) Quack() {
	t.turkey.Gobble()
}

func (t *TurkeyAdapter) Fly() {
	for i := 0; i < 5; i++ {
		t.turkey.Fly()
	}
}
