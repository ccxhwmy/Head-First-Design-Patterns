package home_theater

import "fmt"

type Screen struct {
	description string
}

func NewScreen(description string) *Screen {
	return &Screen{description: description}
}

func (this *Screen) Up() {
	fmt.Println(this.description, " going up")
}

func (this *Screen) Down() {
	fmt.Println(this.description, " going down")
}

func (this *Screen) String() string {
	return this.description
}
