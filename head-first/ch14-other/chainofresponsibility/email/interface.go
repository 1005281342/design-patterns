package email

type Handler interface {
	SetNext(handler Handler)
	Handle(email *Email)
}
