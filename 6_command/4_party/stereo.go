package simple_remote

import "fmt"

type Stereo struct {
	location string
}

func NewStereo(location string) *Stereo {
	return &Stereo{location: location}
}

func (this *Stereo) On() {
	fmt.Println(this.location, " stereo is on")
}

func (this *Stereo) Off() {
	fmt.Println(this.location, " stereo is off")
}

func (this *Stereo) SetCD() {
	fmt.Println(this.location, " stereo is set for CD input")
}

func (this *Stereo) SetDVD() {
	fmt.Println(this.location, " stereo is set for DVD input")
}

func (this *Stereo) SetRadio() {
	fmt.Println(this.location, " stereo is set for Radio")
}

func (this *Stereo) SetVolume(volume int) {
	fmt.Println(this.location, " stereo is set to ", volume)
}
