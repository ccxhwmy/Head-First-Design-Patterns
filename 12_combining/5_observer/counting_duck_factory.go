package observer

type CountingDuckFactory struct {
}

func (this *CountingDuckFactory) CreateMallardDuck() Quackable {
	return NewQuackCounter(NewMallardDuck())
}

func (this *CountingDuckFactory) CreateRedheadDuck() Quackable {
	return NewQuackCounter(NewRedheadDuck())
}

func (this *CountingDuckFactory) CreateDuckCall() Quackable {
	return NewQuackCounter(NewDuckCall())
}

func (this *CountingDuckFactory) CreateRubberDuck() Quackable {
	return NewQuackCounter(NewRubberDuck())
}
