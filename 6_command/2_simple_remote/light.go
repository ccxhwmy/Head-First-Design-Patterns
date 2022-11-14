package simple_remote

import "fmt"

type Light struct {
}

func (this *Light) On() {
	fmt.Println("Light is on")
}

func (this *Light) Off() {
	fmt.Println("Light is off")
}
