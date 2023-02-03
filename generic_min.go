package pandas

func (self *NDFrame) Min() any {
	values := self.Values()
	switch rows := values.(type) {
	case []string:
		min := ""
		i := 0
		for idx, iv := range rows {
			if StringIsNaN(iv) {
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
		if i > 0 {
			return min
		}
		return StringNaN
	case []int64:
		min := int64(0)
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
	case []float64:
		min := float64(0)
		i := 0
		for idx, iv := range rows {
			if Float64IsNaN(iv) {
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
		if i > 0 {
			return min
		}
		return Nil2Float64
	}
	return Nil2Float64
}
