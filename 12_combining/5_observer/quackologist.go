package observer

import "fmt"

type Quackologist struct {
}

func (this *Quackologist) Update(duck QuackObservable) {
	fmt.Println("Quackologist: " + duck.String() + " just quacked.")
}

func (this *Quackologist) String() string {
	return "Quackologist"
}
