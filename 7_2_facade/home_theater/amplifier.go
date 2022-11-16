package home_theater

import "fmt"

type Amplifier struct {
	description string
	tuner       *Tuner
	player      *StreamPlayer
}

func NewAmplifier(description string) *Amplifier {
	return &Amplifier{
		description: description,
	}
}

func (this *Amplifier) On() {
	fmt.Println(this.description, " on")
}

func (this *Amplifier) Off() {
	fmt.Println(this.description, " off")
}

func (this *Amplifier) SetStereoSound() {
	fmt.Println(this.description, " stereo mode on")
}

func (this *Amplifier) SetSurroundSound() {
	fmt.Println(this.description, " surround sound on (5 speakers, 1 subwoofer)")
}

func (this *Amplifier) SetVolume(level int) {
	fmt.Println(this.description, " setting volume to ", level)
}

func (this *Amplifier) SetTuner(tuner *Tuner) {
	fmt.Println(this.description, " setting tuner to ", tuner)
	this.tuner = tuner
}

func (this *Amplifier) SetStreamingPlayer(player *StreamPlayer) {
	fmt.Println(this.description, " setting Streaming player to ", player)
	this.player = player
}

func (this *Amplifier) String() string {
	return this.description
}
