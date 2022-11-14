package simple_remote

type StereoOnCommand struct {
	stereo *Stereo
}

func NewStereoOnCommand(stereo *Stereo) *StereoOnCommand {
	return &StereoOnCommand{stereo: stereo}
}

func (this *StereoOnCommand) Execute() {
	this.stereo.On()
}

func (this *StereoOnCommand) Undo() {
	this.stereo.Off()
}
