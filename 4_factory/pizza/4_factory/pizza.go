package factory

import (
	"bytes"
	"fmt"
)

type iPizza interface {
	Prepare()
	Bake()
	Cut()
	Box()
	GetName() string
	SetName(name string)
}

type Pizza struct {
	name      string
	dough     Dough
	sauce     Sauce
	veggies   []Veggies
	cheese    Cheese
	pepperoni Pepperoni
	clam      Clams
}

func (this *Pizza) Prepare() {
	fmt.Println("Preparing Pizza ...")
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

func (this *Pizza) SetName(name string) {
	this.name = name
}

func (this *Pizza) GetName() string {
	return this.name
}

func (this *Pizza) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("--- " + this.name + " ---")
	if this.dough != nil {
		buffer.WriteString(this.dough.String())
		buffer.WriteString("\n")
	}
	if this.sauce != nil {
		buffer.WriteString(this.sauce.String())
		buffer.WriteString("\n")
	}
	if this.cheese != nil {
		buffer.WriteString(this.cheese.String())
		buffer.WriteString("\n")
	}
	if this.veggies != nil {
		for i, vv := range this.veggies {
			buffer.WriteString(vv.String())
			if i < len(this.veggies)-1 {
				buffer.WriteString(", ")
			}
		}
		buffer.WriteString("\n")
	}
	if this.clam != nil {
		buffer.WriteString(this.clam.String())
		buffer.WriteString("\n")
	}
	if this.pepperoni != nil {
		buffer.WriteString(this.pepperoni.String())
		buffer.WriteString("\n")
	}
	return buffer.String()
}

type CheesePizza struct {
	Pizza
	ingredientFactory PizzaIngredientFactory
}

func NewCheesePizza(ingredientFactory PizzaIngredientFactory) iPizza {
	return &CheesePizza{
		ingredientFactory: ingredientFactory,
	}
}

func (this *CheesePizza) Prepare() {
	fmt.Println("Preparing ", this.name)
	this.dough = this.ingredientFactory.createDough()
	this.sauce = this.ingredientFactory.createSauce()
	this.cheese = this.ingredientFactory.createCheese()
}

type PepperoniPizza struct {
	Pizza
	ingredientFactory PizzaIngredientFactory
}

func NewPepperoniPizza(ingredientFactory PizzaIngredientFactory) iPizza {
	return &PepperoniPizza{
		ingredientFactory: ingredientFactory,
	}
}

func (this *PepperoniPizza) Prepare() {
	fmt.Println("Preparing ", this.name)
	this.dough = this.ingredientFactory.createDough()
	this.sauce = this.ingredientFactory.createSauce()
	this.cheese = this.ingredientFactory.createCheese()
	this.veggies = this.ingredientFactory.createVeggies()
	this.pepperoni = this.ingredientFactory.createPepperoni()
}

type ClamPizza struct {
	Pizza
	ingredientFactory PizzaIngredientFactory
}

func NewClamPizza(ingredientFactory PizzaIngredientFactory) iPizza {
	return &ClamPizza{
		ingredientFactory: ingredientFactory,
	}
}

func (this *ClamPizza) Prepare() {
	fmt.Println("Preparing ", this.name)
	this.dough = this.ingredientFactory.createDough()
	this.sauce = this.ingredientFactory.createSauce()
	this.cheese = this.ingredientFactory.createCheese()
	this.clam = this.ingredientFactory.createClam()
}

type VeggiePizza struct {
	Pizza
	ingredientFactory PizzaIngredientFactory
}

func NewVeggiePizza(ingredientFactory PizzaIngredientFactory) iPizza {
	return &VeggiePizza{
		ingredientFactory: ingredientFactory,
	}
}

func (this *VeggiePizza) Prepare() {
	fmt.Println("Preparing ", this.name)
	this.dough = this.ingredientFactory.createDough()
	this.sauce = this.ingredientFactory.createSauce()
	this.cheese = this.ingredientFactory.createCheese()
	this.veggies = this.ingredientFactory.createVeggies()
}
