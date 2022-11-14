package simple_remote

type CeilingFanHighCommand struct {
	ceilingFan *CeilingFan
	prevSpeed  int
}

func NewCeilingFanHighCommand(ceilingFan *CeilingFan) *CeilingFanHighCommand {
	return &CeilingFanHighCommand{ceilingFan: ceilingFan}
}

func (this *CeilingFanHighCommand) Execute() {
	this.prevSpeed = this.ceilingFan.getSpeed()
	this.ceilingFan.high()
}

func (this *CeilingFanHighCommand) Undo() {
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
