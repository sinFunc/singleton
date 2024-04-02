package singleton

import (
	"reflect"
	"sync"
)

var objs sync.Map

// It is more convinient for user to use
// because user can convert the type which they want
func GetInstance[T any]() interface{} {
	var to *T
	t := reflect.TypeOf(to)

	o, exist := objs.Load(t)
	if exist {
		return o.(*T)
	}

	obj := new(T)
	objs.Store(t, obj)

	return obj

}

func GetInstanceTemplate[T any]() (obj *T) {
	t := reflect.TypeOf(obj)

	o, exist := objs.Load(t)
	if exist {
		return o.(*T)
	}

	obj = new(T)
	objs.Store(t, obj)

	return

}
