package gumbal

import (
	"fmt"
	"testing"
)

func TestNewGumballMachine(t *testing.T) {
	gumballMachine := NewGumballMachine(5)

	fmt.Println(gumballMachine)

	gumballMachine.InsertQuarter()
	gumballMachine.TurnCrank()

	fmt.Println(gumballMachine)

	gumballMachine.InsertQuarter()
	gumballMachine.EjectQuarter()
	gumballMachine.TurnCrank()

	fmt.Println(gumballMachine)

	gumballMachine.InsertQuarter()
	gumballMachine.TurnCrank()
	gumballMachine.InsertQuarter()
	gumballMachine.TurnCrank()
	gumballMachine.EjectQuarter()

	fmt.Println(gumballMachine)

	gumballMachine.InsertQuarter()
	gumballMachine.InsertQuarter()
	gumballMachine.TurnCrank()
	gumballMachine.InsertQuarter()
	gumballMachine.TurnCrank()
	gumballMachine.InsertQuarter()
	gumballMachine.TurnCrank()

	fmt.Println(gumballMachine)
}
