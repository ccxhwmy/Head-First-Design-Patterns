package simple_remote

type LightOnCommand struct {
	light *Light
}

func NewLightOnCommand(light *Light) *LightOnCommand {
	return &LightOnCommand{light: light}
}

func (this *LightOnCommand) Execute() {
	this.light.On()
}

func (this *LightOnCommand) Undo() {
	this.light.Off()
}
