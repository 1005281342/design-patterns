package component

type coffer struct {
	desc string
	size Size
}

func (b *coffer) GetDesc() string {
	return b.desc
}

func (b *coffer) GetSize() Size {
	return b.size
}

func (b *coffer) SetSize(size Size) {
	b.size = size
}
