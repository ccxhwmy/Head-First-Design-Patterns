package observer

import "fmt"

type DuckCall struct {
	observable *Observable
}

func NewDuckCall() *DuckCall {
	res := &DuckCall{}
	res.observable = NewObservable(res)
	return res
}

func (this *DuckCall) RegisterObserver(observer Observer) {
	this.observable.RegisterObserver(observer)
}

func (this *DuckCall) NotifyObservers() {
	this.observable.NotifyObservers()
}

func (this *DuckCall) Quack() {
	fmt.Println("Kwak")
	this.NotifyObservers()
}

func (this *DuckCall) String() string {
	return "Duck Call"
}
