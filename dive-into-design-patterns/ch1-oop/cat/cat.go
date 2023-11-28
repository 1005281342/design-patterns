package cat

// Cat 猫，类
type Cat struct {
	Name    string // 名字
	Sex     string // 性别
	Age     uint8  // 年龄
	Weight  uint8  // 体重
	Color   string // 毛色
	Texture string // 纹理
}

// Breathe 呼吸
func (c Cat) Breathe() {}

// Eat 吃
func (c Cat) Eat() {}

// Meow 喵
func (c Cat) Meow() {}

// Run 跑
func (c Cat) Run() {}

// Sleep 睡觉
func (c Cat) Sleep() {}
