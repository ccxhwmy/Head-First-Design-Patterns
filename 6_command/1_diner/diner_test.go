package diner

import "testing"

func TestDiner(t *testing.T) {
	cook := new(Cook)
	waitress := new(Waitress)
	customer := NewCustomer(waitress)
	customer.createOrder(NewBurgerAndFriesOrder(cook))
	customer.hungry()
}
