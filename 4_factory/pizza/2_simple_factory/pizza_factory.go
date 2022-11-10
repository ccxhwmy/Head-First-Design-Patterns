package simple_factory

type PizzaFactory interface {
	createPizza(typ string) iPizza
}

type SimplePizzaFactory struct {
}

func NewSimplePizzaFactory() PizzaFactory {
	return &SimplePizzaFactory{}
}

func (this *SimplePizzaFactory) createPizza(typ string) iPizza {
	var pizza iPizza

	if typ == "cheese" {
		pizza = NewCheesePizza()
	} else if typ == "pepperoni" {
		pizza = NewPepperoniPizza()
	} else if typ == "clam" {
		pizza = NewClamPizza()
	} else if typ == "veggie" {
		pizza = NewVeggiePizza()
	}

	if pizza == nil {
		pizza = NewCheesePizza()
	}

	return pizza
}
