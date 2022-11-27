package diner_merger

import "fmt"

type Waitress struct {
	pancakeHouseMenu Menu
	dinerMenu        Menu
}

func NewWaitress(pancakeHouseMenu, dinerMenu Menu) *Waitress {
	return &Waitress{
		pancakeHouseMenu: pancakeHouseMenu,
		dinerMenu:        dinerMenu,
	}
}

func (this *Waitress) PrintMenu() {
	pancakeIterator := this.pancakeHouseMenu.CreateIterator()
	dinerIterator := this.dinerMenu.CreateIterator()

	fmt.Println("MENU\n----\nBREAKFAST")
	this.printMenuIterator(pancakeIterator)
	fmt.Println("\nLUNCH")
	this.printMenuIterator(dinerIterator)
}

func (this *Waitress) printMenuIterator(iterator Iterator) {
	for iterator.HasNext() {
		menuItem := iterator.Next()
		fmt.Print(menuItem.GetName() + ", ")
		fmt.Print(fmt.Sprintf("%.2f", menuItem.GetPrice()) + " -- ")
		fmt.Println(menuItem.GetDescription())
	}
}

func (this *Waitress) PrintVegetarianMenu() {
	this.printVegetarianMenuIterator(this.pancakeHouseMenu.CreateIterator())
	this.printVegetarianMenuIterator(this.dinerMenu.CreateIterator())
}

func (this *Waitress) printVegetarianMenuIterator(iterator Iterator) {
	for iterator.HasNext() {
		menuItem := iterator.Next()
		if menuItem.IsVegetarian() {
			fmt.Print(menuItem.GetName())
			fmt.Println("\t\t" + fmt.Sprintf("%.2f", menuItem.GetPrice()))
			fmt.Println("\t" + menuItem.GetDescription())
		}
	}
}

func (this *Waitress) IsItemVegetarian(name string) bool {
	breakfastIterator := this.pancakeHouseMenu.CreateIterator()
	if this.isVegetarian(name, breakfastIterator) {
		return true
	}
	if this.isVegetarian(name, this.dinerMenu.CreateIterator()) {
		return true
	}
	return false
}

func (this *Waitress) isVegetarian(name string, iterator Iterator) bool {
	for iterator.HasNext() {
		menuItem := iterator.Next()
		if menuItem.GetName() == name {
			if menuItem.IsVegetarian() {
				return true
			}
		}
	}
	return false
}
