package pandas

import (
	"gitee.com/quant1x/gox/api"
	"gitee.com/quant1x/gox/util/homedir"
	"gitee.com/quant1x/pandas/internal/csvreader"
	"github.com/gocarina/gocsv"
	"os"
)

const (
	DefaultTagName = "dataframe"
)

func init() {
	gocsv.TagName = DefaultTagName
}

// Csv2Slices CSV文件转struct切片
//
//	deprecated: Use CsvToSlices
func Csv2Slices[S ~[]E, E any](filename string, pointer *S) error {
	filepath, err := homedir.Expand(filename)
	if err != nil {
		return err
	}
	decoder := csvreader.New()
	return decoder.UnMarshalFile(filepath, pointer)
}

// Slices2Csv struct切片保存csv文件
//
//	deprecated: Use SlicesToCsv
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

// CsvToSlices CSV文件转struct切片
func CsvToSlices[S ~[]E, E any](filename string, pointer *S) error {
	filepath, err := homedir.Expand(filename)
	if err != nil {
		return err
	}
	csvFile, err := os.Open(filepath)
	if err != nil {
		return err
	}
	err = gocsv.Unmarshal(csvFile, pointer)
	return err
}

// SlicesToCsv struct切片转csv文件
func SlicesToCsv[S ~[]E, E any](filename string, s S) error {
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
