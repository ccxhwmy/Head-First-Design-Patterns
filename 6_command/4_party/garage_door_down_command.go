package simple_remote

type GarageDoorDownCommand struct {
	garageDoor *GarageDoor
}

func NewGarageDoorDownCommand(garageDoor *GarageDoor) *GarageDoorUpCommand {
	return &GarageDoorUpCommand{garageDoor: garageDoor}
}

func (this *GarageDoorDownCommand) Execute() {
	this.garageDoor.Down()
}
