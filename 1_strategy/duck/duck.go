package duck

import "fmt"

type FlyBehavior interface {
	Fly()
}

type FlyWithWings struct {
}

func (this *FlyWithWings) Fly() {
	fmt.Println("I'm flying!")
}

type FlyNoWay struct {
}

func (this *FlyNoWay) Fly() {
	fmt.Println("I can't Fly")
}

type FlyRocketPowered struct {
}

func (this *FlyRocketPowered) Fly() {
	fmt.Println("I'm flying with a rocket!")
}

type QuackBehavior interface {
	Quack()
}

type MuteQuack struct {
}

func (this *MuteQuack) Quack() {
	fmt.Println("<< Silence >>")
}

type Squeak struct {
}

func (this *Squeak) Quack() {
	fmt.Println("Ga Ga Ga ...")
}

type DuckCall struct {
}

func (this *DuckCall) Quack() {
	fmt.Println("Ga Ga Ga ...(machine display)")
}

type Duck struct {
	flyBehavior   FlyBehavior
	quackBehavior QuackBehavior
}

func (this *Duck) PerformQuack() {
	this.quackBehavior.Quack()
}

func (this *Duck) PerformFly() {
	this.flyBehavior.Fly()
}

func (this *Duck) SetFlyBehavior(fb FlyBehavior) {
	this.flyBehavior = fb
}

func (this *Duck) SetQuackBehavior(qb QuackBehavior) {
	this.quackBehavior = qb
}

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
