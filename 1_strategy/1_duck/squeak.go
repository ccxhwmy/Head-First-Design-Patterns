package duck

import "fmt"

type Squeak struct {
}

func (this *Squeak) Quack() {
	fmt.Println("Ga Ga Ga ...")
}
