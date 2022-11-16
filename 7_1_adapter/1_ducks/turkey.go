package ducks

import "fmt"

type Turkey interface {
	Gobble()
	Fly()
}

type WildTurkey struct {
}

func (this *WildTurkey) Gobble() {
	fmt.Println("Gobble gobble")
}

func (this *WildTurkey) Fly() {
	fmt.Println("I'm flying a short distance")
}
