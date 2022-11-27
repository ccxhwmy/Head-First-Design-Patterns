package decorator

import "sync/atomic"

var (
	numberOfQuacks uint64 = 0
)

type QuackCounter struct {
	duck Quackable
}

func NewQuackCounter(duck Quackable) *QuackCounter {
	return &QuackCounter{duck: duck}
}

func (this *QuackCounter) Quack() {
	this.duck.Quack()
	atomic.AddUint64(&numberOfQuacks, 1)
}

func GetQuacks() uint64 {
	return numberOfQuacks
}

func (this *QuackCounter) String() string {
	return this.duck.String()
}
