package simple_remote

type MacroCommand struct {
	commands []Command
}

func NewMacroCommand(commands []Command) *MacroCommand {
	return &MacroCommand{commands: commands}
}

func (this *MacroCommand) Execute() {
	for i := 0; i < len(this.commands); i++ {
		this.commands[i].Execute()
	}
}

func (this *MacroCommand) Undo() {
	for i := 0; i < len(this.commands); i++ {
		this.commands[i].Undo()
	}
}
