package observer

type GooseAdapter struct {
	goose      *Goose
	observable *Observable
}

func NewGooseAdapter(goose *Goose) *GooseAdapter {
	res := &GooseAdapter{goose: goose}
	res.observable = NewObservable(res)
	return res
}

func (this *GooseAdapter) RegisterObserver(observer Observer) {
	this.observable.RegisterObserver(observer)
}

func (this *GooseAdapter) NotifyObservers() {
	this.observable.NotifyObservers()
}

func (this *GooseAdapter) Quack() {
	this.goose.Honk()
	this.NotifyObservers()
}

func (this *GooseAdapter) String() string {
	return "Goose pretending to be a Duck"
}
