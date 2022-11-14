package simple_remote

type LivingroomLightOnCommand struct {
	light *Light
}

func NewLivingroomLightOnCommand(light *Light) *LivingroomLightOnCommand {
	return &LivingroomLightOnCommand{light: light}
}

func (this *LivingroomLightOnCommand) Execute() {
	this.light.Off()
}
