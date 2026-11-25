//go:build onnx

// ONNX 推理对拍（M8，Go 侧）：跑 `onnx/mlp.onnx`，与 `features.csv` 的 `pred_mlp` 列逐值比对。
//
// cgo 绑定在 `onnx.go`（同为 `onnx` 标签）；构建与运行方式见该文件头部注释。
package features

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
)

var onnxFeatures = []string{"ma5", "ma20", "ema12", "wma10", "hhv20", "llv20",
	"std20", "hhvbars20", "ret1", "bias5", "cross_up"}

// TestONNX 读 features.csv 的 11 个特征跑 mlp.onnx，与 pred_mlp 逐值比对。
//
// 需要 ORT_DYLIB_PATH 指向 onnxruntime 动态库；未设置时跳过（与 Rust 侧同款策略）。
func TestONNX(t *testing.T) {
	dylib := os.Getenv("ORT_DYLIB_PATH")
	if dylib == "" {
		t.Skip("跳过：未设置 ORT_DYLIB_PATH")
	}
	if err := LoadRuntime(dylib); err != nil {
		t.Fatal(err)
	}

	modelPath, err := filepath.Abs("onnx/mlp.onnx")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(modelPath); err != nil {
		t.Fatalf("找不到模型: %v", err)
	}

	inputs, expected, rows, err := loadONNXInputs("features.csv")
	if err != nil {
		t.Fatal(err)
	}

	output, err := InferONNX(modelPath, inputs, rows)
	if err != nil {
		t.Fatal(err)
	}

	for index, want := range expected {
		got := output[index]
		if math.IsNaN(got) && math.IsNaN(want) {
			continue
		}
		if math.Abs(got-want) <= 1e-12*math.Max(1, math.Abs(want)) {
			continue
		}
		t.Fatalf("第 %d 位 %v != %v", index, got, want)
	}
	t.Logf("onnx 推理与 pred_mlp 一致（%d 行）", rows)
}

func loadONNXInputs(path string) ([]float64, []float64, int, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, nil, 0, err
	}
	lines := strings.Split(strings.TrimRight(string(raw), "\n"), "\n")
	if len(lines) < 2 {
		return nil, nil, 0, fmt.Errorf("fixture 没有数据行")
	}
	header := strings.Split(lines[0], ",")
	index := make(map[string]int, len(header))
	for i, name := range header {
		index[name] = i
	}
	var inputs, expected []float64
	for _, line := range lines[1:] {
		if line == "" {
			continue
		}
		fields := strings.Split(line, ",")
		for _, name := range onnxFeatures {
			value, err := parseONNXValue(fields[index[name]])
			if err != nil {
				return nil, nil, 0, err
			}
			inputs = append(inputs, value)
		}
		value, err := parseONNXValue(fields[index["pred_mlp"]])
		if err != nil {
			return nil, nil, 0, err
		}
		expected = append(expected, value)
	}
	return inputs, expected, len(expected), nil
}

func parseONNXValue(text string) (float64, error) {
	if strings.TrimSpace(text) == "nan" {
		return math.NaN(), nil
	}
	return strconv.ParseFloat(strings.TrimSpace(text), 64)
}
