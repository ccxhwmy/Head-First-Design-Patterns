package bumball_state_winner

import "fmt"

type WinnerState struct {
	gumballMachine *GumballMachine
}

func NewWinnerState(gumballMachine *GumballMachine) *WinnerState {
	return &WinnerState{gumballMachine: gumballMachine}
}

func (this *WinnerState) InsertQuarter() {
	fmt.Println("Please wait, we're already giving you a Gumball")
}

func (this *WinnerState) EjectQuarter() {
	fmt.Println("Please wait, we're already giving you a Gumball")
}

func (this *WinnerState) TurnCrank() {
	fmt.Println("Turning again doesn't get you another gumball!")
}

func (this *WinnerState) Dispense() {
	this.gumballMachine.releaseBall()
	if this.gumballMachine.getCount() == 0 {
		this.gumballMachine.setState(this.gumballMachine.GetSoldOutState())
	} else {
		this.gumballMachine.releaseBall()
		fmt.Println("YOU'RE A WINNER! You got two gumballs for your quarter")
		if this.gumballMachine.getCount() > 0 {
			this.gumballMachine.setState(this.gumballMachine.GetNoQuarterState())
		} else {
			this.gumballMachine.setState(this.gumballMachine.GetSoldOutState())
		}
	}
}

func (this *WinnerState) Refill() {

}

func (this *WinnerState) String() string {
	return "despensing two gumballs for your quarter, because YOU'RE A WINNER!"
}
