package home_theater

import "fmt"

type TheaterLights struct {
	description string
}

func NewTheaterLights(description string) *TheaterLights {
	return &TheaterLights{description: description}
}

func (this *TheaterLights) On() {
	fmt.Println(this.description, " on")
}

func (this *TheaterLights) Off() {
	fmt.Println(this.description, " off")
}

func (this *TheaterLights) Dim(level int) {
	fmt.Println(this.description, " dimming to ", level, "%")
}

func (this *TheaterLights) String() string {
	return this.description
}
