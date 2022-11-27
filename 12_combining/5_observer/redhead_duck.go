package observer

import "fmt"

type RedheadDuck struct {
	observable *Observable
}

func NewRedheadDuck() *RedheadDuck {
	res := &RedheadDuck{}
	res.observable = NewObservable(res)
	return res
}

func (this *RedheadDuck) RegisterObserver(observer Observer) {
	this.observable.RegisterObserver(observer)
}

func (this *RedheadDuck) NotifyObservers() {
	this.observable.NotifyObservers()
}

func (this *RedheadDuck) Quack() {
	fmt.Println("Quack")
	this.NotifyObservers()
}

func (this *RedheadDuck) String() string {
	return "Redhead Duck"
}
