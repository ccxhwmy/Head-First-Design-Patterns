package duck

import (
	"fmt"
	"testing"
)

func TestDuckPerform(t *testing.T) {
	fmt.Println("-------------- mallard duck show time -------------")
	mallardDuck := NewMallardDuck()
	mallardDuck.Display()
	mallardDuck.PerformFly()
	mallardDuck.PerformQuack()

	fmt.Println("-------------- model duck show time -------------")
	modelDuck := NewModelDuck()
	modelDuck.Display()
	modelDuck.PerformFly()
	modelDuck.PerformQuack()

	fmt.Println("-------------- model duck set rocket power -------------")
	modelDuck.SetRocketPowered()
	modelDuck.PerformFly()

	fmt.Println("-------------- model duck set buzzer -------------")
	modelDuck.SetBuzzerQuack()
	modelDuck.PerformQuack()
}
