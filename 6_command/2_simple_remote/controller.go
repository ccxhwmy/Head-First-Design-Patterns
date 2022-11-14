package simple_remote

type SimpleRemoteControl struct {
	slot Command
}

func (this *SimpleRemoteControl) SetCommand(command Command) {
	this.slot = command
}

func (this *SimpleRemoteControl) buttonWasPressed() {
	this.slot.Execute()
}
