package bumball_state_winner

import "fmt"

type NoQuarterState struct {
	gumballMachine *GumballMachine
}

func NewNoQuarterState(gumballMachine *GumballMachine) *NoQuarterState {
	return &NoQuarterState{gumballMachine: gumballMachine}
}

func (this *NoQuarterState) InsertQuarter() {
	fmt.Println("You inserted a quarter")
	this.gumballMachine.setState(this.gumballMachine.GetHasQuarterState())
}

func (this *NoQuarterState) EjectQuarter() {
	fmt.Println("You haven't inserted a quarter")
}

func (this *NoQuarterState) TurnCrank() {
	fmt.Println("You turned, but there's no quarter")
}

func (this *NoQuarterState) Dispense() {
	fmt.Println("You need to pay first")
}

func (this *NoQuarterState) Refill() {

}

func (this *NoQuarterState) String() string {
	return "waiting for quarter"
}
