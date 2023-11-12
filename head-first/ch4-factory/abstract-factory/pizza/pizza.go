package pizza

import (
	"log"
	"strings"
)

type Type string

func (t Type) String() string {
	return string(t)
}

const TypeCheese Type = "cheese"

type Pizza interface {
	Prepare()
	Bake()
	Cut()
	Box()

	GetName() string
	Info()
}

type pizza struct {
	name     string
	dough    string   // 面团
	sauce    string   // 酱
	cheese   string   // 芝士
	toppings []string // 配料
}

func (p *pizza) Prepare() {
	log.Println("Preparing " + p.name)
	log.Println("Tossing dough...")
	log.Println("Adding sauce...")
	log.Println("Adding toppings: " + strings.Join(p.toppings, " "))
}

func (p *pizza) Bake() {
	log.Println("Bake for 25 minutes at 350")
}

func (p *pizza) Cut() {
	log.Println("Cutting the pizza into diagonal slices")
}

func (p *pizza) Box() {
	log.Println("Place pizza in official PizzaStore box")
}

func (p *pizza) GetName() string {
	return p.name
}

func (p *pizza) Info() {
	log.Printf("pizza: %+v\n", p)
}
