package stat

func (r RollingAndExpandingMixin) Std() Series {
	s := r.Series.Empty()
	for _, block := range r.GetBlocks() {
		//// 1. 排序处理方式
		//if block.Len() == 0 {
		//	s.Append(s.NaN())
		//	continue
		//}
		//sort.Sort(block)
		//r := RangeFinite(-1)
		//_s := block.Select(r)
		//s.Append(_s.Values())
		// 2. Series.Max方法
		s.Append(block.Std())
	}
	return s
}
