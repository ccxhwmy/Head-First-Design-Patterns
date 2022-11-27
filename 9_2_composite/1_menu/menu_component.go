package menu

type MenuComponent interface {
	Add(menuComponent MenuComponent)
	Remove(menuComponent MenuComponent)
	GetChild(i int) MenuComponent
	GetName() string
	GetDescription() string
	GetPrice() float64
	IsVegetarian() bool
	Print()
	GetUUID() string
}
