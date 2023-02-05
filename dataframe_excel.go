package pandas

import (
	"fmt"
	"github.com/mymmsc/gox/logger"
	"github.com/mymmsc/gox/util/homedir"
	"github.com/tealeg/xlsx"
)

// 读取excel文件
func ReadExcel(filename string, options ...LoadOption) DataFrame {
	if IsEmpty(filename) {
		return DataFrame{Err: fmt.Errorf("filaname is empty")}
	}

	filepath, err := homedir.Expand(filename)
	if err != nil {
		logger.Errorf("%s, error=%+v\n", filename, err)
		return DataFrame{Err: err}
	}
	//filename := "test.xlsx"
	xlFile, err := xlsx.OpenFile(filepath)
	if err != nil {
		return DataFrame{Err: err}
	}
	colnums := make([][]string, 0)
	for _, sheet := range xlFile.Sheets {
		//fmt.Printf("Sheet Name: %s\n", sheet.Name)
		for _, row := range sheet.Rows {
			col := make([]string, 0)
			for _, cell := range row.Cells {
				text := cell.String()
				col = append(col, text)
			}
			colnums = append(colnums, col)
		}
		// 只展示第一个sheet
		break
	}

	return LoadRecords(colnums, options...)
}

// WriteExcel 支持文件名和io两种方式写入数据
func (self DataFrame) WriteExcel(filename string, options ...WriteOption) error {
	filepath, err := homedir.Expand(filename)
	if err != nil {
		return err
	}
	xlFile := xlsx.NewFile()
	sheet, err := xlFile.AddSheet("Sheet(pandas)")
	if err != nil {
		return err
	}
	// Set the default write options
	cfg := writeOptions{
		writeHeader: true,
	}

	// Set any custom write options
	for _, option := range options {
		option(&cfg)
	}

	records := self.Records()
	if !cfg.writeHeader {
		records = records[1:]
	}
	for _, cols := range records {
		row := sheet.AddRow()
		for _, col := range cols {
			cell := row.AddCell()
			cell.SetString(col)
		}
	}

	return xlFile.Save(filepath)
}
