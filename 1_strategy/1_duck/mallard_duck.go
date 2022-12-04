package duck

import "fmt"

type MallardDuck struct {
	Duck
}

func NewMallardDuck() *MallardDuck {
	return &MallardDuck{Duck{
		flyBehavior:   &FlyWithWings{},
		quackBehavior: &Squeak{},
	}}
}

func (this *MallardDuck) Display() {
	fmt.Println("I'm a real Mallard duck")
}
