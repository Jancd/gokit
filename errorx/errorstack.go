package errorx

import (
	"fmt"
	"runtime"

	"github.com/spf13/cast"
)

// WithTrack return the error stack info
func WithTrack(err error) error {
	return fmt.Errorf("err: %s  stack: %s", err.Error(), getStackInfo(2))
}

func getStackInfo(skip int) string {
	line, funcName := 0, "???"
	pc, _, line, ok := runtime.Caller(skip)
	if ok {
		funcName = runtime.FuncForPC(pc).Name()
	}
	return "funcName: " + funcName + "  line: " + cast.ToString(line)
}
