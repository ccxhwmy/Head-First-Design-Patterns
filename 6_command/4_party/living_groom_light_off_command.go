package simple_remote

type LivingRoomLightOffCommand struct {
	light *Light
}

func NewLivingRoomLightOffCommand(light *Light) *LivingRoomLightOffCommand {
	return &LivingRoomLightOffCommand{light: light}
}

func (this *LivingRoomLightOffCommand) Execute() {
	this.light.Off()
}

func (this *LivingRoomLightOffCommand) Undo() {
	this.light.On()
}
