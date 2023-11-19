package flock

import "github.com/1005281342/design-patterns/head-first/ch12-compound/quackable"

type Quackable interface {
	HasNext() bool
	Next() quackable.Quackable
}

type FlockIterator struct {
	flock *Flock
	idx   int
}

func NewFlockIterator(flock *Flock) *FlockIterator {
	return &FlockIterator{flock: flock}
}

func (f *FlockIterator) HasNext() bool {
	return f.idx < len(f.flock.quackers)
}

func (f *FlockIterator) Next() quackable.Quackable {
	f.idx++
	return f.flock.quackers[f.idx-1]
}
