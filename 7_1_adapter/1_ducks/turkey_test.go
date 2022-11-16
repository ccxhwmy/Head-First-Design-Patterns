package ducks

import (
	"fmt"
	"testing"
)

func TestTurkeyAdapter(t *testing.T) {
	duck := new(MallardDuck)
	duckAdapter := NewDuckAdapter(duck)

	for i := 0; i < 10; i++ {
		fmt.Println("The DuckAdapter says...")
		duckAdapter.Gobble()
		duckAdapter.Fly()
	}
}
