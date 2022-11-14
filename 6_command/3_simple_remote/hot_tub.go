package simple_remote

import "fmt"

type HotTub struct {
	on          bool
	temperature int
}

func (this *HotTub) turnOn() {
	this.on = true
}

func (this *HotTub) turnOff() {
	this.on = false
}

func (this *HotTub) bubblesOn() {
	if this.on {
		fmt.Println("HotTub is bubbling!")
	}
}

func (this *HotTub) bubblesOff() {
	if this.on {
		fmt.Println("HotTub is not bubbling")
	}
}

func (this *HotTub) jetsOn() {
	if this.on {
		fmt.Println("HotTub jets are on")
	}
}

func (this *HotTub) jetsOff() {
	if this.on {
		fmt.Println("HotTub jets are off")
	}
}

func (this *HotTub) setTemperature(temperature int) {
	this.temperature = temperature
}

func (this *HotTub) heat() {
	this.temperature = 105
	fmt.Println("HotTub is heating to a steaming 105 degrees")
}

func (this *HotTub) cool() {
	this.temperature = 98
	fmt.Println("HotTub is heating to a steaming 105 degrees")
}
