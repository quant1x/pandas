package pandas

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

// String implements the Stringer interface for DataFrame
func (self DataFrame) String() (str string) {
	return self.print(true, false, true, true, 10, 70, "DataFrame")
}

func (self DataFrame) print(
	shortRows, shortCols, showDims, showTypes bool,
	maxRows int,
	maxCharsTotal int,
	class string) (str string) {

	addRightPadding := func(s string, nchar int) string {
		if utf8.RuneCountInString(s) < nchar {
			return s + strings.Repeat(" ", nchar-utf8.RuneCountInString(s))
		}
		return s
	}

	addLeftPadding := func(s string, nchar int) string {
		if utf8.RuneCountInString(s) < nchar {
			return strings.Repeat(" ", nchar-utf8.RuneCountInString(s)) + s
		}
		return s
	}

	if self.Err != nil {
		str = fmt.Sprintf("%s error: %v", class, self.Err)
		return
	}
	nMinRows := maxRows / 2
	nTotal := 0
	nrows, ncols := self.Dims()
	if nrows == 0 || ncols == 0 {
		str = fmt.Sprintf("Empty %s", class)
		return
	}
	var records [][]string
	shortening := false
	if shortRows && nrows > maxRows {
		shortening = true
		dfHead := self.Subset(0, nMinRows)
		records = dfHead.Records()
		nTotal += dfHead.Nrow()
		if shortening {
			dots := make([]string, ncols)
			for i := 0; i < ncols; i++ {
				dots[i] = "..."
			}
			records = append(records, dots)
		}
		nTotal += 1
		dfTail := self.Subset(nrows-nMinRows, nrows)
		tails := dfTail.Records()
		nTotal += dfTail.Nrow()
		records = append(records, tails[1:]...)
	} else {
		records = self.Records()
		nTotal += self.Nrow()
	}

	if showDims {
		str += fmt.Sprintf("[%dx%d] %s\n\n", nrows, ncols, class)
	}

	// Add the row numbers
	rowNumbersOffset := 0
	for i := 0; i < nTotal+1; /*self.nrows+1*/ i++ {
		add := ""
		if i == 0 {
			// 跳过
		} else if i == nMinRows+1 && shortening {
			// 跳过
			rowNumbersOffset -= 1
		} else if i >= nMinRows+1 && shortening {
			add = strconv.Itoa(nrows-maxRows+i-1+rowNumbersOffset) + ":"
		} else {
			add = strconv.Itoa(i-1+rowNumbersOffset) + ":"
		}
		//fmt.Println(i)
		records[i] = append([]string{add}, records[i]...)
	}
	//if shortening {
	//	dots := make([]string, ncols+1)
	//	for i := 1; i < ncols+1; i++ {
	//		dots[i] = "..."
	//	}
	//	records = append(records, dots)
	//}
	types := self.Types()
	typesrow := make([]string, ncols)
	for i := 0; i < ncols; i++ {
		typesrow[i] = fmt.Sprintf("<%v>", types[i])
	}
	typesrow = append([]string{""}, typesrow...)

	if showTypes {
		records = append(records, typesrow)
	}

	maxChars := make([]int, self.ncols+1)
	for i := 0; i < len(records); i++ {
		for j := 0; j < self.ncols+1; j++ {
			// Escape special characters
			records[i][j] = strconv.Quote(records[i][j])
			records[i][j] = records[i][j][1 : len(records[i][j])-1]

			// Detect maximum number of characters per column
			if len(records[i][j]) > maxChars[j] {
				maxChars[j] = utf8.RuneCountInString(records[i][j])
			}
		}
	}
	maxCols := len(records[0])
	var notShowing []string
	if shortCols {
		maxCharsCum := 0
		for colnum, m := range maxChars {
			maxCharsCum += m
			if maxCharsCum > maxCharsTotal {
				maxCols = colnum
				break
			}
		}
		notShowingNames := records[0][maxCols:]
		notShowingTypes := typesrow[maxCols:]
		notShowing = make([]string, len(notShowingNames))
		for i := 0; i < len(notShowingNames); i++ {
			notShowing[i] = fmt.Sprintf("%s %s", notShowingNames[i], notShowingTypes[i])
		}
	}
	for i := 0; i < len(records); i++ {
		// Add right padding to all elements
		records[i][0] = addLeftPadding(records[i][0], maxChars[0]+1)
		for j := 1; j < self.ncols; j++ {
			records[i][j] = addRightPadding(records[i][j], maxChars[j])
		}
		records[i] = records[i][0:maxCols]
		if shortCols && len(notShowing) != 0 {
			records[i] = append(records[i], "...")
		}
		// Create the final string
		str += strings.Join(records[i], " ")
		str += "\n"
	}
	// 没有显示字段处理逻辑
	if shortCols && len(notShowing) != 0 {
		var notShown string
		var notShownArr [][]string
		cum := 0
		i := 0
		for n, ns := range notShowing {
			cum += len(ns)
			if cum > maxCharsTotal {
				notShownArr = append(notShownArr, notShowing[i:n])
				cum = 0
				i = n
			}
		}
		if i < len(notShowing) {
			notShownArr = append(notShownArr, notShowing[i:])
		}
		for k, ns := range notShownArr {
			notShown += strings.Join(ns, ", ")
			if k != len(notShownArr)-1 {
				notShown += ","
			}
			notShown += "\n"
		}
		str += fmt.Sprintf("\nNot Showing: %s", notShown)
	}
	return str
}
