package simple_remote

import (
	"fmt"
	"testing"
)

func TestRemoteControl(t *testing.T) {
	remoteControl := NewRemoteControl()

	livingRoomLight := NewLight("Living Room")
	kitchenLight := NewLight("Kitchen")
	ceilingFan := NewCeilingFan("Living Room")
	garageDoor := NewGarageDoor("Garage")
	stereo := NewStereo("Living Room")

	livingRoomLightOn := NewLightOnCommand(livingRoomLight)
	livingRoomLightOff := NewLightOffCommand(livingRoomLight)
	kitchenLightOn := NewLightOnCommand(kitchenLight)
	kitchenLightOff := NewLightOffCommand(kitchenLight)
	ceilingFanOn := NewCeilingFanOnCommand(ceilingFan)
	ceilingFanOff := NewCeilingFanOffCommand(ceilingFan)
	garageDoorUp := NewGarageDoorUpCommand(garageDoor)
	garageDoorDown := NewGarageDoorDownCommand(garageDoor)
	stereoOnWithCD := NewStereoOnWithCDCommand(stereo)
	stereoOff := NewStereoOffCommand(stereo)

	remoteControl.SetCommand(0, livingRoomLightOn, livingRoomLightOff)
	remoteControl.SetCommand(1, kitchenLightOn, kitchenLightOff)
	remoteControl.SetCommand(2, ceilingFanOn, ceilingFanOff)
	remoteControl.SetCommand(3, stereoOnWithCD, stereoOff)
	remoteControl.SetCommand(4, garageDoorUp, garageDoorDown)

	fmt.Println(remoteControl)

	remoteControl.onButtonWasPushed(0)
	remoteControl.offButtonWasPushed(0)
	remoteControl.onButtonWasPushed(1)
	remoteControl.offButtonWasPushed(1)
	remoteControl.onButtonWasPushed(2)
	remoteControl.offButtonWasPushed(2)
	remoteControl.onButtonWasPushed(3)
	remoteControl.offButtonWasPushed(3)
	remoteControl.onButtonWasPushed(4)
	remoteControl.offButtonWasPushed(4)
}
