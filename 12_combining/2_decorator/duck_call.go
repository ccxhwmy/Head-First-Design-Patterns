package decorator

import "fmt"

type DuckCall struct {
}

func (this *DuckCall) Quack() {
	fmt.Println("Kwak")
}

func (this *DuckCall) String() string {
	return "Duck Call"
}
