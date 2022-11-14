package simple_remote

type Command interface {
	Execute()
}

type LightOnCommand struct {
	light *Light
}

func NewLightOnCommand(light *Light) *LightOnCommand {
	return &LightOnCommand{light: light}
}

func (this *LightOnCommand) Execute() {
	this.light.On()
}

type LightOffCommand struct {
	light *Light
}

func NewLightOffCommand(light *Light) *LightOffCommand {
	return &LightOffCommand{light: light}
}

func (this *LightOffCommand) Execute() {
	this.light.Off()
}

type GarageDoorOpenCommand struct {
	garageDoor *GarageDoor
}

func NewGarageDoorOpenCommand(garageDoor *GarageDoor) *GarageDoorOpenCommand {
	return &GarageDoorOpenCommand{garageDoor: garageDoor}
}

func (this *GarageDoorOpenCommand) Execute() {
	this.garageDoor.Up()
}
