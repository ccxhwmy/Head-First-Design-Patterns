package bumball_state

import (
	"fmt"
	"testing"
)

func TestGumballMachine(t *testing.T) {
	gumballMachine := NewGumballMachine(2)

	fmt.Println(gumballMachine)

	gumballMachine.InsertQuarter()
	gumballMachine.TurnCrank()

	fmt.Println(gumballMachine)

	gumballMachine.InsertQuarter()
	gumballMachine.TurnCrank()
	gumballMachine.InsertQuarter()
	gumballMachine.TurnCrank()

	gumballMachine.refill(5)
	gumballMachine.InsertQuarter()
	gumballMachine.TurnCrank()

	fmt.Println(gumballMachine)
}
