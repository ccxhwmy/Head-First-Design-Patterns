package home_theater

import "fmt"

type HomeTheaterAdapter struct {
	amp       *Amplifier
	tuner     *Tuner
	player    *StreamPlayer
	cd        *CDPlayer
	projector *Projector
	lights    *TheaterLights
	screen    *Screen
	popper    *PopcornPopper
}

func NewHomeTheaterAdapter(amp *Amplifier,
	tuner *Tuner,
	player *StreamPlayer,
	projector *Projector,
	screen *Screen,
	lights *TheaterLights,
	popper *PopcornPopper) *HomeTheaterAdapter {
	return &HomeTheaterAdapter{
		amp:       amp,
		tuner:     tuner,
		player:    player,
		cd:        nil,
		projector: projector,
		lights:    lights,
		screen:    screen,
		popper:    popper,
	}
}

func (this *HomeTheaterAdapter) WatchMovie(movie string) {
	fmt.Println("Get ready to watch a movie...")
	this.popper.On()
	this.popper.Pop()
	this.lights.Dim(10)
	this.screen.Down()
	this.projector.On()
	this.projector.WideScreenMode()
	this.amp.On()
	this.amp.SetStreamingPlayer(this.player)
	this.amp.SetSurroundSound()
	this.amp.SetVolume(5)
	this.player.On()
	this.player.Play(movie)
}

func (this *HomeTheaterAdapter) EndMovie() {
	fmt.Println("Shutting movie theater down...")
	this.popper.Off()
	this.lights.On()
	this.screen.Up()
	this.projector.Off()
	this.amp.Off()
	this.player.Stop()
	this.player.Off()
}

func (this *HomeTheaterAdapter) ListenToRadio(frequency float64) {
	fmt.Println("Tuning in the airwaves...")
	this.tuner.On()
	this.tuner.SetFrequency(frequency)
	this.amp.On()
	this.amp.SetVolume(5)
	this.amp.SetTuner(this.tuner)
}

func (this *HomeTheaterAdapter) EndRadio() {
	fmt.Println("Shutting down the tuner...")
	this.tuner.Off()
	this.amp.Off()
}
