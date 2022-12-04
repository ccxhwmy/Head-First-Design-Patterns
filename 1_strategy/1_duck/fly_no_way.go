package duck

import "fmt"

type FlyNoWay struct {
}

func (this *FlyNoWay) Fly() {
	fmt.Println("I can't Fly")
}
