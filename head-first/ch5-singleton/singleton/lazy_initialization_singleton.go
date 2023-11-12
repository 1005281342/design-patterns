package singleton

import "sync"

type LazyInitializationSingleton struct{}

var (
	lazyInitializationInstance *LazyInitializationSingleton
	once                       = sync.Once{}
)

func GetLazyInitializationInstance() *LazyInitializationSingleton {
	once.Do(func() {
		lazyInitializationInstance = &LazyInitializationSingleton{}
	})

	return lazyInitializationInstance
}
