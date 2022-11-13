package factory

import "fmt"

type PizzaStore interface {
	createPizza(item string) iPizza
}

type PizzaStoreManger struct {
}

func newPizzaStoreManager() *PizzaStoreManger {
	return &PizzaStoreManger{}
}

func (this *PizzaStoreManger) orderPizza(store PizzaStore, typ string) iPizza {
	var pizza iPizza

	pizza = store.createPizza(typ)
	fmt.Println("--- Make a ", pizza.GetName(), " ---")

	pizza.Prepare()
	pizza.Bake()
	pizza.Cut()
	pizza.Box()

	return pizza
}

type ChicagoPizzaStore struct {
}

func newChicagoPizzaStore() PizzaStore {
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
}

func newNYPizzaStore() PizzaStore {
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
