package simple_remote

import "fmt"

type Light struct {
	location string
}

func NewLight(location string) *Light {
	return &Light{
		location: location,
	}
}

func (this *Light) On() {
	fmt.Println("Light is on")
}

func (this *Light) Off() {
	fmt.Println("Light is off")
}
