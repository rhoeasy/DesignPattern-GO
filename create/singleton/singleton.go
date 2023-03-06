package singleton

import "sync"

type Singleton struct{}

var singleton *Singleton

// 饿汉模式
func init() {
	singleton = &Singleton{}
}

func GetInstance() *Singleton {
	return singleton
}

var (
	// 懒汉模式
	lazySingleton *Singleton
	once          = &sync.Once{}
)

func GetLazyInstance() *Singleton {
	if lazySingleton != nil {
		once.Do(func() {
			lazySingleton = &Singleton{}
		})
	}
	return lazySingleton
}
