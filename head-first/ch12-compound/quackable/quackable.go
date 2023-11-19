package quackable

import "github.com/1005281342/design-patterns/head-first/ch12-compound/quackable/observable"

type Quackable interface {
	Quack()

	observable.QuackObservable
}
