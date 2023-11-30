package related

type Knowledge interface{}

type Course interface {
	GetKnowledge() Knowledge
}

type Student interface {
	Remember(knowledge Knowledge)
}

type Professor struct {
	student Student
}

func (p *Professor) Teach(c Course) {
	p.student.Remember(c.GetKnowledge())
}
