package stat

// Diff 元素的第一个离散差
//
//	First discrete difference of element.
//	Calculates the difference of a {klass} element compared with another
//	element in the {klass} (default is element in previous row).
func Diff[T Number](s []T, param any) []T {
	blocks := Rolling[T](s, param)
	var d []T
	var front = typeDefault[T]()
	for _, block := range blocks {
		vs := block
		vl := len(block)
		if vl == 0 {
			d = append(d, typeDefault[T]())
			continue
		}
		vf := vs[0]
		vc := vs[vl-1]
		if DTypeIsNaN(Any2DType(vc)) || DTypeIsNaN(Any2DType(front)) {
			front = vf
			d = append(d, typeDefault[T]())
			continue
		}
		diff := vc - front
		d = append(d, diff)
		front = vf
	}

	return d
}

func Diff2[T BaseType](s []T, param any) []T {
	var d any
	switch vs := any(s).(type) {
	case []float32:
		d = Diff(vs, param)
	case []float64:
		d = Diff(vs, param)
	case []int:
		d = Diff(vs, param)
	case []int8:
		d = Diff(vs, param)
	case []int16:
		d = Diff(vs, param)
	case []int32:
		d = Diff(vs, param)
	case []int64:
		d = Diff(vs, param)
	//case []uint, []uint8, []uint16, []uint32, []uint64, []uintptr:
	//	d = xv
	default:
		// 其它类型原样返回
		panic(Throw(any(s)))
	}

	return d.([]T)
}
