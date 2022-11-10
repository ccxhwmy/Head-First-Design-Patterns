package factory

import (
	"fmt"
	"testing"
)

func TestPizza(t *testing.T) {
	nyStore := newNYPizzaStore()
	chicagoStore := newChicagoPizzaStore()

	pizza := nyStore.orderPizza("cheese")
	fmt.Println("Ethan ordered a", pizza.GetName())

	pizza = chicagoStore.orderPizza("cheese")
	fmt.Println("Ethan ordered a", pizza.GetName())

	pizza = nyStore.orderPizza("clam")
	fmt.Println("Ethan ordered a", pizza.GetName())

	pizza = chicagoStore.orderPizza("clam")
	fmt.Println("Ethan ordered a", pizza.GetName())

	pizza = nyStore.orderPizza("pepperoni")
	fmt.Println("Ethan ordered a", pizza.GetName())

	pizza = chicagoStore.orderPizza("pepperoni")
	fmt.Println("Ethan ordered a", pizza.GetName())

	pizza = nyStore.orderPizza("veggie")
	fmt.Println("Ethan ordered a", pizza.GetName())

	pizza = chicagoStore.orderPizza("veggie")
	fmt.Println("Ethan ordered a", pizza.GetName())
}
