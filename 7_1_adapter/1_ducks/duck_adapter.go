package ducks

import "math/rand"

type DuckAdapter struct {
	duck Duck
	rand rand.Rand
}

func NewDuckAdapter(duck Duck) *DuckAdapter {
	return &DuckAdapter{
		duck: duck,
		rand: rand.Rand{},
	}
}

func (this *DuckAdapter) Gobble() {
	this.duck.Quack()
}

func (this *DuckAdapter) Fly() {
	if rand.Int()%5 == 0 {
		this.duck.Fly()
	}
}
