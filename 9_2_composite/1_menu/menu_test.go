package menu

import "testing"

func TestMenu(t *testing.T) {
	pancakeHouseMenu := NewMenu("PANCAKE HOUSE MENU", "Breakfast")
	dinerMenu := NewMenu("DINER MENU", "Lunch")
	cafeMenu := NewMenu("CAFE MENU", "Dinner")
	dessertMenu := NewMenu("DESSERT MENU", "Dessert of course!")
	coffeeMenu := NewMenu("COFFEE MENU", "Stuff to go with your afternoon coffee")

	allMenus := NewMenu("ALL MENUS", "All menus combined")

	allMenus.Add(pancakeHouseMenu)
	allMenus.Add(dinerMenu)
	allMenus.Add(cafeMenu)

	pancakeHouseMenu.Add(NewMenuItem("K&B's Pancake Breakfast",
		"Pancakes with scrambled eggs and toast",
		true,
		2.99))
	pancakeHouseMenu.Add(NewMenuItem("Regular Pancake Breakfast",
		"Pancakes with fried eggs, sausage",
		false,
		2.99))
	pancakeHouseMenu.Add(NewMenuItem("Blueberry Pancakes",
		"Pancakes made with fresh blueberries, and blueberry syrup",
		true,
		3.49))
	pancakeHouseMenu.Add(NewMenuItem("Waffles",
		"Waffles with your choice of blueberries or strawberries",
		true,
		3.59))

	dinerMenu.Add(NewMenuItem("BLT",
		"Bacon with lettuce & tomato on whole wheat",
		false,
		2.99))
	dinerMenu.Add(NewMenuItem("Soup of the day",
		"A bowl of the soup of the day, with a side of potato salad",
		false,
		3.29))
	dinerMenu.Add(NewMenuItem("Hot Dog",
		"A hot dog, with saurkraut, relish, onions, topped with cheese",
		false,
		3.05))
	dinerMenu.Add(NewMenuItem("Steamed Veggies and Brown Rice",
		"Steamed vegetables over brown rice",
		true,
		3.99))
	dinerMenu.Add(NewMenuItem("Pasta",
		"Spaghetti with marinara sauce, and a slice of sourdough bread",
		true,
		3.89))

	dinerMenu.Add(dessertMenu)

	dessertMenu.Add(NewMenuItem("Apple Pie",
		"Apple pie with a flakey crust, topped with vanilla icecream",
		true,
		1.59))
	dessertMenu.Add(NewMenuItem("Cheesecake",
		"Creamy New York cheesecake, with a chocolate graham crust",
		true,
		1.99))
	dessertMenu.Add(NewMenuItem("Sorbet",
		"A scoop of raspberry and a scoop of lime",
		true,
		1.89))

	cafeMenu.Add(NewMenuItem("Veggie Burger and Air Fries",
		"Veggie burger on a whole wheat bun, lettuce, tomato, and fries",
		true,
		3.99))
	cafeMenu.Add(NewMenuItem("Soup of the day",
		"A cup of the soup of the day, with a side salad",
		false,
		3.69))
	cafeMenu.Add(NewMenuItem("Burrito",
		"A large burrito, with whole pinto beans, salsa, guacamole",
		true,
		4.29))

	cafeMenu.Add(coffeeMenu)

	coffeeMenu.Add(NewMenuItem("Coffee Cake",
		"Crumbly cake topped with cinnamon and walnuts",
		true,
		1.59))
	coffeeMenu.Add(NewMenuItem("Bagel",
		"Flavors include sesame, poppyseed, cinnamon raisin, pumpkin",
		false,
		0.69))
	coffeeMenu.Add(NewMenuItem("Biscotti",
		"Three almond or hazelnut biscotti cookies",
		true,
		0.89))

	waitress := NewWaitress(allMenus)

	waitress.printMenu()
}
