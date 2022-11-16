package barista

import "fmt"

type CaffeineBeverage interface {
	BoilWater()
	Brew()
	PourInCup()
	AddCondiments()
}

func PrepareRecipe(caffeineBeverage CaffeineBeverage) {
	caffeineBeverage.BoilWater()
	caffeineBeverage.Brew()
	caffeineBeverage.PourInCup()
	caffeineBeverage.AddCondiments()
}

type CaffeineBeverageClass struct {
}

func (this *CaffeineBeverageClass) BoilWater() {
	fmt.Println("Boiling water")
}

func (this *CaffeineBeverageClass) PourInCup() {
	fmt.Println("Pouring into cup")
}
