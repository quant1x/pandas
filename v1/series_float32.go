package v1

// TODO:留给于总的作业
type SeriesFloat32 struct {
	NDFrame
}

func (self *SeriesFloat32) Name() string {
	return self.name
}

func (self *SeriesFloat32) Rename(n string) {
	//TODO implement me
	panic("implement me")
}

func (self *SeriesFloat32) Type() Type {
	//TODO implement me
	panic("implement me")
}

func (self *SeriesFloat32) Len() int {
	//TODO implement me
	panic("implement me")
}

func (self *SeriesFloat32) Values() any {
	//TODO implement me
	panic("implement me")
}

func (self *SeriesFloat32) Empty() Series {
	//TODO implement me
	panic("implement me")
}

func (self *SeriesFloat32) Records() []string {
	//TODO implement me
	panic("implement me")
}

func (self *SeriesFloat32) Copy() Series {
	//TODO implement me
	panic("implement me")
}

func (self *SeriesFloat32) Subset(start, end int, opt ...any) Series {
	//TODO implement me
	panic("implement me")
}

func (self *SeriesFloat32) Repeat(x any, repeats int) Series {
	//TODO implement me
	panic("implement me")
}

func (self *SeriesFloat32) Shift(periods int) Series {
	//TODO implement me
	panic("implement me")
}

func (self *SeriesFloat32) RollingV1(window int) RollingWindowV1 {
	//TODO implement me
	panic("implement me")
}

func (self *SeriesFloat32) Mean() float64 {
	//TODO implement me
	panic("implement me")
}

func (self *SeriesFloat32) StdDev() float64 {
	//TODO implement me
	panic("implement me")
}

func (self *SeriesFloat32) FillNa(v any, inplace bool) {
	//TODO implement me
	panic("implement me")
}

func (self *SeriesFloat32) Max() any {
	//TODO implement me
	panic("implement me")
}
