package simple_remote

type CeilingFanOnCommand struct {
	ceilingFan *CeilingFan
}

func NewCeilingFanOnCommand(ceilingFan *CeilingFan) *CeilingFanOnCommand {
	return &CeilingFanOnCommand{ceilingFan: ceilingFan}
}

func (this *CeilingFanOnCommand) Execute() {
	this.ceilingFan.high()
}
