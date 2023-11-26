package main

import (
	"time"

	"github.com/1005281342/design-patterns/head-first/ch14-other/mediator/futurehouse"
)

func main() {
	var (
		alarm     = &futurehouse.Alarm{}
		calendar  = &futurehouse.Calendar{}
		coffeePot = &futurehouse.CoffeePot{}
		sprinkler = &futurehouse.Sprinkler{}
	)

	var mediator = futurehouse.NewConcreteMediator(alarm, calendar, coffeePot, sprinkler)

	// 模拟调度
	// 闹钟响了
	mediator.Mediate(alarm)

	// 模拟明天
	calendar.MockToday(time.Now().Add(24 * time.Hour))
	mediator.Mediate(alarm)

	// 模拟昨天
	calendar.MockToday(time.Now().Add(-24 * time.Hour))
	mediator.Mediate(alarm)

	//2023/11/26 16:05:18 日期:2023-11-26 闹钟响了
	//2023/11/26 16:05:18 日期:2023-11-27 闹钟响了
	//2023/11/26 16:05:18 开始煮咖啡
	//2023/11/26 16:05:19 咖啡煮好了
	//2023/11/26 16:05:19 日期:2023-11-25 闹钟响了
	//预约明早 9 点开启洒水器成功
}
