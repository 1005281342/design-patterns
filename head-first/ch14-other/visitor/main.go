package main

import "github.com/1005281342/design-patterns/head-first/ch14-other/visitor/healthconscious"

func main() {
	// 创建对象村餐厅和对象村煎饼屋
	restaurant := healthconscious.Restaurant{Name: "Objects Village Restaurant", Menu: []string{"Dish1", "Dish2"}}
	pancakeHouse := healthconscious.PancakeHouse{Name: "Objects Village Pancake House", Menu: []string{"Pancake1", "Pancake2"}}

	// 创建一个健康意识的访问者
	visitor := healthconscious.AVisitor{}

	// 顾客询问营养信息
	restaurant.Accept(visitor)
	pancakeHouse.Accept(visitor)
}
