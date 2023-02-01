package pandas

// FillNa dataframe实现FillNa
func (self DataFrame) FillNa(v any, inplace bool) {
	for _, series := range self.columns {
		if series.Len() > 0 {
			series.FillNa(v, inplace)
		}
	}
}
