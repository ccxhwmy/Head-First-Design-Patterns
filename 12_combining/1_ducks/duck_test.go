package ducks

import (
	"fmt"
	"testing"
)

func TestDuckCall(t *testing.T) {
	simulate()
}

func simulate() {
	mallardDuck := new(MallardDuck)
	redheadDuck := new(RedheadDuck)
	duckCall := new(DuckCall)
	rubberDuck := new(RubberDuck)

	fmt.Println("\nDuck Simulator")

	simulateQuack(mallardDuck)
	simulateQuack(redheadDuck)
	simulateQuack(duckCall)
	simulateQuack(rubberDuck)
}

func simulateQuack(duck Quackable) {
	duck.Quack()
}
