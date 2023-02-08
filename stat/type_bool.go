package stat

const (
	Nil2Bool              = false      // 空指针转bool
	BoolNaN               = false      // bool 无效值
	True2Bool             = true       // true转bool
	False2Bool            = false      // false 转bool
	True2Float32  float32 = float32(1) // true转float32
	False2Float32 float32 = float32(0) // false转float32

	StringBad2Bool   = false // 字符串解析bool异常
	StringTrue2Bool  = true  // 字符串true转bool
	StringFalse2Bool = false // 字符串false转bool
)

func isTrue(s string) bool {
	if s == "true" || s == "TRUE" || s == "True" || s == "1" || s == "真" || s == "对" || s == "好" {
		return true
	} else {
		return false
	}
}

func isFalse(s string) bool {
	if s == "false" || s == "FALSE" || s == "False" || s == "0" || s == "假" || s == "错" || s == "坏" {
		return true
	} else {
		return false
	}
}

func boolToInt64(b bool) int64 {
	if b {
		return True2Int64
	}
	return False2Int64
}

// bool转float32
func boolToFloat32(b bool) float32 {
	if b {
		return True2Float32
	}
	return False2Float32
}

// bool转float64
func boolToFloat64(b bool) float64 {
	if b {
		return True2Float64
	}
	return False2Float64
}
