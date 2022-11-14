package simple_remote

type LivingRoomLightOnCommand struct {
	light *Light
}

func NewLivingRoomLightOnCommand(light *Light) *LivingRoomLightOnCommand {
	return &LivingRoomLightOnCommand{light: light}
}

func (this *LivingRoomLightOnCommand) Execute() {
	this.light.Off()
}

func (this *LivingRoomLightOnCommand) Undo() {
	this.light.On()
}
