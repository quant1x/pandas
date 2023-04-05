package stat

import (
	"fmt"
	"testing"
)

func TestDecimal(t *testing.T) {
	f := float64(1.234567891)
	fmt.Println(Decimal(f, 2))
}
