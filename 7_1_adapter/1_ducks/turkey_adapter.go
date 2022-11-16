package ducks

type TurkeyAdapter struct {
	turkey Turkey
}

func NewTurkeyAdapter(turkey Turkey) *TurkeyAdapter {
	return &TurkeyAdapter{turkey: turkey}
}

func (this *TurkeyAdapter) Quack() {
	this.turkey.Gobble()
}

func (this *TurkeyAdapter) Fly() {
	for i := 0; i < 5; i++ {
		this.turkey.Fly()
	}
}
