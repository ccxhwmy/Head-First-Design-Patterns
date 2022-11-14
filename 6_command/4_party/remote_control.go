package simple_remote

import (
	"bytes"
	"reflect"
	"strconv"
)

type RemoteControl struct {
	onCommands  []Command
	offCommands []Command
	undoCommand Command
}

func NewRemoteControl() *RemoteControl {
	remoteControl := &RemoteControl{
		onCommands:  make([]Command, 7),
		offCommands: make([]Command, 7),
	}
	for i := 0; i < 7; i++ {
		remoteControl.onCommands[i] = &NoCommand{}
		remoteControl.offCommands[i] = &NoCommand{}
	}
	remoteControl.undoCommand = &NoCommand{}
	return remoteControl
}

func (this *RemoteControl) SetCommand(slot int, onCommand, offCommand Command) {
	this.onCommands[slot] = onCommand
	this.offCommands[slot] = offCommand
}

func (this *RemoteControl) onButtonWasPushed(slot int) {
	this.onCommands[slot].Execute()
}

func (this *RemoteControl) offButtonWasPushed(slot int) {
	this.offCommands[slot].Execute()
}

func (this *RemoteControl) UndoButtonWasPushed() {
	this.undoCommand.Undo()
}

func (this *RemoteControl) String() string {
	buffer := bytes.Buffer{}
	buffer.WriteString("\n------ Remote Control -------\n")
	for i := 0; i < len(this.onCommands); i++ {
		buffer.WriteString(`[slot ` + strconv.Itoa(i) + `]`)
		buffer.WriteString(reflect.TypeOf(this.onCommands[i]).Elem().Name())
		buffer.WriteString("    ")
		buffer.WriteString(reflect.TypeOf(this.offCommands[i]).Elem().Name())
		buffer.WriteString("\n")
	}
	return buffer.String()
}
