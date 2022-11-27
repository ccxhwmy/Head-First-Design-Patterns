package observer

type Observable struct {
	observers []Observer
	duck      QuackObservable
}

func NewObservable(duck QuackObservable) *Observable {
	return &Observable{
		observers: []Observer{},
		duck:      duck,
	}
}

func (this *Observable) RegisterObserver(observer Observer) {
	this.observers = append(this.observers, observer)
}

func (this *Observable) NotifyObservers() {
	for _, observer := range this.observers {
		observer.Update(this.duck)
	}
}
