package futurehouse

import "fmt"

// Sprinkler 表示洒水器对象
type Sprinkler struct {
	mediator Mediator
}

func (s *Sprinkler) SetMediator(mediator Mediator) {
	s.mediator = mediator
}

func (s *Sprinkler) Notify() {

}

func (s *Sprinkler) Reserve() {
	fmt.Println("预约明早 9 点开启洒水器成功")
}
