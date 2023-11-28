package main

import (
	"fmt"

	"github.com/1005281342/design-patterns/dive-into-design-patterns/ch1-oop/cat"
)

func main() {
	var kaka = cat.Cat{
		Name:    "卡卡",
		Sex:     "男孩",
		Age:     3,
		Weight:  7,
		Color:   "棕色",
		Texture: "条纹",
	}
	var lulu = cat.Cat{
		Name:    "露露",
		Sex:     "女孩",
		Age:     2,
		Weight:  5,
		Color:   "灰色",
		Texture: "纯色",
	}
	fmt.Printf("KaKa:%+v\n", kaka)
	fmt.Printf("LuLu:%+v\n", lulu)
}
