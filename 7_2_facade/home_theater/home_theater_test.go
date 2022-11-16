package home_theater

import "testing"

func TestNewHomeTheaterAdapter(t *testing.T) {
	amp := NewAmplifier("Amplifier")
	tuner := NewTuner("AM/FM Tuner", amp)
	player := NewStreamPlayer("Streaming Player", amp)
	//cd := NewCDPlayer("Streaming Player", amp)
	projector := NewProjector("Projector", player)
	lights := NewTheaterLights("Theater Ceiling Lights")
	screen := NewScreen("Theater Screen")
	popper := NewPopcornPopper("Popcorn Popper")

	homeTheater := NewHomeTheaterAdapter(amp, tuner, player, projector, screen, lights, popper)

	homeTheater.WatchMovie("Raiders of the Lost Ark")
	homeTheater.EndMovie()
}
