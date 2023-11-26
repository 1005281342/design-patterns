package main

import (
	"fmt"
	"github.com/1005281342/design-patterns/head-first/ch14-other/memento/game"
)

func main() {
	var (
		saved *game.Memento
		obj   = game.NewMasterGameObject()
	)
	currentState(obj)

	// 模拟到达了下个关卡
	obj.SetState("2")
	// 模拟存档
	saved = obj.GetCurrentState()
	currentState(obj)

	// 继续游戏，模拟通过某个关卡
	obj.SetState("x")
	currentState(obj)

	// 模拟回档
	obj.RestoreState(saved)
	currentState(obj)
}

func currentState(obj *game.MasterGameObject) {
	fmt.Printf("当前关卡：%s\n", obj.GetCurrentState().SavedGameState())
}
