package starbucks

type Mocha struct {
	beverage Beverage
}

func (this *Mocha) GetDescription() string {
	return this.beverage.GetDescription() + ", Mocha"
}

func (this *Mocha) Cost() float64 {
	return this.beverage.Cost() + .2
}

type Milk struct {
	beverage Beverage
}

func (this *Milk) GetDescription() string {
	return this.beverage.GetDescription() + ", Milk"
}

func (this *Milk) Cost() float64 {
	return this.beverage.Cost() + .1
}

type Soy struct {
	beverage Beverage
}

func (this *Soy) GetDescription() string {
	return this.beverage.GetDescription() + ", Soy"
}

func (this *Soy) Cost() float64 {
	return this.beverage.Cost() + .15
}

type Whip struct {
	beverage Beverage
}

func (this *Whip) GetDescription() string {
	return this.beverage.GetDescription() + ", Whip"
}

func (this *Whip) Cost() float64 {
	return this.beverage.Cost() + .1
}
