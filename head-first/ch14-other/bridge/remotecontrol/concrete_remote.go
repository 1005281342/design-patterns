package remotecontrol

import "github.com/1005281342/design-patterns/head-first/ch14-other/bridge/tv"

type ConcreteRemote struct {
	tv             tv.TV
	currentStation uint
}

func NewConcreteRemote(tv tv.TV) *ConcreteRemote {
	return &ConcreteRemote{tv: tv}
}

func (c *ConcreteRemote) On() {
	c.tv.On()
}

func (c *ConcreteRemote) Off() {
	c.tv.Off()
}

func (c *ConcreteRemote) SetChannel(index uint) {
	c.tv.TuneChannel(index)
	c.currentStation = index
}

func (c *ConcreteRemote) NextChannel() {
	c.tv.TuneChannel(c.currentStation + 1)
}

func (c *ConcreteRemote) PrevChannel() {
	c.tv.TuneChannel(c.currentStation - 1)
}
