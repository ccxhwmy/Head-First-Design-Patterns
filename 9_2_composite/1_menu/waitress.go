package menu

type Waitress struct {
	allMenus MenuComponent
}

func NewWaitress(allMenus MenuComponent) *Waitress {
	return &Waitress{allMenus: allMenus}
}

func (this *Waitress) printMenu() {
	this.allMenus.Print()
}
