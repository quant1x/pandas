package pandas

import "gitee.com/quant1x/pandas/stat"

// Len 获得行数, 实现sort.Interface接口的获取元素数量方法
func (this *NDFrame) Len() int {
	return this.rows
}

// Less 实现sort.Interface接口的比较元素方法
func (this *NDFrame) Less(i, j int) bool {
	if this.type_ == stat.SERIES_TYPE_BOOL {
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
	} else if this.type_ == stat.SERIES_TYPE_INT64 {
		values := this.Values().([]int64)
		return values[i] < values[j]
	} else if this.type_ == stat.SERIES_TYPE_FLOAT32 {
		values := this.Values().([]float32)
		return values[i] < values[j]
	} else if this.type_ == stat.SERIES_TYPE_FLOAT64 {
		values := this.Values().([]float64)
		return values[i] < values[j]
	} else if this.type_ == stat.SERIES_TYPE_STRING {
		values := this.Values().([]string)
		return values[i] < values[j]
	} else {
		// SERIES_TYPE_INVAILD
		// 应该到不了这里, Len()会返回0
		panic(ErrUnsupportedType)
	}
	return false

}

// Swap 实现sort.Interface接口的交换元素方法
func (this *NDFrame) Swap(i, j int) {
	if this.type_ == stat.SERIES_TYPE_BOOL {
		values := this.Values().([]bool)
		values[i], values[j] = values[j], values[i]
	} else if this.type_ == stat.SERIES_TYPE_INT64 {
		values := this.Values().([]int64)
		values[i], values[j] = values[j], values[i]
	} else if this.type_ == stat.SERIES_TYPE_FLOAT32 {
		values := this.Values().([]float32)
		values[i], values[j] = values[j], values[i]
	} else if this.type_ == stat.SERIES_TYPE_FLOAT64 {
		values := this.Values().([]float64)
		values[i], values[j] = values[j], values[i]
	} else if this.type_ == stat.SERIES_TYPE_STRING {
		values := this.Values().([]string)
		values[i], values[j] = values[j], values[i]
	} else {
		// SERIES_TYPE_INVAILD
		// 应该到不了这里, Len()会返回0
		panic(ErrUnsupportedType)
	}
}
