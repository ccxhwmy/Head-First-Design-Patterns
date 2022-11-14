package simple_remote

type StereoOnWithCDCommand struct {
	stereo *Stereo
}

func NewStereoOnWithCDCommand(stereo *Stereo) *StereoOnWithCDCommand {
	return &StereoOnWithCDCommand{stereo: stereo}
}

func (this *StereoOnWithCDCommand) Execute() {
	this.stereo.On()
	this.stereo.SetCD()
	this.stereo.SetVolume(11)
}
