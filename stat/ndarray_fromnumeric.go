package stat

func (self NDArray[T]) ArgMax() int {
	return ArgMax2(self)
}

func (self NDArray[T]) ArgMin() int {
	return ArgMin2(self)
}
