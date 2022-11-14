package simple_remote

type TVOnCommand struct {
	tv *TV
}

func NewTVOnCommand(tv *TV) *TVOnCommand {
	return &TVOnCommand{tv: tv}
}

func (this *TVOnCommand) Execute() {
	this.tv.Off()
}
