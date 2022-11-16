package barista

import (
	"fmt"
	"testing"
)

func TestBarista(t *testing.T) {
	fmt.Println("\nMaking tea...")
	PrepareRecipe(new(Tea))

	fmt.Println("\nMaking coffee...")
	PrepareRecipe(new(Coffee))
}
