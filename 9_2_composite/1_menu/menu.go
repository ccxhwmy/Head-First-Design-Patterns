package menu

import (
	"fmt"
	"github.com/rs/xid"
)

type Menu struct {
	menuComponents []MenuComponent
	size           int
	name           string
	description    string
	id             string
}

func NewMenu(name, description string) *Menu {
	return &Menu{
		name:        name,
		description: description,
		id:          xid.New().String(),
	}
}

func (this *Menu) Add(menuComponent MenuComponent) {
	this.menuComponents = append(this.menuComponents, menuComponent)
	this.size++
}

func (this *Menu) Remove(menuComponent MenuComponent) {
	for ii, vv := range this.menuComponents {
		if vv.GetUUID() == menuComponent.GetUUID() {
			this.menuComponents = append(this.menuComponents[:ii], this.menuComponents[ii+1:]...)
			this.size--
		}
	}
}

func (this *Menu) GetChild(i int) MenuComponent {
	return this.menuComponents[i]
}

func (this *Menu) GetName() string {
	return this.name
}

func (this *Menu) GetDescription() string {
	return this.description
}

func (this *Menu) GetPrice() float64 {
	return 0
}

func (this *Menu) IsVegetarian() bool {
	return false
}

func (this *Menu) Print() {
	fmt.Println("\n", this.GetName())
	fmt.Println(", ", this.GetDescription())
	fmt.Println("-------------------------")

	for _, vv := range this.menuComponents {
		vv.Print()
	}
}

func (this *Menu) GetUUID() string {
	return this.id
}
