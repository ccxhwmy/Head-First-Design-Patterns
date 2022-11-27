package factory

import "fmt"

type Goose struct {
}

func (this *Goose) Honk() {
	fmt.Println("Honk")
}

func (this *Goose) String() string {
	return "Goose"
}
