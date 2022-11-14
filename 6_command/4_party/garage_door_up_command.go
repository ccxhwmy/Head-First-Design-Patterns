package simple_remote

type GarageDoorUpCommand struct {
	garageDoor *GarageDoor
}

func NewGarageDoorUpCommand(garageDoor *GarageDoor) *GarageDoorUpCommand {
	return &GarageDoorUpCommand{garageDoor: garageDoor}
}

func (this *GarageDoorUpCommand) Execute() {
	this.garageDoor.Up()
}
