package observer

type AbstractDuckFactory interface {
	CreateMallardDuck() Quackable
	CreateRedheadDuck() Quackable
	CreateDuckCall() Quackable
	CreateRubberDuck() Quackable
}

type DuckFactory struct {
}

func (this *DuckFactory) CreateMallardDuck() Quackable {
	return NewMallardDuck()
}

func (this *DuckFactory) CreateRedheadDuck() Quackable {
	return NewRedheadDuck()
}

func (this *DuckFactory) CreateDuckCall() Quackable {
	return NewDuckCall()
}

func (this *DuckFactory) CreateRubberDuck() Quackable {
	return NewRubberDuck()
}
