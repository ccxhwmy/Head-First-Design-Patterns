package decorator

type GooseAdapter struct {
	goose *Goose
}

func NewGooseAdapter(goose *Goose) *GooseAdapter {
	return &GooseAdapter{
		goose: goose,
	}
}

func (this *GooseAdapter) Quack() {
	this.goose.Honk()
}

func (this *GooseAdapter) String() string {
	return "Goose pretending to be a Duck"
}
