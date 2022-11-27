package menu

import (
	"fmt"
	"github.com/rs/xid"
)

type MenuItem struct {
	name        string
	description string
	vegetarian  bool
	price       float64
	id          string
}

func NewMenuItem(name, description string, vegetarian bool, price float64) *MenuItem {
	return &MenuItem{
		name:        name,
		description: description,
		vegetarian:  vegetarian,
		price:       price,
		id:          xid.New().String(),
	}
}

func (this *MenuItem) Add(menuComponent MenuComponent) {
	return
}

func (this *MenuItem) Remove(menuComponent MenuComponent) {

}

func (this *MenuItem) GetChild(i int) MenuComponent {
	return nil
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

func (this *MenuItem) Print() {
	fmt.Print("  ", this.GetName())
	if this.IsVegetarian() {
		fmt.Print("(v)")
	}
	fmt.Println(", ", this.GetPrice())
	fmt.Println("    -- ", this.GetDescription())
}

func (this *MenuItem) GetUUID() string {
	return this.id
}
