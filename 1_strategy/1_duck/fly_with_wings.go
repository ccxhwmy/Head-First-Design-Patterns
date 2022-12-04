package duck

import "fmt"

type FlyWithWings struct {
}

func (this *FlyWithWings) Fly() {
	fmt.Println("I'm flying!")
}
