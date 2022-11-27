package diner_merger

import "container/list"

type PancakeHouseMenu struct {
	menuItems *list.List
}

func NewPancakeHouseMenu() *PancakeHouseMenu {
	pancakeHouseMenu := &PancakeHouseMenu{menuItems: list.New()}
	pancakeHouseMenu.AddItem("K&B's Pancake Breakfast",
		"Pancakes with scrambled eggs and toast",
		true,
		2.99)
	pancakeHouseMenu.AddItem("Regular Pancake Breakfast",
		"Pancakes with fried eggs, sausage",
		false,
		2.99)
	pancakeHouseMenu.AddItem("Blueberry Pancakes",
		"Pancakes made with fresh blueberries",
		true,
		3.49)
	pancakeHouseMenu.AddItem("Waffles",
		"Waffles with your choice of blueberries or strawberries",
		true,
		3.59)
	return pancakeHouseMenu
}

func (this *PancakeHouseMenu) AddItem(name, description string, vegetarian bool, price float64) {
	this.menuItems.PushBack(NewMenuItem(name, description, vegetarian, price))
}

func (this *PancakeHouseMenu) GetMenuItems() *list.List {
	return this.menuItems
}

func (this *PancakeHouseMenu) CreateIterator() Iterator {
	return NewPancakeHouseMenuIterator(this.menuItems)
}
