package stat

func (r RollingAndExpandingMixin) Max() (s Series) {
	s = r.Series.Empty()
	for _, block := range r.GetBlocks() {
		s = s.Append(block.Max())
	}
	return
}
