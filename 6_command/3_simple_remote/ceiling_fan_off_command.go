package simple_remote

type CeilingFanOffCommand struct {
	ceilingFan *CeilingFan
}

func NewCeilingFanOffCommand(ceilingFan *CeilingFan) *CeilingFanOffCommand {
	return &CeilingFanOffCommand{ceilingFan: ceilingFan}
}

func (this *CeilingFanOffCommand) Execute() {
	this.ceilingFan.off()
}
