package tree

type Tree struct {
	X, Y int
	Age  int
	Type *Type
}

func NewTree(x int, y int, age int, treeType *Type) *Tree {
	return &Tree{X: x, Y: y, Age: age, Type: treeType}
}

func (t *Tree) Display() {
	t.Type.Display(t.X, t.Y, t.Age)
}
