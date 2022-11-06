package game

import (
	"fmt"
	"testing"
)

func TestGame(t *testing.T) {
	fmt.Println("-------------- show time -------------")
	king := NewKing()
	king.fight()

	queen := NewQueen()
	queen.fight()

	knight := NewKnight()
	knight.fight()

	troll := NewTroll()
	troll.fight()

	fmt.Println("-------------- troll killed the knight and snatch the bow and arrow -------------")
	troll.SnatchBowAndArrow()
	troll.fight()
}
