package chocolate

import "testing"

func TestChocolate(t *testing.T) {
	boiler := GetInstance()
	boiler.fill()
	boiler.boil()
	boiler.drain()

	boiler2 := GetInstance()
	boiler2.fill()
}
