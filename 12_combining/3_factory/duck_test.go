package factory

import (
	"fmt"
	"testing"
)

func TestDuckCall(t *testing.T) {
	duckFactory := new(CountingDuckFactory)
	simulate(duckFactory)
}

func simulate(duckFactory AbstractDuckFactory) {
	mallardDuck := duckFactory.CreateMallardDuck()
	redheadDuck := duckFactory.CreateRedheadDuck()
	duckCall := duckFactory.CreateDuckCall()
	rubberDuck := NewQuackCounter(new(RubberDuck))
	gooseDuck := NewGooseAdapter(new(Goose))

	fmt.Println("\nDuck Simulator")

	simulateQuack(mallardDuck)
	simulateQuack(redheadDuck)
	simulateQuack(duckCall)
	simulateQuack(rubberDuck)
	simulateQuack(gooseDuck)

	fmt.Println("The ducks quacked " +
		fmt.Sprintf("%d", GetQuacks()) + " times")
}

func simulateQuack(duck Quackable) {
	duck.Quack()
}
