package sort

import (
	"fmt"
	"sort"
	"testing"
)

var (
	oriDucks = DuckSort{
		{
			name:   "Daffy",
			weight: 8,
		},
		{
			name:   "Dewey",
			weight: 2,
		},
		{
			name:   "Howard",
			weight: 7,
		},
		{
			name:   "Louie",
			weight: 2,
		},
		{
			name:   "Donald",
			weight: 10,
		},
		{
			name:   "Huey",
			weight: 2,
		},
	}
)

func TestDuckSort(t *testing.T) {
	fmt.Println("Before sorting:")
	fmt.Println(oriDucks)

	sort.Sort(oriDucks)

	fmt.Println("\nAfter sorting:")
	fmt.Println(oriDucks)
}
