package component

type Size string

const (
	SizeTall  Size = "tall"
	SizeGrand Size = "grand"
	SizeVenti Size = "venti"
)

type Beverage interface {
	GetDesc() string
	Cost() float32
	GetSize() Size
	SetSize(size Size)
}
