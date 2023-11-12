package singleton

type EagerInitializationSingleton struct{}

var eagerInitializationInstance = &EagerInitializationSingleton{}

func GetEagerInitializationInstance() *EagerInitializationSingleton {
	return eagerInitializationInstance
}
