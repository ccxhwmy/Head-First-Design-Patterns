package decorator

import "fmt"

type MallardDuck struct {
}

func (this *MallardDuck) Quack() {
	fmt.Println("Quack")
}

func (this *MallardDuck) String() string {
	return "Mallard Duck"
}
