package tree

type Factory struct {
	treeTypes map[string]*Type
}

func NewTreeFactory() *Factory {
	return &Factory{treeTypes: make(map[string]*Type)}
}

func (tf *Factory) GetType(name, color string) *Type {
	key := name + "_" + color
	if treeType, ok := tf.treeTypes[key]; ok {
		return treeType
	}

	newType := NewType(name, color)
	tf.treeTypes[key] = newType
	return newType
}
