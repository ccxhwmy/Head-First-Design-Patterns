package duck

import "fmt"

type ModelDuck struct {
	Duck
}

func NewModelDuck() *ModelDuck {
	return &ModelDuck{Duck{
		flyBehavior:   &FlyNoWay{},
		quackBehavior: &MuteQuack{},
	}}
}

func (this *ModelDuck) Display() {
	fmt.Println("I'm a Model duck")
}

func (this *ModelDuck) SetRocketPowered() {
	this.SetFlyBehavior(&FlyRocketPowered{})
}

func (this *ModelDuck) SetDuckCallQuack() {
	this.SetQuackBehavior(&DuckCall{})
}
