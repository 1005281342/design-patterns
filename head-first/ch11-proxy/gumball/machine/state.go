package machine

// State 状态接口
type State interface {
	InsertQuarter()
	EjectQuarter()
	TurnCrank()
	Dispense()
	String() string
}
