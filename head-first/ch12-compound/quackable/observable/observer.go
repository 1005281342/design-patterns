package observable

import "fmt"

type Observer interface {
	Update(duck QuackObservable)
}

type Quackologist struct{}

func NewQuackologist() *Quackologist {
	return &Quackologist{}
}

func (q *Quackologist) Update(duck QuackObservable) {
	fmt.Printf("Quackologist %s just quacked\n", duck)
}
