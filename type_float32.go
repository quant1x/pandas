package pandas

import (
	"fmt"
	"github.com/mymmsc/gox/logger"
	"math"
	"strconv"
)

const (
	MaxFloat32                  = float32(math.MaxFloat32)             // float32最大值
	MinFloat32                  = float32(math.SmallestNonzeroFloat32) // float32最小值
	StringTrue2Float32  float32 = float32(1)                           // 字符串true转float32
	StringFalse2Float32 float32 = float32(0)                           // 字符串false转float32
)

// Float32IsNaN 判断float32是否NaN
func Float32IsNaN(f float32) bool {
	return Float64IsNaN(float64(f))
}

// ParseFloat32 字符串转float32
func ParseFloat32(s string, v any) float32 {
	defer func() {
		// 解析失败以后输出日志, 以备检查
		if err := recover(); err != nil {
			logger.Errorf("ParseFloat32 %+v, error=%+v\n", v, err)
		}
	}()

	if IsEmpty(s) {
		// TODO:NaN是针对64位, 这样直接转换应该有问题, 需要进一步确认
		return Nil2Float32
	}
	if isTrue(s) {
		return StringTrue2Float32
	} else if isFalse(s) {
		return StringFalse2Float32
	}

	f, err := strconv.ParseFloat(s, 32)
	if err == nil {
		return float32(f)
	}
	if IgnoreParseExceptions {
		return Nil2Float32
	}
	_ = v.(float32) // Intentionally panic
	return Nil2Float32
}

// float32转string
func float32ToString(v float32) string {
	if Float32IsNaN(v) {
		return StringNaN
	}
	return fmt.Sprintf("%f", v)
}

func AnyToFloat32(v any) float32 {
	if isPoint(v) {
		//vv := reflect.ValueOf(v)
		//if vv.Kind() == reflect.Pointer {
		//	if vv.IsNil() {
		//		return Nil2Float32
		//	}
		//	v = vv.Elem()
		//	vv.Float()
		//}
		return point_to_number[float32](v, Nil2Float32, boolToFloat32, ParseFloat32)
	}
	f := value_to_number[float32](v, Nil2Float32, boolToFloat32, ParseFloat32)
	return f
}
