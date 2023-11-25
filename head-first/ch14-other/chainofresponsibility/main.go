package main

import (
	"fmt"

	"github.com/1005281342/design-patterns/head-first/ch14-other/chainofresponsibility/email"
	"github.com/1005281342/design-patterns/head-first/ch14-other/chainofresponsibility/tag"
)

func main() {
	// 初始化邮件处理器
	var (
		spamHandler      = email.NewSpamHandler()
		fanHandler       = email.NewFanHandler()
		complaintHandler = email.NewComplaintHandler()
		newLocHandler    = email.NewNewLocHandler()
	)

	// 构建责任链
	spamHandler.SetNext(fanHandler)
	fanHandler.SetNext(complaintHandler)
	complaintHandler.SetNext(newLocHandler)

	// 模拟不同类型的邮件
	var emails = []*email.Email{
		{Subject: "Fan Mail", Content: "I'm your biggest fan!", Tag: tag.Fan},
		{Subject: "Complaint", Content: "Unhappy with the service.", Tag: tag.Complaint},
		{Subject: "New Location", Content: "Opening a new branch.", Tag: tag.NewLoc},
		{Subject: "Spam", Content: "Buy now!", Tag: tag.Spam},
		{Subject: "Important Mail", Content: "Urgent message.", Tag: tag.Important},
	}

	// 处理邮件
	for _, e := range emails {
		fmt.Printf("\nHandling Email - Subject: %s\n", e.Subject)
		spamHandler.Handle(e)
	}
}
