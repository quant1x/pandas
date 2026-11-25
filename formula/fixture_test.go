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
		n := parseOptionalInt(t, record[3])
		m := parseOptionalInt(t, record[4])
		want := parseList(t, record[5])

		series := pandas.Vector(x)
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
