package healthconscious

// PancakeHouse 对象村煎饼屋
type PancakeHouse struct {
	Name string
	Menu []string // 假设这是煎饼屋的菜单
}

func (p PancakeHouse) Accept(visitor Visitor) {
	visitor.VisitPancakeHouse(p)
}
