package home_theater

import "fmt"

type Projector struct {
	description string
	player      *StreamPlayer
}

func NewProjector(description string, player *StreamPlayer) *Projector {
	return &Projector{
		description: description,
		player:      player,
	}
}

func (this *Projector) On() {
	fmt.Println(this.description, " on")
}

func (this *Projector) Off() {
	fmt.Println(this.description, " off")
}

func (this *Projector) WideScreenMode() {
	fmt.Println(this.description, " in widescreen mode (16x9 aspect ratio)")
}

func (this *Projector) TvMode() {
	fmt.Println(this.description, " in tv mode (4x3 aspect ratio)")
}

func (this *Projector) String() string {
	return this.description
}
