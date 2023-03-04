package stat

import (
	"fmt"
	"testing"
)

func TestZTable_FindZScore(t *testing.T) {
	// Creating a new z-score table
	zTable := NewZTable(nil)

	// To find the closest z-score given a percentage
	ci := float64(90)
	alpha := 1 - ci/100
	zScore, err := zTable.FindZScore(alpha / 2)
	fmt.Println(-zScore, err)
}
