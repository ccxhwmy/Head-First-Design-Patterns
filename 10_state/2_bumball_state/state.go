package bumball_state

type State interface {
	InsertQuarter()
	EjectQuarter()
	TurnCrank()
	Dispense()
	Refill()
	String() string
}
