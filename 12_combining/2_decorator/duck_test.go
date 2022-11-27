package decorator

import (
	"fmt"
	"testing"
)

func TestDuckCall(t *testing.T) {
	simulate()
}

func simulate() {
	mallardDuck := NewQuackCounter(new(MallardDuck))
	redheadDuck := NewQuackCounter(new(RedheadDuck))
	duckCall := NewQuackCounter(new(DuckCall))
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
