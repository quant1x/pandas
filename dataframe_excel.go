package pandas

import (
	"fmt"
	"gitee.com/quant1x/pandas/stat"
	"github.com/mymmsc/gox/logger"
	"github.com/mymmsc/gox/util/homedir"
	xlsv3 "github.com/tealeg/xlsx/v3"
	"strings"
)

// ReadExcel 读取excel文件
func ReadExcel(filename string, options ...LoadOption) DataFrame {
	if stat.IsEmpty(filename) {
		return DataFrame{Err: fmt.Errorf("filaname is empty")}
	}

	filepath, err := homedir.Expand(filename)
	if err != nil {
		logger.Errorf("%s, error=%+v\n", filename, err)
		return DataFrame{Err: err}
	}
	//filename := "test.xlsx"
	xlFile, err := xlsv3.OpenFile(filepath)
	if err != nil {
		return DataFrame{Err: err}
	}
	colnums := make([][]string, 0)
	for _, sheet := range xlFile.Sheets {
		//fmt.Printf("Sheet Name: %s\n", sheet.Name)
		//for _, row := range sheet.Rows {
		//	col := make([]string, 0)
		//	for _, cell := range row.Cells {
		//		//cell.SetStringFormula("%s")
		//		if cell.IsTime() {
		//			cell.SetFormat("yyyy-mm-dd")
		//		} else if strings.HasPrefix(cell.Value, "0") {
		//			cell.SetFormat("")
		//		}
		//		text := cell.String()
		//		col = append(col, text)
		//	}
		//	colnums = append(colnums, col)
		//}
		_ = sheet.ForEachRow(func(r *xlsv3.Row) error {
			col := make([]string, 0)
			_ = r.ForEachCell(func(c *xlsv3.Cell) error {
				if c.IsTime() {
					c.SetFormat("yyyy-mm-dd")
				} else if strings.HasPrefix(c.Value, "0") {
					c.SetFormat("")
				}
				text := c.String()
				col = append(col, text)
				return nil
			})
			colnums = append(colnums, col)
			return nil
		})
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
	xlFile := xlsv3.NewFile()
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
