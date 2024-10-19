package common

import (
	"fmt"
	"reflect"
	"runtime"

	customError "github.com/GooDu-dev/gd-practical-project-backend/utils/error"
)

func IsDefaultValueOrNil(data any) (output bool) {
	value := reflect.ValueOf(data)
	switch kind := reflect.TypeOf(data).Kind(); kind {
	case reflect.Array:
		// if array contains (nil || default_value) return true
		output = false
		for _, d := range []any{data} {
			value = reflect.ValueOf(d)
			output = value.IsNil() || value.IsZero()
			if output {
				return true
			}
			output = value.IsNil() || value.IsZero()
		}
		return output
	case reflect.Pointer:
		fmt.Println(value.IsNil(), value.IsZero())
		output = value.IsNil() || value.IsZero()
	case reflect.String:
		output = value.String() == ""
	case reflect.Int:
		output = value.IsZero()
	}
	return output
}

func DeepIsDefaultValueOrNil(data interface{}) (err error) {
	kind := reflect.TypeOf(data)
	if kind.Kind() != reflect.Struct {
		return customError.DataTypeIsNotStructError
	}

	value := reflect.ValueOf(data)
	for i := 0; i < value.NumField(); i++ {
		if canNull(kind.Field(i).Type) {
			continue
		}
		if reflect.DeepEqual(value.Field(i).Interface(), reflect.Zero(kind.Field(i).Type).Interface()) {
			return customError.FieldContainsNilOrDefaultValueError
		}
	}
	return nil
}

func canNull(t reflect.Type) bool {
	switch t.Kind() {
	case reflect.Ptr, reflect.Map, reflect.Slice, reflect.Interface, reflect.Chan:
		return true
	}
	return false
}

func GetFunctionWithPackageName() string {
	pc, _, _, _ := runtime.Caller(1)
	funcName := runtime.FuncForPC(pc).Name()
	packageName, functionName := splitFunctionName(funcName)
	return fmt.Sprintf("%s:%s", packageName, functionName)
}

func splitFunctionName(funcName string) (packageName string, functionName string) {
	for i := len(funcName) - 1; i >= 0; i-- {
		if funcName[i] == '.' {
			packageName = funcName[:i]
			functionName = funcName[i+1:]
			break
		}
	}
	return packageName, functionName
}
