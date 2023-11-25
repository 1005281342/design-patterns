package duck

type Context struct{}

func NewContext() *Context {
	return &Context{}
}
