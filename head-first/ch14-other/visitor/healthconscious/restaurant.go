package healthconscious

// Restaurant 对象村餐厅
type Restaurant struct {
	Name string
	Menu []string // 假设这是餐厅的菜单
}

func (r Restaurant) Accept(visitor Visitor) {
	visitor.VisitRestaurant(r)
}
