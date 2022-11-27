package factory

type CountingDuckFactory struct {
}

func (this *CountingDuckFactory) CreateMallardDuck() Quackable {
	return NewQuackCounter(new(MallardDuck))
}

func (this *CountingDuckFactory) CreateRedheadDuck() Quackable {
	return NewQuackCounter(new(RedheadDuck))
}

func (this *CountingDuckFactory) CreateDuckCall() Quackable {
	return NewQuackCounter(new(DuckCall))
}

func (this *CountingDuckFactory) CreateRubberDuck() Quackable {
	return NewQuackCounter(new(RubberDuck))
}
