package futurehouse

// Alarm 表示闹钟对象
type Alarm struct {
	mediator Mediator
}

func (a *Alarm) SetMediator(mediator Mediator) {
	a.mediator = mediator
}

func (a *Alarm) Notify() {
	a.mediator.Mediate(a)
}
