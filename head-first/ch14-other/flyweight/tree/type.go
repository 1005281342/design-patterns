package tree

import "fmt"

// Type 表示树的类型，包含了共享的属性
type Type struct {
	Name  string
	Color string
}

func NewType(name, color string) *Type {
	return &Type{Name: name, Color: color}
}

func (t *Type) Display(x, y, age int) {
	fmt.Printf("Drawing a %s %s tree at position (%d, %d) with age %d\n", t.Color, t.Name, x, y, age)
}
