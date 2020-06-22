package core

import (
	"fmt"
	"reflect"
)

func Call(f interface{}, args ...interface{}) ([]interface{}, error) {
	fType := reflect.TypeOf(f)
	if fType.Kind() != reflect.Func {
		return nil, fmt.Errorf("f is not func: %+v", f)
	}
	if fType.NumIn() != len(args) {
		return nil, fmt.Errorf("length of args is expected %d, but actual %d", fType.NumIn(), len(args))
	}

	argValues := make([]reflect.Value, 0, fType.NumIn())
	for i := 0; i < fType.NumIn(); i++ {
		paramType := fType.In(i)
		argType := reflect.ValueOf(args[i]).Type()
		if !paramType.AssignableTo(argType) {
			return nil, fmt.Errorf("i-th arg's type is %+v and is not assignable to %+v", argType, paramType)
		}
		argValues = append(argValues, reflect.ValueOf(args[i]))
	}

	resValues := reflect.ValueOf(f).Call(argValues)

	res := make([]interface{}, 0, fType.NumOut())
	for i := 0; i < fType.NumOut(); i++ {
		res = append(res, resValues[i].Interface())
	}
	return res, nil
}
