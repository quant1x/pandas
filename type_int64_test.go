package pandas

import (
	"fmt"
	"math"
	"testing"
)

func TestInt64(t *testing.T) {
	fmt.Println(math.MaxInt64 == MaxInt64)
	fmt.Println(math.MinInt64 == MinInt64)
}
