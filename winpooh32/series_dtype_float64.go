//go:build !series_f32

package winpooh32

import (
	"gitee.com/quant1x/pandas/winpooh32/math"
)

type DType = float64

const (
	EpsFp32 = 1e-7
	EpsFp64 = 1e-14
	Eps     = EpsFp64

	EnabledFloat32 = false
)

const maxFloat = math.MaxFloat64
