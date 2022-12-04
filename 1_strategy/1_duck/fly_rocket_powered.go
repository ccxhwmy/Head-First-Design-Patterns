package duck

import "fmt"

type FlyRocketPowered struct {
}

func (this *FlyRocketPowered) Fly() {
	fmt.Println("I'm flying with a rocket!")
}
