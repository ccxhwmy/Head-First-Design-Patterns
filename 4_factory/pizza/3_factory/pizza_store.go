package factory

type iPizzaStore interface {
	orderPizza(typ string) iPizza
	createPizza(item string) iPizza
}

type PizzaStore struct {
}

func NewPizzaStore() *PizzaStore {
	return &PizzaStore{}
}

func (this *PizzaStore) orderPizza(typ string) iPizza {
	var pizza iPizza

	pizza = this.createPizza(typ)

	pizza.Prepare()
	pizza.Bake()
	pizza.Cut()
	pizza.Box()

	return pizza
}

func (this *PizzaStore) createPizza(item string) iPizza {
	var pizza iPizza

	if item == "cheese" {
		pizza = NewCheesePizza()
	} else if item == "pepperoni" {
		pizza = NewPepperoniPizza()
	} else if item == "clam" {
		pizza = NewClamPizza()
	} else if item == "veggie" {
		pizza = NewVeggiePizza()
	}

	if pizza == nil {
		pizza = NewCheesePizza()
	}

	return pizza
}

type ChicagoPizzaStore struct {
	*PizzaStore
}

func newChicagoPizzaStore() iPizzaStore {
	return &ChicagoPizzaStore{}
}

func (this *ChicagoPizzaStore) createPizza(item string) iPizza {
	if item == "cheese" {
		return newChicagoCheesePizza()
	} else if item == "veggie" {
		return newChicagoVeggiePizza()
	} else if item == "clam" {
		return newChicagoClamPizza()
	} else if item == "pepperoni" {
		return newChicagoPepperoniPizza()
	} else {
		return nil
	}
}

type NYPizzaStore struct {
	*PizzaStore
}

func newNYPizzaStore() iPizzaStore {
	return &NYPizzaStore{}
}

func (this *NYPizzaStore) createPizza(item string) iPizza {
	if item == "cheese" {
		return newNYCheesePizza()
	} else if item == "veggie" {
		return newNYVeggiePizza()
	} else if item == "clam" {
		return newNYClamPizza()
	} else if item == "pepperoni" {
		return newNYPepperoniPizza()
	} else {
		return nil
	}
}
