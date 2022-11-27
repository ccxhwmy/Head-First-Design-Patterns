package composite

import "fmt"

type RubberDuck struct {
}

func (this *RubberDuck) Quack() {
	fmt.Println("Squeak")
}

func (this *RubberDuck) String() string {
	return "Rubber Duck"
}
