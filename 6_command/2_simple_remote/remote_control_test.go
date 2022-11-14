package simple_remote

import "testing"

func TestRemoteControl(t *testing.T) {
	remote := new(SimpleRemoteControl)
	light := new(Light)
	garageDoor := new(GarageDoor)

	lightOn := NewLightOnCommand(light)
	remote.SetCommand(lightOn)
	remote.buttonWasPressed()

	lightOff := NewLightOffCommand(light)
	remote.SetCommand((lightOff))
	remote.buttonWasPressed()

	garageDoorOpen := NewGarageDoorOpenCommand(garageDoor)
	remote.SetCommand(garageDoorOpen)
	remote.buttonWasPressed()
}
