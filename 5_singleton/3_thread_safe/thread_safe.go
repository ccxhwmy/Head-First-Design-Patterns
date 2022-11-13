package thread_safe

import "sync"

var (
	uniqueInstance *SingletonThreadSafe
	once           = sync.Once{}
)

type SingletonThreadSafe struct {
}

func GetInstance() *SingletonThreadSafe {
	if uniqueInstance == nil {
		once.Do(func() {
			uniqueInstance = &SingletonThreadSafe{}
		})
	}
	return uniqueInstance
}

func (this *SingletonThreadSafe) GetDescription() string {
	return `I'm a thread safe Singleton!`
}
