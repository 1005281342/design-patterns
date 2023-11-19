package duck

import (
	"fmt"

	"github.com/1005281342/design-patterns/head-first/ch12-compound/quackable/observable"
)

type DuckCall struct {
	observable.QuackObservable
}

func NewDuckCall() *DuckCall {
	var d = &DuckCall{}
	d.QuackObservable = observable.NewObservable(d)
	return d
}

func (d *DuckCall) Quack() {
	fmt.Println("Kwak")

	d.Notify()
}

func (d *DuckCall) String() string {
	return "DuckCall"
}
