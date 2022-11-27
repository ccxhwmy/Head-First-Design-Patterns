package bumball_state

import "fmt"

type SoldState struct {
	gumballMachine *GumballMachine
}

func NewSoldState(gumballMachine *GumballMachine) *SoldState {
	return &SoldState{gumballMachine: gumballMachine}
}

func (this *SoldState) InsertQuarter() {
	fmt.Println("Please wait, we're already giving you a gumball")
}

func (this *SoldState) EjectQuarter() {
	fmt.Println("Sorry, you already turned the crank")
}

func (this *SoldState) TurnCrank() {
	fmt.Println("Turning twice doesn't get you another gumball!")
}

func (this *SoldState) Dispense() {
	this.gumballMachine.releaseBall()
	if this.gumballMachine.getCount() > 0 {
		this.gumballMachine.setState(this.gumballMachine.GetNoQuarterState())
	} else {
		this.gumballMachine.setState(this.gumballMachine.GetSoldOutState())
	}
}

func (this *SoldState) Refill() {

}

func (this *SoldState) String() string {
	return "SoldState"
}
