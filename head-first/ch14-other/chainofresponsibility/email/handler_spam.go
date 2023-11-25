package email

import (
	"fmt"

	"github.com/1005281342/design-patterns/head-first/ch14-other/chainofresponsibility/tag"
)

// SpamHandler 处理垃圾邮件
type SpamHandler struct {
	nextHandler Handler
}

func NewSpamHandler() *SpamHandler {
	return &SpamHandler{}
}

func (s *SpamHandler) SetNext(handler Handler) {
	s.nextHandler = handler
}

func (s *SpamHandler) Handle(email *Email) {
	if email.Tag == tag.Spam {
		fmt.Println("Handling spam mail...")
	} else if s.nextHandler != nil {
		s.nextHandler.Handle(email)
	}
}
