package singleton

import (
	"reflect"
	"sync"
)

type Singleton struct {
	mutex sync.Mutex
	objs  map[string]interface{}
}

// make sure there is only one obj
var (
	once      sync.Once
	singleton *Singleton
)

func getInstance() *Singleton {
	once.Do(func() {
		singleton = &Singleton{
			objs: make(map[string]interface{}),
		}
	})
	return singleton

}

// user interface

// If you dont care about the performance.it will be ok
func GetInstance[T any]() interface{} {
	s := getInstance()
	s.mutex.Lock()
	defer s.mutex.Unlock()

	obj := new(T)

	tn := reflect.TypeOf(obj).String()

	if o, exist := s.objs[tn]; exist {
		return o
	}

	s.objs[tn] = obj

	return obj

}
