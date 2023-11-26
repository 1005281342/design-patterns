package futurehouse

import (
	"log"
	"time"
)

// CoffeePot 表示咖啡壶对象
type CoffeePot struct {
	mediator Mediator
}

func (c *CoffeePot) SetMediator(mediator Mediator) {
	c.mediator = mediator
}

func (c *CoffeePot) Notify() {
	c.mediator.Mediate(c)
}

func (c *CoffeePot) StartBrewingCoffee() {
	log.Println("开始煮咖啡")
	time.Sleep(time.Second) // 简单 sleep 模拟制作咖啡耗时
	c.Notify()
}
