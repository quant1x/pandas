package stat

import (
	"bytes"
	"fmt"
	"runtime"
)

// PanicTrace panic 堆栈信息
func PanicTrace(err interface{}) string {
	buf := new(bytes.Buffer)
	_, _ = fmt.Fprintf(buf, "%v\n", err)
	for i := 1; ; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		_, _ = fmt.Fprintf(buf, "%s:%d (0x%x)\n", file, line, pc)
	}
	return buf.String()
}
