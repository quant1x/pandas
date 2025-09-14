package pandas

import "gitee.com/quant1x/num"

func (this vector[T]) Len() int {
	return len(this)
}

// Less 实现sort.Interface接口的比较元素方法
func (this vector[T]) Less(i, j int) bool {
	type_ := this.Type()
	if type_ == SERIES_TYPE_BOOL {
		values := this.Values().([]bool)
		var (
			a = int(0)
			b = int(0)
		)
		if values[i] {
			a = 1
		}
		if values[j] {
			b = 1
		}
		return a < b
	} else if type_ == SERIES_TYPE_INT64 {
		values := this.Values().([]int64)
		return values[i] < values[j]
	} else if type_ == SERIES_TYPE_FLOAT32 {
		values := this.Values().([]float32)
		return values[i] < values[j]
	} else if type_ == SERIES_TYPE_FLOAT64 {
		values := this.Values().([]float64)
		return values[i] < values[j]
	} else if type_ == SERIES_TYPE_STRING {
		values := this.Values().([]string)
		return values[i] < values[j]
	} else {
		// SERIES_TYPE_INVAILD
		// 应该到不了这里, Len()会返回0
		panic(num.ErrUnsupportedType)
	}
	return false

}

// Swap 实现sort.Interface接口的交换元素方法
func (this vector[T]) Swap(i, j int) {
	type_ := this.Type()
	if type_ == SERIES_TYPE_BOOL {
		values := this.Values().([]bool)
		values[i], values[j] = values[j], values[i]
	} else if type_ == SERIES_TYPE_INT64 {
		values := this.Values().([]int64)
		values[i], values[j] = values[j], values[i]
	} else if type_ == SERIES_TYPE_FLOAT32 {
		values := this.Values().([]float32)
		values[i], values[j] = values[j], values[i]
	} else if type_ == SERIES_TYPE_FLOAT64 {
		values := this.Values().([]float64)
		values[i], values[j] = values[j], values[i]
	} else if type_ == SERIES_TYPE_STRING {
		values := this.Values().([]string)
		values[i], values[j] = values[j], values[i]
	} else {
		// SERIES_TYPE_INVAILD
		// 应该到不了这里, Len()会返回0
		panic(num.ErrUnsupportedType)
	}
}
