package email

import "github.com/1005281342/design-patterns/head-first/ch14-other/chainofresponsibility/tag"

// Email 结构体表示邮件的基本信息
type Email struct {
	Subject string
	Content string
	Tag     tag.Tag
}
