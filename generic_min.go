package pandas

func (self *NDFrame) Min() any {
	values := self.Values()
	switch rows := values.(type) {
	case []bool:
		min := true
		i := 0
		for idx, iv := range rows {
			if !iv && min {
				min = iv
				i += 1
			}
			_ = idx
		}
		if i > 0 {
			return min
		}
		return false
	case []string:
		min := ""
		hasNaN := false
		i := 0
		for idx, iv := range rows {
			if StringIsNaN(iv) {
				hasNaN = true
				break
			} else if i < 1 {
				min = iv
				i += 1
			}
			if iv < min {
				min = iv
				i += 1
			}
			_ = idx
		}
		if hasNaN {
			return StringNaN
		} else if i > 0 {
			return min
		}
		return StringNaN
	case []int64:
		min := MaxInt64
		i := 0
		for idx, iv := range rows {
			if Float64IsNaN(float64(iv)) {
				continue
			} else if i < 1 {
				min = iv
				i += 1
			}
			if iv < min {
				min = iv
				i += 1
			}
			_ = idx
		}
		return min
	case []float32:
		min := MaxFloat32
		hasNan := false
		i := 0
		for idx, iv := range rows {
			if Float32IsNaN(iv) {
				hasNan = true
				break
			}
			if iv < min {
				min = iv
				i += 1
			}
			_ = idx
		}
		if hasNan {
			return Nil2Float32
		} else if i > 0 {
			return min
		}
		return Nil2Float32
	//case []float32:
	//	if self.Len() == 0 {
	//		return Nil2Float32
	//	}
	//	return stat.Min(rows)
	case []float64:
		min := MaxFloat64
		hasNaN := false
		i := 0
		for idx, iv := range rows {
			if Float64IsNaN(iv) {
				hasNaN = true
				break
			}
			if iv < min {
				min = iv
				i += 1
			}
			_ = idx
		}
		if hasNaN {
			return Nil2Float64
		} else if i > 0 {
			return min
		}
		return Nil2Float64
	//case []float64:
	//	if self.Len() == 0 {
	//		return Nil2Float64
	//	}
	//	return stat.Min(rows)
	default:
		panic(ErrUnsupportedType)
	}
	return Nil2Float64
}
