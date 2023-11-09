package subject

import "github.com/1005281342/design-patterns/head-first/ch2-observer/observer"

type Subject interface {
	Register(name string, observer observer.Observer) error
	Unregister(name string)
	Notify()
}
