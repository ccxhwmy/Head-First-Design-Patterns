package bumball_state_winner

type State interface {
	InsertQuarter()
	EjectQuarter()
	TurnCrank()
	Dispense()
	Refill()
	String() string
}
