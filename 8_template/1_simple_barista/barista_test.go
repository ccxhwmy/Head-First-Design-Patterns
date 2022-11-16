package simple_barista

import (
	"fmt"
	"testing"
)

func TestBarista(t *testing.T) {
	tea := new(Tea)
	coffee := new(Coffee)
	fmt.Println("Making tea...")
	tea.PrepareRecipe()
	fmt.Println("Making coffee...")
	coffee.PrepareRecipe()
}
