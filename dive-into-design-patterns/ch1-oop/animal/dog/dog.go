package dog

import "github.com/1005281342/design-patterns/dive-into-design-patterns/ch1-oop/animal"

type Dog struct {
	*animal.Animal

	bestFriend string
}

func (d *Dog) Bark() {}
