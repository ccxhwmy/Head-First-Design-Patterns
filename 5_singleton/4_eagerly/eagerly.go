package eagerly

var (
	uniqueSingleton *SingletonEagerly
)

func init() {
	uniqueSingleton = new(SingletonEagerly)
}

type SingletonEagerly struct {
}

func GetInstance() *SingletonEagerly {
	return uniqueSingleton
}

func (this *SingletonEagerly) GetDescription() string {
	return `I'm a eager Singleton!`
}
