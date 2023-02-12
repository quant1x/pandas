package stat

import "gitee.com/quant1x/pandas/exception"

const (
	errorTypeBase = 0
)

var (
	// ErrUnsupportedType 不支持的类型
	ErrUnsupportedType = exception.New(errorTypeBase+0, "Unsupported type")
)
