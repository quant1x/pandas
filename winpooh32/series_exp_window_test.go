package winpooh32

import (
	"testing"
)

func TestExpWindow_Mean(t *testing.T) {
	type fields struct {
		data     Series
		atype    AlphaType
		param    DType
		adjust   bool
		ignoreNA bool
	}
	tests := []struct {
		name   string
		fields fields
		want   Series
	}{
		{
			"simple",
			fields{
				data: MakeData(
					1,
					[]int64{1, 2, 3, 4, 5},
					[]DType{0, 1, 2, NaN, 4},
				),
				atype:    AlphaCom,
				param:    0.5,
				adjust:   true,
				ignoreNA: false,
			},
			MakeData(
				1,
				[]int64{1, 2, 3, 4, 5},
				[]DType{0, 0.6923077, 1.575, 1.575, 3.198347},
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := ExponentialMovingWindow{
				data:     tt.fields.data,
				atype:    tt.fields.atype,
				param:    tt.fields.param,
				adjust:   tt.fields.adjust,
				ignoreNA: tt.fields.ignoreNA,
			}

			if got := w.Mean(); !got.Equals(tt.want, 10e-4) {
				t.Errorf("ExponentialMovingWindow.Mean() = %v, want %v", got, tt.want)
			}
		})
	}
}
