package observable

type QuackObservable interface {
	String() string
	Register(observer Observer)
	Notify()
}

type Observable struct {
	observers []Observer
	duck      QuackObservable
}

func NewObservable(duck QuackObservable) *Observable {
	return &Observable{duck: duck}
}

func (o *Observable) String() string {
	return o.duck.String()
}

func (o *Observable) Register(observer Observer) {
	o.observers = append(o.observers, observer)
}

func (o *Observable) Notify() {
	for _, observers := range o.observers {
		observers.Update(o.duck)
	}
}
