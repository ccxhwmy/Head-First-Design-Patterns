package simple_remote

type StereoOffCommand struct {
	stereo *Stereo
}

func NewStereoOffCommand(stereo *Stereo) *StereoOffCommand {
	return &StereoOffCommand{stereo: stereo}
}

func (this *StereoOffCommand) Execute() {
	this.stereo.Off()
}

func (this *StereoOffCommand) Undo() {
	this.stereo.On()
}
