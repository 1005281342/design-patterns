package main

import (
	"fmt"

	"github.com/1005281342/design-patterns/head-first/ch14-other/prototype/monster"
)

func main() {
	var maker = &monster.Maker{}
	fmt.Println(maker.MakeRandomMonster().Type())
	fmt.Println(maker.MakeRandomMonster().Type())
	fmt.Println(maker.MakeRandomMonster().Type())
	fmt.Println(maker.MakeRandomMonster().Type())
}
