package barista

import "fmt"

type Coffee struct {
	CaffeineBeverageClass
}

func (this *Coffee) Brew() {
	fmt.Println("Dripping Coffee through filter")
}

func (this *Coffee) AddCondiments() {
	fmt.Println("Adding Sugar and Milk")
}
