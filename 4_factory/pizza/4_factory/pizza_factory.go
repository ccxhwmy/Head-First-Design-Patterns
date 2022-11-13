package factory

type PizzaIngredientFactory interface {
	createDough() Dough
	createSauce() Sauce
	createCheese() Cheese
	createVeggies() []Veggies
	createPepperoni() Pepperoni
	createClam() Clams
}

type ChicagoPizzaIngredientFactory struct {
}

func (this *ChicagoPizzaIngredientFactory) createDough() Dough {
	return new(ThickCrustDough)
}

func (this *ChicagoPizzaIngredientFactory) createSauce() Sauce {
	return new(PlumTomatoSauce)
}

func (this *ChicagoPizzaIngredientFactory) createCheese() Cheese {
	return new(MozzarellaCheese)
}

func (this *ChicagoPizzaIngredientFactory) createVeggies() []Veggies {
	veggies := []Veggies{
		new(BlackOlives),
		new(Spinach),
		new(Eggplant),
	}
	return veggies
}

func (this *ChicagoPizzaIngredientFactory) createPepperoni() Pepperoni {
	return new(SlicedPepperoni)
}

func (this *ChicagoPizzaIngredientFactory) createClam() Clams {
	return new(FrozenClams)
}

type NYPizzaIngredientFactory struct {
}

func (this *NYPizzaIngredientFactory) createDough() Dough {
	return new(ThinCrustDough)
}

func (this *NYPizzaIngredientFactory) createSauce() Sauce {
	return new(MarinaraSauce)
}

func (this *NYPizzaIngredientFactory) createCheese() Cheese {
	return new(ReggianoCheese)
}

func (this *NYPizzaIngredientFactory) createVeggies() []Veggies {
	veggies := []Veggies{
		new(Garlic),
		new(Onion),
		new(Mushroom),
		new(RedPepper),
	}
	return veggies
}

func (this *NYPizzaIngredientFactory) createPepperoni() Pepperoni {
	return new(SlicedPepperoni)
}

func (this *NYPizzaIngredientFactory) createClam() Clams {
	return new(FrozenClams)
}
