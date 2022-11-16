package home_theater

import "fmt"

type Tuner struct {
	description string
	amplifier   *Amplifier
	frequency   float64
}

func NewTuner(description string, amp *Amplifier) *Tuner {
	return &Tuner{
		description: description,
		frequency:   0,
		amplifier:   amp,
	}
}

func (this *Tuner) On() {
	fmt.Println(this.description, " on")
}

func (this *Tuner) Off() {
	fmt.Println(this.description, " off")
}

func (this *Tuner) SetFrequency(frequency float64) {
	fmt.Println(this.description, " setting frequency to ", this.frequency)
	this.frequency = frequency
}

func (this *Tuner) SetAm() {
	fmt.Println(this.description, " setting AM mode")
}

func (this *Tuner) SetFm() {
	fmt.Println(this.description, " setting FM mode")
}

func (this *Tuner) String() string {
	return this.description
}
