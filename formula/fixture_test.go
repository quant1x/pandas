package formula

import (
	"encoding/csv"
	"math"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/quant1x/pandas"
)

func TestFormulaFixtures(t *testing.T) {
	path := "../tests/formula/cases.csv"
	file, err := os.Open(path)
	if err != nil {
		t.Fatalf("打开 fixture: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		t.Fatalf("解析 fixture: %v", err)
	}
	if len(records) < 2 {
		t.Fatal("fixture 没有数据行")
	}
	for line, record := range records[1:] {
		if len(record) != 6 {
			t.Fatalf("第 %d 行有 %d 列", line+2, len(record))
		}
		name := record[0]
		x := parseList(t, record[1])
		y := parseList(t, record[2])
		n := parseOptionalInt(t, record[3])
		m := parseOptionalInt(t, record[4])
		want := parseList(t, record[5])

		series := pandas.Vector(x)
		pair := pandas.Vector(y)
		var got []float64
		switch name {
		case "abs":
			got = ABS(series).DTypes()
		case "ref":
			got = REF(series, n).DTypes()
		case "ma":
			got = MA(series, n).DTypes()
		case "ema":
			got = EMA(series, n).DTypes()
		case "sma":
			got = SMA(series, n, m).DTypes()
		case "hhv":
			got = HHV(series, n).DTypes()
		case "llv":
			got = LLV(series, n).DTypes()
		case "sum":
			got = SUM(series, n).DTypes()
		case "std":
			got = STD(series, n).DTypes()
		case "wma":
			got = WMA(series, n).DTypes()
		case "sqrt":
			got = SQRT(series)
		case "count":
			got = COUNT(series, n).DTypes()
		case "filter":
			got = FILTER(series, n).DTypes()
		case "hhvbars":
			got = HHVBARS(series, n).DTypes()
		case "llvbars":
			got = LLVBARS(series, n).DTypes()
		case "cross":
			got = CROSS(series, pair).DTypes()
		case "dma":
			got = DMA(series, pair)
		case "max":
			got = MAX(series, pair).DTypes()
		case "min":
			got = MIN(series, pair).DTypes()
		case "gt":
			got = boolsToFloat(CompareGt(x, y))
		case "gte":
			got = boolsToFloat(CompareGte(x, y))
		case "lt":
			got = boolsToFloat(CompareLt(x, y))
		case "lte":
			got = boolsToFloat(CompareLte(x, y))
		case "eq":
			got = boolsToFloat(EQ(series, pair))
		case "neq":
			got = boolsToFloat(NEQ(series, pair))
		case "and":
			got = boolsToFloat(AND(floatsToBools(x), floatsToBools(y)))
		case "or":
			got = boolsToFloat(OR(floatsToBools(x), floatsToBools(y)))
		case "not":
			got = NOT(series).DTypes()
		case "barslast":
			got = BARSLAST(series).DTypes()
		case "barslastcount":
			got = BARSLASTCOUNT(series).DTypes()
		case "barslasts":
			got = BARSLASTS(series, int(n)).DTypes()
		case "barssincen":
			got = BARSSINCEN(series, n).DTypes()
		case "if":
			got = IF(series, pair, int(n)).DTypes()
		case "ifn":
			got = IFN(series, pair, int(n)).DTypes()
		case "avedev":
			got = AVEDEV(series, n).DTypes()
		case "slope":
			got = SLOPE(series, n).([]float64)
		case "forcast":
			got = FORCAST(series, n).([]float64)
		case "const":
			got = CONST(series).DTypes()
		case "last":
			got = LAST(series, int(n), int(m)).DTypes()
		default:
			t.Fatalf("%s: 未覆盖的公式", name)
		}
		if len(got) != len(want) {
			t.Fatalf("%s: 长度 %d != %d", name, len(got), len(want))
		}
		for i := range want {
			if equalFixture(got[i], want[i]) {
				continue
			}
			if math.Abs(got[i]-want[i]) <= 1e-12*math.Max(1, math.Abs(want[i])) {
				continue
			}
			t.Fatalf("%s: 第 %d 位 %v != %v", name, i, got[i], want[i])
		}
	}
}

// boolsToFloat 把布尔序列落成 0/1，供 fixture 比对。
func boolsToFloat(flags []bool) []float64 {
	out := make([]float64, len(flags))
	for i, flag := range flags {
		if flag {
			out[i] = 1
		}
	}
	return out
}

// floatsToBools 把 0/1 序列转成布尔序列，非零即真。
func floatsToBools(values []float64) []bool {
	out := make([]bool, len(values))
	for i, value := range values {
		out[i] = value != 0
	}
	return out
}

func parseList(t *testing.T, text string) []float64 {
	t.Helper()
	if text == "" {
		return nil
	}
	parts := strings.Split(text, ";")
	out := make([]float64, len(parts))
	for i, part := range parts {
		part = strings.TrimSpace(part)
		if part == "nan" {
			out[i] = math.NaN()
			continue
		}
		value, err := strconv.ParseFloat(part, 64)
		if err != nil {
			t.Fatalf("解析数值 %q: %v", part, err)
		}
		out[i] = value
	}
	return out
}

func parseOptionalInt(t *testing.T, text string) int {
	t.Helper()
	if text == "" {
		return 0
	}
	value, err := strconv.Atoi(text)
	if err != nil {
		t.Fatalf("解析整数 %q: %v", text, err)
	}
	return value
}

func equalFixture(a, b float64) bool {
	if math.IsNaN(a) && math.IsNaN(b) {
		return true
	}
	return a == b
}
