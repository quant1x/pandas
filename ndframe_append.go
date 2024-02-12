package pandas

import (
	"gitee.com/quant1x/gox/api"
	"gitee.com/quant1x/num"
	"reflect"
)

func (this *NDFrame) Reverse() Series {
	s := this.Empty()

	if this.type_ == SERIES_TYPE_BOOL {
		values := api.Reverse(this.values.([]bool))
		s = s.Append(values)
	} else if this.type_ == SERIES_TYPE_INT64 {
		values := api.Reverse(this.values.([]int64))
		s = s.Append(values)
	} else if this.type_ == SERIES_TYPE_FLOAT32 {
		values := api.Reverse(this.values.([]float32))
		s = s.Append(values)
	} else if this.type_ == SERIES_TYPE_FLOAT64 {
		values := api.Reverse(this.values.([]float64))
		s = s.Append(values)
	} else {
		values := api.Reverse(this.values.([]string))
		s = s.Append(values)
	}
	return s
}

// 插入一条记录
func (this *NDFrame) insert(idx, size int, v any) {
	if this.type_ == SERIES_TYPE_BOOL {
		val := num.AnyToBool(v)
		ndFrameAssign[bool](this, idx, size, val)
	} else if this.type_ == SERIES_TYPE_INT64 {
		val := num.AnyToInt64(v)
		ndFrameAssign[int64](this, idx, size, val)
	} else if this.type_ == SERIES_TYPE_FLOAT32 {
		val := num.AnyToFloat32(v)
		ndFrameAssign[float32](this, idx, size, val)
	} else if this.type_ == SERIES_TYPE_FLOAT64 {
		val := num.AnyToFloat64(v)
		ndFrameAssign[float64](this, idx, size, val)
	} else {
		val := num.AnyToString(v)
		ndFrameAssign[string](this, idx, size, val)
	}
}

// Append 批量增加记录
func (this *NDFrame) Append(values ...any) Series {
	size := 0
	for idx, v := range values {
		switch val := v.(type) {
		case nil, int8, uint8, int16, uint16, int32, uint32, int64, uint64, int, uint, float32, float64, bool, string:
			// 基础类型
			this.insert(idx, size, val)
		default:
			vv := reflect.ValueOf(val)
			vk := vv.Kind()
			switch vk {
			//case reflect.Invalid: // {interface} nil
			//	series.ndFrameAssign(idx, size, Nil2Float64)
			case reflect.Slice, reflect.Array: // 切片或数组
				for i := 0; i < vv.Len(); i++ {
					tv := vv.Index(i).Interface()
					this.insert(idx, size, tv)
				}
			case reflect.Struct: // 忽略结构体
				continue
			default:
				this.insert(idx, size, nil)
			}
		}
	}
	return this
}

func (this *NDFrame) Concat(x Series) Series {
	y := this.Copy()
	y = y.Append(x.Values())
	return y
}
