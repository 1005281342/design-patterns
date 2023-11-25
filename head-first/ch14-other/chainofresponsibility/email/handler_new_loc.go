package email

import (
	"fmt"

	"github.com/1005281342/design-patterns/head-first/ch14-other/chainofresponsibility/tag"
)

// NewLocHandler 处理垃圾邮件
type NewLocHandler struct {
	nextHandler Handler
}

func NewNewLocHandler() *NewLocHandler {
	return &NewLocHandler{}
}

func (n *NewLocHandler) SetNext(handler Handler) {
	n.nextHandler = handler
}

func (n *NewLocHandler) Handle(email *Email) {
	if email.Tag == tag.NewLoc {
		fmt.Println("Handling new loc mail...")
	} else if n.nextHandler != nil {
		n.nextHandler.Handle(email)
	}
}
