package main

import "fmt"

type FlyBehavior interface {
	fly()
}

type FlyWithWings struct {
}

func (this *FlyWithWings) fly() {
	fmt.Println("I'm flying!")
}

type FlyNoWay struct {
}

func (this *FlyNoWay) fly() {
	fmt.Println("I can't fly")
}

type FlyRocketPowered struct {
}

func (this *FlyRocketPowered) fly() {
	fmt.Println("I'm flying with a rocket!")
}

type QuackBehavior interface {
	quack()
}

type MuteQuack struct {
}

func (this *MuteQuack) quack() {
	fmt.Println("<< Silence >>")
}

type Squeak struct {
}

func (this *Squeak) quack() {
	fmt.Println("Ga Ga Ga ...")
}

type Buzzer struct {
}

func (this *Buzzer) quack() {
	fmt.Println("eng eng eng ...")
}

type Duck struct {
	flyBehavior   FlyBehavior
	quackBehavior QuackBehavior
}

func (this *Duck) performQuack() {
	this.quackBehavior.quack()
}

func (this *Duck) performFly() {
	this.flyBehavior.fly()
}

func (this *Duck) setFlyBehavior(fb FlyBehavior) {
	this.flyBehavior = fb
}

func (this *Duck) setQuackBehavior(qb QuackBehavior) {
	this.quackBehavior = qb
}

type MallardDuck struct {
	Duck
}

func newMallardDuck() *MallardDuck {
	return &MallardDuck{Duck{
		flyBehavior:   &FlyWithWings{},
		quackBehavior: &Squeak{},
	}}
}

func (this *MallardDuck) display() {
	fmt.Println("I'm a real Mallard duck")
}

type ModelDuck struct {
	Duck
}

func newModelDuck() *ModelDuck {
	return &ModelDuck{Duck{
		flyBehavior:   &FlyNoWay{},
		quackBehavior: &MuteQuack{},
	}}
}

func (this *ModelDuck) display() {
	fmt.Println("I'm a Model duck")
}

func (this *ModelDuck) setRocketPowered() {
	this.setFlyBehavior(&FlyRocketPowered{})
}

func (this *ModelDuck) setBuzzerQuack() {
	this.setQuackBehavior(&Buzzer{})
}

func main() {
	fmt.Println("-------------- mallard duck show time -------------")
	mallardDuck := newMallardDuck()
	mallardDuck.display()
	mallardDuck.performFly()
	mallardDuck.performQuack()

	fmt.Println("-------------- model duck show time -------------")
	modelDuck := newModelDuck()
	modelDuck.display()
	modelDuck.performFly()
	modelDuck.performQuack()

	fmt.Println("-------------- model duck set rocket power -------------")
	modelDuck.setRocketPowered()
	modelDuck.performFly()

	fmt.Println("-------------- model duck set buzzer -------------")
	modelDuck.setBuzzerQuack()
	modelDuck.performQuack()
}
