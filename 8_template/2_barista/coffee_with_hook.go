package barista

import (
	"fmt"
)

type CoffeeWithHook struct {
	CaffeineBeverageWithHookClass
}

func (this *CoffeeWithHook) Brew() {
	fmt.Println("Dripping Coffee through filter")
}

func (this *CoffeeWithHook) AddCondiments() {
	fmt.Println("Adding Sugar and Milk")
}

func (this *CoffeeWithHook) UserPrompt() string {
	return "Would you like milk and sugar with your coffee"
}

//
//func (this *CoffeeWithHook) CustomerWantsCondiments() bool {
//	if strings.HasPrefix(strings.ToLower(this.getUserInput()), "y") {
//		return true
//	}
//	return false
//}
//
//func (this *CoffeeWithHook) getUserInput() string {
//	inputReader := bufio.NewReader(os.Stdin)
//	input, err := inputReader.ReadString('\n')
//	if err != nil {
//		return "no"
//	}
//	return input
//}
