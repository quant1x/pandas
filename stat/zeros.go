package stat

// Zeros Return a new array of given shape and type, filled with zeros.
//
//	args[0] dtype 基础数据类型
func Zeros[T BaseType](shape int) []T {
	//var __dtype reflect.Kind = reflect.Invalid
	//if len(args) > 0 {
	//	// 第一个参数为是否copy
	//	if _cp, ok := args[0].(reflect.Kind); ok {
	//		__dtype = _cp
	//	}
	//}
	//switch __dtype {
	//case reflect.Invalid:
	//case reflect.Int:
	//
	//default:
	//}
	var t T
	return Repeat(t, shape)
}
