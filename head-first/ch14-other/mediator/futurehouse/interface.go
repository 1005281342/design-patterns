package futurehouse

// Mediator 接口定义了中介者的行为
type Mediator interface {
	Mediate(colleague Colleague)
}

// Colleague 表示参与交互的对象的通用接口
type Colleague interface {
	SetMediator(mediator Mediator)
	Notify()
}
