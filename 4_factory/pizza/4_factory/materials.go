package factory

type Dough interface {
	String() string
}

type ThickCrustDough struct {
}

func (this *ThickCrustDough) String() string {
	return "ThickCrust style extra thick crust dough"
}

type ThinCrustDough struct {
}

func (this *ThinCrustDough) String() string {
	return "Thin Crust Dough"
}

type Sauce interface {
	String() string
}

type PlumTomatoSauce struct {
}

func (this *PlumTomatoSauce) String() string {
	return "Tomato sauce with plum tomatoes"
}

type MarinaraSauce struct {
}

func (this *MarinaraSauce) String() string {
	return "Marinara Sauce"
}

type Cheese interface {
	String() string
}

type MozzarellaCheese struct {
}

func (this *MozzarellaCheese) String() string {
	return "Shredded Mozzarella"
}

type ParmesanCheese struct {
}

func (this *ParmesanCheese) String() string {
	return "Shredded Parmesan"
}

type ReggianoCheese struct {
}

func (this *ReggianoCheese) String() string {
	return "Reggiano Cheese"
}

type Veggies interface {
	String() string
}

type Onion struct {
}

func (this *Onion) String() string {
	return "Onion"
}

type BlackOlives struct {
}

func (this *BlackOlives) String() string {
	return "Black Olives"
}

type Garlic struct {
}

func (this *Garlic) String() string {
	return "Garlic"
}

type RedPepper struct {
}

func (this *RedPepper) String() string {
	return "Garlic"
}

type Spinach struct {
}

func (this *Spinach) String() string {
	return "Spinach"
}

type Mushroom struct {
}

func (this *Mushroom) String() string {
	return "Mushrooms"
}

type Eggplant struct {
}

func (this *Eggplant) String() string {
	return "Eggplant"
}

type Pepperoni interface {
	String() string
}

type SlicedPepperoni struct {
}

func (this *SlicedPepperoni) String() string {
	return "Sliced Pepperoni"
}

type Clams interface {
	String() string
}

type FreshClams struct {
}

func (this *FreshClams) String() string {
	return "Fresh Clams from Long Island Sound"
}

type FrozenClams struct {
}

func (this *FrozenClams) String() string {
	return "Frozen Clams from Chesapeake Bay"
}
