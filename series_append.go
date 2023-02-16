package pandas

import (
	"gitee.com/quant1x/pandas/stat"
	"reflect"
)

func (self *NDFrame) Reverse() stat.Series {
	s := self.Empty()

	if self.type_ == stat.SERIES_TYPE_BOOL {
		values := stat.Reverse(self.values.([]bool))
		s = s.Append(values)
	} else if self.type_ == stat.SERIES_TYPE_INT64 {
		values := stat.Reverse(self.values.([]int64))
		s = s.Append(values)
	} else if self.type_ == stat.SERIES_TYPE_FLOAT32 {
		values := stat.Reverse(self.values.([]float32))
		s = s.Append(values)
	} else if self.type_ == stat.SERIES_TYPE_FLOAT64 {
		values := stat.Reverse(self.values.([]float64))
		s = s.Append(values)
	} else {
		values := stat.Reverse(self.values.([]string))
		s = s.Append(values)
	}
	return s
}

// 插入一条记录
func (self *NDFrame) insert(idx, size int, v any) {
	if self.type_ == stat.SERIES_TYPE_BOOL {
		val := stat.AnyToBool(v)
		assign[bool](self, idx, size, val)
	} else if self.type_ == stat.SERIES_TYPE_INT64 {
		val := stat.AnyToInt64(v)
		assign[int64](self, idx, size, val)
	} else if self.type_ == stat.SERIES_TYPE_FLOAT32 {
		val := stat.AnyToFloat32(v)
		assign[float32](self, idx, size, val)
	} else if self.type_ == stat.SERIES_TYPE_FLOAT64 {
		val := stat.AnyToFloat64(v)
		assign[float64](self, idx, size, val)
	} else {
		val := stat.AnyToString(v)
		assign[string](self, idx, size, val)
	}
}

// Append 批量增加记录
func (self *NDFrame) Append(values ...any) stat.Series {
	size := 0
	for idx, v := range values {
		switch val := v.(type) {
		case nil, int8, uint8, int16, uint16, int32, uint32, int64, uint64, int, uint, float32, float64, bool, string:
			// 基础类型
			self.insert(idx, size, val)
		default:
			vv := reflect.ValueOf(val)
			vk := vv.Kind()
			switch vk {
			//case reflect.Invalid: // {interface} nil
			//	series.assign(idx, size, Nil2Float64)
			case reflect.Slice, reflect.Array: // 切片或数组
				for i := 0; i < vv.Len(); i++ {
					tv := vv.Index(i).Interface()
					self.insert(idx, size, tv)
				}
			case reflect.Struct: // 忽略结构体
				continue
			default:
				self.insert(idx, size, nil)
			}
		}
	}
	return self
}
