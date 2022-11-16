package barista

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type CaffeineBeverageWithHook interface {
	CaffeineBeverage
	CustomerWantsCondiments() bool
	UserPrompt() string
}

func PrepareRecipeWithHook(caffeineBeverageWithHook CaffeineBeverageWithHook) {
	caffeineBeverageWithHook.BoilWater()
	caffeineBeverageWithHook.Brew()
	caffeineBeverageWithHook.PourInCup()
	fmt.Print(caffeineBeverageWithHook.UserPrompt(), " (y/n)?: ")
	if caffeineBeverageWithHook.CustomerWantsCondiments() {
		caffeineBeverageWithHook.AddCondiments()
	}
}

type CaffeineBeverageWithHookClass struct {
}

func (this *CaffeineBeverageWithHookClass) BoilWater() {
	fmt.Println("Boiling water")
}

func (this *CaffeineBeverageWithHookClass) PourInCup() {
	fmt.Println("Pouring into cup")
}

func (this *CaffeineBeverageWithHookClass) CustomerWantsCondiments() bool {
	if strings.HasPrefix(strings.ToLower(this.getUserInput()), "y") {
		return true
	}
	return false
}

func (this *CaffeineBeverageWithHookClass) getUserInput() string {
	inputReader := bufio.NewReader(os.Stdin)
	input, err := inputReader.ReadString('\n')
	if err != nil {
		return "no"
	}
	return input
}
