package main

import (
	"log"

	"github.com/1005281342/design-patterns/head-first/ch3-decorator/component"
	"github.com/1005281342/design-patterns/head-first/ch3-decorator/condiment"
)

func cat(beverage component.Beverage) {
	log.Printf("desc: %s, size: %s, cost: %f\n",
		beverage.GetDesc(), beverage.GetSize(), beverage.Cost())
}

func main() {
	var beverage1 component.Beverage = component.NewEspresso()
	cat(beverage1) // 2023/11/11 01:02:01 desc: 浓缩咖啡, size: tall, cost: 1.990000

	var beverage2 component.Beverage = component.NewDarkRoast()
	beverage2 = condiment.AddMocha(beverage2)
	beverage2 = condiment.AddMocha(beverage2)
	beverage2 = condiment.AddWhip(beverage2)
	cat(beverage2) // 2023/11/11 01:02:01 desc: 深度烘焙, 摩卡, 摩卡, 奶泡, size: tall, cost: 1.490000

	var beverage3 component.Beverage = component.NewHouseBlend()
	beverage3 = condiment.AddSoy(beverage3)
	beverage3 = condiment.AddMocha(beverage3)
	beverage3 = condiment.AddWhip(beverage3)
	cat(beverage3) // 2023/11/11 01:02:01 desc: 家常综合, 豆奶, 摩卡, 奶泡, size: tall, cost: 1.340000
}
