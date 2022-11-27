package composite

type Flock struct {
	quackers []Quackable
}

func NewFlock() *Flock {
	return &Flock{quackers: []Quackable{}}
}

func (this *Flock) Add(quacker Quackable) {
	this.quackers = append(this.quackers, quacker)
}

func (this *Flock) Quack() {
	for _, quacker := range this.quackers {
		quacker.Quack()
	}
}

func (this *Flock) String() string {
	return "Flock of Quackers"
}
