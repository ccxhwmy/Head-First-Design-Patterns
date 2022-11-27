package observer

import "fmt"

type DecoyDuck struct {
	observable *Observable
}

func NewDecoyDuck() *DecoyDuck {
	res := &DecoyDuck{}
	res.observable = NewObservable(res)
	return res
}

func (this *DecoyDuck) RegisterObserver(observer Observer) {
	this.observable.RegisterObserver(observer)
}

func (this *DecoyDuck) NotifyObservers() {
	this.observable.NotifyObservers()
}

func (this *DecoyDuck) Quack() {
	fmt.Println("<< Silence >>")
	this.NotifyObservers()
}

func (this *DecoyDuck) String() string {
	return "Decoy Duck"
}
