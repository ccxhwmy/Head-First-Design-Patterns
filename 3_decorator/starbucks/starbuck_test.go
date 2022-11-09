package starbucks

import (
	"fmt"
	"testing"
)

func TestStarbucks(t *testing.T) {
	beverage := &Espresso{}
	fmt.Printf("%s $%.2f\n", beverage.GetDescription(), beverage.Cost())

	darkRoast := &DarkRoast{}
	singleMocha := &Mocha{beverage: darkRoast}

	doubleMocha := &Mocha{beverage: singleMocha}

	doubleMochaWhip := &Whip{beverage: doubleMocha}

	fmt.Printf("%s $%.2f\n", doubleMochaWhip.GetDescription(), doubleMochaWhip.Cost())

	soyMochaWhipHouseBlend := &Whip{beverage: &Mocha{beverage: &HouseBlend{}}}

	fmt.Printf("%s $%.2f\n", soyMochaWhipHouseBlend.GetDescription(), soyMochaWhipHouseBlend.Cost())
}
