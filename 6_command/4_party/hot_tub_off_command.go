package simple_remote

type HotTubOffCommand struct {
	hotTub *HotTub
}

func NewHotTubOffCommand(hotTub *HotTub) *HotTubOffCommand {
	return &HotTubOffCommand{hotTub: hotTub}
}

func (this *HotTubOffCommand) Execute() {
	this.hotTub.cool()
	this.hotTub.turnOff()
}

func (this *HotTubOffCommand) Undo() {
	this.hotTub.turnOn()
}
