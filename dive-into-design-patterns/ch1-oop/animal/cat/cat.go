package cat

import "github.com/1005281342/design-patterns/dive-into-design-patterns/ch1-oop/animal"

type Cat struct {
	*animal.Animal

	isNasty bool
}

func (c *Cat) Meow() {}
