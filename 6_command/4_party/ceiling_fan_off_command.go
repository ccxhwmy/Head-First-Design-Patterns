package simple_remote

type CeilingFanOffCommand struct {
	ceilingFan *CeilingFan
	prevSpeed  int
}

func NewCeilingFanOffCommand(ceilingFan *CeilingFan) *CeilingFanOffCommand {
	return &CeilingFanOffCommand{ceilingFan: ceilingFan}
}

func (this *CeilingFanOffCommand) Execute() {
	this.prevSpeed = this.ceilingFan.getSpeed()
	this.ceilingFan.off()
}

func (this *CeilingFanOffCommand) Undo() {
	switch this.prevSpeed {
	case HIGH:
		this.ceilingFan.high()
		break
	case MEDIUM:
		this.ceilingFan.medium()
		break
	case LOW:
		this.ceilingFan.low()
		break
	default:
		this.ceilingFan.off()
		break
	}
}
