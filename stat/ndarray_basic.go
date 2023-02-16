package stat

func (self NDArray[T]) Name() string {
	return "x"
}

func (self NDArray[T]) Rename(name string) {

}

func (self NDArray[T]) Type() Type {
	return checkoutRawType(self)
}

func (self NDArray[T]) Values() any {
	return []T(self)
}

func (self NDArray[T]) Reverse() Series {
	return Reverse(self)
}
