package bumball_state

import (
	"bytes"
	"fmt"
	"strconv"
)

type GumballMachine struct {
	soldOutState    State
	noQuarterState  State
	hasQuarterState State
	soldState       State
	state           State
	count           int
}

func NewGumballMachine(numberGumball int) *GumballMachine {
	res := &GumballMachine{
		count: numberGumball,
	}
	res.soldOutState = NewSoldOutState(res)
	res.noQuarterState = NewNoQuarterState(res)
	res.hasQuarterState = NewHasQuarterState(res)
	res.soldState = NewSoldState(res)
	if numberGumball > 0 {
		res.state = res.noQuarterState
	} else {
		res.state = res.soldOutState
	}
	return res
}

func (this *GumballMachine) InsertQuarter() {
	this.state.InsertQuarter()
}

func (this *GumballMachine) EjectQuarter() {
	this.state.EjectQuarter()
}

func (this *GumballMachine) TurnCrank() {
	this.state.TurnCrank()
	this.state.Dispense()
}

func (this *GumballMachine) releaseBall() {
	fmt.Println("A gumball comes rolling out the slot...")
	if this.count > 0 {
		this.count--
	}
}

func (this *GumballMachine) getCount() int {
	return this.count
}

func (this *GumballMachine) refill(count int) {
	this.count += count
	fmt.Println("The gumball machine was just refilled; its new count is: ", this.count)
	this.state.Refill()
}

func (this *GumballMachine) setState(state State) {
	this.state = state
}

func (this *GumballMachine) GetState() State {
	return this.state
}

func (this *GumballMachine) GetSoldOutState() State {
	return this.soldOutState
}

func (this *GumballMachine) GetNoQuarterState() State {
	return this.noQuarterState
}

func (this *GumballMachine) GetHasQuarterState() State {
	return this.hasQuarterState
}

func (this *GumballMachine) GetSoldState() State {
	return this.soldState
}

func (this *GumballMachine) String() string {
	res := bytes.Buffer{}
	res.WriteString("\nMighty Gumball, Inc.")
	res.WriteString("\nInventory: " + strconv.Itoa(this.count) + " gumball")
	if this.count != 1 {
		res.WriteString("s")
	}
	res.WriteString("\n")
	res.WriteString("Machine is " + this.state.String() + "\n")
	return res.String()
}
