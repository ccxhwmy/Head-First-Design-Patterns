package original_pizza

import "fmt"

type iPizza interface {
	Prepare()
	Bake()
	Cut()
	Box()
}

type Pizza struct {
	name     string
	dough    string
	sauce    string
	toppings []string
}

func (this *Pizza) Prepare() {
	fmt.Println("Preparing ", this.name)
	fmt.Println("Tossing dough...")
	fmt.Println("Adding sauce...")
	fmt.Println("Adding toppings: ")
	for i := 0; i < len(this.toppings); i++ {
		fmt.Println("    ", this.toppings[i])
	}
}

func (this *Pizza) Bake() {
	fmt.Println("Bake for 25 minute at 350")
}

func (this *Pizza) Cut() {
	fmt.Println("Cutting the pizza into diagonal slices")
}

func (this *Pizza) Box() {
	fmt.Println("Place pizza in official PizzaStore box")
}

func (this *Pizza) GetName() string {
	return this.name
}

func newCheesePizza() *Pizza {
	return &Pizza{
		name:  "CheesePizza",
		dough: "CheeseDough",
		sauce: "CheeseSauce",
		toppings: []string{
			"CheeseCumin", "CheesePepper",
		},
	}
}

func newPepperoniPizza() *Pizza {
	return &Pizza{
		name:  "PepperoniPizza",
		dough: "PepperoniDough",
		sauce: "PepperoniSauce",
		toppings: []string{
			"PepperoniCumin", "PepperoniPepper",
		},
	}
}

func newClamPizza() *Pizza {
	return &Pizza{
		name:  "ClamPizza",
		dough: "ClamDough",
		sauce: "ClamSauce",
		toppings: []string{
			"ClamCumin", "ClamPepper",
		},
	}
}

func newVeggiePizza() *Pizza {
	return &Pizza{
		name:  "VeggiePizza",
		dough: "VeggieDough",
		sauce: "VeggieSauce",
		toppings: []string{
			"VeggieCumin", "VeggiePepper",
		},
	}
}

func newNYCheesePizza() *Pizza {
	return &Pizza{
		name:  "NYCheesePizza",
		dough: "NYCheeseDough",
		sauce: "NYCheeseSauce",
		toppings: []string{
			"NYCheeseCumin", "NYCheesePepper",
		},
	}
}

func newNYPepperoniPizza() *Pizza {
	return &Pizza{
		name:  "NYPepperoniPizza",
		dough: "NYPepperoniDough",
		sauce: "NYPepperoniSauce",
		toppings: []string{
			"NYPepperoniCumin", "NYPepperoniPepper",
		},
	}
}

func newNYClamPizza() *Pizza {
	return &Pizza{
		name:  "NYClamPizza",
		dough: "NYClamDough",
		sauce: "NYClamSauce",
		toppings: []string{
			"NYClamCumin", "NYClamPepper",
		},
	}
}

func newNYVeggiePizza() *Pizza {
	return &Pizza{
		name:  "NYVeggiePizza",
		dough: "NYVeggieDough",
		sauce: "NYVeggieSauce",
		toppings: []string{
			"NYVeggieCumin", "NYVeggiePepper",
		},
	}
}

func newChicagoCheesePizza() *Pizza {
	return &Pizza{
		name:  "ChicagoCheesePizza",
		dough: "ChicagoCheeseDough",
		sauce: "ChicagoCheeseSauce",
		toppings: []string{
			"ChicagoCheeseCumin", "ChicagoCheesePepper",
		},
	}
}

func newChicagoPepperoniPizza() *Pizza {
	return &Pizza{
		name:  "ChicagoPepperoniPizza",
		dough: "ChicagoPepperoniDough",
		sauce: "ChicagoPepperoniSauce",
		toppings: []string{
			"ChicagoPepperoniCumin", "ChicagoPepperoniPepper",
		},
	}
}

func newChicagoClamPizza() *Pizza {
	return &Pizza{
		name:  "ChicagoClamPizza",
		dough: "ChicagoClamDough",
		sauce: "ChicagoClamSauce",
		toppings: []string{
			"ChicagoClamCumin", "ChicagoClamPepper",
		},
	}
}

func newChicagoVeggiePizza() *Pizza {
	return &Pizza{
		name:  "ChicagoVeggiePizza",
		dough: "ChicagoVeggieDough",
		sauce: "ChicagoVeggieSauce",
		toppings: []string{
			"ChicagoVeggieCumin", "ChicagoVeggiePepper",
		},
	}
}

func OrderPizza(typ string) *Pizza {
	var pizza *Pizza

	if typ == "cheese" {
		pizza = newCheesePizza()
	} else if typ == "pepperoni" {
		pizza = newPepperoniPizza()
	} else if typ == "clam" {
		pizza = newClamPizza()
	} else if typ == "veggie" {
		pizza = newVeggiePizza()
	}

	if pizza == nil {
		pizza = newCheesePizza()
	}

	pizza.Prepare()
	pizza.Bake()
	pizza.Cut()
	pizza.Box()

	return pizza
}
