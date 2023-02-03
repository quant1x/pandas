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

func NewSeries2(name string, values ...interface{}) Series {
	frame := NDFrame{
		formatter: DefaultFormatter,
		name:      name,
		type_:     SERIES_TYPE_INVAILD,
		nilCount:  0,
		rows:      0,
		//values:    []E{},
	}
	_type, err := detectTypeBySlice(values)
	//fmt.Println(_type, err)
	if err != nil {
		return nil
	}
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
	size := 0
	for idx, v := range values {
		if frame.type_ == SERIES_TYPE_BOOL {
			val := AnyToBool(v)
			assign[bool](&frame, idx, size, val)
		} else if frame.type_ == SERIES_TYPE_INT {
			val := AnyToInt64(v)
			assign[int64](&frame, idx, size, val)
		} else if frame.type_ == SERIES_TYPE_FLOAT {
			val := AnyToFloat64(v)
			assign[float64](&frame, idx, size, val)
		} else {
			val := AnyToString(v)
			assign[string](&frame, idx, size, val)
		}
	}

	return &frame
}
