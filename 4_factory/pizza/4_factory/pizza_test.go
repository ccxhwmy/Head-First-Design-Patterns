package factory

import (
	"fmt"
	"testing"
)

func TestPizza(t *testing.T) {
	storeMgr := newPizzaStoreManager()
	nyStore := newNYPizzaStore()
	chicagoStore := newChicagoPizzaStore()

	pizza := storeMgr.orderPizza(nyStore, "cheese")
	fmt.Println("Ethan ordered a", pizza.GetName())

	pizza = storeMgr.orderPizza(chicagoStore, "cheese")
	fmt.Println("Ethan ordered a", pizza.GetName())

	pizza = storeMgr.orderPizza(nyStore, "clam")
	fmt.Println("Ethan ordered a", pizza.GetName())

	pizza = storeMgr.orderPizza(chicagoStore, "clam")
	fmt.Println("Ethan ordered a", pizza.GetName())

	pizza = storeMgr.orderPizza(nyStore, "pepperoni")
	fmt.Println("Ethan ordered a", pizza.GetName())

	pizza = storeMgr.orderPizza(chicagoStore, "pepperoni")
	fmt.Println("Ethan ordered a", pizza.GetName())

	pizza = storeMgr.orderPizza(nyStore, "veggie")
	fmt.Println("Ethan ordered a", pizza.GetName())

	pizza = storeMgr.orderPizza(chicagoStore, "veggie")
	fmt.Println("Ethan ordered a", pizza.GetName())
}
