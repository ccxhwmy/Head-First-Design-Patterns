package simple_barista

import "fmt"

type Coffee struct {
}

func (this *Coffee) PrepareRecipe() {
	this.BoilWater()
	this.BrewCoffeeGrinds()
	this.PourInCup()
	this.AddSugarAndMilk()
}

func (this *Coffee) BoilWater() {
	fmt.Println("Boiling water")
}

func (this *Coffee) BrewCoffeeGrinds() {
	fmt.Println("Dripping Coffee through filter")
}

func (this *Coffee) PourInCup() {
	fmt.Println("Pouring into cup")
}

func (this *Coffee) AddSugarAndMilk() {
	fmt.Println("Adding Sugar and Milk")
}
