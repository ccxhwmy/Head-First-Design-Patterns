package duck

import "fmt"

type DuckCall struct {
}

func (this *DuckCall) Quack() {
	fmt.Println("Ga Ga Ga ...(machine display)")
}
