package animal

type Animal struct {
	Name   string // 名字
	Sex    string // 性别
	Age    uint8  // 年龄
	Weight uint8  // 体重
	Color  string // 毛色
}

// Breathe 呼吸
func (a *Animal) Breathe() {}

// Eat 吃
func (a *Animal) Eat() {}

// Run 跑
func (a *Animal) Run() {}

// Sleep 睡觉
func (a *Animal) Sleep() {}
