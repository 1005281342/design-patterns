package behavior

import "log"

type QuackBehavior interface {
	Quack()
}

type Quack struct{}

func (q *Quack) Quack() {
	log.Println("嘎嘎嘎！")
}

type Squeak struct{}

func (s *Squeak) Quack() {
	log.Println("吱吱吱！")
}

type QuackMute struct{}

func (q *QuackMute) Quack() {}
