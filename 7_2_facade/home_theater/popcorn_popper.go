package home_theater

import "fmt"

type PopcornPopper struct {
	description string
}

func NewPopcornPopper(description string) *PopcornPopper {
	return &PopcornPopper{description: description}
}

func (this *PopcornPopper) On() {
	fmt.Println(this.description, " on")
}

func (this *PopcornPopper) Off() {
	fmt.Println(this.description, " off")
}

func (this *PopcornPopper) Pop() {
	fmt.Println(this.description, " popping popcorn!")
}

func (this *PopcornPopper) String() string {
	return this.description
}
