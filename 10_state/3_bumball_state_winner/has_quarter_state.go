package bumball_state_winner

import (
	"fmt"
	"math/rand"
)

type HasQuarterState struct {
	gumballMachine *GumballMachine
}

func NewHasQuarterState(gumballMachine *GumballMachine) *HasQuarterState {
	return &HasQuarterState{gumballMachine: gumballMachine}
}

func (this *HasQuarterState) InsertQuarter() {
	fmt.Println("You can't insert another quarter")
}

func (this *HasQuarterState) EjectQuarter() {
	fmt.Println("Quarter returned")
	this.gumballMachine.setState(this.gumballMachine.GetNoQuarterState())
}

func (this *HasQuarterState) TurnCrank() {
	fmt.Println("You turned...")
	winner := rand.Int31n(10)
	if winner == 1 && this.gumballMachine.getCount() > 1 {
		this.gumballMachine.setState(this.gumballMachine.GetWinnerState())
	} else {
		this.gumballMachine.setState(this.gumballMachine.GetSoldState())
	}
}

func (this *HasQuarterState) Dispense() {
	fmt.Println("No gumball dispensed")
}

func (this *HasQuarterState) Refill() {

}

func (this *HasQuarterState) String() string {
	return "waiting for turn of crank"
}
