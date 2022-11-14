package simple_remote

import "fmt"

type TV struct {
	location string
	channel  int
}

func NewTV(location string) *TV {
	return &TV{
		location: location,
	}
}

func (this *TV) On() {
	fmt.Println(this.location, " TV is on")
}

func (this *TV) Off() {
	fmt.Println(this.location, " TV is off")
}

func (this *TV) SetInputChannel() {
	this.channel = 3
	fmt.Println(this.location, " TV channel is set for DVD")
}
