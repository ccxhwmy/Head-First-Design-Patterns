package decorator

import "fmt"

type DecoyDuck struct {
}

func (this *DecoyDuck) Quack() {
	fmt.Println("<< Silence >>")
}

func (this *DecoyDuck) String() string {
	return "Decoy Duck"
}
