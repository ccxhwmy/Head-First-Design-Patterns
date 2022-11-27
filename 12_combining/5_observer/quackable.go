package observer

type Quackable interface {
	QuackObservable
	Quack()
	String() string
}
