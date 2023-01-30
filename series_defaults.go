// Copyright 2018-20 PJ Engineering and Business Solutions Pty. Ltd. All rights reserved.

package pandas

import (
	"fmt"
	"github.com/viterin/vek"

	"github.com/google/go-cmp/cmp"
)

// 初始化 vek
func init() {
	// 开启加速选项
	vek.SetAcceleration(true)
}

// DefaultIsEqualFunc is the default comparitor to determine if
// two values in the series are the same.
func DefaultIsEqualFunc(a, b interface{}) bool {
	return cmp.Equal(a, b)
}

// DefaultValueFormatter will return a string representation
// of the data in a particular row.
func DefaultValueFormatter(v interface{}) string {
	if v == nil {
		return "NaN"
	}
	return fmt.Sprintf("%v", v)
}
