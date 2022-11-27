package bumball_state_winner

import "fmt"

type SoldOutState struct {
	gumballMachine *GumballMachine
}

func NewSoldOutState(gumballMachine *GumballMachine) *SoldOutState {
	return &SoldOutState{gumballMachine: gumballMachine}
}

func (this *SoldOutState) InsertQuarter() {
	fmt.Println("You can't insert a quarter, the machine is sold out")
}

func (this *SoldOutState) EjectQuarter() {
	fmt.Println("You can't eject, you haven't inserted a quarter yet")
}

func (this *SoldOutState) TurnCrank() {
	fmt.Println("You turned, but there are no gumballs")
}

func (this *SoldOutState) Dispense() {
	fmt.Println("No gumball dispensed")
}

func (this *SoldOutState) Refill() {
	this.gumballMachine.setState(this.gumballMachine.GetNoQuarterState())
}

func (this *SoldOutState) String() string {
	return "sold out"
}
