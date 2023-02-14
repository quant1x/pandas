package v1

func (r RollingAndExpandingMixin) Std() Series {
	s := r.series.Empty()
	for _, block := range r.getBlocks() {
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
