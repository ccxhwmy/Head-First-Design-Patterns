package diner

import "fmt"

type Order interface {
	orderUp()
}

type Cook struct {
}

func (this *Cook) MakeBurger() {
	fmt.Println("Making a burger")
}

func (this *Cook) MakeFries() {
	fmt.Println("Making fries")
}

type Waitress struct {
	order Order
}

func (this *Waitress) takeOrder(order Order) {
	this.order = order
	order.orderUp()
}

type BurgerAndFriesOrder struct {
	cook *Cook
}

func NewBurgerAndFriesOrder(cook *Cook) *BurgerAndFriesOrder {
	return &BurgerAndFriesOrder{cook: cook}
}

func (this *BurgerAndFriesOrder) orderUp() {
	this.cook.MakeBurger()
	this.cook.MakeFries()
}

type Customer struct {
	waitress *Waitress
	Order    Order
}

func NewCustomer(waitress *Waitress) *Customer {
	return &Customer{
		waitress: waitress,
		Order:    nil,
	}
}

func (this *Customer) createOrder(order Order) {
	this.Order = order
}

func (this *Customer) hungry() {
	this.waitress.takeOrder(this.Order)
}
