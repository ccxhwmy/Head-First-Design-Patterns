package simple_remote

import (
	"fmt"
	"testing"
)

func TestRemoteControl(t *testing.T) {
	remoteControl := NewRemoteControl()

	light := NewLight("Living Room")
	tv := NewTV("Living Room")
	stereo := NewStereo("Living Room")
	hotTub := new(HotTub)

	lightOn := NewLightOnCommand(light)
	stereoOn := NewStereoOnCommand(stereo)
	tvOn := NewTVOnCommand(tv)
	hotHubOn := NewHotTubOnCommand(hotTub)
	lightOff := NewLightOffCommand(light)
	stereoOff := NewStereoOffCommand(stereo)
	tvOff := NewTVOffCommand(tv)
	hotHubOff := NewHotTubOffCommand(hotTub)

	partyOn := []Command{lightOn, stereoOn, tvOn, hotHubOn}
	partyOff := []Command{lightOff, stereoOff, tvOff, hotHubOff}

	partyOnMacro := NewMacroCommand(partyOn)
	partyOffMacro := NewMacroCommand(partyOff)

	remoteControl.SetCommand(0, partyOnMacro, partyOffMacro)

	fmt.Println(remoteControl)
	fmt.Println("--- Pushing Macro On---")
	remoteControl.onButtonWasPushed(0)
	fmt.Println("--- Pushing Macro Off---")
	remoteControl.offButtonWasPushed(0)
}
