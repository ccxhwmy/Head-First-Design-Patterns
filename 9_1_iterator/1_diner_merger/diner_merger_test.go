package diner_merger

import "testing"

func TestDinerMerger(t *testing.T) {
	pancakeHouseMenu := NewPancakeHouseMenu()
	dinerMenu := NewDinerMenu()

	waitress := NewWaitress(pancakeHouseMenu, dinerMenu)

	waitress.PrintMenu()
}
