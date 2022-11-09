package starbucks

type Beverage interface {
	GetDescription() string
	Cost() float64
}

type Espresso struct {
}

func (this *Espresso) GetDescription() string {
	return "Espresso"
}

func (this *Espresso) Cost() float64 {
	return 1.99
}

type HouseBlend struct {
}

func (this *HouseBlend) GetDescription() string {
	return "House Blend Coffee"
}

func (this *HouseBlend) Cost() float64 {
	return .89
}

type DarkRoast struct {
}

func (this *DarkRoast) GetDescription() string {
	return "Dark Roast Coffee"
}

func (this *DarkRoast) Cost() float64 {
	return .99
}

type Decaf struct {
}

func (this *Decaf) GetDescription() string {
	return "Decaf"
}

func (this *Decaf) Cost() float64 {
	return 1.05
}
