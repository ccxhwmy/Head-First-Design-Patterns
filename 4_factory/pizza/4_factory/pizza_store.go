package factory

import "fmt"

type iPizzaStore interface {
	createPizza(item string) iPizza
}

type PizzaStoreManger struct {
}

func newPizzaStoreManager() *PizzaStoreManger {
	return &PizzaStoreManger{}
}

func (this *PizzaStoreManger) orderPizza(store iPizzaStore, typ string) iPizza {
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

func newChicagoPizzaStore() *ChicagoPizzaStore {
	return &ChicagoPizzaStore{}
}

func (this *ChicagoPizzaStore) createPizza(item string) iPizza {
	ingredientFactory := new(ChicagoPizzaIngredientFactory)
	var pizza iPizza
	if item == "cheese" {
		pizza = NewCheesePizza(ingredientFactory)
		pizza.SetName("Chicago Style Cheese Pizza")
	} else if item == "veggie" {
		pizza = NewVeggiePizza(ingredientFactory)
		pizza.SetName("Chicago Style Veggie Pizza")
	} else if item == "clam" {
		pizza = NewClamPizza(ingredientFactory)
		pizza.SetName("Chicago Style Clam Pizza")
	} else if item == "pepperoni" {
		pizza = NewPepperoniPizza(ingredientFactory)
		pizza.SetName("Chicago Style Pepperoni Pizza")
	} else {
		return nil
	}
	return pizza
}

type NYPizzaStore struct {
}

func newNYPizzaStore() *NYPizzaStore {
	return &NYPizzaStore{}
}

func (this *NYPizzaStore) createPizza(item string) iPizza {
	ingredientFactory := new(NYPizzaIngredientFactory)
	var pizza iPizza
	if item == "cheese" {
		pizza = NewCheesePizza(ingredientFactory)
		pizza.SetName("New York Style Cheese Pizza")
	} else if item == "veggie" {
		pizza = NewVeggiePizza(ingredientFactory)
		pizza.SetName("New York Style Veggie Pizza")
	} else if item == "clam" {
		pizza = NewClamPizza(ingredientFactory)
		pizza.SetName("New York Style Clam Pizza")
	} else if item == "pepperoni" {
		pizza = NewPepperoniPizza(ingredientFactory)
		pizza.SetName("New York Style Pepperoni Pizza")
	} else {
		return nil
	}
	return pizza
}
