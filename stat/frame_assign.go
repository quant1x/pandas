package stat

import (
	"reflect"
)

// 赋值
func (self *GenericFrame[T]) assign(idx, size int, val any) {
	var v any
	if self.type_ == SERIES_TYPE_BOOL {
		v = AnyToBool(val)
	} else if self.type_ == SERIES_TYPE_INT64 {
		v = AnyToInt64(val)
	} else if self.type_ == SERIES_TYPE_FLOAT64 {
		v = AnyToFloat64(val)
	} else {
		v = AnyToString(val)
	}
	//// 检测类型
	//if self.type_ == SERIES_TYPE_INVAILD {
	//	_type, _ := detectTypes(v)
	//	if _type != SERIES_TYPE_INVAILD {
	//		self.type_ = _type
	//	}
	//}
	_vv := reflect.ValueOf(v)
	_vi := _vv.Interface()
	// float和string类型有可能是NaN, 对nil和NaN进行计数
	if self.type_ == SERIES_TYPE_FLOAT64 && Float64IsNaN(_vi.(float64)) {
		self.nilCount++
	} else if self.type_ == SERIES_TYPE_STRING && StringIsNaN(_vi.(string)) {
		self.nilCount++
		// 以下修正string的NaN值, 统一为"NaN"
		//_rv := reflect.ValueOf(StringNaN)
		//_vv.Set(_rv) // 这样赋值会崩溃
		// TODO:值可修改条件之一: 可被寻址
		// 通过反射修改变量值的前提条件之一: 这个值必须可以被寻址, 简单地说就是这个变量必须能被修改.
		// 第一步: 通过变量v反射(v的地址)
		_vp := reflect.ValueOf(&v)
		// 第二步: 取出v地址的元素(v的值)
		_vv := _vp.Elem()
		// 判断_vv是否能被修改
		if _vv.CanSet() {
			// 修改v的值为新值
			_vv.SetString(StringNaN)
			// 执行之后, 通过debug可以看到assign入参的v已经变成了"NaN"
		}
	}
	// 确保只添加了1个元素
	if idx < size {
		self.values[idx] = v.(T)
	} else {
		self.values = append(self.values, v.(T))
	}
	// 行数+1
	self.rows += 1
}
