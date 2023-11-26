package healthconscious

// Element 定义元素接口
type Element interface {
	Accept(visitor Visitor)
}
