package pandas

import (
	"gitee.com/quant1x/gox/api"
	"gitee.com/quant1x/gox/util/homedir"
	"gitee.com/quant1x/pandas/internal/csvreader"
	"github.com/gocarina/gocsv"
	"os"
)

// Csv2Slices CSV文件转struct切片
func Csv2Slices[S ~[]E, E any](filename string, pointer *S) error {
	filepath, err := homedir.Expand(filename)
	if err != nil {
		return err
	}
	decoder := csvreader.New()
	return decoder.UnMarshalFile(filepath, pointer)
}

// Slices2Csv struct切片保存csv文件
func Slices2Csv[S ~[]E, E any](filename string, s S) error {
	filepath, err := homedir.Expand(filename)
	if err != nil {
		return err
	}
	// 检查目录, 不存在就创建
	_ = api.CheckFilepath(filepath, true)
	csvFile, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer api.CloseQuietly(csvFile)
	err = gocsv.MarshalFile(s, csvFile)
	return err
}
