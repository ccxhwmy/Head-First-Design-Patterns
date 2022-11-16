package barista

import "fmt"

type Tea struct {
	CaffeineBeverageClass
}

func (this *Tea) Brew() {
	fmt.Println("Steeping the tea")
}

func (this *Tea) AddCondiments() {
	fmt.Println("Steeping the tea")
}
