package areflect

import (
	"reflect"
	"runtime"
	"strings"
)

//判读是否为nil
func IsNil(i interface{}) bool {
	vi := reflect.ValueOf(i)
	if vi.Kind() == reflect.Ptr {
		return vi.IsNil()
	}
	return false
}

func GetFuncAllName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func GetFuncName(i interface{}) string {
	allName := GetFuncAllName(i)
	names := strings.Split(allName, ".")
	return names[len(names)-1]
}
