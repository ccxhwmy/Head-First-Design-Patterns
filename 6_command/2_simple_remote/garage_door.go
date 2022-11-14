package simple_remote

import "fmt"

type GarageDoor struct {
}

func (this *GarageDoor) Up() {
	fmt.Println("Garage Door is Open")
}

func (this *GarageDoor) Down() {
	fmt.Println("Garage Door is Closed")
}

func (this *GarageDoor) Stop() {
	fmt.Println("Garage Door is Stopped")
}

func (this *GarageDoor) LightOn() {
	fmt.Println("Garage light is on")
}

func (this *GarageDoor) LightOff() {
	fmt.Println("Garage light is off")
}
