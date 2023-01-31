package pandas

import "gitee.com/quant1x/pandas/algorithms"

// Mean gonum.org/v1/gonum/stat不支持整型, 每次都要转换有点难受啊
func Mean[T algorithms.Number](x []T) float64 {
	d := algorithms.Mean_Go(x)
	return float64(d)
}
