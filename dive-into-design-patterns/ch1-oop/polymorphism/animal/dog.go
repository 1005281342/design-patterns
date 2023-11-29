package animal

import "fmt"

type Dog struct{}

func (d *Dog) MakeSound() {
	fmt.Println("汪汪！")
}
