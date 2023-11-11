package sample_pizza_factory

type Type string

const (
	TypeCheese    = "cheese"
	TypeVeggie    = "veggie"
	TypeClam      = "clam"
	TypePepperoni = "pepperoni"
)

type Pizza interface {
	Prepare()
	Bake()
	Cut()
	Box()
}

type pizza struct{}

func (p *pizza) Prepare() {}

func (p *pizza) Bake() {}

func (p *pizza) Cut() {}

func (p *pizza) Box() {}
