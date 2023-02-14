package stat

import (
	"gitee.com/quant1x/pandas/exception"
	"reflect"
)

const (
	errorTypeBase = 0
)

var (
	// ErrUnsupportedType 不支持的类型
	ErrUnsupportedType = exception.New(errorTypeBase+0, "Unsupported type")
)

func Throw(tv any) *exception.Exception {
	typeName := reflect.TypeOf(tv).String()
	return exception.New(errorTypeBase+1, "Unsupported type: "+typeName)
}
