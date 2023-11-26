package monster

type WellKnown struct{}

func (w *WellKnown) clone() Monster {
	return &WellKnown{}
}

func (w *WellKnown) Type() string {
	return "WellKnownMonster"
}
