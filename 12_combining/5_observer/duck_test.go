package observer

import (
	"fmt"
	"testing"
)

func TestDuckCall(t *testing.T) {
	duckFactory := new(CountingDuckFactory)
	simulate(duckFactory)
}

func simulate(duckFactory AbstractDuckFactory) {
	redheadDuck := duckFactory.CreateRedheadDuck()
	duckCall := duckFactory.CreateDuckCall()
	rubberDuck := duckFactory.CreateRubberDuck()
	gooseDuck := NewGooseAdapter(new(Goose))

	fmt.Println("\nDuck Simulator: With Composite - Flocks")

	flockOfDucks := NewFlock()

	flockOfDucks.Add(redheadDuck)
	flockOfDucks.Add(duckCall)
	flockOfDucks.Add(rubberDuck)
	flockOfDucks.Add(gooseDuck)

	flockOfMallards := NewFlock()

	mallardOne := duckFactory.CreateMallardDuck()
	mallardTwo := duckFactory.CreateMallardDuck()
	mallardThree := duckFactory.CreateMallardDuck()
	mallardFour := duckFactory.CreateMallardDuck()

	flockOfMallards.Add(mallardOne)
	flockOfMallards.Add(mallardTwo)
	flockOfMallards.Add(mallardThree)
	flockOfMallards.Add(mallardFour)

	flockOfDucks.Add(flockOfMallards)

	fmt.Println("\nDuck Simulator: With Observer")

	quackologist := new(Quackologist)
	flockOfDucks.RegisterObserver(quackologist)

	simulateQuack(flockOfDucks)

	fmt.Println("The ducks quacked " +
		fmt.Sprintf("%d", GetQuacks()) + " times")
}

func simulateQuack(duck Quackable) {
	duck.Quack()
}
