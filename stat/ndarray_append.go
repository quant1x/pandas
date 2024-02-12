package stat

import (
	"gitee.com/quant1x/num"
	"reflect"
)

// 赋值
func ndArrayAssign[T num.BaseType](type_ Type, array Series, idx, size int, v T) Series {
	_vv := reflect.ValueOf(v)
	_vi := _vv.Interface()
	// float和string类型有可能是NaN, 对nil和NaN进行计数
	if type_ == SERIES_TYPE_FLOAT32 && num.Float32IsNaN(_vi.(float32)) {
		//array.nilCount++
	} else if type_ == SERIES_TYPE_FLOAT64 && num.Float64IsNaN(_vi.(float64)) {
		//array.nilCount++
	} else if type_ == SERIES_TYPE_STRING && num.StringIsNaN(_vi.(string)) {
		//array.nilCount++
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
			_vv.SetString(num.StringNaN)
			// 执行之后, 通过debug可以看到assign入参的v已经变成了"NaN"
		}
	}
	values := (array).Values().([]T)
	// 确保只添加了1个元素
	if idx < size {
		values[idx] = v
	} else {
		values = append(values, v)
	}
	// 行数+1
	//array.rows += 1
	return NDArray[T](values)
}

// 插入一条记录
func (self NDArray[T]) insert(idx, size int, v any) NDArray[T] {
	type_ := checkoutRawType(self)
	if type_ == SERIES_TYPE_BOOL {
		val := num.AnyToBool(v)
		an := ndArrayAssign[bool](type_, self, idx, size, val)
		self = an.(NDArray[T])
	} else if type_ == SERIES_TYPE_INT64 {
		val := num.AnyToInt64(v)
		an := ndArrayAssign[int64](type_, self, idx, size, val)
		self = an.(NDArray[T])
	} else if type_ == SERIES_TYPE_FLOAT32 {
		val := num.AnyToFloat32(v)
		an := ndArrayAssign[float32](type_, self, idx, size, val)
		self = an.(NDArray[T])
	} else if type_ == SERIES_TYPE_FLOAT64 {
		val := num.AnyToFloat64(v)
		an := ndArrayAssign[float64](type_, self, idx, size, val)
		self = an.(NDArray[T])
	} else {
		val := num.AnyToString(v)
		an := ndArrayAssign[string](type_, self, idx, size, val)
		self = an.(NDArray[T])
	}
	return self
}

func (self NDArray[T]) Append(values ...any) Series {
	size := 0
	for idx, v := range values {
		switch val := v.(type) {
		case nil, int8, uint8, int16, uint16, int32, uint32, int64, uint64, int, uint, float32, float64, bool, string:
			// 基础类型
			self = self.insert(idx, size, val)
		default:
			vv := reflect.ValueOf(val)
			vk := vv.Kind()
			switch vk {
			case reflect.Slice, reflect.Array: // 切片或数组
				for i := 0; i < vv.Len(); i++ {
					tv := vv.Index(i).Interface()
					self = self.insert(idx, size, tv)
				}
			case reflect.Struct: // 忽略结构体
				continue
			default:
				self = self.insert(idx, size, nil)
			}
		}
	}
	return self
}

func (self NDArray[T]) Concat(x Series) Series {
	y := self.Copy()
	y = y.Append(x.Values())
	return y
}
