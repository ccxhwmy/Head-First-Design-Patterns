package observer

import "fmt"

type MallardDuck struct {
	observable *Observable
}

func NewMallardDuck() *MallardDuck {
	res := &MallardDuck{}
	res.observable = NewObservable(res)
	return res
}

func (this *MallardDuck) RegisterObserver(observer Observer) {
	this.observable.RegisterObserver(observer)
}

func (this *MallardDuck) NotifyObservers() {
	this.observable.NotifyObservers()
}

func (this *MallardDuck) Quack() {
	fmt.Println("Quack")
	this.NotifyObservers()
}

func (this *MallardDuck) String() string {
	return "Mallard Duck"
}
