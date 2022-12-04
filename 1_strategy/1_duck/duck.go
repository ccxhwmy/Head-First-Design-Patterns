package duck

type Duck struct {
	flyBehavior   FlyBehavior
	quackBehavior QuackBehavior
}

func (this *Duck) PerformQuack() {
	this.quackBehavior.Quack()
}

func (this *Duck) PerformFly() {
	this.flyBehavior.Fly()
}

func (this *Duck) SetFlyBehavior(fb FlyBehavior) {
	this.flyBehavior = fb
}

func (this *Duck) SetQuackBehavior(qb QuackBehavior) {
	this.quackBehavior = qb
}
