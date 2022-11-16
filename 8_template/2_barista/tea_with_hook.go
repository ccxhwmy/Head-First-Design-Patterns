package barista

import (
	"fmt"
)

type TeaWithHook struct {
	CaffeineBeverageWithHookClass
}

func (this *TeaWithHook) Brew() {
	fmt.Println("Steeping the tea")
}

func (this *TeaWithHook) AddCondiments() {
	fmt.Println("Steeping the tea")
}

func (this *TeaWithHook) UserPrompt() string {
	return "Would you like lemon with your tea"
}

//func (this *TeaWithHook) CustomerWantsCondiments() bool {
//	if strings.HasPrefix(strings.ToLower(this.getUserInput()), "y") {
//		return true
//	}
//	return false
//}

//
//func (this *TeaWithHook) getUserInput() string {
//	inputReader := bufio.NewReader(os.Stdin)
//	input, err := inputReader.ReadString('\n')
//	if err != nil {
//		return "no"
//	}
//	return input
//}
