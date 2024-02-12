package pandas

func (r RollingAndExpandingMixin) Std() Series {
	s := r.Series.Empty()
	for _, block := range r.GetBlocks() {
		s = s.Append(block.Std())
	}
	return s
}
