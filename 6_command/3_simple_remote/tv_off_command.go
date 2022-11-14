package simple_remote

type TVOffCommand struct {
	tv *TV
}

func NewTVOffCommand(tv *TV) *TVOffCommand {
	return &TVOffCommand{tv: tv}
}

func (this *TVOffCommand) Execute() {
	this.tv.Off()
}
