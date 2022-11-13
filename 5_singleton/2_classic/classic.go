package classic

// NOTE: This is not thread safe!

var (
	uniqueInstance *SingletonClassic
)

type SingletonClassic struct {
}

func GetInstance() *SingletonClassic {
	if uniqueInstance == nil {
		uniqueInstance = new(SingletonClassic)
	}
	return uniqueInstance
}

func (this *SingletonClassic) getDescription() string {
	return `I'm a classic Singleton!`
}
