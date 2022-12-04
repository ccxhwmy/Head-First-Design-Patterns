package duck

import "fmt"

type MuteQuack struct {
}

func (this *MuteQuack) Quack() {
	fmt.Println("<< Silence >>")
}
