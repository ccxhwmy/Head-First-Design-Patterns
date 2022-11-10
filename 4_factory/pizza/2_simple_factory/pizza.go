package simple_factory

import "fmt"

type iPizza interface {
	Prepare()
	Bake()
	Cut()
	Box()
	GetName() string
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

type CheesePizza struct {
	*Pizza
}

func NewCheesePizza() iPizza {
	return &CheesePizza{&Pizza{
		name:  "CheesePizza",
		dough: "CheeseDough",
		sauce: "CheeseSauce",
		toppings: []string{
			"CheeseCumin", "CheesePepper",
		},
	}}
}

type PepperoniPizza struct {
	*Pizza
}

func NewPepperoniPizza() iPizza {
	return &PepperoniPizza{&Pizza{
		name:  "PepperoniPizza",
		dough: "PepperoniDough",
		sauce: "PepperoniSauce",
		toppings: []string{
			"PepperoniCumin", "PepperoniPepper",
		},
	}}
}

func (this *PepperoniPizza) Cut() {
	fmt.Println("I am Pepperoni Pizza, I will cut Pizza in a different way!")
}

type ClamPizza struct {
	*Pizza
}

func NewClamPizza() iPizza {
	return &PepperoniPizza{&Pizza{
		name:  "ClamPizza",
		dough: "ClamDough",
		sauce: "ClamSauce",
		toppings: []string{
			"ClamCumin", "ClamPepper",
		},
	}}
}

type VeggiePizza struct {
	*Pizza
}

func NewVeggiePizza() iPizza {
	return &PepperoniPizza{&Pizza{
		name:  "VeggiePizza",
		dough: "VeggieDough",
		sauce: "VeggieSauce",
		toppings: []string{
			"VeggieCumin", "VeggiePepper",
		},
	}}
}
