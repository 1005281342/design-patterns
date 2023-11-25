package forest

import "github.com/1005281342/design-patterns/head-first/ch14-other/flyweight/tree"

type Forest struct {
	trees []*tree.Tree
}

func NewForest() *Forest {
	return &Forest{}
}

func (f *Forest) PlantTree(x int, y int, age int, name string, color string) {
	treeFactory := tree.NewTreeFactory()
	treeType := treeFactory.GetType(name, color)
	t := tree.NewTree(x, y, age, treeType)
	f.trees = append(f.trees, t)
}

func (f *Forest) Display() {
	for _, t := range f.trees {
		t.Display()
	}
}
