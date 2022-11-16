package simple_barista

import "fmt"

type Tea struct {
}

func (this *Tea) PrepareRecipe() {
	this.BoilWater()
	this.steepTeaBag()
	this.pourInCup()
	this.addLemon()
}

func (this *Tea) BoilWater() {
	fmt.Println("Boiling water")
}

func (this *Tea) steepTeaBag() {
	fmt.Println("Steeping the tea")
}

func (this *Tea) pourInCup() {
	fmt.Println("Pouring into cup")
}

func (this *Tea) addLemon() {
	fmt.Println("Adding Lemon")
}
