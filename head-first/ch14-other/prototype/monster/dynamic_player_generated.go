package monster

type DynamicPlayerGenerated struct{}

func (d *DynamicPlayerGenerated) clone() Monster {
	return &DynamicPlayerGenerated{}
}

func (d *DynamicPlayerGenerated) Type() string {
	return "DynamicPlayerGeneratedMonster"
}
