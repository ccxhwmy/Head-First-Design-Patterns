package simple_remote

type CeilingFanMediumCommand struct {
	ceilingFan *CeilingFan
	prevSpeed  int
}

func NewCeilingFanMediumCommand(ceilingFan *CeilingFan) *CeilingFanMediumCommand {
	return &CeilingFanMediumCommand{ceilingFan: ceilingFan}
}

func (this *CeilingFanMediumCommand) Execute() {
	this.prevSpeed = this.ceilingFan.getSpeed()
	this.ceilingFan.medium()
}

func (this *CeilingFanMediumCommand) Undo() {
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
