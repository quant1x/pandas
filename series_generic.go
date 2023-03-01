package pandas

import (
	"fmt"
	"gitee.com/quant1x/pandas/stat"
	"reflect"
)

// NewSeriesWithoutType 不带类型创新一个新series
func NewSeriesWithoutType(name string, values ...any) stat.Series {
	_type, err := stat.DetectTypeBySlice(values...)
	if err != nil {
		return nil
	}
	return NewSeriesWithType(_type, name, values...)
}

// NewSeriesWithType 通过类型创新一个新series
func NewSeriesWithType(_type stat.Type, name string, values ...any) stat.Series {
	frame := NDFrame{
		formatter: stat.DefaultFormatter,
		name:      name,
		type_:     stat.SERIES_TYPE_INVAILD,
		nilCount:  0,
		rows:      0,
		//values:    []E{},
	}

	frame.type_ = _type
	if frame.type_ == stat.SERIES_TYPE_BOOL {
		// bool
		//frame.values = reflect.MakeSlice(stat.TypeBool, 0, 0).Interface()
		frame.values = make([]bool, 0)
	} else if frame.type_ == stat.SERIES_TYPE_INT64 {
		// int64
		//frame.values = reflect.MakeSlice(stat.TypeInt64, 0, 0).Interface()
		frame.values = make([]int64, 0)
	} else if frame.type_ == stat.SERIES_TYPE_FLOAT32 {
		// float32
		//frame.values = reflect.MakeSlice(stat.TypeFloat32, 0, 0).Interface()
		frame.values = make([]float32, 0)
	} else if frame.type_ == stat.SERIES_TYPE_FLOAT64 {
		// float64
		//frame.values = reflect.MakeSlice(stat.TypeFloat64, 0, 0).Interface()
		frame.values = make([]float64, 0)
	} else {
		// string, 字符串最后容错使用
		//frame.values = reflect.MakeSlice(stat.TypeString, 0, 0).Interface()
		frame.values = make([]string, 0)
	}
	frame.Append(values...)

	return &frame
}

// NewSeries 指定类型创建序列
func NewSeries(t stat.Type, name string, vals any) stat.Series {
	var series stat.Series
	if t == stat.SERIES_TYPE_BOOL {
		series = NewSeriesWithType(stat.SERIES_TYPE_BOOL, name, vals)
	} else if t == stat.SERIES_TYPE_INT64 {
		series = NewSeriesWithType(stat.SERIES_TYPE_INT64, name, vals)
	} else if t == stat.SERIES_TYPE_STRING {
		series = NewSeriesWithType(stat.SERIES_TYPE_STRING, name, vals)
	} else if t == stat.SERIES_TYPE_FLOAT64 {
		series = NewSeriesWithType(stat.SERIES_TYPE_FLOAT64, name, vals)
	} else {
		// 默认全部强制转换成float32
		series = NewSeriesWithType(stat.SERIES_TYPE_FLOAT32, name, vals)
	}
	return series
}

// GenericSeries 泛型方法, 构造序列, 比其它方式对类型的统一性要求更严格
func GenericSeries[T stat.GenericType](name string, values ...T) stat.Series {
	// 第一遍, 确定类型, 找到第一个非nil的值
	var _type stat.Type = stat.SERIES_TYPE_STRING
	for _, v := range values {
		// 泛型处理这里会出现一个错误, invalid operation: v == nil (mismatched types T and untyped nil)
		//if v == nil {
		//	continue
		//}
		vv := reflect.ValueOf(v)
		vk := vv.Kind()
		switch vk {
		case reflect.Bool:
			_type = stat.SERIES_TYPE_BOOL
		case reflect.Int64:
			_type = stat.SERIES_TYPE_INT64
		case reflect.Float32:
			_type = stat.SERIES_TYPE_FLOAT32
		case reflect.Float64:
			_type = stat.SERIES_TYPE_FLOAT64
		case reflect.String:
			_type = stat.SERIES_TYPE_STRING
		default:
			panic(fmt.Errorf("unknown type, %+v", v))
		}
		break
	}
	return NewSeries(_type, name, values)
}
