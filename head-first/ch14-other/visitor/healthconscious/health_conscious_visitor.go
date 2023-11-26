package healthconscious

import "fmt"

// AVisitor 具体的访问者类型
type AVisitor struct{}

func (a AVisitor) VisitRestaurant(restaurant Restaurant) {
	fmt.Println("Visiting Restaurant:", restaurant.Name)
	fmt.Println("Providing nutrition information for the menu items.")
	// 这里可以提供餐厅菜单的营养信息
}

func (a AVisitor) VisitPancakeHouse(pancake PancakeHouse) {
	fmt.Println("Visiting Pancake House:", pancake.Name)
	fmt.Println("Providing nutrition information for the menu items.")
	// 这里可以提供煎饼屋菜单的营养信息
}
