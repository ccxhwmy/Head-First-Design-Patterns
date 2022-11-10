package original_pizza

type DependentPizzaStore struct {
}

func (this *DependentPizzaStore) createPizza(style, typ string) iPizza {
	var pizza iPizza
	if style == "NY" {
		if typ == "cheese" {
			pizza = newNYCheesePizza()
		} else if typ == "pepperoni" {
			pizza = newNYPepperoniPizza()
		} else if typ == "clam" {
			pizza = newNYClamPizza()
		} else if typ == "veggie" {
			pizza = newNYVeggiePizza()
		}
	} else if style == "Chicago" {
		if typ == "cheese" {
			pizza = newChicagoCheesePizza()
		} else if typ == "pepperoni" {
			pizza = newChicagoPepperoniPizza()
		} else if typ == "clam" {
			pizza = newChicagoClamPizza()
		} else if typ == "veggie" {
			pizza = newChicagoVeggiePizza()
		}
	} else {
		return nil
	}

	pizza.Prepare()
	pizza.Bake()
	pizza.Cut()
	pizza.Box()

	return pizza
}
