package pandas

import "fmt"

func ExampleConvect() {
	v1 := []float64{1.1, 2.2, 3.3}
	v2 := Convect[float64](v1)
	fmt.Println(v2)
	// Output
	// dtype[int]: 1,2,3
}
