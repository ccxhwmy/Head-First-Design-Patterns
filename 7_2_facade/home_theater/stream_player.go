package home_theater

import "fmt"

type StreamPlayer struct {
	description    string
	currentChapter int
	amplifier      *Amplifier
	movie          string
}

func NewStreamPlayer(description string, amplifier *Amplifier) *StreamPlayer {
	return &StreamPlayer{
		description:    description,
		currentChapter: 0,
		amplifier:      amplifier,
		movie:          "",
	}
}

func (this *StreamPlayer) On() {
	fmt.Println(this.description, " on")
}

func (this *StreamPlayer) Off() {
	fmt.Println(this.description, " off")
}

func (this *StreamPlayer) Play(movie string) {
	this.movie = movie
	this.currentChapter = 0
	fmt.Println(this.description, " playing chapter ", this.currentChapter, " of \"", this.movie, "\"")
}

func (this *StreamPlayer) Stop() {
	this.currentChapter = 0
	fmt.Println(this.description, " stopped \"", this.movie, "\"")
}

func (this *StreamPlayer) Pause() {
	fmt.Println(this.description, " paused \"", this.movie, "\"")
}

func (this *StreamPlayer) SetTwoChannelAudio() {
	fmt.Println(this.description, " set two channel audio")
}

func (this *StreamPlayer) SetSurroundAudio() {
	fmt.Println(this.description, " set surround audio")
}

func (this *StreamPlayer) String() string {
	return this.description
}
