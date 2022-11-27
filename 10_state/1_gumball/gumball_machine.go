package gumbal

import (
	"bytes"
	"fmt"
	"strconv"
)

type State int

const (
	SoldOut State = iota
	NoQuarter
	HasQuarter
	Sold
)

type GumballMachine struct {
	state State
	count int
}

func NewGumballMachine(count int) *GumballMachine {
	return &GumballMachine{
		state: NoQuarter,
		count: count,
	}
}

func (this *GumballMachine) InsertQuarter() {
	if this.state == HasQuarter {
		fmt.Println("You can't insert another quarter")
	} else if this.state == NoQuarter {
		this.state = HasQuarter
		fmt.Println("You inserted a quarter")
	} else if this.state == SoldOut {
		fmt.Println("You can't insert a quarter, the machine is sold out")
	} else if this.state == Sold {
		fmt.Println("Please wait, we're already giving you a gumball")
	}
}

func (this *GumballMachine) EjectQuarter() {
	if this.state == HasQuarter {
		fmt.Println("Quarter returned")
	} else if this.state == NoQuarter {
		fmt.Println("You haven't inserted a quarter")
	} else if this.state == SoldOut {
		fmt.Println("You can't eject, you haven't inserted a quarter yet")
	} else if this.state == Sold {
		fmt.Println("Sorry, you already turned the crank")
	}
}

func (this *GumballMachine) TurnCrank() {
	if this.state == HasQuarter {
		fmt.Println("You turned...")
		this.state = Sold
		this.dispense()
	} else if this.state == NoQuarter {
		fmt.Println("You haven't inserted a quarter")
	} else if this.state == SoldOut {
		fmt.Println("Sorry, you already turned the crank")
	} else if this.state == Sold {
		fmt.Println("You can't eject, you haven't inserted a quarter yet")
	}
}

func (this *GumballMachine) dispense() {
	if this.state == HasQuarter {
		fmt.Println("No gumball dispensed")
	} else if this.state == NoQuarter {
		fmt.Println("You need to pay first")
	} else if this.state == SoldOut {
		fmt.Println("No gumball dispensed")
	} else if this.state == Sold {
		fmt.Println("A gumball comes rolling out the slot")
		this.count--
		if this.count == 0 {
			fmt.Println("Oops, out of gumballs!")
			this.state = SoldOut
		} else {
			this.state = NoQuarter
		}
	}
}

func (this *GumballMachine) Refill(numberGumballs int) {
	this.count = numberGumballs
	this.state = NoQuarter
}

func (this *GumballMachine) String() string {
	res := bytes.Buffer{}
	res.WriteString("\nMighty Gumball, Inc.")
	res.WriteString("\nInventory: " + strconv.Itoa(this.count) + " gumball")
	if this.count != 1 {
		res.WriteString("s")
	}
	res.WriteString("\n")
	res.WriteString("Machine is ")
	if this.state == HasQuarter {
		res.WriteString("waiting for turn of crank")
	} else if this.state == NoQuarter {
		fmt.Println("waiting for quarter")
	} else if this.state == SoldOut {
		fmt.Println("sold out")
	} else if this.state == Sold {
		fmt.Println("delivering a gumball")
	}
	res.WriteString("\n")
	return res.String()
}
