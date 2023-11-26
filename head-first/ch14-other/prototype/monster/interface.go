package monster

type Monster interface {
	clone() Monster
	Type() string
}
