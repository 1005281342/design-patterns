package healthconscious

// Visitor 定义访问者接口
type Visitor interface {
	VisitRestaurant(restaurant Restaurant)
	VisitPancakeHouse(pancake PancakeHouse)
}
