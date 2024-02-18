package pandas

// FillNa dataframe实现FillNa
func (this DataFrame) FillNa(v any, inplace bool) {
	for _, series := range this.columns {
		if series.Len() > 0 {
			series.FillNa(v, inplace)
		}
	}
}
