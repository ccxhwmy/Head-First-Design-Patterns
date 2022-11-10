package simple_factory

import (
	"fmt"
	"testing"
)

func TestSimplePizzaFactory(t *testing.T) {
	factory := NewSimplePizzaFactory()
	store := NewPizzaStore(factory)

	pizza := store.orderPizza("cheese")
	fmt.Println("We order a ", pizza.GetName())

	pizza = store.orderPizza("veggie")
	fmt.Println("We order a ", pizza.GetName())

	pizza = store.orderPizza("pepperoni")
	fmt.Println("We order a ", pizza.GetName())
}
