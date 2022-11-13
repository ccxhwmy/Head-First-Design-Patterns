package chocolate

import "fmt"

var (
	uniqueInstance *ChocolateBoiler
)

type ChocolateBoiler struct {
	empty  bool
	boiled bool
}

func GetInstance() *ChocolateBoiler {
	if uniqueInstance == nil {
		fmt.Println("Creating unique instance of Chocolate Boiler")
		uniqueInstance = new(ChocolateBoiler)
	}
	fmt.Println("Returning instance of Chocolate Boiler")
	return uniqueInstance
}

func (this *ChocolateBoiler) fill() {
	if this.isEmpty() {
		this.empty = false
		this.boiled = false
	}
}

func (this *ChocolateBoiler) drain() {
	if !this.isEmpty() && this.isBoiled() {
		this.empty = true
	}
}

func (this *ChocolateBoiler) boil() {
	if !this.isEmpty() && !this.isBoiled() {
		this.boiled = true
	}
}

func (this *ChocolateBoiler) isEmpty() bool {
	return this.empty
}

func (this *ChocolateBoiler) isBoiled() bool {
	return this.boiled
}
