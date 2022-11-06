package game

import "fmt"

type WeaponBehavior interface {
	useWeapon()
}

type SwordBehavior struct {
}

func (this *SwordBehavior) useWeapon() {
	fmt.Println("swing the sword ...")
}

type KnifeBehavior struct {
}

func (this *KnifeBehavior) useWeapon() {
	fmt.Println("stabbing knife ...")
}

type BowAndArrowBehavior struct {
}

func (this *BowAndArrowBehavior) useWeapon() {
	fmt.Println("archery ...")
}

type AxeBehavior struct {
}

func (this *AxeBehavior) useWeapon() {
	fmt.Println("hack ...")
}

type Character struct {
	weaponBehavior WeaponBehavior
}

func (this *Character) fight() {
	this.weaponBehavior.useWeapon()
}

func (this *Character) setWeapon(w WeaponBehavior) {
	this.weaponBehavior = w
}

type King struct {
	Character
}

func NewKing() *King {
	return &King{Character{weaponBehavior: &SwordBehavior{}}}
}

type Queen struct {
	Character
}

func NewQueen() *Queen {
	return &Queen{Character{weaponBehavior: &KnifeBehavior{}}}
}

type Knight struct {
	Character
}

func NewKnight() *Knight {
	return &Knight{Character{weaponBehavior: &BowAndArrowBehavior{}}}
}

type Troll struct {
	Character
}

func NewTroll() *Troll {
	return &Troll{Character{weaponBehavior: &AxeBehavior{}}}
}

func (this *Troll) SnatchBowAndArrow() {
	this.setWeapon(&BowAndArrowBehavior{})
}
