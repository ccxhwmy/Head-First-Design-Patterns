package simple_remote

import "fmt"

const (
	HIGH   = 2
	MEDIUM = 1
	LOW    = 0
)

type CeilingFan struct {
	location string
	level    int
}

func NewCeilingFan(location string) *CeilingFan {
	return &CeilingFan{
		location: location,
	}
}

func (this *CeilingFan) high() {
	this.level = HIGH
	fmt.Println(this.location, " ceiling fan is on high")
}

func (this *CeilingFan) medium() {
	this.level = MEDIUM
	fmt.Println(this.location, " ceiling fan is on medium")
}

func (this *CeilingFan) low() {
	this.level = LOW
	fmt.Println(this.location, " ceiling fan is on low")
}

func (this *CeilingFan) off() {
	this.level = 0
	fmt.Println(this.location, " ceiling fan is off")
}

func (this *CeilingFan) getSpeed() int {
	return this.level
}
