package pandas

import (
	"fmt"
	"gitee.com/quant1x/pandas/stat"
)

func bool2String(v bool) string {
	if v {
		return stat.True2String
	}
	return stat.False2String
}

// float64转string
func float64ToString(v float64) string {
	if isNaN(v) {
		return stat.StringNaN
	}
	return fmt.Sprintf("%f", v)
}

func int64ToString(v int64) string {
	if stat.Float64IsNaN(float64(v)) {
		return stat.StringNaN
	}
	return fmt.Sprint(v)
}

// bool转float32
func boolToFloat32(b bool) float32 {
	if b {
		return stat.True2Float32
	}
	return stat.False2Float32
}
