package diner_merger

import "fmt"

const MAX_ITEMS int = 6

type DinerMenu struct {
	numberOfItems int
	menuItems     []MenuItem
}

func NewDinerMenu() *DinerMenu {
	dinerMenu := &DinerMenu{
		numberOfItems: 0,
		menuItems:     []MenuItem{},
	}

	dinerMenu.AddItem("Vegetarian BLT",
		"(Fakin') Bacon with lettuce & tomato on whole wheat", true, 2.99)
	dinerMenu.AddItem("BLT",
		"Bacon with lettuce & tomato on whole wheat", false, 2.99)
	dinerMenu.AddItem("Soup of the day",
		"Soup of the day, with a side of potato salad", false, 3.29)
	dinerMenu.AddItem("Hotdog",
		"A hot dog, with sauerkraut, relish, onions, topped with cheese",
		false, 3.05)
	dinerMenu.AddItem("Steamed Veggies and Brown Rice",
		"Steamed vegetables over brown rice", true, 3.99)
	dinerMenu.AddItem("Pasta",
		"Spaghetti with Marinara Sauce, and a slice of sourdough bread",
		true, 3.89)
	return dinerMenu
}

func (this *DinerMenu) AddItem(name, description string, vegetarian bool, price float64) {
	if this.numberOfItems >= MAX_ITEMS {
		fmt.Println("Sorry, menu is full!  Can't add item to menu")
	} else {
		this.menuItems = append(this.menuItems, *NewMenuItem(name, description, vegetarian, price))
		this.numberOfItems++
	}
}

func (this *DinerMenu) GetMenuItems() []MenuItem {
	return this.menuItems
}

func (this *DinerMenu) CreateIterator() Iterator {
	return NewDinerMenuIterator(this.menuItems)
}
