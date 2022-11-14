package simple_remote

type HotTubOnCommand struct {
	hotTub *HotTub
}

func NewHotTubOnCommand(hotTub *HotTub) *HotTubOnCommand {
	return &HotTubOnCommand{hotTub: hotTub}
}

func (this *HotTubOnCommand) Execute() {
	this.hotTub.turnOn()
	this.hotTub.heat()
	this.hotTub.bubblesOn()
}
