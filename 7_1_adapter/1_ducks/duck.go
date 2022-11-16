package ducks

import "fmt"

type Duck interface {
	Quack()
	Fly()
}

type MallardDuck struct {
}

func (this *MallardDuck) Quack() {
	fmt.Println("Quack")
}

func (this *MallardDuck) Fly() {
	fmt.Println("I'm flying")
}
