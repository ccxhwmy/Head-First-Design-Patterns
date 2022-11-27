package ducks

import "fmt"

type DecoyDuck struct {
}

func (this *DecoyDuck) Quack() {
	fmt.Println("<< Silence >>")
}
