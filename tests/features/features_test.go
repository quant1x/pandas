// Package features 读取 tests/features/features.csv，用 formula 包重算特征列并逐值对拍。
//
// fixture 由 gen_features.py 生成：输入 OHLCV 与期望特征同表，缺失写 nan。
package features

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/quant1x/pandas"
	"github.com/quant1x/pandas/formula"
)

// featureColumns 是参与对拍的特征列，顺序即 CSV 表头顺序（也是模型求和顺序）。
var featureColumns = []string{"ma5", "ma20", "ema12", "wma10", "hhv20", "llv20", "std20", "hhvbars20", "ret1", "bias5", "cross_up"}

type table struct {
	close  []float64
	high   []float64
	low    []float64
	column map[string][]float64
}

func loadTable(t *testing.T) *table {
	t.Helper()
	file, err := os.Open("features.csv")
	if err != nil {
		t.Fatalf("打开 fixture: %v", err)
	}
	defer file.Close()

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		t.Fatalf("解析 fixture: %v", err)
	}
	if len(records) < 2 {
		t.Fatal("fixture 没有数据行")
	}
	index := make(map[string]int, len(records[0]))
	for i, name := range records[0] {
		index[name] = i
	}
	out := &table{column: make(map[string][]float64)}
	for _, record := range records[1:] {
		out.close = append(out.close, parseValue(t, record[index["close"]]))
		out.high = append(out.high, parseValue(t, record[index["high"]]))
		out.low = append(out.low, parseValue(t, record[index["low"]]))
		for _, name := range append(featureColumns, "pred", "pred_mlp") {
			out.column[name] = append(out.column[name], parseValue(t, record[index[name]]))
		}
	}
	return out
}

func parseValue(t *testing.T, text string) float64 {
	t.Helper()
	text = strings.TrimSpace(text)
	if text == "nan" {
		return math.NaN()
	}
	value, err := strconv.ParseFloat(text, 64)
	if err != nil {
		t.Fatalf("非法数值 %q: %v", text, err)
	}
	return value
}

func expectColumn(t *testing.T, name string, got, want []float64) {
	t.Helper()
	if len(got) != len(want) {
		t.Fatalf("%s: 长度 %d != %d", name, len(got), len(want))
	}
	for i := range want {
		if math.IsNaN(got[i]) && math.IsNaN(want[i]) {
			continue
		}
		if math.Abs(got[i]-want[i]) <= 1e-12*math.Max(1, math.Abs(want[i])) {
			continue
		}
		t.Fatalf("%s: 第 %d 位 %v != %v", name, i, got[i], want[i])
	}
}

// loadModel 读取模型权重（term,weight）。
func loadModel(t *testing.T, name string) map[string]float64 {
	t.Helper()
	file, err := os.Open(name)
	if err != nil {
		t.Fatalf("打开模型: %v", err)
	}
	defer file.Close()

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		t.Fatalf("解析模型: %v", err)
	}
	weights := make(map[string]float64, len(records))
	for _, record := range records[1:] {
		value, err := strconv.ParseFloat(strings.TrimSpace(record[1]), 64)
		if err != nil {
			t.Fatalf("非法权重 %q: %v", record[1], err)
		}
		weights[strings.TrimSpace(record[0])] = value
	}
	return weights
}

// predict 按表头顺序线性组合，最后加 bias；任一特征缺失则为 NaN。
func predict(weights map[string]float64, features map[string][]float64, index int) float64 {
	total := weights["bias"]
	for _, name := range featureColumns {
		value := features[name][index]
		if math.IsNaN(value) {
			return math.NaN()
		}
		total += weights[name] * value
	}
	return total
}

// predictMLP 两级小网络：h_j = b1_j + Σ_i W1[i][j]*f_i，y = b2 + Σ_j W2[j]*tanh(h_j)。
func predictMLP(weights map[string]float64, features map[string][]float64, index int, hidden int) float64 {
	values := make([]float64, len(featureColumns))
	for position, name := range featureColumns {
		value := features[name][index]
		if math.IsNaN(value) {
			return math.NaN()
		}
		values[position] = value
	}
	activations := make([]float64, hidden)
	for j := 0; j < hidden; j++ {
		total := weights[fmt.Sprintf("b1_%d", j)]
		for i, value := range values {
			total += weights[fmt.Sprintf("w1_%d_%d", i, j)] * value
		}
		activations[j] = math.Tanh(total)
	}
	total := weights["b2"]
	for j := 0; j < hidden; j++ {
		total += weights[fmt.Sprintf("w2_%d", j)] * activations[j]
	}
	return total
}

func TestFeatureFixtures(t *testing.T) {
	data := loadTable(t)
	series := pandas.Vector(data.close)

	ma5 := formula.MA(series, 5).DTypes()
	ma20 := formula.MA(series, 20).DTypes()
	ema12 := formula.EMA(series, 12).DTypes()
	hhv20 := formula.HHV(pandas.Vector(data.high), 20).DTypes()
	llv20 := formula.LLV(pandas.Vector(data.low), 20).DTypes()
	wma10 := formula.WMA(series, 10).DTypes()
	std20 := formula.STD(series, 20).DTypes()
	hhvbars20 := formula.HHVBARS(pandas.Vector(data.high), 20).DTypes()
	crossUp := formula.CROSS(pandas.Vector(ma5), pandas.Vector(ma20)).DTypes()

	ret1 := make([]float64, len(data.close))
	ret1[0] = math.NaN()
	for i := 1; i < len(data.close); i++ {
		ret1[i] = data.close[i]/data.close[i-1] - 1
	}
	bias5 := make([]float64, len(data.close))
	for i := range data.close {
		if math.IsNaN(ma5[i]) {
			bias5[i] = math.NaN()
			continue
		}
		bias5[i] = data.close[i]/ma5[i] - 1
	}

	computed := map[string][]float64{
		"ma5": ma5, "ma20": ma20, "ema12": ema12, "wma10": wma10,
		"hhv20": hhv20, "llv20": llv20, "std20": std20, "hhvbars20": hhvbars20,
		"ret1": ret1, "bias5": bias5, "cross_up": crossUp,
	}
	for _, name := range featureColumns {
		expectColumn(t, name, computed[name], data.column[name])
	}

	weights := loadModel(t, "model.csv")
	pred := make([]float64, len(data.close))
	for i := range pred {
		pred[i] = predict(weights, computed, i)
	}
	expectColumn(t, "pred", pred, data.column["pred"])

	mlpWeights := loadModel(t, "model_mlp.csv")
	predMLP := make([]float64, len(data.close))
	for i := range predMLP {
		predMLP[i] = predictMLP(mlpWeights, computed, i, 4)
	}
	expectColumn(t, "pred_mlp", predMLP, data.column["pred_mlp"])
}
