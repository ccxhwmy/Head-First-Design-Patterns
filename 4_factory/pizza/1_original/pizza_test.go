package original_pizza

import "testing"

func TestOrderPizza(t *testing.T) {
	OrderPizza("cheese")

	OrderPizza("pepperoni")

	OrderPizza("clam")

	OrderPizza("veggie")
}
