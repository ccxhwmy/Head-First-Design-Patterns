package simple_remote

type LivingroomLightOffCommand struct {
	light *Light
}

func NewLivingroomLightOffCommand(light *Light) *LivingroomLightOffCommand {
	return &LivingroomLightOffCommand{light: light}
}

func (this *LivingroomLightOffCommand) Execute() {
	this.light.Off()
}
