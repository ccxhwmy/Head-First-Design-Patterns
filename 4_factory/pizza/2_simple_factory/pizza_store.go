package simple_factory

type iPizzaStore interface {
	orderPizza(typ string) iPizza
}

type PizzaStore struct {
	factory PizzaFactory
}

func NewPizzaStore(factory PizzaFactory) *PizzaStore {
	return &PizzaStore{factory: factory}
}

func (this *PizzaStore) orderPizza(typ string) iPizza {
	var pizza iPizza

	pizza = this.factory.createPizza(typ)

	pizza.Prepare()
	pizza.Bake()
	pizza.Cut()
	pizza.Box()

	return pizza
}
