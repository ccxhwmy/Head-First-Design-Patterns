package observer

type Flock struct {
	quackers []Quackable
}

func NewFlock() *Flock {
	return &Flock{quackers: []Quackable{}}
}

func (this *Flock) Add(quacker Quackable) {
	this.quackers = append(this.quackers, quacker)
}

func (this *Flock) RegisterObserver(observer Observer) {
	for _, quacker := range this.quackers {
		quacker.RegisterObserver(observer)
	}
}

func (this *Flock) NotifyObservers() {
	//for _, quacker := range this.quackers {
	//	quacker.NotifyObservers()
	//}
}

func (this *Flock) Quack() {
	for _, quacker := range this.quackers {
		quacker.Quack()
	}
}

func (this *Flock) String() string {
	return "Flock of Quackers"
}
