package stat

type NDArray []DType

type Frame interface {
	Len() int
}

type Array[T Number] []T

func (a Array[T]) Len() int {
	return len(a)
}
