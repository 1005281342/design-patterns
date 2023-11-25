package email

import (
	"fmt"

	"github.com/1005281342/design-patterns/head-first/ch14-other/chainofresponsibility/tag"
)

// ComplaintHandler 处理垃圾邮件
type ComplaintHandler struct {
	nextHandler Handler
}

func NewComplaintHandler() *ComplaintHandler {
	return &ComplaintHandler{}
}

func (c *ComplaintHandler) SetNext(handler Handler) {
	c.nextHandler = handler
}

func (c *ComplaintHandler) Handle(email *Email) {
	if email.Tag == tag.Complaint {
		fmt.Println("Handling complaint mail...")
	} else if c.nextHandler != nil {
		c.nextHandler.Handle(email)
	}
}
