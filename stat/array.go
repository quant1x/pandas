package stat

type NDArray []DType

type Array[T Number] []T

func (a Array[T]) Len() int {
	return len(a)
}

//type FloatX interface {
//	~float64 | []float64 | int64 | []int64 | int | []int | int32 | []int32 | [][]float64 | [][]int | [][]int64 | [][]int32
//}
//
//// IsVector checks if a variable is a slice
//func IsVector[T FloatX](obj T) bool {
//	switch reflect.TypeOf(obj).Kind() {
//	case reflect.Slice:
//		return true
//	default:
//		return false
//	}
//}
//
//// AsSlice converts a variable to a slice
//func AsSlice[T FloatX](obj T) []float64 {
//	switch reflect.TypeOf(obj).Kind() {
//	case reflect.Slice:
//		return any(obj).([]float64)
//	default:
//		return []float64{any(obj).(float64)}
//	}
//}
//
//// asFloat64 converts a variable to a float64
//func asFloat64[T FloatX](obj T) float64 {
//	switch reflect.TypeOf(obj).Kind() {
//	case reflect.Slice:
//		return any(obj).([]float64)[0]
//	case reflect.Int:
//		return float64(any(obj).(int))
//	case reflect.Int64:
//		return float64(any(obj).(int))
//	case reflect.Int32:
//		return float64(any(obj).(int))
//	default:
//		return any(obj).(float64)
//	}
//}
//
//// Round rounds a slice of numbers to a given decimal
//func Round[T FloatX](element T, decimals int) any {
//	if IsVector(element) {
//		var roundedArray []float64
//		array := AsSlice(element)
//		for i := range array {
//			roundedNum := math.Round(array[i]*math.Pow(10, float64(decimals))) / math.Pow(10, float64(decimals))
//			roundedArray = append(roundedArray, roundedNum)
//		}
//		return roundedArray
//	} else {
//		number := asFloat64(element)
//		return math.Round(number*math.Pow(10, float64(decimals))) / math.Pow(10, float64(decimals))
//	}
//
//}
