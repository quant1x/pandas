package pandas

import (
	"encoding/csv"
	"io"
	"os"

	"github.com/quant1x/num"
	"github.com/quant1x/x/api"
	"github.com/quant1x/x/exception"
	"github.com/quant1x/x/util/homedir"
)

// ReadCSV reads a CSV file from a io.Reader and builds a DataFrame with the
//
//	resulting records.
//	支持文件名和io两种方式读取数据
func ReadCSV(in any, options ...LoadOption) DataFrame {
	var (
		reader   io.Reader
		filename string
	)
	switch param := in.(type) {
	case io.Reader:
		reader = param
	case string:
		filename = param
	}

	if !num.IsEmpty(filename) {
		filepath, err := homedir.Expand(filename)
		if err != nil {
			return DataFrame{Err: err}
		}
		csvFile, err := os.Open(filepath)
		if err != nil {
			return DataFrame{Err: err}
		}
		defer api.CloseQuietly(csvFile)
		reader = csvFile
	}
	if reader == nil {
		return DataFrame{Err: exception.New(1, "csv文件操作失败")}
	}

	csvReader := csv.NewReader(reader)
	cfg := loadOptions{
		delimiter:  ',',
		lazyQuotes: false,
		comment:    0,
	}
	for _, option := range options {
		option(&cfg)
	}

	csvReader.Comma = cfg.delimiter
	csvReader.LazyQuotes = cfg.lazyQuotes
	csvReader.Comment = cfg.comment

	records, err := csvReader.ReadAll()
	if err != nil {
		return DataFrame{Err: err}
	}
	return LoadRecords(records, options...)
}

// WriteCSV writes the DataFrame to the given io.Writer as a CSV file.
//
//	支持文件名和io两种方式写入数据
func (this DataFrame) WriteCSV(out any, options ...WriteOption) error {
	var (
		writer   io.Writer
		filename string
	)
	switch param := out.(type) {
	case io.Writer:
		writer = param
	case string:
		filename = param
	}

	if !num.IsEmpty(filename) {
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
		writer = csvFile
	}
	// Set the default write options
	cfg := writeOptions{
		writeHeader: true,
	}

	// Set any custom write options
	for _, option := range options {
		option(&cfg)
	}

	records := this.Records()
	if !cfg.writeHeader {
		records = records[1:]
	}

	return csv.NewWriter(writer).WriteAll(records)
}
