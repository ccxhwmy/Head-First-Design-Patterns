package ducks

import "fmt"

type RubberDuck struct {
}

func (this *RubberDuck) Quack() {
	fmt.Println("Squeak")
}
