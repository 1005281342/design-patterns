package email

import (
	"fmt"

	"github.com/1005281342/design-patterns/head-first/ch14-other/chainofresponsibility/tag"
)

// FanHandler 处理垃圾邮件
type FanHandler struct {
	nextHandler Handler
}

func NewFanHandler() *FanHandler {
	return &FanHandler{}
}

func (f *FanHandler) SetNext(handler Handler) {
	f.nextHandler = handler
}

func (f *FanHandler) Handle(email *Email) {
	if email.Tag == tag.Fan {
		fmt.Println("Handling fan mail...")
	} else if f.nextHandler != nil {
		f.nextHandler.Handle(email)
	}
}
