package pandas

func (self *NDFrame) Max() any {
	values := self.Values()
	switch rows := values.(type) {
	case []string:
		max := ""
		i := 0
		for idx, iv := range rows {
			if StringIsNaN(iv) {
				continue
			}
			if iv > max {
				max = iv
				i += 1
			}
			_ = idx
		}
		if i > 0 {
			return max
		}
		return StringNaN
	case []int64:
		max := int64(0)
		//i := 0
		for idx, iv := range rows {
			if Float64IsNaN(float64(iv)) {
				continue
			}
			if iv > max {
				max = iv
				//i = idx
			}
			_ = idx
		}
		return max
	case []float64:
		max := float64(0)
		i := 0
		for idx, iv := range rows {
			if Float64IsNaN(iv) {
				continue
			}
			if iv > max {
				max = iv
				i += 1
			}
			_ = idx
		}
		if i > 0 {
			return max
		}
		return Nil2Float64
	}
	return Nil2Float64
}
