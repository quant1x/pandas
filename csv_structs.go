package pandas

import "gitee.com/quant1x/pandas/internal/csvreader"

// Csv2Slices CSV文件转struct切片
func Csv2Slices[S ~[]E, E any](filename string, pointer *S) error {
	decoder := csvreader.New()
	return decoder.UnMarshalFile(filename, pointer)
}
