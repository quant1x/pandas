package pandas

import (
	"reflect"
)

// 初始化全局的私有变量
var (
	rawBool     bool    = true
	typeBool            = reflect.TypeOf([]bool{})
	rawInt32    int32   = int32(0)
	typeInt32           = reflect.TypeOf([]int32{})
	rawInt64    int64   = int64(0)
	typeInt64           = reflect.TypeOf([]int64{})
	rawFloat32  float32 = float32(0)
	typeFloat32         = reflect.TypeOf([]float32{})
	rawFloat64  float64 = float64(0)
	typeFloat64         = reflect.TypeOf([]float64{})
	typeString          = reflect.TypeOf([]string{})
)

func CreateSeries(t Type, name string, v ...any) Series {
	frame := NDFrame{
		formatter: DefaultFormatter,
		name:      name,
		type_:     SERIES_TYPE_INVAILD,
		nilCount:  0,
		rows:      0,
		//values:    []E{},
	}
	return &frame
}

// NewSeriesWithoutType 不带类型创新一个新series
func NewSeriesWithoutType(name string, values ...interface{}) Series {
	_type, err := detectTypeBySlice(values)
	if err != nil {
		return nil
	}
	return NewSeriesWithType(_type, name, values...)
}

// NewSeriesWithType 通过类型创新一个新series
func NewSeriesWithType(_type Type, name string, values ...interface{}) Series {
	frame := NDFrame{
		formatter: DefaultFormatter,
		name:      name,
		type_:     SERIES_TYPE_INVAILD,
		nilCount:  0,
		rows:      0,
		//values:    []E{},
	}
	//_type, err := detectTypeBySlice(values)
	//if err != nil {
	//	return nil
	//}
	frame.type_ = _type
	if frame.type_ == SERIES_TYPE_BOOL {
		// bool
		frame.values = reflect.MakeSlice(typeBool, 0, 0).Interface()
	} else if frame.type_ == SERIES_TYPE_INT {
		// int64
		frame.values = reflect.MakeSlice(typeInt64, 0, 0).Interface()
	} else if frame.type_ == SERIES_TYPE_FLOAT {
		// float64
		frame.values = reflect.MakeSlice(typeFloat64, 0, 0).Interface()
	} else {
		// string, 字符串最后容错使用
		frame.values = reflect.MakeSlice(typeString, 0, 0).Interface()
	}
	//series.Data = make([]float64, 0) // Warning: filled with 0.0 (not NaN)
	//size := len(series.values)
	//size := 0
	//for idx, v := range values {
	//	switch val := v.(type) {
	//	case nil, int8, uint8, int16, uint16, int32, uint32, int64, uint64, int, uint, float32, float64, bool, string:
	//		// 基础类型
	//		series_append(&frame, idx, size, val)
	//	default:
	//		vv := reflect.ValueOf(val)
	//		vk := vv.Kind()
	//		switch vk {
	//		//case reflect.Invalid: // {interface} nil
	//		//	series.assign(idx, size, Nil2Float64)
	//		case reflect.Slice, reflect.Array: // 切片或数组
	//			for i := 0; i < vv.Len(); i++ {
	//				tv := vv.Index(i).Interface()
	//				//series.assign(idx, size, str)
	//				series_append(&frame, idx, size, tv)
	//			}
	//		case reflect.Struct: // 忽略结构体
	//			continue
	//		default:
	//			series_append(&frame, idx, size, nil)
	//		}
	//	}
	//}
	frame.Append(values...)

	return &frame
}
