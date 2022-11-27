package factory

import "fmt"

type RedheadDuck struct {
}

func (this *RedheadDuck) Quack() {
	fmt.Println("Quack")
}

func (this *RedheadDuck) String() string {
	return "Redhead Duck"
}
