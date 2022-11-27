package observer

import "fmt"

type RubberDuck struct {
	observable *Observable
}

func NewRubberDuck() *RubberDuck {
	res := &RubberDuck{}
	res.observable = NewObservable(res)
	return res
}

func (this *RubberDuck) RegisterObserver(observer Observer) {
	this.observable.RegisterObserver(observer)
}

func (this *RubberDuck) NotifyObservers() {
	this.observable.NotifyObservers()
}

func (this *RubberDuck) Quack() {
	fmt.Println("Squeak")
	this.NotifyObservers()
}

func (this *RubberDuck) String() string {
	return "Rubber Duck"
}
