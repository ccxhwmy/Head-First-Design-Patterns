package ducks

import "fmt"

type RedheadDuck struct {
}

func (this *RedheadDuck) Quack() {
	fmt.Println("Quack")
}
