package ducks

import "fmt"

type MallardDuck struct {
}

func (this *MallardDuck) Quack() {
	fmt.Println("Quack")
}
