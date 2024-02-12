package pandas

func (r RollingAndExpandingMixin) Min() (s Series) {
	s = r.Series.Empty()
	for _, block := range r.GetBlocks() {
		s = s.Append(block.Min())
	}
	return
}
