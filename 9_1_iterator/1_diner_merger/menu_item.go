package diner_merger

import (
	"fmt"
)

type MenuItem struct {
	name        string
	description string
	vegetarian  bool
	price       float64
}

func NewMenuItem(name, description string, vegetarian bool, price float64) *MenuItem {
	return &MenuItem{
		name:        name,
		description: description,
		vegetarian:  vegetarian,
		price:       price,
	}
}

func (this *MenuItem) GetName() string {
	return this.name
}

func (this *MenuItem) GetDescription() string {
	return this.description
}

func (this *MenuItem) GetPrice() float64 {
	return this.price
}

func (this *MenuItem) IsVegetarian() bool {
	return this.vegetarian
}

func (this *MenuItem) String() string {
	return this.name + ", $" + fmt.Sprintf("%.2f", this.price) + "\n " + this.description
}
