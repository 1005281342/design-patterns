package cheese

type ReggianoCheese struct{}

func NewReggianoCheese() *ReggianoCheese {
	return &ReggianoCheese{}
}

func (r *ReggianoCheese) Name() string {
	return "Reggiano 芝士"
}
