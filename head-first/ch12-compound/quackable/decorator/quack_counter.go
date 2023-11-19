package decorator

import (
	"sync/atomic"

	"github.com/1005281342/design-patterns/head-first/ch12-compound/quackable"
	"github.com/1005281342/design-patterns/head-first/ch12-compound/quackable/observable"
)

var numberOfQuacks uint32

type QuackCounter struct {
	duck quackable.Quackable

	observable.QuackObservable
}

func NewQuackCounter(duck quackable.Quackable) *QuackCounter {
	var q = &QuackCounter{duck: duck}
	q.QuackObservable = observable.NewObservable(q)
	return q
}

func (q *QuackCounter) Quack() {
	q.duck.Quack()
	atomic.AddUint32(&numberOfQuacks, 1)
}

func (q *QuackCounter) String() string {
	return q.duck.String()
}

func GetQuacks() uint32 {
	return numberOfQuacks
}
