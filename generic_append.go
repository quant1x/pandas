package pandas

import "reflect"

// 插入一条记录
func (self *NDFrame) insert(idx, size int, v any) {
	if self.type_ == SERIES_TYPE_BOOL {
		val := AnyToBool(v)
		assign[bool](self, idx, size, val)
	} else if self.type_ == SERIES_TYPE_INT64 {
		val := AnyToInt64(v)
		assign[int64](self, idx, size, val)
	} else if self.type_ == SERIES_TYPE_FLOAT32 {
		val := AnyToFloat32(v)
		assign[float32](self, idx, size, val)
	} else if self.type_ == SERIES_TYPE_FLOAT64 {
		val := AnyToFloat64(v)
		assign[float64](self, idx, size, val)
	} else {
		val := AnyToString(v)
		assign[string](self, idx, size, val)
	}
}

// Append 批量增加记录
func (self *NDFrame) Append(values ...interface{}) {
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
}
