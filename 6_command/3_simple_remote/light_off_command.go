package simple_remote

type LightOffCommand struct {
	light *Light
}

func NewLightOffCommand(light *Light) *LightOffCommand {
	return &LightOffCommand{light: light}
}

func (this *LightOffCommand) Execute() {
	this.light.Off()
}
