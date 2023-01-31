package pandas

import (
	"encoding/csv"
	"github.com/mymmsc/gox/api"
	"github.com/mymmsc/gox/util/homedir"
	"io"
	"os"
)

// ReadCSV reads a CSV file from a io.Reader and builds a DataFrame with the
// resulting records.
func ReadCSV(r io.Reader, options ...LoadOption) DataFrame {
	csvReader := csv.NewReader(r)
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

// WriteOption is the type used to configure the writing of elements
type WriteOption func(*writeOptions)

type writeOptions struct {
	// Specifies whether the header is also written
	writeHeader bool
}

// WriteHeader sets the writeHeader option for writeOptions.
func WriteHeader(b bool) WriteOption {
	return func(c *writeOptions) {
		c.writeHeader = b
	}
}

// WriteCSV writes the DataFrame to the given io.Writer as a CSV file.
func (df DataFrame) WriteCSV(out any, options ...WriteOption) error {
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
	//if df.Err != nil {
	//	return df.Err
	//}
	if !IsEmpty(filename) {
		filepath, err := homedir.Expand(filename)
		if err != nil {
			return err
		}
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

	records := df.Records()
	if !cfg.writeHeader {
		records = records[1:]
	}

	return csv.NewWriter(writer).WriteAll(records)
}

// ToCSV 写csv格式文件
func (self DataFrame) oldToCSV(filename string, options ...WriteOption) error {
	filepath, err := homedir.Expand(filename)
	if err != nil {
		return err
	}
	csvFile, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer api.CloseQuietly(csvFile)
	err = self.WriteCSV(csvFile, options...)
	return err
}
