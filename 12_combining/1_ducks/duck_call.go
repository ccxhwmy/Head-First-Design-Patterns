package ducks

import "fmt"

type DuckCall struct {
}

func (this *DuckCall) Quack() {
	fmt.Println("Kwak")
}
