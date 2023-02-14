package v1

// Len 获得行数, 实现sort.Interface接口的获取元素数量方法
func (self *NDFrame) Len() int {
	return self.rows
}

// Less 实现sort.Interface接口的比较元素方法
func (self *NDFrame) Less(i, j int) bool {
	if self.type_ == SERIES_TYPE_BOOL {
		values := self.Values().([]bool)
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
	} else if self.type_ == SERIES_TYPE_INT64 {
		values := self.Values().([]int64)
		return values[i] < values[j]
	} else if self.type_ == SERIES_TYPE_FLOAT32 {
		values := self.Values().([]float32)
		return values[i] < values[j]
	} else if self.type_ == SERIES_TYPE_FLOAT64 {
		values := self.Values().([]float64)
		return values[i] < values[j]
	} else if self.type_ == SERIES_TYPE_STRING {
		values := self.Values().([]string)
		return values[i] < values[j]
	} else {
		// SERIES_TYPE_INVAILD
		// 应该到不了这里, Len()会返回0
		panic(ErrUnsupportedType)
	}
	return false

}

// Swap 实现sort.Interface接口的交换元素方法
func (self *NDFrame) Swap(i, j int) {
	if self.type_ == SERIES_TYPE_BOOL {
		values := self.Values().([]bool)
		values[i], values[j] = values[j], values[i]
	} else if self.type_ == SERIES_TYPE_INT64 {
		values := self.Values().([]int64)
		values[i], values[j] = values[j], values[i]
	} else if self.type_ == SERIES_TYPE_FLOAT32 {
		values := self.Values().([]float32)
		values[i], values[j] = values[j], values[i]
	} else if self.type_ == SERIES_TYPE_FLOAT64 {
		values := self.Values().([]float64)
		values[i], values[j] = values[j], values[i]
	} else if self.type_ == SERIES_TYPE_STRING {
		values := self.Values().([]string)
		values[i], values[j] = values[j], values[i]
	} else {
		// SERIES_TYPE_INVAILD
		// 应该到不了这里, Len()会返回0
		panic(ErrUnsupportedType)
	}
}
