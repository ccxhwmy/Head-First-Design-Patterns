package factory

type AbstractDuckFactory interface {
	CreateMallardDuck() Quackable
	CreateRedheadDuck() Quackable
	CreateDuckCall() Quackable
	CreateRubberDuck() Quackable
}

type DuckFactory struct {
}

func (this *DuckFactory) CreateMallardDuck() Quackable {
	return new(MallardDuck)
}

func (this *DuckFactory) CreateRedheadDuck() Quackable {
	return new(RedheadDuck)
}

func (this *DuckFactory) CreateDuckCall() Quackable {
	return new(DuckCall)
}

func (this *DuckFactory) CreateRubberDuck() Quackable {
	return new(RubberDuck)
}
