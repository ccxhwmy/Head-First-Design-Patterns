package ducks

import (
	"fmt"
	"testing"
)

func TestDuckAdapter(t *testing.T) {
	duck := new(MallardDuck)
	turkey := new(WildTurkey)
	turkeyAdapter := NewTurkeyAdapter(turkey)

	fmt.Println("The Turkey says...")
	turkey.Gobble()
	turkey.Fly()

	fmt.Println("\nThe Duck says...")
	duck.Quack()
	duck.Fly()

	fmt.Println("\nThe TurkeyAdapter says...")
	turkeyAdapter.Quack()
	turkeyAdapter.Fly()
}
