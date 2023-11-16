package menu

type Iterator interface {
	HasNext() bool
	Next() Component
}
