package home_theater

import "fmt"

type CDPlayer struct {
	description  string
	currentTrack int
	amplifier    *Amplifier
	title        string
}

func NewCDPlayer(description string, amplifier *Amplifier) *CDPlayer {
	return &CDPlayer{
		description:  description,
		currentTrack: 0,
		amplifier:    amplifier,
		title:        "",
	}
}

func (this *CDPlayer) On() {
	fmt.Println(this.description, " on")
}

func (this *CDPlayer) Off() {
	fmt.Println(this.description, " off")
}

func (this *CDPlayer) Eject() {
	this.title = ""
	fmt.Println(this.description, " eject")
}

func (this *CDPlayer) PlayTitle(title string) {
	this.title = title
	this.currentTrack = 0
	fmt.Println(this.description, " playing \"", title, "\"")
}

func (this *CDPlayer) PlayTrack(track int) {
	if this.title == "" {
		fmt.Println(this.description, " can't play track ", this.currentTrack, ", no cd inserted")
	} else {
		this.currentTrack = track
		fmt.Println(this.description, " playing track ", this.currentTrack)
	}
}

func (this *CDPlayer) Stop() {
	this.currentTrack = 0
	fmt.Println(this.description, " stopped")
}

func (this *CDPlayer) Pause() {
	this.currentTrack = 0
	fmt.Println(this.description, " paused \"", this.title, "\"")
}

func (this *CDPlayer) String() string {
	return this.description
}
