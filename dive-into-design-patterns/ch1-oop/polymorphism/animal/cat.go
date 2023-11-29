package animal

import "fmt"

type Cat struct{}

func (c Cat) MakeSound() {
	fmt.Println("喵喵！")
}
