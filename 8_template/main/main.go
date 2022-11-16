package main

import (
	"fmt"
	barista "headfirstdesign/8_template/2_barista"
)

func baristaWithHook() {
	fmt.Println("\nMaking tea...")
	barista.PrepareRecipeWithHook(new(barista.TeaWithHook))

	fmt.Println("\nMaking coffee...")
	barista.PrepareRecipeWithHook(new(barista.CoffeeWithHook))
}

func main() {
	baristaWithHook()
}
